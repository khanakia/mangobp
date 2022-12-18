//go:build wireinject
// +build wireinject

package wireapp

import (
	"github.com/google/wire"
	"github.com/khanakia/mangobp/mango/cache_nats"
	"github.com/khanakia/mangobp/mango/cacherdbms"
	"github.com/khanakia/mangobp/mango/cli"
	"github.com/khanakia/mangobp/mango/configmgr"
	"github.com/khanakia/mangobp/mango/dbconn"
	"github.com/khanakia/mangobp/mango/geo"
	"github.com/khanakia/mangobp/mango/gormdb"
	"github.com/khanakia/mangobp/mango/interfaces"
	"github.com/khanakia/mangobp/mango/logdb"
	"github.com/khanakia/mangobp/mango/natso"
	"github.com/khanakia/mangobp/mango/xmail/xmail_app"
	"github.com/khanakia/mangobp/pkg/auth/auth_app"
	"github.com/khanakia/mangobp/pkg/cache_natsapi"
	"github.com/khanakia/mangobp/pkg/dapp"
	"github.com/ubgo/gofm/cache"
	"github.com/ubgo/gofm/logger"
	"gorm.io/gorm"
)

func NewGormConfig(dbConn *dbconn.DbConn) gormdb.Config {
	return gormdb.Config{
		DB: dbConn.SqlDb,
	}
}

func NewConfigMgrConfig() configmgr.Config {
	return configmgr.Config{}
}

func NewGormDb(gormdb gormdb.GormDB) *gorm.DB {
	return gormdb.DB
}

// func NewConfigMgr(config configmgr.Config) *configmgr.ConfigMgr {
// 	return &configmgr.ConfigMgr{}
// }

func Init() Plugin {
	wire.Build(
		NewConfigMgrConfig,
		// NewConfigMgr,
		configmgr.New,
		wire.Bind(new(interfaces.IConfig), new(*configmgr.ConfigMgr)),
		cli.New,
		logger.New,
		wire.Struct(new(dbconn.Config), "*"),
		dbconn.New,
		NewGormConfig,
		gormdb.New,
		NewGormDb,
		wire.Struct(new(cacherdbms.Config), "*"),
		cacherdbms.New,
		// wire.Bind(new(cache.Store), new(*cacherdbms.Rdbms)),
		wire.Struct(new(cache_natsapi.Config), "*"),
		cache_natsapi.New,
		// wire.Struct(new(cache.Config), "*"),
		// cache.New,
		wire.Struct(new(natso.Config), "*"),
		natso.New,
		wire.Struct(new(cache_nats.Config), "*"),
		cache_nats.New,
		wire.Bind(new(cache.Store), new(*cache_nats.CacheNats)),
		wire.Struct(new(logdb.Config), "*"),
		logdb.New,
		// wire.Struct(new(logdb_nats_client.Config), "*"),
		// logdb_nats_client.New,
		wire.Struct(new(xmail_app.Config), "*"),
		xmail_app.New,
		wire.Struct(new(auth_app.Config), "*"),
		auth_app.New,
		wire.Struct(new(geo.Config), "*"),
		geo.New,
		wire.Struct(new(dapp.Config), "*"),
		dapp.New,
		wire.Struct(new(Plugin), "*"),
	)
	return Plugin{}
}
