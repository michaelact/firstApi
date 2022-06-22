package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"

	"github.com/michaelact/firstApi/repository"
	"github.com/michaelact/firstApi/controller"
	"github.com/michaelact/firstApi/service"
	"github.com/michaelact/firstApi/app"
)

func main() {
	db := app.ConnectDB()

	validate := validator.New()

	// Activity
	activityRepository := repository.NewActivityRepository()
	activityService := service.NewActivityService(activityRepository, db, validate)
	activityController := controller.NewActivityController(activityService)

	// Todo
	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(todoRepository, db, validate)
	todoController := controller.NewTodoController(todoService)

	router := httprouter.New()

	router.GET("/activity_groups", activityController.FindAll)
	router.POST("/activity_groups", activityController.Create)
	router.GET("/activity_groups/:activityId", activityController.FindById)
	router.PATCH("/activity_groups/:activityId", activityController.Update)
	router.DELETE("/activity_groups/:activityId", activityController.Delete)

	router.GET("/todo_items", todoController.FindAll)
	router.POST("/todo_items", todoController.Create)
	router.GET("/todo_items/:todoId", todoController.FindById)
	router.PATCH("/todo_items/:todoId", todoController.Update)
	router.DELETE("/todo_items/:todoId", todoController.Delete)
}
