package auth_app

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/khanakia/mangobp/mango/cache_nats_client"
	"github.com/khanakia/mangobp/mango/cli"
	"github.com/khanakia/mangobp/mango/gormdb"
	"github.com/khanakia/mangobp/mango/logdb/logdb_nats_client"
	"github.com/khanakia/mangobp/mango/natso"
	"github.com/khanakia/mangobp/pkg/auth/auth_domain"
	"github.com/khanakia/mangobp/pkg/auth/auth_nats"
	"github.com/khanakia/mangobp/pkg/auth/auth_repo"
	"github.com/spf13/cobra"
)

type Auth struct {
	Config
	UserRepo auth_repo.UserRepo
}

func (pkg Auth) Version() string {
	return "0.01"
}

func (pkg Auth) Say(name string) {
	fmt.Println("Hello " + name)
}

func (pkg Auth) MigrateDb() {
	// fmt.Println("ddd", pkg.GormDB.DB)
	pkg.GormDB.DB.AutoMigrate(&auth_domain.User{})
}

type Config struct {
	Cli             cli.Cli
	GormDB          gormdb.GormDB
	Natso           natso.Natso
	CacheNatsClient cache_nats_client.CacheNatsClient
	LogDbNatsClient logdb_nats_client.LogDbNatsClient
}

func New(config Config) Auth {
	pkg := Auth{
		Config: config,
		UserRepo: auth_repo.UserRepo{
			DB: config.GormDB.DB,
		},
	}

	AuthCommands(pkg)

	// fmt.Println(config.CacheNatsClient)

	auth_nats.New(auth_nats.Config{
		// UserRepo: pkg.UserRepo,
		Natso:           config.Natso,
		DB:              config.GormDB.DB,
		CacheNatsClient: config.CacheNatsClient,
	})
	return pkg
}

func AuthCommands(auth Auth) {
	rootcmd := auth.Config.Cli.RootCmd
	var Main = &cobra.Command{
		Use:   "auth",
		Short: "Auth plugin",
		Run: func(cmd *cobra.Command, args []string) {
			color.Yellow("Run below command to see all the child commands.")
			color.Cyan("go run . auth --help")
		},
	}

	rootcmd.AddCommand(Main)
	Main.AddCommand(&cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(auth.Version())
		},
	})
}
