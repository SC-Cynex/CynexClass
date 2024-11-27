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

type ClassHandler struct {
	Service *services.ClassService
}

func NewClassHandler(service *services.ClassService) *ClassHandler {
	return &ClassHandler{Service: service}
}

func (h *ClassHandler) CreateClass(w http.ResponseWriter, r *http.Request) {
	var classDTO dto.ClassRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&classDTO); err != nil {
		response.Send(w, response.NewBadRequestResponse("Failed to decode request body"))
		return
	}
	
	class := &models.Class{
		Name:      classDTO.Name,
		Semester:  classDTO.Semester,
		TeacherID: classDTO.TeacherID,
		SubjectID: classDTO.SubjectID,
		Day:       classDTO.Day,
		CourseID:  classDTO.CourseID,
		Period:    classDTO.Period,
	}

	err := h.Service.CreateClass(class)
	if err != nil {
		response.Send(w, response.NewInternalServerErrorResponse())
		return
	}

	response.Send(w, response.NewCreatedResponse(class))
}

func (h *ClassHandler) GetAllClasses(w http.ResponseWriter, r *http.Request) {
	classes, err := h.Service.GetAllClasses()
	if err != nil {
		response.Send(w, response.NewInternalServerErrorResponse())
		return
	}

	response.Send(w, response.NewSuccessResponse(classes))
}

func (h *ClassHandler) GetClassesByUserID(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		response.Send(w, response.NewBadRequestResponse("Invalid user ID"))
		return
	}

	classes, err := h.Service.GetClassesByUserID(userID)
	if err != nil {
		response.Send(w, response.NewInternalServerErrorResponse())
		return
	}

	response.Send(w, response.NewSuccessResponse(classes))
}

func (h *ClassHandler) GetClassesByClassID(w http.ResponseWriter, r *http.Request) {
	classIDStr := r.URL.Query().Get("class_id")
	classID, err := strconv.Atoi(classIDStr)
	if err != nil {
		response.Send(w, response.NewBadRequestResponse("Invalid user ID"))
		return
	}

	classes, err := h.Service.GetClassesByClassID(classID)
	if err != nil {
		response.Send(w, response.NewInternalServerErrorResponse())
		return
	}

	response.Send(w, response.NewSuccessResponse(classes))
}
