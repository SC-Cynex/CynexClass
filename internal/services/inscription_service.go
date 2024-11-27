package services

import (
    "github.com/SC-Cynex/cynex-class-service/internal/models"
    "github.com/SC-Cynex/cynex-class-service/internal/repository"
)

type InscriptionService struct {
    Repo *repository.InscriptionRepository
}

func NewInscriptionService(repo *repository.InscriptionRepository) *InscriptionService {
    return &InscriptionService{Repo: repo}
}

func (s *InscriptionService) CreateInscription(inscription *models.Inscription) error {
    return s.Repo.CreateInscription(inscription)
}

func (s *InscriptionService) GetInscriptionsByUserID(userID int) ([]models.Inscription, error) {
    return s.Repo.GetInscriptionsByUserID(userID)
}

func (s *InscriptionService) GetAllInscriptions() ([]models.Inscription, error) {
	return s.Repo.GetAllInscriptions()
}