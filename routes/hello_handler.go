package routes

import (
	"encoding/json"
	// "fmt"
	"net/http"
)

type HelloNameModel struct {
	Name string `json:"name"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"msg": "Hello, World!"})
}

func HelloNameHandler(w http.ResponseWriter, r *http.Request) {
	name :=  r.PathValue("name")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	res := HelloNameModel{Name: name}
	json.NewEncoder(w).Encode(res)
}