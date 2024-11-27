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

func (h *Handler) GetAddresses(w http.ResponseWriter, r *http.Request) {
	addresses, err := h.Service.GetAddresses()
	if err != nil {
		response.Send(w, response.NewInternalServerErrorResponse())
		return
	}

	response.Send(w, response.NewSuccessResponse(addresses))
}
