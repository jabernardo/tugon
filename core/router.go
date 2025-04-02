package core

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Routes map[string]http.Handler
type Router struct {
	routes      Routes
	middlewares []Middleware
	path        string
}

func NewRouter() *Router {
	return &Router{
		routes:      make(Routes),
		middlewares: make([]Middleware, 0),
	}
}

func createRoute(router *Router, pattern string, handler http.HandlerFunc, middleware ...Middleware) http.Handler {
	route := http.NewServeMux()
	route.HandleFunc(pattern, handler)

	groupMiddlewares := CreateMiddlewareStack(router.middlewares...)
	routeSpecificMiddlewares := CreateMiddlewareStack(middleware...)
	return groupMiddlewares(routeSpecificMiddlewares(route))
}

func (router *Router) SetGroup(path string) {
	if len(router.routes) > 0 {
		log.Fatalln("[core.router] could not set group when routes already exists")
	}
	router.path = "/" + strings.Trim(path, "/")
}

func (router *Router) GetRoutes() Routes {
	return router.routes
}

func (router *Router) Use(middleware ...Middleware) {
	router.middlewares = append(router.middlewares, middleware...)
}

func (router *Router) Add(method string, pattern string, handler http.HandlerFunc, middleware ...Middleware) {
	if len(method) > 0 {
		method = method + " "
	}

	grouped := fmt.Sprintf("%s %s/%s", method, router.path, strings.Trim(pattern, "/"))

	if _, ok := router.routes[grouped]; ok {
		log.Fatalf("[core.router] duplicated route: `%s`\n", pattern)
	}

	router.routes[grouped] = createRoute(router, grouped, handler, middleware...)
}

func (router *Router) Get(pattern string, handler http.HandlerFunc, middleware ...Middleware) {
	router.Add("GET", pattern, handler, middleware...)
}

func (router *Router) Post(pattern string, handler http.HandlerFunc, middleware ...Middleware) {
	router.Add("POST", pattern, handler, middleware...)
}

func (router *Router) PUT(pattern string, handler http.HandlerFunc, middleware ...Middleware) {
	router.Add("PUT", pattern, handler, middleware...)
}

func (router *Router) Patch(pattern string, handler http.HandlerFunc, middleware ...Middleware) {
	router.Add("PATCH", pattern, handler, middleware...)
}

func (router *Router) Delete(pattern string, handler http.HandlerFunc, middleware ...Middleware) {
	router.Add("DELETE", pattern, handler, middleware...)
}

func (router *Router) Options(pattern string, handler http.HandlerFunc, middleware ...Middleware) {
	router.Add("OPTIONS", pattern, handler, middleware...)
}

func (router *Router) Head(pattern string, handler http.HandlerFunc, middleware ...Middleware) {
	router.Add("HEAD", pattern, handler, middleware...)
}

func (router *Router) All(pattern string, handler http.HandlerFunc, middleware ...Middleware) {
	router.Add("", pattern, handler, middleware...)
}
