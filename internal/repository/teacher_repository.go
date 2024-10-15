package repository

import (
	"database/sql"

	"github.com/SC-Cynex/cynex-class-service/internal/models"
)

type TeacherRepository struct {
	DB *sql.DB
}

func NewTeacherRepository(db *sql.DB) *TeacherRepository {
	return &TeacherRepository{DB: db}
}

func (r *TeacherRepository) Create(teacher *models.Teacher) error {
	query := "INSERT INTO teachers (name, email) VALUES ($1, $2) RETURNING id"
	return r.DB.QueryRow(query, teacher.Name, teacher.Email).Scan(&teacher.ID)
}

func (r *TeacherRepository) GetByID(id int) (*models.Teacher, error) {
	teacher := &models.Teacher{}
	query := "SELECT id, name, email FROM teachers WHERE id = $1"
	err := r.DB.QueryRow(query, id).Scan(&teacher.ID, &teacher.Name, &teacher.Email)
	if err != nil {
		return nil, err
	}
	return teacher, nil
}
