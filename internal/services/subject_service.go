package services

import (
	"github.com/SC-Cynex/cynex-class-service/internal/models"
	"github.com/SC-Cynex/cynex-class-service/internal/repository"
)

type SubjectService struct {
	Repo *repository.SubjectRepository
}

func NewSubjectService(repo *repository.SubjectRepository) *SubjectService {
	return &SubjectService{Repo: repo}
}

func (s *SubjectService) GetSubjects() ([]models.Subject, error) {
	return s.Repo.GetSubjects()
}
