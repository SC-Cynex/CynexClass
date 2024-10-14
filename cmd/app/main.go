package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := mux.NewRouter()

	PORT := os.Getenv("APP_PORT")
	if PORT == "" {
		PORT = "8080"
	}

	log.Println("Server running on port", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, r))
}
