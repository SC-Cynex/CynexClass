package auth

import (
	"encoding/json"
	"net/http"

	"github.com/SC-Cynex/cynex-class-service/internal/models"
	"github.com/SC-Cynex/cynex-class-service/internal/services"
)

type Handler struct {
	Service *services.AuthService
}

func NewHandler(service *services.AuthService) *Handler {
	return &Handler{Service: service}
}

// POST /api/v1/register.
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.Service.CreateUser(&user); err != nil {
		http.Error(w, "Could not create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
