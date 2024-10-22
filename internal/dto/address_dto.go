package dto

// AddressRequestDTO represents the data required to create a new address.
// @Description Address data for creating a new address
type AddressRequestDTO struct {
	// Street is the name of the street for the address.
	// It is a required field with a maximum length of 100 characters.
	Street string `json:"street" validate:"required,max=100" example:"Rua das Flores"`

	// Number is the house or building number.
	// It is a required field with a minimum value of 1.
	Number int `json:"number" validate:"required,min=1" example:"72"`

	// Neighborhood is the name of the neighborhood.
	// It is a required field with a maximum length of 100 characters.
	Neighborhood string `json:"neighborhood" validate:"required,max=100" example:"Centro"`

	// City is the name of the city.
	// It is a required field with a maximum length of 100 characters.
	City string `json:"city" validate:"required,max=100" example:"Joinville"`

	// State is the two-letter state code.
	// It is a required field with a fixed length of 2 characters.
	State string `json:"state" validate:"required,len=2" example:"SC"`

	// ZipCode is the postal code.
	// It is a required field with a fixed length of 8 characters.
	ZipCode string `json:"zip_code" validate:"required,len=8" example:"89201304"`
}

// AddressResponseDTO represents the data returned after creating or retrieving an address.
type AddressResponseDTO struct {
	// ID is the unique identifier for the address.
	ID int `json:"id" example:"1"`

	// Street is the name of the street for the address.
	Street string `json:"street" example:"Rua das Flores"`

	// Number is the house or building number.
	Number int `json:"number" example:"72"`

	// Neighborhood is the name of the neighborhood.
	Neighborhood string `json:"neighborhood" example:"Centro"`

	// City is the name of the city.
	City string `json:"city" example:"Joinville"`

	// State is the two-letter state code.
	State string `json:"state" example:"SC"`

	// ZipCode is the postal code.
	ZipCode string `json:"zip_code" example:"89201304"`
}
