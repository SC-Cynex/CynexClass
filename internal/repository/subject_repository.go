package repository

import (
	"database/sql"
	"github.com/SC-Cynex/cynex-class-service/internal/models"
)

type SubjectRepository struct {
	DB *sql.DB
}

func NewSubjectRepository(db *sql.DB) *SubjectRepository {
	return &SubjectRepository{DB: db}
}

func (r *SubjectRepository) GetSubjects() ([]models.Subject, error) {
	var subjects []models.Subject

	query := "SELECT id_subject, name_subject, created_at, updated_at FROM TBLSubject"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var subject models.Subject
		if err := rows.Scan(
			&subject.ID,
			&subject.Name,
			&subject.CreatedAt,
			&subject.UpdatedAt,
		); err != nil {
			return nil, err
		}
		subjects = append(subjects, subject)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return subjects, nil
}
