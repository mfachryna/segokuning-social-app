package dto

type CommentCreate struct {
	Comment string `json:"comment" validate:"required,min=2,max=500"`
	PostId  string `json:"postId" validate:"required"`
}
