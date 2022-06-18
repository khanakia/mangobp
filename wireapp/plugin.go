package wireapp

import (
	"github.com/khanakia/mangobp/mango/cli"
	"github.com/khanakia/mangobp/mango/dbconn"
	"github.com/khanakia/mangobp/mango/gormdb"
	"github.com/khanakia/mangobp/mango/interfaces"
	"github.com/khanakia/mangobp/mango/natso"
	"github.com/ubgo/gofm/logger"
)

type Plugin struct {
	ConfigMgr interfaces.IConfig
	Cli       cli.Cli
	Logger    logger.Logger
	DbConn    dbconn.DbConn
	GormDB    gormdb.GormDB
	Natso     natso.Natso
	// Auth      auth_app.Auth
}
