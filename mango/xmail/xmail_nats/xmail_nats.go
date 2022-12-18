package xmail_nats

import (
	"github.com/khanakia/mangobp/mango/natso"
	"gorm.io/gorm"
)

type Config struct {
	DB    *gorm.DB
	Natso natso.Natso
}

type AuthNats struct {
	Config
}

func New(config Config) {
	// ec := config.Natso.GetEncodedConn()

}
