package dto

type UserUpdate struct {
	Name     string `json:"name" validate:"required,min=5,max=50"`
	ImageUrl string `json:"imageUrl" validate:"required,url"`
}
