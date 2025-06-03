package routes

import (
	"main/middleware"
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
  mux.Handle("/admin/", http.StripPrefix("/admin", AdminRouter()))
  mux.Handle("/hello", stack(http.HandlerFunc(HelloHandler)))
  mux.Handle("/login", stack(http.HandlerFunc(LoginHandler))) 
  mux.Handle("/hello/all", stack(http.HandlerFunc(ShowHelloHandler)))
  mux.Handle("/hello/{name}", middleware.ExampleMiddleware(http.HandlerFunc(HelloNameHandler)))
  mux.Handle("GET /pokemon", stack(http.HandlerFunc(PokemonAllHandler)))
  mux.Handle("POST /pokemon", stack(http.HandlerFunc(PokemonInsertHandler)))
  mux.Handle("GET /pokemon/{name}", stack(http.HandlerFunc(PokemonHandlerByID)))
  return mux
}