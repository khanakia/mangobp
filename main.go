package main

import (
	"github.com/khanakia/mangobp/apigql"
	"github.com/khanakia/mangobp/mango/configmgr"
	"github.com/khanakia/mangobp/mango/dbconn"
	"github.com/khanakia/mangobp/mango/gormdb"
	"github.com/khanakia/mangobp/mango/interfaces"
	"github.com/khanakia/mangobp/mango/plug"
	"github.com/khanakia/mangobp/wireapp"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

func NewConfigMgrConfig() configmgr.Config {
	return configmgr.Config{}
}

func NewInterfaceConfig(c *configmgr.ConfigMgr) interfaces.IConfig {
	return c
}

func NewGormConfig(dbConn *dbconn.DbConn) gormdb.Config {
	return gormdb.Config{
		DB: dbConn.SqlDb,
	}
}

func NewGormDb(gormdb gormdb.GormDB) *gorm.DB {
	return gormdb.DB
}

func main() {
	// USING WIRE
	plugin := wireapp.Init()
	plug.InitPlugins(plugin.Cli.RootCmd, plugin)

	var ServerCommand = &cobra.Command{
		Use:   "server",
		Short: "starts the HTTP server",
		Run: func(cmd *cobra.Command, args []string) {
			apigql.Boot(plugin)

			//// Only wait if we do not start the http server
			// plugin.Natso.Wait()

		},
	}

	plugin.Cli.RootCmd.AddCommand(ServerCommand)

	plugin.Cli.Execute()

}
