package server

import (
	"log/slog"
	"net/http"
	"reflect"
	"runtime"
	"time"
)

func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		end := time.Now()
		name := runtime.FuncForPC(reflect.ValueOf(next).Pointer()).Name()
		slog.Debug(name, "response time", end.Sub(start))
	}
}
