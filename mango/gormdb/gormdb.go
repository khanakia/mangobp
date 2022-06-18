package gormdb

import (
	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormDB struct {
	DB *gorm.DB
}

type Config struct {
	GormConfig *gorm.Config
	DB         *sql.DB
}

func New(config Config) GormDB {
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		// Logger: logger,
	}

	if config.GormConfig != nil {
		gormConfig = config.GormConfig
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: config.DB,
	}), gormConfig)

	if err != nil {
		// fmt.Println("error", err)
		panic(err)
	}
	return GormDB{
		DB: db,
	}
}
