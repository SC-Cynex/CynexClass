package response

import (
	"net/http"

	"github.com/SC-Cynex/cynex-class-service/internal/dto"
)

// NewSuccessResponse creates a Response indicating a successful operation with HTTP 200 status.
// Parameters:
//   - data: The data to be included in the response.
func NewSuccessResponse(data interface{}) dto.Response {
	return dto.Response{
		Success: true,
		Status:  http.StatusOK,
		Data:    data,
		Title:   http.StatusText(http.StatusOK),
	}
}

// NewCreatedResponse creates a Response indicating a resource was successfully created with HTTP 201 status.
// Parameters:
//   - data: The data to be included in the response.
func NewCreatedResponse(data interface{}) dto.Response {
	return dto.Response{
		Success: true,
		Status:  http.StatusCreated,
		Data:    data,
		Title:   http.StatusText(http.StatusCreated),
	}
}

// NewErrorResponse creates a Response indicating an error with a custom status, data, and title.
// Parameters:
//   - status: The HTTP status code for the error.
//   - data: The data to be included in the response.
//   - title: A short description of the error.
func NewErrorResponse(status int, data interface{}, title string) dto.Response {
	return dto.Response{
		Success: false,
		Status:  status,
		Data:    data,
		Title:   title,
	}
}

// NewBadRequestResponse creates a Response for a bad request error with HTTP 400 status.
// Parameters:
//   - data: The data to be included in the response.
func NewBadRequestResponse(data interface{}) dto.Response {
	return NewErrorResponse(http.StatusBadRequest, data, "Bad Request")
}

// NewUnprocessableEntityResponse creates a Response for an unprocessable entity error with HTTP 422 status.
// Parameters:
//   - data: The data to be included in the response.
func NewUnprocessableEntityResponse(data interface{}) dto.Response {
	return NewErrorResponse(http.StatusUnprocessableEntity, data, "Unprocessable Entity")
}

// NewInternalServerErrorResponse creates a Response for an internal server error with HTTP 500 status.
func NewInternalServerErrorResponse() dto.Response {
	return NewErrorResponse(http.StatusInternalServerError, nil, "Internal Server Error")
}
