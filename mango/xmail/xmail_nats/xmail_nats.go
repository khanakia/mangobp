package xmail_nats

import (
	"github.com/khanakia/mangobp/mango/cache_nats_client"
	"github.com/khanakia/mangobp/mango/natso"
	"gorm.io/gorm"
)

type Config struct {
	DB              *gorm.DB
	Natso           natso.Natso
	CacheNatsClient cache_nats_client.CacheNatsClient
}

type AuthNats struct {
	Config
}

func New(config Config) {
	// ec := config.Natso.GetEncodedConn()

}
