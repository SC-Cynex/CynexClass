package handlers

import (
	"net/http"

	"github.com/SC-Cynex/cynex-class-service/internal/dto"
	"github.com/SC-Cynex/cynex-class-service/internal/response"
	"github.com/SC-Cynex/cynex-class-service/internal/services"
)

type SubjectHandler struct {
	Service *services.SubjectService
}

func NewSubjectHandler(service *services.SubjectService) *SubjectHandler {
	return &SubjectHandler{Service: service}
}

func (h *SubjectHandler) GetSubjects(w http.ResponseWriter, r *http.Request) {
	subjects, err := h.Service.GetSubjects()
	if err != nil {
		response.Send(w, response.NewInternalServerErrorResponse())
		return
	}

	var subjectDTOs []dto.SubjectResponseDTO
	for _, subject := range subjects {
		subjectDTOs = append(subjectDTOs, dto.SubjectResponseDTO{
			ID:   subject.ID,
			Name: subject.Name,
		})
	}

	response.Send(w, response.NewSuccessResponse(subjectDTOs))
}
