package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func Chain(mdws ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for _, m := range mdws {
			next = m(next)
		}
		return next
	}
}
