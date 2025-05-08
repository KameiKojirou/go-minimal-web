package main

import (
	"main/middleware"
	"net/http"
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

func Router() {
	err := godotenv.Load()
	if err != nil {
		os.Exit(1)
	}
	router := http.NewServeMux()
	router.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	router.Handle("/", SpaRoutes())
	router.Handle("/api/", http.StripPrefix("/api", APIRouter()))
	server := middleware.CorsMiddleware(router)
	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}
	fmt.Println("Listening on port " + port)
	http.ListenAndServe(fmt.Sprint(":", port), server)
}