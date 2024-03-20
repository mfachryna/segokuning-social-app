package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/shafaalafghany/segokuning-social-app/config"
	interfaces "github.com/shafaalafghany/segokuning-social-app/internal/interfaces"
	"github.com/shafaalafghany/segokuning-social-app/pkg/jwt"
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

	r.Route("/user", func(r chi.Router) {
		r.Post("/register", uh.Register)
		r.Post("/login", uh.Login)

		r.Route("/link", func(r chi.Router) {
			r.Use(jwt.JwtMiddleware)
			r.Post("/phone", uh.LinkPhone)
			r.Post("/email", uh.LinkEmail)
		})
	})
}
