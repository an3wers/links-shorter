package middleware

import (
	"context"
	"go/links-shorter/configs"
	"go/links-shorter/pkg/jwt"
	"go/links-shorter/pkg/resp"
	"net/http"
	"strings"
)

type key string

const (
	ContextAuthKey key = "contextAuthKey"
)

func Auth(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		bearerToken := r.Header.Get("Authorization")

		if bearerToken == "" {
			resp.Json(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		if !strings.HasPrefix(bearerToken, "Bearer ") {
			resp.Json(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		token := strings.TrimPrefix(bearerToken, "Bearer ")

		isValid, data := jwt.NewJWT(config.Auth.SecretKey).Parse(token)

		if !isValid {
			resp.Json(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		reqContext := r.Context()
		ctx := context.WithValue(reqContext, ContextAuthKey, data.Email)

		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
