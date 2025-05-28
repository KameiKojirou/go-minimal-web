package routes

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}

func HelloName(w http.ResponseWriter, r *http.Request) {
	name :=  r.PathValue("name")
	fmt.Fprintf(w, "Hello %s", name)
}