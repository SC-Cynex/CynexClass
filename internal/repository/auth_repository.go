package repository

import (
	"database/sql"

	"github.com/SC-Cynex/cynex-class-service/internal/models"
)

type AuthRepository struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

func (r *AuthRepository) Create(user *models.User) error {
	return nil
}

func (r *AuthRepository) GetByID(id int) (*models.User, error) {
	user := &models.User{}
	return user, nil
}
