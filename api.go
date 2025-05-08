package main

import (
	"main/middleware"
	"main/routes"
	"net/http"
)

func APIRouter() *http.ServeMux {
	apirouter := http.NewServeMux()
	apirouter.Handle("/health/", http.StripPrefix("/health", middleware.Logger(routes.HealthRoutes())))
	apirouter.Handle("/hello/", http.StripPrefix("/hello", middleware.Logger(routes.HelloRoutes())))
	return apirouter
}