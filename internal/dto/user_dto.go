package dto

type UserRegistrationRequestDTO struct {
	Username string `json:"username" validate:"required,min=3,max=50" example:"Israel"`
	Email string `json:"email" validate:"required,email" example:"israel@gmail.com"`
	Role int `json:"role" example:"student"`
	Password string `json:"password,omitempty" validate:"required,min=6" example:"123456"`
}

type UserResponseDTO struct {
	ID int `json:"id" example:"1"`
	Username string `json:"username" example:"Israel Schroeder"`
	Email string `json:"email" example:"israel@gmail.com"`
	Role int `json:"role" example:"student"`
}

type UserCredentialsDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponseDTO struct {
	Token string `json:"token"`
	UserID int `json:"user_id"`
	RoleID int `json:"role_id"`
}
