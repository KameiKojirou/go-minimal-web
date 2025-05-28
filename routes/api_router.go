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
  stack := middleware.MiddlewareStack(
    middleware.ExampleOneMiddleware,
    middleware.ExampleTwoMiddleware,
    middleware.ExampleMiddleware,
  )
  mux := http.NewServeMux()
  mux.Handle("/hello", stack(http.HandlerFunc(HelloHandler)))
  mux.Handle("/hello/{name}", middleware.ExampleMiddleware(http.HandlerFunc(HelloName)))
  mux.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
    log.Println("itemsHandler sees:", r.URL.Path)
    fmt.Fprintln(w, "items endpoint")
  })

  return mux
}