package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCors(t *testing.T) {
	allowedOrigin := map[string]bool{
		"http://localhost:3000": true,
	}
	allowedMethods := []string{"GET", "POST"}
	allowedHeaders := []string{"Content-Type", "Origin"}

	corsMiddleware := NewCors(allowedOrigin, allowedMethods, allowedHeaders, true)

	handler := corsMiddleware.Cors(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	corsMiddlewareOriginOnly := NewCors(allowedOrigin, []string{}, []string{}, false)

	handlerOriginOnly := corsMiddlewareOriginOnly.Cors(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	t.Run("Allow origin", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://localhost:3000", nil)
		req.Header.Set("Origin", "http://localhost:3000")
		resp := httptest.NewRecorder()

		handler.ServeHTTP(resp, req)

		if got := resp.Header().Get("Access-Control-Allow-Origin"); got != "http://localhost:3000" {
			t.Errorf("[middlewares.cors] expected `Access-Control-Allow-Origin` to be 'http://localhost:3000', got '%s'", got)
		}

		if got := resp.Header().Get("Access-Control-Allow-Methods"); got != "GET,POST" {
			t.Errorf("[middlewares.cors] expected `Access-Control-Allow-Method` to be 'GET,POST', got '%s'", got)
		}

		if got := resp.Header().Get("Access-Control-Allow-Headers"); got != "Content-Type,Origin" {
			t.Errorf("[middlewares.cors] expected `Access-Control-Allow-Headers` to be 'Content-Type,Origin', got '%s'", got)
		}

		if got := resp.Header().Get("Access-Control-Allow-Credentials"); got != "true" {
			t.Errorf("[middlewares.cors] expected `Access-Control-Allow-Credentials` to be 'true', got '%s'", got)
		}
	})

	t.Run("Allow origin without defined methods and headers", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://localhost:3000", nil)
		req.Header.Set("Origin", "http://localhost:3000")
		resp := httptest.NewRecorder()

		handlerOriginOnly.ServeHTTP(resp, req)

		if got := resp.Header().Get("Access-Control-Allow-Origin"); got != "http://localhost:3000" {
			t.Errorf("[middlewares.cors] expected `Access-Control-Allow-Origin` to be 'http://localhost:3000', got '%s'", got)
		}

		if got := resp.Header().Get("Access-Control-Allow-Methods"); got != "*" {
			t.Errorf("[middlewares.cors] expected `Access-Control-Allow-Method` to be 'GET,POST', got '%s'", got)
		}

		if got := resp.Header().Get("Access-Control-Allow-Headers"); got != "*" {
			t.Errorf("[middlewares.cors] expected `Access-Control-Allow-Headers` to be 'Content-Type,Origin', got '%s'", got)
		}

		if got := resp.Header().Get("Access-Control-Allow-Credentials"); got != "" {
			t.Errorf("[middlewares.cors] expected `Access-Control-Allow-Credentials` to be 'true', got '%s'", got)
		}
	})

	t.Run("Disallowed origin", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://localhost:3001", nil)
		req.Header.Set("Origin", "http://localhost:3001")
		resp := httptest.NewRecorder()

		handler.ServeHTTP(resp, req)

		if got := resp.Header().Get("Access-Control-Allow-Origin"); got != "" {
			t.Errorf("[middlewares.cors] expected `Access-Control-Allow-Origin` to be empty, got '%s'", got)
		}

		if got := resp.Header().Get("Access-Control-Allow-Methods"); got != "" {
			t.Errorf("[middlewares.cors] expected `Access-Control-Allow-Methods` to be empty, got '%s'", got)
		}

		if got := resp.Header().Get("Access-Control-Allow-Headers"); got != "" {
			t.Errorf("[middlewares.cors] expected `Access-Control-Allow-Headers` to be empty, got '%s'", got)
		}
	})
}
