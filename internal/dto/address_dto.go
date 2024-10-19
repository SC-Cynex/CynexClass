package dto

type AddressRequestDTO struct {
	Street       string `json:"street" validate:"required,max=100"`
	Number       int    `json:"number" validate:"required,min=1"`
	Neighborhood string `json:"neighborhood" validate:"required,max=100"`
	City         string `json:"city" validate:"required,max=100"`
	State        string `json:"state" validate:"required,len=2"`
	ZipCode      string `json:"zip_code" validate:"required,len=8"`
}
