package web

import (
	"time"
)

type TodoCreateRequest struct {
	ActivityGroupId int    `validate:"required"`
	Title           string `validate:"required,max=30"`
}

type TodoUpdateRequest struct {
	Id              int    `validate:"required"`
	ActivityGroupId int    `validate:"required"`
	Title           string `validate:"required,max=30"`
}

type TodoResponse struct {
	Id              int
	ActivityGroupId int
	Title           string
	IsActive        bool
	Priority        string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
}
