package main

import (
	"log/slog"
	"net/http"
)

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		slog.Info(
			"request",
			"method", r.Method,
			"url", r.RequestURI,
		)
	})
}
