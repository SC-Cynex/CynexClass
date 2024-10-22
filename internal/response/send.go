package response

import (
	"encoding/json"
	"net/http"

	"github.com/SC-Cynex/cynex-class-service/internal/dto"
)

func Send(w http.ResponseWriter, resp dto.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Status)
	json.NewEncoder(w).Encode(resp)
}
