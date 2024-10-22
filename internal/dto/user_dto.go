package dto

// UserRegistrationRequestDTO represents the data required for user registration.
type UserRegistrationRequestDTO struct {
	// Username is the unique identifier for the user.
	Username string `json:"username" validate:"required,min=3,max=50" example:"Israel"`

	// Email is the user's email address.
	Email string `json:"email" validate:"required,email" example:"israel@gmail.com"`

	// Role defines the user's role in the system.
	Role string `json:"role" validate:"required,oneof=teacher student" example:"student"`

	// Password is the user's password.
	Password string `json:"password,omitempty" validate:"required,min=6" example:"123456"`

	// Address contains the user's address details.
	Address AddressRequestDTO `json:"address" validate:"required"`
}

// UserResponseDTO represents the data returned for a user.
type UserResponseDTO struct {
	// ID is the unique identifier for the user.
	ID int `json:"id" example:"1"`

	// Username is the unique identifier for the user.
	Username string `json:"username" example:"Israel Schroeder"`

	// Email is the user's email address.
	Email string `json:"email" example:"israel@gmail.com"`

	// Role defines the user's role in the system.
	Role string `json:"role" example:"student"`

	// Address contains the user's address details.
	Address AddressResponseDTO `json:"address" validate:"required"`
}
