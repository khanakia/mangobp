package logdb_nats

import (
	"github.com/khanakia/mangobp/mango/logdb/logdb_domain"
	"github.com/khanakia/mangobp/mango/logdb/logdb_fn"
	"github.com/khanakia/mangobp/mango/natso"
	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
)

type Config struct {
	Natso natso.Natso
	DB    *gorm.DB
}

type LogDbNats struct {
	Config
}

func New(config Config) {
	ec := config.Natso.GetEncodedConn()
	LogSubs(ec, config)
}

func LogSubs(ec *nats.EncodedConn, config Config) {
	ec.Subscribe(logdb_domain.NATS_LOGDB_LOG, func(subj, reply string, msg logdb_domain.CreateArgs) {
		// fmt.Println("MSG", msg)
		// goutil.PrintToJSON(msg)
		db := config.DB
		logdb_fn.Create(msg, db)
	})
}
