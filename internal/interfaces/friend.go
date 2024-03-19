package interfaces

import (
	"context"
)

// Translation -.
type (
	FriendRepository interface {
		FindByRelation(context.Context, string, string) (int, error)
		Insert(context.Context, string, string) error
		Delete(context.Context, string, string) error
	}
)
