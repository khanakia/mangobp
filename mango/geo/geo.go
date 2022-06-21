package geo

import (
	"github.com/khanakia/mangobp/mango/cli"
	"github.com/khanakia/mangobp/mango/geo/geo_domain"
	"github.com/khanakia/mangobp/mango/gormdb"
)

type Config struct {
	Cli    cli.Cli
	GormDB gormdb.GormDB
}

type Geo struct {
	Config
}

func (pkg Geo) Version() string {
	return "0.01"
}

func (pkg Geo) MigrateDb() {
	pkg.GormDB.DB.AutoMigrate(&geo_domain.Country{}, &geo_domain.State{})
}

func New(config Config) Geo {
	pkg := Geo{
		Config: config,
	}
	return pkg
}
