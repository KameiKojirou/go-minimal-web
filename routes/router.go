package routes

import (
	"main/middleware"
	"net/http"
)

func Router() {
	routes := http.NewServeMux()
	routes.Handle("/health/", http.StripPrefix("/health", middleware.Logger(HealthRoutes())))
	routes.Handle("/hello/", http.StripPrefix("/hello", HelloRoutes()))
	server := middleware.CorsMiddleware(routes)
	port := ":1323"
	http.ListenAndServe(port, server)
	HelloRoutes()
}