package routes

import (
	"github.com/SC-Cynex/cynex-class-service/internal/initialization"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func RegisterRoutes(router *mux.Router, deps *initialization.AppDependencies) {
	// Auth Routes
	router.HandleFunc("/api/v1/auth/register", deps.AuthHandler.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/auth/users", deps.AuthHandler.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/v1/auth/users/role", deps.AuthHandler.GetUsersByRole).Methods("GET")
	router.HandleFunc("/api/v1/auth/user", deps.AuthHandler.GetUserByID).Methods("GET")
	router.HandleFunc("/api/v1/auth/login", deps.AuthHandler.Login).Methods("POST")

	// Address Routes
	router.HandleFunc("/api/v1/address", deps.AddressHandler.CreateAddress).Methods("POST")
	router.HandleFunc("/api/v1/address", deps.AddressHandler.GetAddresses).Methods("GET")

	// Subject Routes
	router.HandleFunc("/api/v1/subjects", deps.SubjectHandler.GetSubjects).Methods("GET")

	// Class Routes
	router.HandleFunc("/api/v1/classes", deps.ClassHandler.GetAllClasses).Methods("GET")
	router.HandleFunc("/api/v1/class", deps.ClassHandler.CreateClass).Methods("POST")
	router.HandleFunc("/api/v1/classes/user", deps.ClassHandler.GetClassesByUserID).Methods("GET")
	router.HandleFunc("/api/v1/classes/class", deps.ClassHandler.GetClassesByClassID).Methods("GET")
	
	// Inscription Routes
	router.HandleFunc("/api/v1/inscription", deps.InscriptionHandler.CreateInscription).Methods("POST")
    router.HandleFunc("/api/v1/inscription", deps.InscriptionHandler.GetAllInscriptions).Methods("GET")
    router.HandleFunc("/api/v1/inscription/user", deps.InscriptionHandler.GetInscriptionByUserID).Methods("GET")

	// Swagger Routes
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler())
}
