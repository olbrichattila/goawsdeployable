package db

import (
	"database/sql"
	"fmt"
	"sharedconfig"

	_ "github.com/go-sql-driver/mysql"
)

func Db(config sharedconfig.SharedConfiger) (*sql.DB, error) {
	dbConf := config.GetDBConfig()
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		dbConf.Username,
		dbConf.Password,
		dbConf.Host,
		dbConf.Port,
		dbConf.Database,
	)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}
