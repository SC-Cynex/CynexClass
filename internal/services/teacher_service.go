package services

import (
	"github.com/SC-Cynex/cynex-class-service/internal/models"
	"github.com/SC-Cynex/cynex-class-service/internal/repository"
)

type TeacherService struct {
	Repo *repository.TeacherRepository
}

func NewTeacherService(repo *repository.TeacherRepository) *TeacherService {
	return &TeacherService{Repo: repo}
}

func (s *TeacherService) CreateTeacher(teacher *models.Teacher) error {
	return s.Repo.Create(teacher)
}

func (s *TeacherService) GetTeacher(id int) (*models.Teacher, error) {
	return s.Repo.GetByID(id)
}
