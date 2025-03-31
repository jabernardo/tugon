package app

import (
	"github.com/jabernardo/aapi/app/handlers/v1"
	"github.com/jabernardo/aapi/app/middlewares"
	"github.com/jabernardo/aapi/core"
)

func GetRouter() *core.Router {
	router := core.NewRouter()
	router.SetGroup("/v1")

	router.Use(middleware.Logger)

	router.All("/hello", v1.Hello)

	return router
}
