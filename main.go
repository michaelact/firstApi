package main

import (
	"os"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"

	"github.com/michaelact/firstApi/repository"
	"github.com/michaelact/firstApi/controller"
	"github.com/michaelact/firstApi/middleware"
	"github.com/michaelact/firstApi/exception"
	"github.com/michaelact/firstApi/service"
	"github.com/michaelact/firstApi/helper"
	"github.com/michaelact/firstApi/app"
)

func main() {
	db := app.ConnectDB()

	validate := validator.New()
	router := httprouter.New()
	router.PanicHandler = exception.ErrorHandler

	// Activity
	activityRepository := repository.NewActivityRepository()
	activityService := service.NewActivityService(activityRepository, db, validate)
	activityController := controller.NewActivityController(activityService)
	
	router.GET("/activity-groups", activityController.FindAll)
	router.POST("/activity-groups", activityController.Create)
	router.GET("/activity-groups/:activityId", activityController.FindById)
	router.PATCH("/activity-groups/:activityId", activityController.Update)
	router.DELETE("/activity-groups/:activityId", activityController.Delete)

	// Todo
	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(todoRepository, db, validate)
	todoController := controller.NewTodoController(todoService)

	router.GET("/todo-items", todoController.FindAll)
	router.POST("/todo-items", todoController.Create)
	router.GET("/todo-items/:todoId", todoController.FindById)
	router.PATCH("/todo-items/:todoId", todoController.Update)
	router.DELETE("/todo-items/:todoId", todoController.Delete)

	address := fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
	server := http.Server{
		Addr: address, 
		Handler: middleware.NewAuthMiddleware(router), 
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
