package middleware

import (
	"go/links-shorter/pkg/resp"
	"net/http"
	"strings"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		bearerToken := r.Header.Get("Authorization")

		if bearerToken == "" {
			resp.Json(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		// token := strings.Split(bearerToken, " ")

		// if len(token) < 2 {
		// 	resp.Json(w, http.StatusUnauthorized, "Unauthorized")
		// 	return
		// }

		// r.Header.Set("AuthToken", token[1])
		// Так проще
		token := strings.TrimPrefix(bearerToken, "Bearer ")
		r.Header.Set("AuthToken", token)

		next.ServeHTTP(w, r)
	})
}
