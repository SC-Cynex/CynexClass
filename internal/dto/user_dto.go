package dto

type UserRegistrationRequestDTO struct {
	Username string            `json:"username" validate:"required,min=3,max=50"`
	Email    string            `json:"email" validate:"required,email"`
	Role     string            `json:"role" validate:"required,oneof=teacher student"`
	Password string            `json:"password,omitempty" validate:"required,min=6"`
	Address  AddressRequestDTO `json:"address"`
}
