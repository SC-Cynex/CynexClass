package services

import (
	"github.com/SC-Cynex/cynex-class-service/internal/dto"
	"github.com/SC-Cynex/cynex-class-service/internal/models"
	"github.com/SC-Cynex/cynex-class-service/internal/repository"
)

type AddressService struct {
	Repo *repository.AddressRepository
}

func NewAddressService(repo *repository.AddressRepository) *AddressService {
	return &AddressService{Repo: repo}
}

func (s *AddressService) CreateAddress(address *dto.AddressRequestDTO) error {
	return s.Repo.Create(address)
}

func (s *AddressService) GetAddresses() ([]models.Address, error) {
	addresses, err := s.Repo.GetAddresses()
	if err != nil {
		return nil, err
	}

	return addresses, nil
}

func (s *AddressService) GetAddress(id int) (*models.Address, error) {
	return nil, nil
}
