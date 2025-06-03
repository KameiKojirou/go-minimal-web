package routes

import (
	"encoding/json"
	"main/utils"
	"net/http"
)

type Pokemon struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}


func PokemonAllHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, _ := utils.TursoDB()
	rows, _ := db.Query("SELECT * FROM pokemon")
	var pokemon []Pokemon
	for rows.Next() {
		var p Pokemon
		rows.Scan(&p.ID, &p.Name)
		pokemon = append(pokemon, p)
	}
	if pokemon == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"msg": "pokemon not found"})
		return
	}
	defer db.Close()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pokemon)
}

func PokemonInsertHandler(w http.ResponseWriter, r *http.Request) {
	db, _ := utils.TursoDB()
	w.Header().Set("Content-Type", "application/json")
	body := r.Body
	var pokemon Pokemon
	json.NewDecoder(body).Decode(&pokemon)
	res, _ := db.Exec("INSERT INTO pokemon (name) VALUES (?)", pokemon.Name)
	defer db.Close()
	if res == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"msg": "pokemon not inserted"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"msg": "inserted "+pokemon.Name})
}

func PokemonHandlerByID(w http.ResponseWriter, r *http.Request) {
	db, _ := utils.TursoDB()
	name := r.PathValue("name")
	res := Pokemon{}
	db.QueryRow("SELECT * FROM pokemon WHERE name = ?", name).Scan(&res.ID, &res.Name)
	defer db.Close()
	if res.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"msg": "pokemon not found"})
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}