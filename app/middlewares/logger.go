package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Log struct {
	Timestamp    string  `json:"timestamp"`
	ClientIP     string  `json:"client_ip"`
	Method       string  `json:"method"`
	Url          string  `json:"url"`
	StatusCode   int     `json:"status_code"`
	ResponseTime float32 `json:"response_time"`
	UserAgent    string  `json:"user_agent"`
	Referrer     string  `json:"referrer"`
}

func getClientIP(r *http.Request) string {
	ip := r.Header.Get("x-real-ip")

	if len(ip) == 0 {
		ip = r.Header.Get("x-forwarded-for")
	}

	if len(ip) == 0 {
		ip = r.RemoteAddr
	}

	return ip
}

func NewLogObject(r *http.Request) *Log {
	return &Log{
		Timestamp: time.Now().Local().String(),
		ClientIP:  getClientIP(r),
		Method:    r.Method,
		Url:       r.URL.String(),
		UserAgent: r.UserAgent(),
		Referrer:  r.Referer(),
	}
}

func (log *Log) SetElapsed(elapsed time.Duration) *Log {
	log.ResponseTime = float32(elapsed.Microseconds()) / 1000

	return log
}

func (log *Log) SetStatus(code int) *Log {
	log.StatusCode = code

	return log
}

func (log *Log) String() string {
	response, err := json.Marshal(log)

	if err != nil {
		return "{}"
	}

	return string(response)
}

type ResponseWriterWithLogging struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriterWithLogging(w http.ResponseWriter) *ResponseWriterWithLogging {
	return &ResponseWriterWithLogging{w, http.StatusOK}
}

func (w *ResponseWriterWithLogging) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rwl := NewResponseWriterWithLogging(w)
		next.ServeHTTP(rwl, r)
		elapsed := time.Since(start)

		log := NewLogObject(r).SetElapsed(elapsed).SetStatus(rwl.statusCode).String()

		fmt.Println(log)
	})
}

