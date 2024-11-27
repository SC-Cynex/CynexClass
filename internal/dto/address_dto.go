package dto

type AddressRequestDTO struct {
	Street string `json:"street" validate:"required,max=100" example:"Rua das Flores"`
	Number int `json:"number" validate:"required,min=1" example:"72"`
	Neighborhood string `json:"neighborhood" validate:"required,max=100" example:"Centro"`
	City string `json:"city" validate:"required,max=100" example:"Joinville"`
	State string `json:"state" validate:"required,len=2" example:"SC"`
	ZipCode string `json:"zip_code" validate:"required,len=8" example:"89201304"`
}

type AddressResponseDTO struct {
	ID int `json:"id" example:"1"`
	Street string `json:"street" example:"Rua das Flores"`
	Number int `json:"number" example:"72"`
	Neighborhood string `json:"neighborhood" example:"Centro"`
	City string `json:"city" example:"Joinville"`
	State string `json:"state" example:"SC"`
	ZipCode string `json:"zip_code" example:"89201304"`
}
