package main

import (
	"log"
	"net/http"
	"os"

	"github.com/SC-Cynex/cynex-class-service/configs"
	"github.com/SC-Cynex/cynex-class-service/internal/api/teachers"
	"github.com/SC-Cynex/cynex-class-service/internal/repository"
	"github.com/SC-Cynex/cynex-class-service/internal/services"
	"github.com/SC-Cynex/cynex-class-service/pkg/migrate"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Using environment variables.")
	}

	db := configs.GetDB()
	defer db.Close()

	migrate.RunMigrations()

	teacherRepo := repository.NewTeacherRepository(db)
	teacherService := services.NewTeacherService(teacherRepo)
	teacherHandler := teachers.NewHandler(teacherService)

	r := mux.NewRouter()

	PORT := os.Getenv("APP_PORT")
	if PORT == "" {

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8080"
	}

	log.Printf("Server is running on port %s", appPort)
	log.Fatal(http.ListenAndServe(":"+appPort, r))
}
