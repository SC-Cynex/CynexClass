package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/SC-Cynex/cynex-class-service/internal/dto"
	"github.com/SC-Cynex/cynex-class-service/internal/response"
	"github.com/SC-Cynex/cynex-class-service/internal/services"
	"github.com/SC-Cynex/cynex-class-service/internal/validators"
)

type AuthHandler struct {
	Service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{Service: service}
}

// @Summary Register a new user
// @Description Register a new user with email and password
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   user  body  dto.UserRegistrationRequestDTO  true  "User registration data"
// @Success 201 {object} dto.CreatedResponse{data=dto.UserResponseDTO} "Created"
// @Failure 400 {object} dto.BadRequestResponse "Bad Request"
// @Failure 422 {object} dto.UnprocessableEntityResponse "Unprocessable Entity"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal Server Error"
// @Router /api/v1/auth/register [post]
func (h *AuthHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userDTO dto.UserRegistrationRequestDTO

	if err := json.NewDecoder(r.Body).Decode(&userDTO); err != nil {
		response.Send(w, response.NewBadRequestResponse("Failed to decode request body"))
	}

	validationErrors := validators.ValidateStruct(userDTO)
	if validationErrors != nil {
		response.Send(w, response.NewUnprocessableEntityResponse(validationErrors))
	}

	err := h.Service.CreateUser(&userDTO)
	if err != nil {
		response.Send(w, response.NewInternalServerErrorResponse())
	}

	response.Send(w, response.NewCreatedResponse(userDTO))
}
