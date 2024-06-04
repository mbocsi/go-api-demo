package middleware

import (
	"log/slog"
	"net/http"
)

type RequestLogger struct {
	handler http.Handler
}

func NewRequestLogger(next http.Handler) *RequestLogger {
	return &RequestLogger{next}
}

func (rl *RequestLogger) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	slog.Info(
		"Incoming request",
		"method",
		req.Method,
		"url",
		req.URL,
		"user",
		req.UserAgent(),
	)
	rl.handler.ServeHTTP(res, req)
}
