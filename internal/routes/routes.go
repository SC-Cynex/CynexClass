package routes

import (
	"github.com/SC-Cynex/cynex-class-service/internal/api/auth"
	"github.com/gorilla/mux"
)

func NewRouter(authHandler *auth.Handler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/register", authHandler.CreateUser).Methods("POST")

	return r
}
