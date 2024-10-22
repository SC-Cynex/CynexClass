package routes

import (
	"github.com/SC-Cynex/cynex-class-service/internal/initialization"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func RegisterRoutes(router *mux.Router, deps *initialization.AppDependencies) {
	// Auth Routes
	router.HandleFunc("/api/v1/auth/register", deps.AuthHandler.CreateUser).Methods("POST")

	// Address Routes
	router.HandleFunc("/api/v1/address", deps.AddressHandler.CreateAddress).Methods("POST")

	// Swagger Routes
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler())
}
