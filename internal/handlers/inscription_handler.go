package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "github.com/SC-Cynex/cynex-class-service/internal/dto"
    "github.com/SC-Cynex/cynex-class-service/internal/models"
    "github.com/SC-Cynex/cynex-class-service/internal/response"
    "github.com/SC-Cynex/cynex-class-service/internal/services"
)

type InscriptionHandler struct {
    Service *services.InscriptionService
}

func NewInscriptionHandler(service *services.InscriptionService) *InscriptionHandler {
    return &InscriptionHandler{Service: service}
}

func (h *InscriptionHandler) CreateInscription(w http.ResponseWriter, r *http.Request) {
    var inscriptionDTO dto.InscriptionRequestDTO
    if err := json.NewDecoder(r.Body).Decode(&inscriptionDTO); err != nil {
        response.Send(w, response.NewBadRequestResponse("Failed to decode request body"))
        return
    }

    inscription := &models.Inscription{
        UserID:  inscriptionDTO.UserID,
        ClassID: inscriptionDTO.ClassID,
    }

    err := h.Service.CreateInscription(inscription)
    if err != nil {
        response.Send(w, response.NewInternalServerErrorResponse())
        return
    }

    response.Send(w, response.NewCreatedResponse(inscription))
}

func (h *InscriptionHandler) GetInscriptionByUserID(w http.ResponseWriter, r *http.Request) {
    userIDStr := r.URL.Query().Get("user_id")
    userID, err := strconv.Atoi(userIDStr)
    if err != nil {
        response.Send(w, response.NewBadRequestResponse("Invalid user ID"))
        return
    }

    inscriptions, err := h.Service.GetInscriptionsByUserID(userID)
    if err != nil {
        response.Send(w, response.NewInternalServerErrorResponse())
        return
    }

    response.Send(w, response.NewSuccessResponse(inscriptions))
}

func (h *InscriptionHandler) GetAllInscriptions(w http.ResponseWriter, r *http.Request) {
	inscriptions, err := h.Service.GetAllInscriptions()
	if err != nil {
		response.Send(w, response.NewInternalServerErrorResponse())
		return
	}

	response.Send(w, response.NewSuccessResponse(inscriptions))
}