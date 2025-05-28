package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func MiddlewareStack(mw ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for _, m := range mw {
			next = m(next)
		}
		return next
	}
}