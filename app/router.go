package app

import (
	"github.com/jabernardo/tugon/app/handlers/v1"
	"github.com/jabernardo/tugon/app/handlers/v1/todo"
	"github.com/jabernardo/tugon/app/middlewares"
	"github.com/jabernardo/tugon/core"
)

func GetTodoRouter() *core.Router {
	// Todo Example
	todoGroup := core.NewRouter()

	todoGroup.All("/", todo.Create)
	todoGroup.Get("/{id}", todo.Get)
	todoGroup.Delete("/{id}", todo.Delete)
	todoGroup.Patch("/{id}", todo.Update)
	todoGroup.Get("/all", todo.GetAll)

	return todoGroup
}

func GetRouter() *core.Router {
	router := core.NewRouter()

	cors := middlewares.NewCors(map[string]bool{"http://localhost:44720": true}, []string{}, []string{"Content-Type", "X-Custom-Header"}, true)
	router.Use(middlewares.Logger, cors.Cors)

	route1 := core.NewRouter()
	route1.Get("/hello", v1.Hello)
	route1.Get("/ping", v1.Ping)
	route1.Group("/todo", GetTodoRouter())

	router.Group("/v1", route1)

	return router
}
