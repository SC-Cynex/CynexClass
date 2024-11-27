package repository

import (
	"database/sql"
	"github.com/SC-Cynex/cynex-class-service/internal/models"
)

type ClassRepository struct {
	DB *sql.DB
}

func NewClassRepository(db *sql.DB) *ClassRepository {
	return &ClassRepository{DB: db}
}

func (r *ClassRepository) CreateClass(class *models.Class) error {
	query := `
		INSERT INTO TBLClass (name_class, semester, teacher_id, subject, day_subject, course, period)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id_class, created_at, updated_at
	`
	err := r.DB.QueryRow(
		query,
		class.Name,
		class.Semester,
		class.TeacherID,
		class.SubjectID,
		class.Day,
		class.CourseID,
		class.Period,
	).Scan(&class.ID, &class.CreatedAt, &class.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *ClassRepository) GetAllClasses() ([]models.Class, error) {
	var classes []models.Class
	query := `SELECT id_class, name_class, semester, teacher_id, subject, day_subject, course, period, TBC.created_at, TBC.updated_at, TBU.username AS teacher_name 
				FROM TBLClass TBC
				INNER JOIN TBLUser TBU ON TBU.user_id = TBC.teacher_id`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var class models.Class
		if err := rows.Scan(
			&class.ID,
			&class.Name,
			&class.Semester,
			&class.TeacherID,
			&class.SubjectID,
			&class.Day,
			&class.CourseID,
			&class.Period,
			&class.CreatedAt,
			&class.UpdatedAt,
			&class.TeacherName,
		); err != nil {
			return nil, err
		}
		classes = append(classes, class)
	}

	return classes, nil
}

func (r *ClassRepository) GetClassesByUserID(userID int) ([]models.Class, error) {
	var classes []models.Class
	query := "SELECT id_class, name_class, semester, teacher_id, subject, day_subject, course, period, created_at, updated_at FROM TBLClass WHERE teacher_id = $1"
	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var class models.Class
		if err := rows.Scan(
			&class.ID,
			&class.Name,
			&class.Semester,
			&class.TeacherID,
			&class.SubjectID,
			&class.Day,
			&class.CourseID,
			&class.Period,
			&class.CreatedAt,
			&class.UpdatedAt,
		); err != nil {
			return nil, err
		}
		classes = append(classes, class)
	}

	return classes, nil
}

func (r *ClassRepository) GetClassesByClassID(classID int) ([]models.Class, error) {
	var classes []models.Class
	query := `SELECT id_class, name_class, semester, teacher_id, subject, day_subject, course, period, TBC.created_at, TBC.updated_at, TBU.username AS teacher_name
				FROM TBLClass TBC
				INNER JOIN TBLUser TBU ON TBU.user_id = TBC.teacher_id
				WHERE id_class = $1;
				`
	rows, err := r.DB.Query(query, classID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var class models.Class
		if err := rows.Scan(
			&class.ID,
			&class.Name,
			&class.Semester,
			&class.TeacherID,
			&class.SubjectID,
			&class.Day,
			&class.CourseID,
			&class.Period,
			&class.CreatedAt,
			&class.UpdatedAt,
			&class.TeacherName,
		); err != nil {
			return nil, err
		}
		classes = append(classes, class)
	}

	return classes, nil
}
