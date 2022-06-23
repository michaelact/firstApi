package app

import (
	"net/http"
	"fmt"

	"github.com/michaelact/firstApi/model/environment"
	"github.com/michaelact/firstApi/middleware"
)

func NewServer(middleware *middleware.AuthMiddleware, c *environment.Global) *http.Server {
	address := fmt.Sprintf("%s:%s", c.Server.Host, c.Server.Port)
	return &http.Server{
		Addr:    address, 
		Handler: middleware, 
	}
}
