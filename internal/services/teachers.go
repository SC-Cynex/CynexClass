package services

import (
	"github.com/SC-Cynex/cynex-class-service/internal/models"
	"github.com/SC-Cynex/cynex-class-service/internal/repository"
)

// Add a new teacher
func AddTeacher(teacher models.Teacher) error {
	return repository.CreateTeacher(teacher)
}

// Fetch a teacher by ID
func GetTeacherByID(id string) (models.Teacher, error) {
	return repository.FindTeacherByID(id)
}
