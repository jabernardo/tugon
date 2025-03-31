package core

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func CreateMiddlewareStack(middleware ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(middleware) - 1; i > -1; i-- {
			x := middleware[i]
			next = x(next)
		}

		return next
	}
}
