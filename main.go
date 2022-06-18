package main

import (
	"github.com/khanakia/mangobp/apigql"
	"github.com/khanakia/mangobp/mango/plug"
	"github.com/khanakia/mangobp/wireapp"
	"github.com/spf13/cobra"
)

func main() {
	// USING WIRE
	plugin := wireapp.Init()
	plug.InitPlugins(plugin.Cli.RootCmd, plugin)

	var ServerCommand = &cobra.Command{
		Use:   "server",
		Short: "starts the HTTP server",
		Run: func(cmd *cobra.Command, args []string) {
			apigql.Boot(plugin)
		},
	}

	plugin.Cli.RootCmd.AddCommand(ServerCommand)

	plugin.Cli.Execute()
}
