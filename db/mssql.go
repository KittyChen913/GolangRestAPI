package db

import (
	"database/sql"

	_ "github.com/microsoft/go-mssqldb"
)

var Db *sql.DB

func InitDb() {
	var err error
	Db, err = sql.Open("sqlserver", "sqlserver://admin:admin@localhost:1433?database=DemoDb&encrypt=disable")
	if err != nil {
		panic("Unable to connect to the database.")
	}
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(5)
}
