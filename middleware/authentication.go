package middleware

import (
	"os"
	"net/http"

	"github.com/michaelact/firstApi/helper"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler, 
	}
}

func (self *AuthMiddleware) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if os.Getenv("SERVER_API_KEY") == req.Header.Get("x-api-key") {
		self.Handler.ServeHTTP(res, req)
	} else {
		helper.WriteToResponseBodyError(res, http.StatusUnauthorized, "UNAUTHORIZED")
	}
}
