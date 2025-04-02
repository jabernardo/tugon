package middlewares

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	handler := Logger(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	t.Run("Log is working", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://localhost:3000/hello", nil)
		req.Header.Set("Referer", "http://localhost:5000")
		req.Header.Set("User-Agent", "Test")
		resp := httptest.NewRecorder()

		originalStdout := os.Stdout

		defer func() { os.Stdout = originalStdout }()

		r, w, _ := os.Pipe()
		os.Stdout = w

		handler.ServeHTTP(resp, req)

		w.Close()

		var buf bytes.Buffer
		_, _ = buf.ReadFrom(r)

		os.Stdout = originalStdout

		var log Log
		err := json.Unmarshal(buf.Bytes(), &log)

		if err != nil {
			t.Errorf("[middlewares.log] Expected log to be in JSON, got '%s'", buf.String())
		}

		if log.Timestamp == "" {
			t.Errorf("[middlewares.log] Expected `Timestamp` to exists, got nothing")
		}

		if log.ClientIP == "" {
			t.Errorf("[middlewares.log] Expected `ClientIP` to exists, got nothing")
		}

		if log.Referer != "http://localhost:5000" {
			t.Errorf("[middlewares.log] Expected `Referer` to be 'http://localhost:5000', got '%s'", log.Referer)
		}

		if log.UserAgent != "Test" {
			t.Errorf("[middlewares.log] Expected `UserAgent` to be 'Test', got '%s'", log.UserAgent)
		}

		if log.Method != "GET" {
			t.Errorf("[middlewares.log] Expected `Method` to be 'GET', got '%s'", log.Method)
		}

		if log.Url != "http://localhost:3000/hello" {
			t.Errorf("[middlewares.log] Expected `Url` to be 'http://localhost:3000/hello', got '%s'", log.Url)
		}

		if log.StatusCode != 200 {
			t.Errorf("[middlewares.log] Expected `StatusCode` to be '200', got '%s'", log.Url)
		}
	})
}
