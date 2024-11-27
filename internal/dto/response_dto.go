package dto

type Response struct {
	Success bool        `json:"success"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Title   string      `json:"title"`
}

type CreatedResponse struct {
	Success bool   `json:"success" example:"true"`
	Status  int    `json:"status" example:"201"`
	Data    any    `json:"data"`
	Title   string `json:"title" example:"Created"`
}

type BadRequestResponse struct {
	Success bool   `json:"success" example:"false"`
	Status  int    `json:"status" example:"400"`
	Data    string `json:"data" example:"null"`
	Title   string `json:"title" example:"Bad Request"`
}

type UnprocessableEntityResponse struct {
	Success bool   `json:"success" example:"false"`
	Status  int    `json:"status" example:"422"`
	Data    string `json:"data" example:"name is required"`
	Title   string `json:"title" example:"Unprocessable Entity"`
}

type InternalServerErrorResponse struct {
	Success bool   `json:"success" example:"false"`
	Status  int    `json:"status" example:"500"`
	Data    string `json:"data" example:"null"`
	Title   string `json:"title" example:"Internal Server Error"`
}
