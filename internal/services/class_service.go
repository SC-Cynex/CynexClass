package services

import (
	"github.com/SC-Cynex/cynex-class-service/internal/models"
	"github.com/SC-Cynex/cynex-class-service/internal/repository"
)

type ClassService struct {
	ClassRepo *repository.ClassRepository
}

func NewClassService(classRepo *repository.ClassRepository) *ClassService {
	return &ClassService{
		ClassRepo: classRepo,
	}
}

func (s *ClassService) CreateClass(class *models.Class) error {
	return s.ClassRepo.CreateClass(class)
}

func (s *ClassService) GetAllClasses() ([]models.Class, error) {
	return s.ClassRepo.GetAllClasses()
}

func (s *ClassService) GetClassesByUserID(userID int) ([]models.Class, error) {
	return s.ClassRepo.GetClassesByUserID(userID)
}

func (s *ClassService) GetClassesByClassID(classID int) ([]models.Class, error) {
	return s.ClassRepo.GetClassesByClassID(classID)
}
