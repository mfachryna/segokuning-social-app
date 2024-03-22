package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/shafaalafghany/segokuning-social-app/config"
	interfaces "github.com/shafaalafghany/segokuning-social-app/internal/interfaces"
	"github.com/shafaalafghany/segokuning-social-app/pkg/jwt"
	"go.uber.org/zap"
)

type UserHandler struct {
	ur  interfaces.UserRepository
	val *validator.Validate
	cfg config.Configuration
	log *zap.Logger
}

func NewUserHandler(
	r chi.Router,
	ur interfaces.UserRepository,
	val *validator.Validate,
	cfg config.Configuration,
	log *zap.Logger,
) {
	uh := &UserHandler{
		ur:  ur,
		val: val,
		cfg: cfg,
		log: log,
	}

	r.Route("/user", func(r chi.Router) {
		r.Post("/register", uh.Register)
		r.Post("/login", uh.Login)

		r.Route("/", func(r chi.Router) {
			r.Use(jwt.JwtMiddleware)
			r.Patch("/", uh.Update)
		})

		r.Route("/link", func(r chi.Router) {
			r.Use(jwt.JwtMiddleware)
			r.Post("/phone", uh.LinkPhone)
			r.Post("/", uh.LinkEmail)
		})
	})
}
