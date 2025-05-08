// routes/hello.go
package routes

import (
	"encoding/json" // For robust JSON handling
	"net/http"
)

func HelloRoutes() *http.ServeMux {
	hellor := http.NewServeMux()
	// Paths are now relative to the stripped /api/hello prefix
	hellor.HandleFunc("GET /", Hello)             // Was "GET /hello"
	hellor.HandleFunc("GET /{name}", HelloName)   // Was "GET /hello/{name}"
	return hellor
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"status":  "Hello OK",
		"code":    "200",
		"message": "Hello World",
	}
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		http.Error(
			w,
			"Failed to marshal JSON",
			http.StatusInternalServerError,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func HelloName(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	// fmt.Println("HelloName received name:", name) // For debugging

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"name": name}
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		http.Error(
			w,
			"Failed to marshal JSON",
			http.StatusInternalServerError,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
