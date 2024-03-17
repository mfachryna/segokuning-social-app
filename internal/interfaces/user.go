package interfaces

import (
	"context"

	"github.com/shafaalafghany/segokuning-social-app/internal/entity"
)

// Translation -.
type (
	UserRepository interface {
		Get(context.Context, entity.User) error
		FindById(context.Context, string) error
		Insert(context.Context, entity.User) error
		Delete(context.Context, string) error
		Update(context.Context, entity.User) error
	}
)
