package repository

import (
	"database/sql"
	"errors"
	"github.com/SC-Cynex/cynex-class-service/internal/models"
)

type AuthRepository struct {
	DB *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

func (r *AuthRepository) Create(user *models.User) error {
	query := `
		INSERT INTO TBLUser (username, role_id, address_id, email, password, is_active)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING user_id`
	err := r.DB.QueryRow(query, user.Username, user.Role, 1, user.Email, user.Password, user.IsActive).Scan(&user.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthRepository) GetAll() ([]models.User, error) {
	var users []models.User
	query := "SELECT user_id, role_id, address_id, username, email, is_active, created_at, updated_at FROM TBLUser"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(
			&user.UserID,
			&user.RoleID,
			&user.AddressID,
			&user.Username,
			&user.Email,
			&user.IsActive,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *AuthRepository) GetByRole(roleID int) ([]models.User, error) {
	var users []models.User
	query := "SELECT user_id, role_id, address_id, username, email, is_active, created_at, updated_at FROM TBLUser WHERE role_id = $1"
	rows, err := r.DB.Query(query, roleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(
			&user.UserID,
			&user.RoleID,
			&user.AddressID,
			&user.Username,
			&user.Email,
			&user.IsActive,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *AuthRepository) GetByID(id int) (*models.User, error) {
	user := &models.User{}
	query := "SELECT user_id, role_id, address_id, username, password, email, is_active, created_at, updated_at FROM TBLUser WHERE user_id = $1"
	err := r.DB.QueryRow(query, id).Scan(
		&user.UserID,
		&user.RoleID,
		&user.AddressID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	} else if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *AuthRepository) GetByEmail(email string) (*models.User, error) {
	user := &models.User{}

	query := "SELECT user_id, role_id, address_id, username, email, is_active, created_at, updated_at FROM TBLUser WHERE email = $1"
	err := r.DB.QueryRow(query, email).Scan(
		&user.UserID,
		&user.RoleID,
		&user.AddressID,
		&user.Username,
		&user.Email,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	} else if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *AuthRepository) CheckCredentials(email, password string) (*models.User, error) {
	user := &models.User{}
	query := "SELECT user_id, role_id, address_id, username, email, is_active, created_at, updated_at FROM TBLUser WHERE email = $1 AND password = $2"
	err := r.DB.QueryRow(query, email, password).Scan(
		&user.UserID,
		&user.RoleID,
		&user.AddressID,
		&user.Username,
		&user.Email,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("invalid credentials")
	} else if err != nil {
		return nil, err
	}

	return user, nil
}
