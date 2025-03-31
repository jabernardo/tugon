package main

import (
	"github.com/jabernardo/aapi/app"
	"github.com/jabernardo/aapi/core"
)

func main() {
	api := core.New("1.0")

	api.Use(app.GetRouter())
	api.ListenAndServe(":3000")
}
