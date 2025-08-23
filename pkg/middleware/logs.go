package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		wrap := &WrapperWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrap, r)

		// after request
		log.Println("StatusCode: ", wrap.StatusCode, "Time: ", time.Since(start).Milliseconds(), "ms", "Path: ", r.URL.Path, "Method: ", r.Method)

	})
}
