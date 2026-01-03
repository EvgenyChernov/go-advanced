package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		wrapper := &WrapperWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}
		next.ServeHTTP(wrapper, r)
		duration := time.Since(startTime)
		log.Printf("%s %s %s %d %s\n", r.Method, r.URL.Path, r.Proto, wrapper.statusCode, duration)
	})
}
