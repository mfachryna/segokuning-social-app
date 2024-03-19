package interfaces

import (
	"context"

	dto "github.com/shafaalafghany/segokuning-social-app/internal/domain/dto/friend"
)

// Translation -.
type (
	FriendRepository interface {
		Get(context.Context, string) (dto.Friend, error)
		FindById(context.Context, string, string) error
		FindByRelation(context.Context, string, string) (int, error)
		Insert(context.Context, string, string) error
		Delete(context.Context, string) error
	}
)
