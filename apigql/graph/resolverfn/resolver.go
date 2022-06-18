package resolverfn

import (
	"github.com/khanakia/mangobp/mango/gormdb"

	"github.com/ubgo/gofm/logger"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	GormDB gormdb.GormDB
	Logger logger.Logger
	// Cache  cache.Store
	// GinStore *gcache.RedisStore
	// Auth   auth.Auth
}
