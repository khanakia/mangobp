package apigql

import (
	"github.com/khanakia/mangobp/apigql/graph/resolverfn"
	"github.com/khanakia/mangobp/apigql/pkg/gql"
	"github.com/khanakia/mangobp/apigql/pkg/middleware"
	"github.com/khanakia/mangobp/apigql/pkg/server_handler"
	"github.com/ubgo/gofm/ginserver"
	"github.com/ubgo/gqlgenfn"

	"github.com/khanakia/mangobp/wireapp"
)

func Boot(plugin wireapp.Plugin) {
	serverServer := ginserver.New(ginserver.Config{
		BeforeHandler: middleware.CORSMiddleware(),
	})

	serverServer.Router.Use(middleware.BlockHostsMiddleware())
	serverServer.Router.Use(gqlgenfn.GinContextToContextMiddleware())

	resolver := &resolverfn.Resolver{
		GormDB: plugin.GormDB,
		Logger: plugin.Logger,
	}
	gqlConfig := gql.Config{
		GormDB:   plugin.GormDB,
		Server:   serverServer,
		Resolver: resolver,
	}
	gql.New(gqlConfig)

	server_handler.New(server_handler.Config{
		GormDB: plugin.GormDB,
		Logger: plugin.Logger,
		Server: serverServer,
	})

	serverServer.Start()
}
