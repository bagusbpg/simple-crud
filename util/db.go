package util

import (
	"database/sql"
	"fmt"
	"rest-api/design-pattern/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func GetDBInstance(config *config.AppConfig) *sql.DB {
	if db == nil {
		path := fmt.Sprintf("%v:%v@/%v", config.Username, config.Password, config.DBName)
		dbNewInstance, err := sql.Open(config.Driver, path)

		if err != nil {
			panic(err)
		}

		db = dbNewInstance
	}
	return db
}
