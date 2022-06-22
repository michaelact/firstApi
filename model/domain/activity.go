package domain

import (
	"database/sql"
	"time"
)

type Activity struct {
	Id        int
	Email     string
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
