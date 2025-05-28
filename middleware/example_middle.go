package middleware

import (
	"fmt"
	"net/http"
)
// ExampleMiddleware logs each request and then calls the next handler.
func ExampleMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("example middleware: %s %s\n", r.Method, r.URL.Path)
    next.ServeHTTP(w, r)
  })
}