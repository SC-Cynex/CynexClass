package repository

import (
	"database/sql"

	"github.com/SC-Cynex/cynex-class-service/internal/models"
)

type CourseRepository struct {
	DB *sql.DB
}

func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{DB: db}
}

func (r *CourseRepository) GetAllCourses() ([]models.Course, error) {
	var courses []models.Course

	query := "SELECT id_course, name_course, created_at, updated_at FROM TBLCourse"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var course models.Course
		if err := rows.Scan(&course.ID, &course.Name, &course.CreatedAt, &course.UpdatedAt); err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return courses, nil
}
