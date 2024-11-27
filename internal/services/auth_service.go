package services

import (
	"errors"
	"github.com/SC-Cynex/cynex-class-service/internal/dto"
	"github.com/SC-Cynex/cynex-class-service/internal/models"
	"github.com/SC-Cynex/cynex-class-service/internal/repository"
)

type AuthService struct {
	AuthRepo       *repository.AuthRepository
	AddressService *AddressService
}

func NewAuthService(authRepo *repository.AuthRepository, addressService *AddressService) *AuthService {
	return &AuthService{
		AuthRepo:       authRepo,
		AddressService: addressService,
	}
}

func (s *AuthService) CreateUser(userDTO *dto.UserRegistrationRequestDTO) error {
	user := &models.User{
		Username: userDTO.Username,
		Role:     userDTO.Role,
		Email:    userDTO.Email,
		Password: userDTO.Password,
		IsActive: true,
	}

	err := s.AuthRepo.Create(user)
	if err != nil {
		return errors.New("failed to create user")
	}

	return nil
}

func (s *AuthService) GetAllUsers() ([]models.User, error) {
	return s.AuthRepo.GetAll()
}

func (s *AuthService) GetUsersByRole(roleID int) ([]models.User, error) {
	return s.AuthRepo.GetByRole(roleID)
}

func (s *AuthService) GetUserByID(id int) (*models.User, error) {
	return s.AuthRepo.GetByID(id)
}

func (s *AuthService) CheckCredentials(email string, password string) (*models.User, error) {
	return s.AuthRepo.CheckCredentials(email, password)
}
