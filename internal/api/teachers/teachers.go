package teachers

import (
	"encoding/json"
	"net/http"

	"github.com/SC-Cynex/cynex-class-service/internal/models"
	"github.com/SC-Cynex/cynex-class-service/internal/services"
	"github.com/gorilla/mux"
)

// Handler to create a new teacher
func CreateTeacher(w http.ResponseWriter, r *http.Request) {
	var teacher models.Teacher
	err := json.NewDecoder(r.Body).Decode(&teacher)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the service to create the teacher
	err = services.AddTeacher(teacher)
	if err != nil {
		http.Error(w, "Error creating teacher", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Teacher created"})
}

// Handler to fetch a teacher by ID
func GetTeacher(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	teacher, err := services.GetTeacherByID(id)
	if err != nil {
		http.Error(w, "Teacher not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(teacher)
}
