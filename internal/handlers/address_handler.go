package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/SC-Cynex/cynex-class-service/internal/dto"
	"github.com/SC-Cynex/cynex-class-service/internal/response"
	"github.com/SC-Cynex/cynex-class-service/internal/services"
	"github.com/SC-Cynex/cynex-class-service/internal/validators"
)

type Handler struct {
	Service *services.AddressService
}

func NewHandler(service *services.AddressService) *Handler {
	return &Handler{Service: service}
}

// @Summary Create a new address
// @Description Create a new address
// @Tags address
// @Accept  json
// @Produce  json
// @Param   address  body  dto.AddressRequestDTO  true  "Address data"
// @Success 201 {object} dto.CreatedResponse{data=dto.AddressResponseDTO} "Created"
// @Failure 400 {object} dto.BadRequestResponse "Bad Request"
// @Failure 422 {object} dto.UnprocessableEntityResponse "Unprocessable Entity"
// @Failure 500 {object} dto.InternalServerErrorResponse "Internal Server Error"
// @Router /api/v1/address [post]
func (h *Handler) CreateAddress(w http.ResponseWriter, r *http.Request) {
	var addressDTO dto.AddressRequestDTO

	if err := json.NewDecoder(r.Body).Decode(&addressDTO); err != nil {
		response.Send(w, response.NewBadRequestResponse("Failed to decode request body"))
		return
	}

	validationErrors := validators.ValidateStruct(addressDTO)
	if validationErrors != nil {
		response.Send(w, response.NewUnprocessableEntityResponse(validationErrors))
		return
	}

	err := h.Service.CreateAddress(&addressDTO)
	if err != nil {
		response.Send(w, response.NewInternalServerErrorResponse())
		return
	}

	response.Send(w, response.NewCreatedResponse(addressDTO))
}
