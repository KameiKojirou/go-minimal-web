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

func ExampleOneMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("example one middleware: %s %s\n", r.Method, r.URL.Path)
    next.ServeHTTP(w, r)
  })
}

func ExampleTwoMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("example two middleware: %s %s\n", r.Method, r.URL.Path)
    next.ServeHTTP(w, r)
  })
}