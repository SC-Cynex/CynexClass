package response

import (
	"net/http"

	"github.com/SC-Cynex/cynex-class-service/internal/dto"
)

func NewSuccessResponse(data interface{}) dto.Response {
	return dto.Response{
		Success: true,
		Status:  http.StatusOK,
		Data:    data,
		Title:   http.StatusText(http.StatusOK),
	}
}

func NewCreatedResponse(data interface{}) dto.Response {
	return dto.Response{
		Success: true,
		Status:  http.StatusCreated,
		Data:    data,
		Title:   http.StatusText(http.StatusCreated),
	}
}

func NewErrorResponse(status int, data interface{}, title string) dto.Response {
	return dto.Response{
		Success: false,
		Status:  status,
		Data:    data,
		Title:   title,
	}
}

func NewBadRequestResponse(data interface{}) dto.Response {
	return NewErrorResponse(http.StatusBadRequest, data, "Bad Request")
}

func NewUnprocessableEntityResponse(data interface{}) dto.Response {
	return NewErrorResponse(http.StatusUnprocessableEntity, data, "Unprocessable Entity")
}

func NewInternalServerErrorResponse() dto.Response {
	return NewErrorResponse(http.StatusInternalServerError, nil, "Internal Server Error")
}
