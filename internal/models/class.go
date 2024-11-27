package models

import "time"

type Class struct {
	ID        int       `db:"id_class"`
	Name      string    `db:"name_class"`
	Semester  string    `db:"semester"`
	TeacherID int       `db:"teacher_id"`
	SubjectID string    `db:"subject"`
	Day       string    `db:"day_subject"`
	CourseID  string    `db:"course"`
	Period    string    `db:"period"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	TeacherName string  `db:"teacher_name"`
}
