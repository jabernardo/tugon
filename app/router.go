package app

import (
	"github.com/jabernardo/tugon/app/handlers/v1"
	"github.com/jabernardo/tugon/app/middlewares"
	"github.com/jabernardo/tugon/core"
)

func GetRouter() *core.Router {
	router := core.NewRouter()
	router.SetGroup("/v1")

	cors := middlewares.NewCors(map[string]bool{"http://localhost:44720": true}, []string{}, []string{"Content-Type", "X-Custom-Header"}, true)
	router.Use(middlewares.Logger, cors.Cors)

	router.Get("/hello", v1.Hello)
	router.Get("/ping", v1.Ping)

	return router
}
