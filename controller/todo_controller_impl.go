package controller

import (
	"strconv"
    "net/http"
    "github.com/julienschmidt/httprouter"

    "github.com/michaelact/firstApi/model/web"
    "github.com/michaelact/firstApi/service"
    "github.com/michaelact/firstApi/helper"
)

type TodoControllerImpl struct {
	TodoService service.TodoService
}

func NewTodoController(todoService service.TodoService) TodoController {
	return &TodoControllerImpl{
		TodoService: todoService, 
	}
}

func (self *TodoControllerImpl) Create(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	request := web.TodoCreateRequest{}
	helper.ReadFromRequestBody(req, &request)

	serviceResponse := self.TodoService.Create(req.Context(), request)
	webResponse := web.WebResponse{
		Status:  "Success", 
		Message: "Success", 
		Data:    serviceResponse, 
	}

	helper.WriteToResponseBody(res, &webResponse)
}

func (self *TodoControllerImpl) Update(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	request := web.TodoUpdateRequest{}
	helper.ReadFromRequestBody(req, &request)

	// Bind ID from Route Parameter
	id, err := strconv.Atoi(params.ByName("todoId"))
	helper.PanicIfError(err)
	request.Id = id

	serviceResponse := self.TodoService.Update(req.Context(), request)
	webResponse := web.WebResponse{
		Status:  "Success", 
		Message: "Success", 
		Data:    serviceResponse, 
	}

	helper.WriteToResponseBody(res, &webResponse)
}

func (self *TodoControllerImpl) Delete(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	requestId, err := strconv.Atoi(params.ByName("todoId"))
	helper.PanicIfError(err)

	self.TodoService.Delete(req.Context(), requestId)
	webResponse := web.WebResponse{
		Status:  "Success", 
		Message: "Success", 
	}

	helper.WriteToResponseBody(res, &webResponse)
}

func (self *TodoControllerImpl) FindById(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	requestId, err := strconv.Atoi(params.ByName("todoId"))
	helper.PanicIfError(err)

	serviceResponse := self.TodoService.FindById(req.Context(), requestId)
	webResponse := web.WebResponse{
		Status:  "Success", 
		Message: "Success", 
		Data:    serviceResponse, 
	}

	helper.WriteToResponseBody(res, &webResponse)
}

func (self *TodoControllerImpl) FindAll(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	serviceResponse := self.TodoService.FindAll(req.Context())
	webResponse := web.WebResponse{
		Status:  "Success", 
		Message: "Success", 
		Data:    serviceResponse, 
	}

	helper.WriteToResponseBody(res, &webResponse)
}
