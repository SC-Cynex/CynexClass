package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
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

func (h *AuthHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userDTO dto.UserRegistrationRequestDTO

	if err := json.NewDecoder(r.Body).Decode(&userDTO); err != nil {
		response.Send(w, response.NewBadRequestResponse("Failed to decode request body"))
		return
	}

	validationErrors := validators.ValidateStruct(userDTO)
	if validationErrors != nil {
		response.Send(w, response.NewUnprocessableEntityResponse(validationErrors))
		return
	}

	err := h.Service.CreateUser(&userDTO)
	if err != nil {
		response.Send(w, response.NewInternalServerErrorResponse())
		return
	}

	response.Send(w, response.NewCreatedResponse(userDTO))
}


func (h *AuthHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.Service.GetAllUsers()
	if err != nil {
		response.Send(w, response.NewInternalServerErrorResponse())
		return
	}

	response.Send(w, response.NewSuccessResponse(users))
}

func (h *AuthHandler) GetUsersByRole(w http.ResponseWriter, r *http.Request) {
	roleIDStr := r.URL.Query().Get("role_id")
	roleID, err := strconv.Atoi(roleIDStr)
	if err != nil || (roleID != 1 && roleID != 2) {
		response.Send(w, response.NewBadRequestResponse("Invalid role ID"))
		return
	}

	users, err := h.Service.GetUsersByRole(roleID)
	if err != nil {
		response.Send(w, response.NewInternalServerErrorResponse())
		return
	}

	response.Send(w, response.NewSuccessResponse(users))
}

func (h *AuthHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("user_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Send(w, response.NewBadRequestResponse("Invalid user ID"))
		return
	}

	user, err := h.Service.GetUserByID(id)
	if err != nil {
		response.Send(w, response.NewBadRequestResponse("User not found"))
		return
	}

	response.Send(w, response.NewSuccessResponse(user))
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var credentials dto.UserCredentialsDTO
	
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		response.Send(w, response.NewBadRequestResponse("Failed to decode request body"))
		return
	}
	
	data, err := h.Service.CheckCredentials(credentials.Email, credentials.Password)
	if err != nil {
		response.Send(w, response.NewBadRequestResponse("Invalid credentials"))
		return
	}

	response.Send(w, response.NewSuccessResponse(dto.UserLoginResponseDTO{
		Token: "simpletoken123",
		UserID: data.UserID,
		RoleID: data.RoleID,
	}))
}
