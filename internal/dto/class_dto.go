package dto

type ClassRequestDTO struct {
	Name      string `json:"name_class" validate:"required,min=3,max=100" example:"Math 101"`
	Semester  string `json:"semester" validate:"required" example:"Fall 2024"`
	TeacherID int    `json:"teacher" validate:"required" example:"1"`
	SubjectID string `json:"subject" validate:"required" example:"2"`
	Day       string `json:"day_subject" validate:"required" example:"Monday"`
	CourseID  string `json:"course" validate:"required" example:"3"`
	Period    string `json:"period" validate:"required" example:"Morning"`
	TeacherName string  `json:"teacher_name"`
}
