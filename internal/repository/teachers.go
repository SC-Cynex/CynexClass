// internal/repository/teachers.go
package repository

import (
	"github.com/SC-Cynex/cynex-class-service/configs"
	"github.com/SC-Cynex/cynex-class-service/internal/models"
)

func CreateTeacher(teacher models.Teacher) error {
	query := `INSERT INTO teachers (name, email) VALUES ($1, $2)`
	_, err := configs.DB.Exec(query, teacher.Name, teacher.Email)
	return err
}

func FindTeacherByID(id string) (models.Teacher, error) {
	var teacher models.Teacher
	query := `SELECT id, name, email FROM teachers WHERE id = $1`
	err := configs.DB.QueryRow(query, id).Scan(&teacher.ID, &teacher.Name, &teacher.Email)
	return teacher, err
}
