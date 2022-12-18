package dapp

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/khanakia/mangobp/mango/cli"
	"github.com/khanakia/mangobp/mango/gormdb"
	"github.com/khanakia/mangobp/mango/natso"
	"github.com/spf13/cobra"
	"github.com/ubgo/gofm/cache"
)

type Dapp struct {
}

func (pkg Dapp) Version() string {
	return "0.01"
}

type Config struct {
	Cli    cli.Cli
	GormDB gormdb.GormDB
	Natso  natso.Natso
	Cache  cache.Store
}

func New(config Config) Dapp {
	// cache_natsapi.New(cache_natsapi.Config{
	// 	Natso: config.Natso,
	// 	Cache: config.Cache,
	// })

	AddCommands(config)
	return Dapp{}
}

func AddCommands(config Config) {
	var Main = &cobra.Command{
		Use:   "dapp",
		Short: "Dapp Pkg - Use `go run . ana --help` to see child commands",
		Run: func(cmd *cobra.Command, args []string) {
			color.Yellow("Run below command to see all the child commands.")
			color.Cyan("go run . amzn --help")
		},
	}

	rootCmd := config.Cli.RootCmd
	rootCmd.AddCommand(Main)

	// go run . core migrate_old_tables
	Main.AddCommand(&cobra.Command{
		Use:   "migrate_geo",
		Short: "Migrate geo table id to string",
		Run: func(cmd *cobra.Command, args []string) {
			// db := config.GormDB.DB
			// var countries []geo_domain.Country
			// db.Find(&countries)
			// for _, country := range countries {
			// 	fmt.Println(country.ID)
			// 	newId := "ctry_" + publicid.Must12()
			// 	db.Where("country_id = ?", country.ID).Model(&geo_domain.State{}).Updates(geo_domain.State{CountryID: newId})
			// 	country.NewID = newId
			// 	err := db.Save(&country).Error
			// 	fmt.Println(err)
			// }

			// var states []geo_domain.State
			// db.Find(&states)
			// for _, state := range states {
			// 	newId := "state_" + publicid.Must12()

			// 	state.NewID = newId
			// 	err := db.Save(&state).Error
			// 	fmt.Println(err)
			// }

			fmt.Println("All Done!!!")

		},
	})
}
