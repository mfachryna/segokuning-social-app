package dto

type UserEmail struct {
	Email string `json:"email" validate:"required"`
}
