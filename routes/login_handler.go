package routes

import (
	"encoding/json"
	"fmt"
	"main/utils"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	utils.CreateCookie(w)
	myCookie := utils.GetCookie(r)
	if myCookie == "" {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"msg": "cookie not found"})
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"msg": "Hello, World!", "cookie": fmt.Sprintf("%v", myCookie)})
}

