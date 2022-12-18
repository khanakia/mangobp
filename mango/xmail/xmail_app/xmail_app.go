package xmail_app

import (
	"github.com/khanakia/mangobp/mango/cli"
	"github.com/khanakia/mangobp/mango/gormdb"
	"github.com/khanakia/mangobp/mango/natso"
	"github.com/khanakia/mangobp/mango/xmail/xmail_dm"
	"github.com/khanakia/mangobp/mango/xmail/xmail_nats"
)

type Config struct {
	Cli    cli.Cli
	GormDB gormdb.GormDB
	Natso  natso.Natso
}

type Xmail struct {
	Config
}

func (pkg Xmail) Version() string {
	return "0.01"
}

func (pkg Xmail) MigrateDb() {
	pkg.GormDB.DB.AutoMigrate(&xmail_dm.Channel{})
}

func New(config Config) Xmail {
	pkg := Xmail{
		Config: config,
	}

	xmail_nats.New(xmail_nats.Config{
		Natso: config.Natso,
		DB:    config.GormDB.DB,
	})
	return pkg
}
