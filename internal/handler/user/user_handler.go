package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/shafaalafghany/segokuning-social-app/config"
	interfaces "github.com/shafaalafghany/segokuning-social-app/internal/interfaces"
)

type UserHandler struct {
	ur  interfaces.UserRepository
	val *validator.Validate
	cfg config.Configuration
}

func NewUserHandler(r chi.Router, ur interfaces.UserRepository, val *validator.Validate, cfg config.Configuration) {
	uh := &UserHandler{
		ur:  ur,
		val: val,
		cfg: cfg,
	}

	r.Post("/user/register", uh.Register)
	r.Post("/user/login", uh.Login)
}
