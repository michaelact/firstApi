package main

import (
	"github.com/michaelact/firstApi/helper"
)

func main() {
	server := InitializeServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
