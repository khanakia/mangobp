package auth_fn

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/khanakia/mangobp/mango/publicid"
	"github.com/khanakia/mangobp/mango/util"
	"github.com/khanakia/mangobp/pkg/auth/auth_domain"
	"github.com/ubgo/gofm/cache"
	"github.com/ubgo/goutil"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// GeneratePassword - Create Bcrypt from string
func GeneratePassword(password string) string {
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(passwordHash)
}

// PasswordMatch - Compare two passwords are equal
func PasswordMatch(password string, password1 string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(password1))
	return err == nil
}

func RandomPass() string {
	rand.Seed(time.Now().UnixNano())

	chars := []rune("0123456789")
	length := 8
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String() // E.g. "ExcbsVQs"
	return str
}

type UserMethod struct {
	User *auth_domain.User
}

func (tm UserMethod) FillDefaults() {
	user := tm.User
	if user.RoleID == 0 {
		user.RoleID = auth_domain.RoleMemberID
	}

	if !user.Status {
		user.Status = true
	}

	if len(user.Password) <= 0 {
		randomPass := RandomPass()
		user.Password = GeneratePassword(randomPass)
	} else {
		user.Password = GeneratePassword(user.Password)
	}

	if len(user.Secret) <= 0 {
		user.Secret = publicid.Must12()
	}

	// if len(user.APIKey) <= 0 {
	// 	user.APIKey = uuid.New().String()
	// }
}

func (tm UserMethod) ValidateBeforeInsert() ([]interface{}, error) {
	user := *tm.User

	errorValidate := validation.ValidateStruct(&user,
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.RoleID, validation.Required),
		validation.Field(&user.Password, validation.Required),
	)

	fmt.Println(errorValidate)
	// data := util.ErroObjToArray(errorValidate)
	return nil, errorValidate
}

func GetUserSecret(user auth_domain.User, db *gorm.DB) string {
	if len(user.Secret) == 0 {
		user.Secret = publicid.Must24()

		db.Model(&user).Updates(&auth_domain.User{
			Secret: user.Secret,
		})
		return user.Secret
	}
	return user.Secret
}

/*
 * This signature is used to create a JWT token
 * Benefits - By adding the user secret with the appsecret it makes the app more secure
 * Let say if appSecret is compromised then stil nobody can generate tokens without user secret
 * If userSecret compromised it will affect only single user not all the users
 * If user wants to focefully logout for all the applications we simply update his userSecret
 * FUTURE CONSIDERATION - Add jwt to token to the blacklist if users logout
 */
func GetSignature(user auth_domain.User, db *gorm.DB) string {
	userSecret := GetUserSecret(user, db)

	appSecret := goutil.Env("appSecret", "IBIrewORShiVReBASTer")
	signature := appSecret + "_" + userSecret
	return signature
}

func CreateToken(user auth_domain.User, db *gorm.DB) (string, error) {
	// expirationTime := time.Now().Add(500 * time.Minute) // 500 minute

	// expire after 30 days
	expires := time.Now().AddDate(0, 0, 30)

	claims := &auth_domain.Claims{
		Sub:     user.ID,
		Email:   user.Email,
		Name:    user.FirstName + " " + user.LastName,
		Company: user.Company,
		// ID:    user.ID,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expires.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signature := []byte(GetSignature(user, db))
	tokenString, err := token.SignedString(signature)

	return tokenString, err
}

func ValidateJWT(token string, db *gorm.DB) (auth_domain.User, *auth_domain.Claims, error) {

	token = strings.TrimSpace(token)

	claims := &auth_domain.Claims{}

	_, _, err := new(jwt.Parser).ParseUnverified(token, claims)
	if err != nil {
		return auth_domain.User{}, nil, err
	}

	if len(claims.Sub) == 0 {
		return auth_domain.User{}, nil, errors.New("userID is empty")
	}

	var user auth_domain.User
	res := db.Where("id = ?", claims.Sub).First(&user)
	if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return auth_domain.User{}, nil, errors.New("auth_domain.User not found")
	}

	if !user.Status {
		return auth_domain.User{}, nil, errors.New("user is disabled")
	}

	// auth := ctr.getAuth()

	signature := []byte(GetSignature(user, db))

	tokenWC, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return signature, nil
	})

	if err != nil {
		return auth_domain.User{}, nil, err
	}

	if !tokenWC.Valid {
		return auth_domain.User{}, nil, errors.New("Unauthorized")
	}

	return user, claims, nil
}

func GetUserFromContext(c *gin.Context) (auth_domain.User, error) {
	userc, _ := c.Get("user")

	if userc == nil {
		return auth_domain.User{}, errors.New("User not found")
	}
	user, _ := userc.(auth_domain.User)
	return user, nil
}

func CheckEmailExists(email string, db *gorm.DB) bool {
	var count int64
	db.Model(auth_domain.User{}).Where("email = ?", email).Count(&count)
	if count > 0 {
		return true
	}

	return false
}

func FindByEmail(email string, db *gorm.DB) (*auth_domain.User, error) {
	var user auth_domain.User
	email = strings.ToLower(email)
	res := db.First(&user, &auth_domain.User{Email: email})

	if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}

	if res.Error != nil {
		return nil, errors.New("server error")
	}

	return &user, nil
}

func GetUser(id string, db *gorm.DB) (*auth_domain.User, error) {
	var user auth_domain.User
	res := db.Where("id = ?", id).First(&user)
	if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, res.Error
	}
	return &user, nil
}

func Register(user *auth_domain.User, db *gorm.DB) (*auth_domain.User, error) {
	// goutil.PrintToJSON(user)
	um := UserMethod{User: user}
	um.FillDefaults()

	_, err := um.ValidateBeforeInsert()
	if err != nil {
		return nil, err
	}

	// db := pkg.DB

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	isEmailExists := CheckEmailExists(user.Email, db)
	if isEmailExists {
		return nil, errors.New("email already exists")
	}

	err = db.Create(user).Error
	if err != nil {
		return nil, errors.New("server error. Please contact support")
	}

	return user, nil
}

func Login(userName string, password string, db *gorm.DB) (string, error) {
	user, err := FindByEmail(userName, db)
	if err != nil {
		return "", err
	}

	if !user.Status {
		return "", errors.New("user is disabled")
	}

	matched := PasswordMatch(user.Password, password)
	if !matched {
		return "", errors.New("password didn't match")
	}

	token, errToken := CreateToken(*user, db)
	if errToken != nil {
		return "", errors.New("cannot create login token")
	}

	return token, nil
}

type ForgotPasswordRequest struct {
	UserName string `json:"userName"`
}

// Validate ...
func (a ForgotPasswordRequest) Validate() ([]interface{}, error) {
	errorValidate := validation.ValidateStruct(&a,
		validation.Field(&a.UserName, validation.Required),
	)
	data := util.ErroObjToArray(errorValidate)
	return data, errorValidate
}

func ForgotPassword(req ForgotPasswordRequest, cache cache.Store, db *gorm.DB) error {

	user, err := FindByEmail(req.UserName, db)
	if err != nil {
		return errors.New("User not found.")
	}

	secret := publicid.Must24()
	token := "fp:" + secret

	cache.Put(token, user.ID, 1000)

	// YTD Send email
	// SendForgotPasswordEmail(*user, secret)

	return nil
}

type ResetPasswordRequest struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}

// Validate ...
func (a ResetPasswordRequest) Validate() ([]interface{}, error) {
	errorValidate := validation.ValidateStruct(&a,
		validation.Field(&a.Token, validation.Required),
		validation.Field(&a.Password, validation.Required),
	)
	data := util.ErroObjToArray(errorValidate)
	return data, errorValidate
}

func ResetPassword(req ResetPasswordRequest, cache cache.Store, db *gorm.DB) (auth_domain.User, error) {
	cacheKey := "fp:" + req.Token
	userID := cache.Get(cacheKey)

	if userID == nil {
		return auth_domain.User{}, errors.New("wrong token supplied")
	}

	var user auth_domain.User
	res := db.Where("id = ?", userID).First(&user)
	if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return auth_domain.User{}, errors.New("user does not exists")
	}

	if res.Error != nil {
		return auth_domain.User{}, errors.New("server error")
	}

	err := db.Model(&user).Updates(&auth_domain.User{
		Password: GeneratePassword(req.Password),
	}).Error

	if err != nil {
		return auth_domain.User{}, errors.New("server error")
	}

	cache.Del(cacheKey)

	return user, nil
}
