package middlewares

import (
	"net/http"
	"strings"
)

type CorsOptions struct {
	AllowedOrigins   map[string]bool
	AllowedMethods   []string
	AllowedHeaders   []string
	AllowCredentials bool
}

func NewCors(origins map[string]bool, methods []string, headers []string, allowCredentials bool) *CorsOptions {
	return &CorsOptions{
		AllowedOrigins:   origins,
		AllowedMethods:   methods,
		AllowedHeaders:   headers,
		AllowCredentials: allowCredentials,
	}
}

func (opts *CorsOptions) Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowedOrigins := "*"
		allowedMethods := "*"
		allowedHeaders := "*"

		origin := r.Header.Get("Origin")

		if len(opts.AllowedMethods) > 0 {
			allowedMethods = strings.Join(opts.AllowedMethods, ",")
		}

		if len(opts.AllowedHeaders) > 0 {
			allowedHeaders = strings.Join(opts.AllowedHeaders, ",")
		}

		if opts.AllowedOrigins[origin] || opts.AllowedOrigins["*"] {
			allowedOrigins = origin
			w.Header().Set("Access-Control-Allow-Origin", allowedOrigins)
			w.Header().Set("Access-Control-Allow-Methods", allowedMethods)
			w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)

			if opts.AllowCredentials {
				w.Header().Set("Access-Control-Allow-Credentials", "true")
			}

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
