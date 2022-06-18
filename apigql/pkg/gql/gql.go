package gql

import (
	"context"
	"errors"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/khanakia/mangobp/apigql/graph/generated"
	"github.com/khanakia/mangobp/apigql/graph/resolverfn"
	"github.com/khanakia/mangobp/mango/gormdb"
	"github.com/khanakia/mangobp/mango/recaptcha"
	"github.com/khanakia/mangobp/mango/util"
	"github.com/spf13/viper"
	"github.com/ubgo/gofm/ginserver"

	"github.com/ubgo/gofm/gqlgin"
	"github.com/ubgo/gqlgenfn"
)

type Gql struct {
	Config
}

type Config struct {
	Server   ginserver.Server
	Resolver *resolverfn.Resolver
	GormDB   gormdb.GormDB
}

func New(config Config) Gql {
	c := generated.Config{Resolvers: config.Resolver}
	c.Directives.HasCaptcha = func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {

		// Bypass captcha for the playground
		isAllowedPlayground := gqlgin.IsPlaygroundAllwedForContext(ctx)
		if isAllowedPlayground {
			return next(ctx)
		}

		gc, err := gqlgenfn.GinContextFromContext(ctx)
		if err != nil {
			return nil, errors.New("server Error")
		}

		// Bypass captcha for next.js app
		captchaBypass := gc.Request.Header.Get("x-captcha-bypass")
		if len(captchaBypass) > 0 {
			captchaBypassSecret := viper.GetString("secret.captcha_bypass")
			if captchaBypass == captchaBypassSecret {
				return next(ctx)
			}
		}

		captchaRes := gc.Request.Header.Get("x-captcha-res")

		if len(captchaRes) <= 0 {
			return nil, errors.New("access Denied (1001)")
		}

		ip := util.GetIP(gc.Request)
		res, _ := recaptcha.Confirm(ip, captchaRes)

		fmt.Printf("%+v", res)

		if !res.Success {
			return nil, errors.New("access Denied (1002)")
		}

		// or let it pass through
		return next(ctx)
	}
	gserver := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	gqlgin.New(gqlgin.Config{
		Server:    config.Server,
		GqlServer: gserver,
		// Middleware: []gin.HandlerFunc{middleware.JwtMiddlewarePublic(config.GormDB.DB)},
	})
	gql := Gql{
		Config: config,
	}

	return gql
}
