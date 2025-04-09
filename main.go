package main

import (
	"os"

	"github.com/jabernardo/tugon/app"
	"github.com/jabernardo/tugon/core"
)

// @title         Tugon
// @version       1.0
// @description   This is a simple REST API for the Boiler Plate API

func main() {
	bootstrap()

	addr := ":5000"

	if addrEnv, ok := os.LookupEnv("ADDR"); ok {
		addr = addrEnv
	}

	router := app.GetRouter()
	api := core.New(router, addr)
	api.ListenAndServe()
}
