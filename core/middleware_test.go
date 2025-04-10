package core

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddleware(t *testing.T) {
	hello := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello")
			next.ServeHTTP(w, r)
		})
	}

	world := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "World")
			next.ServeHTTP(w, r)
		})
	}

	route := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "!")
	})

	t.Run("Should stack function", func(t *testing.T) {
		middlewareStack := CreateMiddlewareStack(hello, world)
		handler := middlewareStack(route)

		req := httptest.NewRequest("GET", "http://localhost:3000", nil)
		resp := httptest.NewRecorder()

		handler.ServeHTTP(resp, req)

		output, err := io.ReadAll(resp.Result().Body)

		if err != nil {
			t.Errorf("[core.middleware] Could not read result of middleware stack")
		}

		if string(output) != "HelloWorld!" {
			t.Errorf("[core.middleware] Response expected to be `HelloWorld!`, got `%s`", output)
		}
	})
}
