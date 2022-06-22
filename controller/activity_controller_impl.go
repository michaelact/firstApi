package controller

import (
	"strconv"
    "net/http"
    "github.com/julienschmidt/httprouter"

    "github.com/michaelact/firstApi/model/web"
    "github.com/michaelact/firstApi/service"
    "github.com/michaelact/firstApi/helper"
)

type ActivityControllerImpl struct {
	ActivityService service.ActivityService
}

func NewActivityController(activityService service.ActivityService) ActivityController {
	return &ActivityControllerImpl{
		ActivityService: activityService, 
	}
}

func (self *ActivityControllerImpl) Create(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	request := web.ActivityCreateRequest{}
	helper.ReadFromRequestBody(req, &request)

	serviceResponse := self.ActivityService.Create(req.Context(), request)
	webResponse := web.WebResponse{
		Status:  "Success", 
		Message: "Success", 
		Data:    serviceResponse, 
	}

	helper.WriteToResponseBody(res, &webResponse)
}

func (self *ActivityControllerImpl) Update(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	request := web.ActivityUpdateRequest{}
	helper.ReadFromRequestBody(req, &request)

	// Bind ID from Route Parameter
	id, err := strconv.Atoi(params.ByName("activityId"))
	helper.PanicIfError(err)
	request.Id = id

	serviceResponse := self.ActivityService.Update(req.Context(), request)
	webResponse := web.WebResponse{
		Status:  "Success", 
		Message: "Success", 
		Data:    serviceResponse, 
	}

	helper.WriteToResponseBody(res, &webResponse)
}

func (self *ActivityControllerImpl) Delete(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	requestId, err := strconv.Atoi(params.ByName("activityId"))
	helper.PanicIfError(err)

	self.ActivityService.Delete(req.Context(), requestId)
	webResponse := web.WebResponse{
		Status:  "Success", 
		Message: "Success", 
	}

	helper.WriteToResponseBody(res, &webResponse)
}

func (self *ActivityControllerImpl) FindById(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	requestId, err := strconv.Atoi(params.ByName("activityId"))
	helper.PanicIfError(err)

	serviceResponse := self.ActivityService.FindById(req.Context(), requestId)
	webResponse := web.WebResponse{
		Status:  "Success", 
		Message: "Success", 
		Data:    serviceResponse, 
	}

	helper.WriteToResponseBody(res, &webResponse)
}

func (self *ActivityControllerImpl) FindAll(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	serviceResponse := self.ActivityService.FindAll(req.Context())
	webResponse := web.WebResponse{
		Status:  "Success", 
		Message: "Success", 
		Data:    serviceResponse, 
	}

	helper.WriteToResponseBody(res, &webResponse)
}
