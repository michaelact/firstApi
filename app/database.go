package app

import (
	"database/sql"
	"time"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
						os.Getenv("MYSQL_USER"),
						os.Getenv("MYSQL_PASSWORD"),
						os.Getenv("MYSQL_HOST"),
						os.Getenv("MYSQL_PORT"),
						os.Getenv("MYSQL_DATABASE"))
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
