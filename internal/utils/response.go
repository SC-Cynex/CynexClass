package utils

import (
	"encoding/json"
	"net/http"

	"github.com/SC-Cynex/cynex-class-service/internal/dto"
)

func SendAPIResponse(w http.ResponseWriter, status int, title string, success bool, data interface{}) {
	response := dto.APIResponse{
		Title:   title,
		Success: success,
		Status:  status,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
