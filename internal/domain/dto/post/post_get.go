package dto

import (
	dtocomment "github.com/shafaalafghany/segokuning-social-app/internal/domain/dto/comment"
	dtometa "github.com/shafaalafghany/segokuning-social-app/internal/domain/dto/meta"
	"github.com/shafaalafghany/segokuning-social-app/internal/entity"
)

type PostData struct {
	Message string       `json:"message"`
	Data    []Post       `json:"data"`
	Meta    dtometa.Meta `json:"meta"`
}

type Post struct {
	ID       string               `json:"postId"`
	Post     entity.Post          `json:"post"`
	Comments []dtocomment.Comment `json:"commments"`
	Creator  entity.User          `json:"creator"`
}
