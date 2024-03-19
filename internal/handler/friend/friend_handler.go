package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/shafaalafghany/segokuning-social-app/config"
	interfaces "github.com/shafaalafghany/segokuning-social-app/internal/interfaces"
	"github.com/shafaalafghany/segokuning-social-app/pkg/jwt"
)

type FriendHandler struct {
	ur  interfaces.UserRepository
	fr  interfaces.FriendRepository
	val *validator.Validate
	cfg config.Configuration
}

func NewFriendHandler(r chi.Router, ur interfaces.UserRepository, fr interfaces.FriendRepository, val *validator.Validate, cfg config.Configuration) {
	fh := &FriendHandler{
		ur:  ur,
		fr:  fr,
		val: val,
		cfg: cfg,
	}

	r.Route("/friend", func(r chi.Router) {
		r.Use(jwt.JwtMiddleware)
		r.Post("/", fh.CreateFriend)
	})
}
