package teachers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/SC-Cynex/cynex-class-service/internal/models"
	"github.com/SC-Cynex/cynex-class-service/internal/services"
	"github.com/gorilla/mux"
)

type Handler struct {
	Service *services.TeacherService
}

func NewHandler(service *services.TeacherService) *Handler {
	return &Handler{Service: service}
}

// POST /teachers.
func (h *Handler) CreateTeacher(w http.ResponseWriter, r *http.Request) {
	var teacher models.Teacher
	if err := json.NewDecoder(r.Body).Decode(&teacher); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.Service.CreateTeacher(&teacher); err != nil {
		http.Error(w, "Could not create teacher", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(teacher)
}

// GET /teachers/{id}.
func (h *Handler) GetTeacher(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	teacher, err := h.Service.GetTeacher(id)
	if err != nil {
		http.Error(w, "Teacher not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(teacher)
}
