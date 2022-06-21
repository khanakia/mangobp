package logdb

import (
	"github.com/khanakia/mangobp/mango/cache_nats_client"
	"github.com/khanakia/mangobp/mango/cli"
	"github.com/khanakia/mangobp/mango/gormdb"
	"github.com/khanakia/mangobp/mango/logdb/logdb_domain"
	"github.com/khanakia/mangobp/mango/logdb/logdb_nats"
	"github.com/khanakia/mangobp/mango/natso"
)

type Config struct {
	Cli             cli.Cli
	GormDB          gormdb.GormDB
	Natso           natso.Natso
	CacheNatsClient cache_nats_client.CacheNatsClient
}

type LogDb struct {
	Config
}

func (pkg LogDb) Version() string {
	return "0.01"
}

func (pkg LogDb) MigrateDb() {
	pkg.GormDB.DB.AutoMigrate(&logdb_domain.Log{})
}

func New(config Config) LogDb {
	pkg := LogDb{
		Config: config,
	}

	logdb_nats.New(logdb_nats.Config{
		Natso:           config.Natso,
		DB:              config.GormDB.DB,
		CacheNatsClient: config.CacheNatsClient,
	})

	return pkg
}
