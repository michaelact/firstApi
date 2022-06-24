package main

import (
	"github.com/go-playground/validator/v10"
	"net/http"

	"github.com/michaelact/firstApi/controller"
	"github.com/michaelact/firstApi/middleware"
	"github.com/michaelact/firstApi/repository"
	"github.com/michaelact/firstApi/util/log"
	"github.com/michaelact/firstApi/service"
	"github.com/michaelact/firstApi/config"
	"github.com/michaelact/firstApi/helper"
	"github.com/michaelact/firstApi/app"
	
)

func InitializeServer() *http.Server {
	conf := config.NewConfig()
	db := app.NewDB(conf)
	validate := validator.New()

	activityRepository := repository.NewActivityRepository()
	activityService := service.NewActivityService(activityRepository, db, validate)
	activityController := controller.NewActivityController(activityService)

	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(todoRepository, db, validate)
	todoController := controller.NewTodoController(todoService)

	router := app.NewRouter(activityController, todoController)
	authMiddleware := middleware.NewAuthMiddleware(router, conf)

	logger := log.NewLogger(conf)
	recorderMiddleware := middleware.NewRecorderMiddleware(authMiddleware, logger)
	server := app.NewServer(recorderMiddleware, conf)
	return server
}

func main() {
	server := InitializeServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
