package libs

import (
	"database/sql"
	"time"

	"os"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	LoadENV()

	db, err := sql.Open(os.Getenv("DRIVER_NAME"), os.Getenv("URI"))
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
