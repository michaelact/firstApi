package exception

import (
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/michaelact/firstApi/model/web"
	"github.com/michaelact/firstApi/helper"
)

func ErrorHandler(res http.ResponseWriter, req *http.Request, err interface{}) {
	res.Header().Set("Content-Type", "application/json")
	if notFoundError(res, req, err) { return }
	if validationError(res, req, err) { return }

	// If error doesn't match the error above
	internalServerError(res, req, err)
	
}

func validationError(res http.ResponseWriter, req *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		helper.WriteToResponseBodyError(res, http.StatusBadRequest, exception)
		return true
	} else {
		return false
	}
}

func notFoundError(res http.ResponseWriter, req *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		helper.WriteToResponseBodyError(res, http.StatusNotFound, exception)
		return true
	} else {
		return false
	}
}

func internalServerError(res http.ResponseWriter, req *http.Request, err interface{}) {
	res.WriteHeader(http.StatusInternalServerError)
	webResponse := web.WebResponse{
		Status:  http.StatusText(http.StatusInternalServerError),
		Message: "Please contact administrator!\n" + "Email: michael.4ct@gmail.com",
	}

	helper.WriteToResponseBody(res, webResponse)
}
