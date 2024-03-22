package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/shafaalafghany/segokuning-social-app/config"
	interfaces "github.com/shafaalafghany/segokuning-social-app/internal/interfaces"
	"github.com/shafaalafghany/segokuning-social-app/pkg/jwt"
	"go.uber.org/zap"
)

type FriendHandler struct {
	ur  interfaces.UserRepository
	fr  interfaces.FriendRepository
	val *validator.Validate
	cfg config.Configuration
	log *zap.Logger
}

func NewFriendHandler(
	r chi.Router,
	ur interfaces.UserRepository,
	fr interfaces.FriendRepository,
	val *validator.Validate,
	cfg config.Configuration,
	log *zap.Logger,
) {
	fh := &FriendHandler{
		ur:  ur,
		fr:  fr,
		val: val,
		cfg: cfg,
		log: log,
	}

	r.Route("/friend", func(r chi.Router) {
		r.Use(jwt.JwtMiddleware)
		r.Get("/", fh.GetFriend)
		r.Post("/", fh.CreateFriend)
		r.Delete("/", fh.DeleteFriend)
	})
}
