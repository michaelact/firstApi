package app

import (
	"net/http"
	"fmt"
	"os"

	"github.com/michaelact/firstApi/middleware"
)

func NewServer(middleware *middleware.AuthMiddleware) *http.Server {
	address := fmt.Sprintf("%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT"))
	return &http.Server{
		Addr:    address, 
		Handler: middleware, 
	}
}
