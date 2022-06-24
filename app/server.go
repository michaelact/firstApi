package app

import (
	"net/http"
	"fmt"

	"github.com/michaelact/firstApi/model/config"
)

func NewServer(middleware http.Handler, c *config.Global) *http.Server {
	address := fmt.Sprintf("%s:%s", c.Server.Host, c.Server.Port)
	return &http.Server{
		Addr:    address, 
		Handler: middleware, 
	}
}
