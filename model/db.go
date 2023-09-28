package model

import (
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func DB() (*sql.DB, error) {

	cfg := mysql.Config{
		User:                 os.Getenv("DB_USERNAME"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
		ParseTime:            true,
		MultiStatements:      true,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())

	return db, err
}
