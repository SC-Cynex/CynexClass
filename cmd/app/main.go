// cmd/app/main.go
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
