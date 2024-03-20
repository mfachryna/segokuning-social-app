package interfaces

import (
	"context"

	dto "github.com/shafaalafghany/segokuning-social-app/internal/domain/dto/post"
	"github.com/shafaalafghany/segokuning-social-app/internal/entity"
)

// Translation -.
type (
	PostRepository interface {
		Insert(context.Context, entity.Post, string) error
		GetPostWithFilter(context.Context, dto.PostFilter) ([]dto.Post, int64, error)
		FindById(context.Context, string) (entity.Post, error)
	}
)
