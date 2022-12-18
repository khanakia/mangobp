package dbconn

import (
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/khanakia/mangobp/mango/interfaces"
)

type DbConn struct {
	SqlDb *sql.DB
}

type Config struct {
	ConfigMgr interfaces.IConfig
}

func New(config1 Config) *DbConn {
	username := config1.ConfigMgr.GetString("database.user")
	password := config1.ConfigMgr.GetString("database.password")
	databaseName := config1.ConfigMgr.GetString("database.name")
	databaseHost := config1.ConfigMgr.GetString("database.host")
	databasePort := config1.ConfigMgr.GetString("database.port")
	sslmode := config1.ConfigMgr.GetString("database.sslmode")

	dbDSN := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s port=%s", databaseHost, username, databaseName, sslmode, password, databasePort)

	config, err := pgx.ParseConfig(dbDSN)
	if err != nil {
		panic(err)
	}

	db := stdlib.OpenDB(*config)

	// defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	return &DbConn{
		SqlDb: db,
	}
}
