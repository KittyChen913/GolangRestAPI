package db

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/microsoft/go-mssqldb"
)

var Db *sql.DB
var retryCount int = 3

func InitDb() {
	var err error
	Db, err = sql.Open("sqlserver", os.Getenv("DemoDb"))
	if err != nil {
		panic("Unable to parse database config: " + err.Error())
	}
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(5)

	// 連線測試
	for i := 1; i <= retryCount; i++ {
		err = Db.Ping()
		if err == nil {
			break
		}
		if i == retryCount {
			panic("Unable to connect to the database: " + err.Error())
		}
		time.Sleep(2 * time.Second)
	}
}
