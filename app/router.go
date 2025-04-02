package app

import (
	"github.com/jabernardo/aapi/app/handlers/v1"
	"github.com/jabernardo/aapi/app/middlewares"
	"github.com/jabernardo/aapi/core"
)

func GetRouter() *core.Router {
	router := core.NewRouter()
	router.SetGroup("/v1")

	cors := middlewares.NewCors(map[string]bool{"http://localhost:44720": true}, []string{}, []string{"Content-Type", "X-Custom-Header"}, true)
	router.Use(middlewares.Logger, cors.Cors)

	router.All("/hello", v1.Hello)
	router.All("/ping", v1.Ping)

	return router
}
