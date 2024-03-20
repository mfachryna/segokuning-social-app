package dto

type UserPhone struct {
	Phone string `json:"phone" validate:"required,min=7,max=13"`
}
