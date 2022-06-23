package app

import (
	"github.com/julienschmidt/httprouter"

	"github.com/michaelact/firstApi/controller"
	"github.com/michaelact/firstApi/exception"
)

func NewRouter(activityController controller.ActivityController, todoController controller.TodoController) *httprouter.Router {
	router := httprouter.New()
	router.PanicHandler = exception.ErrorHandler

	router.GET("/activity-groups", activityController.FindAll)
	router.POST("/activity-groups", activityController.Create)
	router.GET("/activity-groups/:activityId", activityController.FindById)
	router.PATCH("/activity-groups/:activityId", activityController.Update)
	router.DELETE("/activity-groups/:activityId", activityController.Delete)

	router.GET("/todo-items", todoController.FindAll)
	router.POST("/todo-items", todoController.Create)
	router.GET("/todo-items/:todoId", todoController.FindById)
	router.PATCH("/todo-items/:todoId", todoController.Update)
	router.DELETE("/todo-items/:todoId", todoController.Delete)

	return router
}
