package db

import (
	"database/sql"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/microsoft/go-mssqldb"
)

var Db *sql.DB

func InitDb() {
	var err error
	Db, err = sql.Open("sqlserver", os.Getenv("DemoDb"))
	if err != nil {
		panic("Unable to connect to the database.")
	}
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(5)
}
