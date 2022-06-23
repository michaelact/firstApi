package middleware

import (
	"net/http"

	"github.com/michaelact/firstApi/helper"
	"github.com/michaelact/firstApi/model/environment"
)

type AuthMiddleware struct {
	Handler http.Handler
	Key     string
}

func NewAuthMiddleware(handler http.Handler, c *environment.Global) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler, 
		Key:     c.Api.Key, 
	}
}

func (self *AuthMiddleware) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if self.Key == req.Header.Get("x-api-key") {
		self.Handler.ServeHTTP(res, req)
	} else {
		helper.WriteToResponseBodyError(res, http.StatusUnauthorized, "UNAUTHORIZED")
	}
}
