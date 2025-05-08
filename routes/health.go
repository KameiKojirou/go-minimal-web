package routes

import (
	"net/http"
)

func HealthRoutes() *http.ServeMux {
	health := http.NewServeMux()
	health.HandleFunc("GET /", HealthGet)
	health.HandleFunc("POST /", HealthPost)
	return health
}

func HealthPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`
		{
			"status": "OK",
			"code": "200",
			"message": "Health OK",
			"method": "POST"
		}
	`))
}

func HealthGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`
		{
			"status": "OK",
			"code": "200",
			"message": "Health OK"
			"method": "GET"
		}
	`))
}