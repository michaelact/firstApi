//go:build wireinject
// +build wireinject

// This file is ignored by `go build`

package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/google/wire"

	"github.com/michaelact/firstApi/middleware"
	"github.com/michaelact/firstApi/controller"
	"github.com/michaelact/firstApi/repository"
	"github.com/michaelact/firstApi/service"
	"github.com/michaelact/firstApi/app"
)

func InitializeServer() *http.Server {
	wire.Build(
		// Back
		app.NewDB, 
		validator.New, 

		// Activity
		repository.NewActivityRepository, 
		service.NewActivityService,
		controller.NewActivityController, 

		// Todo
		repository.NewTodoRepository, 
		service.NewTodoService,
		controller.NewTodoController, 

		// Front
		app.NewRouter, 
		middleware.NewAuthMiddleware, 
		app.NewServer, 
		wire.Bind(new(http.Handler), new(*httprouter.Router)), 
	)

	// 'nil' value will be changed by Google Wire
	return nil
}
