package wireapp

import (
	"github.com/khanakia/mangobp/mango/cache_nats"
	"github.com/khanakia/mangobp/mango/cacherdbms"
	"github.com/khanakia/mangobp/mango/cli"
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
)

type Plugin struct {
	ConfigMgr    interfaces.IConfig
	Cli          cli.Cli
	Logger       logger.Logger
	DbConn       *dbconn.DbConn
	GormDB       gormdb.GormDB
	Natso        natso.Natso
	CacheRdbms   *cacherdbms.Rdbms
	Cache        cache.Store
	LogDb        logdb.LogDb
	Xmail        xmail_app.Xmail
	Auth         auth_app.Auth
	Geo          geo.Geo
	Dapp         dapp.Dapp
	CacheNats    *cache_nats.CacheNats
	CacheNatsApi *cache_natsapi.CacheNatsApi
}
