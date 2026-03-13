package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	
	"backend/internal/api"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	api.RegisterRoutes()

	log.Println("Server running on :8080")

	http.ListenAndServe(":8080", nil)
}