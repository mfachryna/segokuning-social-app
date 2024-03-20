package interfaces

import (
	"context"

	"github.com/shafaalafghany/segokuning-social-app/internal/entity"
)

// Translation -.
type (
	CommentRepository interface {
		Insert(context.Context, entity.Comment) error
	}
)
