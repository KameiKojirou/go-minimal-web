package routes

import (
	"net/http"
)

func HelloRoutes () *http.ServeMux {
	hellor := http.NewServeMux()
	hellor.HandleFunc("GET /hello", Hello)
	return hellor
}


func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`
		{
			'status': 'OK'
			'code': '200'
			'message': 'Hello World'
		}
	`))
}