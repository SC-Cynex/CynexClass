package services

import (
	"github.com/SC-Cynex/cynex-class-service/internal/models"
	"github.com/SC-Cynex/cynex-class-service/internal/repository"
)

type AuthService struct {
	Repo *repository.AuthRepository
}

func NewAuthService(repo *repository.AuthRepository) *AuthService {
	return &AuthService{Repo: repo}
}

func (s *AuthService) CreateUser(user *models.User) error {
	return s.Repo.Create(user)
}

func (s *AuthService) GetUser(id int) (*models.User, error) {
	return s.Repo.GetByID(id)
}
