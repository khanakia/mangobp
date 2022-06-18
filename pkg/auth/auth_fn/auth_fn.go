package auth_fn

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/khanakia/mangobp/mango/publicid"
	"github.com/khanakia/mangobp/pkg/auth/auth_domain"
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

func CheckEmailExists(email string, db *gorm.DB) bool {
	var count int64
	db.Model(auth_domain.User{}).Where("email = ?", email).Count(&count)
	if count > 0 {
		return true
	}

	return false
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
