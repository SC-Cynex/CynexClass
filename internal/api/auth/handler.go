package auth

import (
	"encoding/json"
	"net/http"

	"github.com/SC-Cynex/cynex-class-service/internal/dto"
	"github.com/SC-Cynex/cynex-class-service/internal/services"
	"github.com/SC-Cynex/cynex-class-service/internal/utils"
	"github.com/SC-Cynex/cynex-class-service/internal/validators"
)

type Handler struct {
	Service *services.AuthService
}

func NewHandler(service *services.AuthService) *Handler {
	return &Handler{Service: service}
}

// POST /api/v1/register.
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userDTO dto.UserRegistrationRequestDTO

	if err := json.NewDecoder(r.Body).Decode(&userDTO); err != nil {
		utils.SendAPIResponse(w, http.StatusUnprocessableEntity, utils.TitleUnprocessableEntity, false, "Failed to decode request body")
		return
	}

	validationErrors := validators.ValidateStruct(userDTO)
	if validationErrors != nil {
		utils.SendAPIResponse(w, http.StatusUnprocessableEntity, utils.TitleUnprocessableEntity, false, validationErrors)
		return
	}

	/*
		if err := h.Service.CreateUser(&userDTO); err != nil {
			utils.SendErrorResponse(w, http.StatusInternalServerError, utils.TitleInternalServerError, "Could not create user")
			return
		}
	*/

	utils.SendAPIResponse(w, http.StatusCreated, utils.TitleCreated, true, userDTO)
}
