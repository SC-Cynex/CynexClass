package handlers

import (
	"net/http"

	"github.com/SC-Cynex/cynex-class-service/internal/dto"
	"github.com/SC-Cynex/cynex-class-service/internal/response"
	"github.com/SC-Cynex/cynex-class-service/internal/services"
)

type CourseHandler struct {
	Service *services.CourseService
}

func NewCourseHandler(service *services.CourseService) *CourseHandler {
	return &CourseHandler{Service: service}
}

func (h *CourseHandler) GetCourses(w http.ResponseWriter, r *http.Request) {
	courses, err := h.Service.GetAllCourses()
	if err != nil {
		response.Send(w, response.NewInternalServerErrorResponse())
		return
	}

	var courseDTOs []dto.CourseResponseDTO
	for _, course := range courses {
		courseDTOs = append(courseDTOs, dto.CourseResponseDTO{
			ID:   course.ID,
			Name: course.Name,
		})
	}

	response.Send(w, response.NewSuccessResponse(courseDTOs))
}
