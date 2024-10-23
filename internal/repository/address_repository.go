package repository

import (
	"database/sql"

	"github.com/SC-Cynex/cynex-class-service/internal/dto"
	"github.com/SC-Cynex/cynex-class-service/internal/models"
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

func (r *AddressRepository) GetAddresses() ([]models.Address, error) {
	var addresses []models.Address

	query := "SELECT * FROM TBLAddress"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var address models.Address
		if err := rows.Scan(
			&address.ID,
			&address.Street,
			&address.Number,
			&address.Neighborhood,
			&address.City,
			&address.State,
			&address.ZipCode,
			&address.CreatedAt,
			&address.UpdatedAt,
		); err != nil {
			return nil, err
		}
		addresses = append(addresses, address)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return addresses, nil
}

func (r *AddressRepository) GetByID(id int) (*dto.AddressResponseDTO, error) {
	return nil, nil
}
