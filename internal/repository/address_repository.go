package repository

import (
	"database/sql"

	"github.com/SC-Cynex/cynex-class-service/internal/dto"
)

type AddressRepository struct {
	DB *sql.DB
}

func NewAddressRepository(db *sql.DB) *AddressRepository {
	return &AddressRepository{DB: db}
}

func (r *AddressRepository) Create(address *dto.AddressRequestDTO) error {
	query := `
		INSERT INTO TBLAddress (street, city, state, zip_code)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	var id int
	err := r.DB.QueryRow(query, address.Street, address.City, address.State, address.ZipCode).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}

func (r *AddressRepository) GetByID(id int) (*dto.AddressResponseDTO, error) {
	return nil, nil
}
