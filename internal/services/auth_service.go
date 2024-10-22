package services

import (
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

func (s *AuthService) CreateUser(user *dto.UserRegistrationRequestDTO) error {
	// create the address first
	//addressId := s.AddressService.CreateAddress(&user.Address)
	// create the user
	return nil
}

func (s *AuthService) GetUser(id int) (*models.User, error) {
	return s.AuthRepo.GetByID(id)
}
