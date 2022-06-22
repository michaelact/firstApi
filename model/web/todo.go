package web

import (
	"time"
)

type TodoCreateRequest struct {
	ActivityGroupId int    `validate:"required"`
	Title           string `validate:"required,max=30" json:"title"`
}

type TodoUpdateRequest struct {
	Id              int    `validate:"required"`
	ActivityGroupId int    `validate:"required" json:"activity_group_id"`
	Title           string `validate:"required,max=30" json:"title"`
}

type TodoResponse struct {
	Id              int       `json:"id"`
	ActivityGroupId int       `json:"activity_group_id"`
	Title           string    `json:"title"`
	IsActive        bool      `json:"is_active"`
	Priority        string    `json:"priority"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at"`
}
