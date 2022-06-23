package app

import (
	"database/sql"
	"time"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/michaelact/firstApi/model/environment"
)

func NewDB(c *environment.Global) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
						c.Database.User,
						c.Database.Password,
						c.Database.Host,
						c.Database.Port,
						c.Database.Name)

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
