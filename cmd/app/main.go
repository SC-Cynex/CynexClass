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
	"github.com/rs/cors"

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
	// Carregar as variáveis de ambiente
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Using environment variables.")
	}

	// Conectar ao banco de dados
	db := configs.GetDB()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Println("Error closing db connection")
		}
	}(db)

	// Rodar as migrações automaticamente
	migrate.RunMigrations()

	// Inicializar as dependências do aplicativo
	dependencies, err := initialization.InitAppDependencies(db)
	if err != nil {
		log.Fatalf("Failed to initialize app dependencies: %v", err)
	}

	// Configurar o roteador
	r := mux.NewRouter()

	// Registrar as rotas
	routes.RegisterRoutes(r, dependencies)

	// Configuração do CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	// Aplica o middleware CORS
	handler := corsHandler.Handler(r)

	// Recuperar a porta do aplicativo a partir das variáveis de ambiente
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8080"
	}

	log.Printf("Server is running on port %s", appPort)
	log.Fatal(http.ListenAndServe(":"+appPort, handler))
}
