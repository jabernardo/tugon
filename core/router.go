package core

import (
	"fmt"
	"net/http"
	"strings"
)

type Router struct {
	Mux         http.ServeMux
	Middlewares []Middleware
}

func NewRouter() *Router {
	return &Router{
		Mux:         *http.NewServeMux(),
		Middlewares: make([]Middleware, 0),
	}
}

func createRoute(router *Router, pattern string, handler http.HandlerFunc, middleware ...Middleware) http.Handler {
	route := http.NewServeMux()
	route.HandleFunc(pattern, handler)

	groupMiddlewares := CreateMiddlewareStack(router.Middlewares...)
	routeSpecificMiddlewares := CreateMiddlewareStack(middleware...)
	return groupMiddlewares(routeSpecificMiddlewares(route))
}

func (router *Router) Use(middleware ...Middleware) {
	router.Middlewares = append(router.Middlewares, middleware...)
}

func (router *Router) Group(path string, sub *Router) {
	cleanedPath := strings.TrimRight(path, "/")
	groupMiddlewares := CreateMiddlewareStack(router.Middlewares...)
	groupHandler := groupMiddlewares(http.StripPrefix(cleanedPath, &sub.Mux))
	router.Mux.Handle(cleanedPath+"/", groupHandler)
}

func (router *Router) Add(method string, pattern string, handler http.HandlerFunc, middleware ...Middleware) {
	if len(method) > 0 {
		method = method + " "
	}

	grouped := fmt.Sprintf("%s %s", method, pattern)

	customHandler := createRoute(router, grouped, handler, middleware...)
	router.Mux.Handle(grouped, customHandler)
}

func (router *Router) Get(pattern string, handler http.HandlerFunc, middleware ...Middleware) {
	router.Add("GET", pattern, handler, middleware...)
}

func (router *Router) Post(pattern string, handler http.HandlerFunc, middleware ...Middleware) {
	router.Add("POST", pattern, handler, middleware...)
}

func (router *Router) Put(pattern string, handler http.HandlerFunc, middleware ...Middleware) {
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
