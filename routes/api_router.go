package routes

import (
	"fmt"
	"net/http"
)


func APIRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "users endpoint")
	})
	mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		path := r.PathValue("id")
		fmt.Fprintln(w, path)
	})
	mux.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "items endpoint")
	})
	return mux
}