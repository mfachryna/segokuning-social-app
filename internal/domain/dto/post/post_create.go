package dto

type PostCreate struct {
	PostInHtml string   `json:"postInHtml" validate:"required,min=2,max=500"`
	Tags       []string `json:"tags" validate:"required,dive,min=1"`
}
