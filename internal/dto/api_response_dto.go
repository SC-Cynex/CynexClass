package dto

type APIResponse struct {
	Title   string      `json:"title"`
	Success bool        `json:"success"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}
