package routes

import (
	"encoding/json"
	"fmt"
	"main/utils"
	"net/http"
	// "strings"
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

func ShowHelloHandler(w http.ResponseWriter, r *http.Request) {
	db, _ := utils.TursoDB()
	defer db.Close()
	res := []string{}
	rows, err := db.Query("SELECT name FROM hello")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			fmt.Println("Error: ", err)
		}
		fmt.Println(name)
		res = append(res, name)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}