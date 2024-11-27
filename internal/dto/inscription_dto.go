package dto

type InscriptionRequestDTO struct {
    UserID  int `json:"user_id" validate:"required"`
    ClassID int `json:"class_id" validate:"required"`
}
