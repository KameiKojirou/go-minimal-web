package routes

import (
	"fmt"
	"log"
	"main/middleware"

	// "main/middleware"
	"net/http"
)

// APIRouter builds its own ServeMux just for /api.
func APIRouter() *http.ServeMux {
  mux := http.NewServeMux()
  mux.Handle("/hello", middleware.ExampleMiddleware(http.HandlerFunc(HelloHandler)))
  mux.Handle("/hello/{name}", middleware.ExampleMiddleware(http.HandlerFunc(HelloName)))
  mux.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
    log.Println("itemsHandler sees:", r.URL.Path)
    fmt.Fprintln(w, "items endpoint")
  })

  return mux
}