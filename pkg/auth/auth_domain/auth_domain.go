package auth_domain

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/khanakia/mangobp/mango/util"
)

const (
	RoleSaID     = 1
	RoleMemberID = 2
)

type User struct {
	util.ModelStr
	IdNum            uint   `json:"idNum" gorm:"->;autoIncrement;unique;not_null"`
	Email            string `json:"email" gorm:"type:varchar(255);unique;"`
	FirstName        string `json:"firstName" gorm:"type:varchar(255)"`
	LastName         string `json:"lastName" gorm:"type:varchar(255)"`
	Company          string `json:"company" gorm:"type:varchar(255)"`
	Status           bool   `json:"status" gorm:"type:boolean;default:true"`
	RoleID           uint   `json:"roleId"`
	Password         string `json:"-" gorm:"type:varchar(250)"`
	Secret           string `json:"-" gorm:"type:varchar(50)"` // Will be used for Login or Other function this will be internal and must never shared to frotend
	WelcomeEmailSent bool   `json:"welcomeEmailSent" gorm:"type:boolean;default:false"`
	// ChgbeeCustomerID string `json:"-" gorm:"type:varchar(50)"`
}

type Claims struct {
	Sub   string `json:"sub"`
	Email string `json:"email"`
	// Shop    string    `json:"shop"`
	Name    string `json:"name"`
	Company string `json:"company"`
	// UID   string    `json:"uid"`
	jwt.StandardClaims
}
