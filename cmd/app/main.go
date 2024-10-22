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
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	_ "github.com/SC-Cynex/cynex-class-service/docs"
)

// @title           Cynex Class Service
// @version         1.0
// @description     Cynex Class Service API Documentation

// @contact.email  team.cynexsc@gmail.com

// schemes http https
// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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

	r := mux.NewRouter()

	routes.RegisterRoutes(r, dependencies)

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8080"
	}

	log.Printf("Server is running on port %s", appPort)
	log.Fatal(http.ListenAndServe(":"+appPort, r))
}
