package auth_repo

import (
	"errors"
	"strings"

	"github.com/khanakia/mangobp/pkg/auth/auth_domain"
	"github.com/khanakia/mangobp/pkg/auth/auth_fn"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func (pkg UserRepo) FindByEmail(email string) (*auth_domain.User, error) {
	db := pkg.DB
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

func (pkg UserRepo) Register(user *auth_domain.User) (*auth_domain.User, error) {
	// goutil.PrintToJSON(user)
	um := auth_fn.UserMethod{User: user}
	um.FillDefaults()

	_, err := um.ValidateBeforeInsert()
	if err != nil {
		return nil, err
	}

	db := pkg.DB

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	isEmailExists := auth_fn.CheckEmailExists(user.Email, db)
	if isEmailExists {
		return nil, errors.New("email already exists")
	}

	err = db.Create(user).Error
	if err != nil {
		return nil, errors.New("server error. Please contact support")
	}

	return user, nil
}
