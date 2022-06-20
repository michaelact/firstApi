package web

import (
	"time"
)

type ActivityCreateRequest struct {
	Email string `validate:"required,max=30"`
	Title string `validate:"required,max=30"`
}

type ActivityUpdateRequest struct {
	Id    int    `validate:"required"`
	Email string `validate:"required,max=30"`
	Title string `validate:"required,max=30"`
}

type ActivityResponse struct {
	Id        int
	Email     string
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
