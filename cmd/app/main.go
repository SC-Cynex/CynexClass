package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/SC-Cynex/cynex-class-service/configs"
	"github.com/SC-Cynex/cynex-class-service/internal/initialization"
	"github.com/SC-Cynex/cynex-class-service/internal/routes"
	"github.com/SC-Cynex/cynex-class-service/pkg/migrate"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Using environment variables.")
	}

	db := configs.GetDB()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println("Error closing db connection")
		}
	}(db)

	// Auto migrate database
	migrate.RunMigrations()

	dependencies, err := initialization.InitAppDependencies(db)
	if err != nil {
		log.Fatalf("Failed to initialize app dependencies: %v", err)
	}

	r := routes.NewRouter(dependencies.AuthHandler)

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8080"
	}

	log.Printf("Server is running on port %s", appPort)
	log.Fatal(http.ListenAndServe(":"+appPort, r))
}
