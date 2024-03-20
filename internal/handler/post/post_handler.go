package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/shafaalafghany/segokuning-social-app/config"
	interfaces "github.com/shafaalafghany/segokuning-social-app/internal/interfaces"
	"github.com/shafaalafghany/segokuning-social-app/pkg/jwt"
)

type PostHandler struct {
	ur  interfaces.UserRepository
	pr  interfaces.PostRepository
	val *validator.Validate
	cfg config.Configuration
}

func NewPostHandler(r chi.Router, ur interfaces.UserRepository, pr interfaces.PostRepository, val *validator.Validate, cfg config.Configuration) {
	fh := &PostHandler{
		ur:  ur,
		pr:  pr,
		val: val,
		cfg: cfg,
	}

	r.Route("/post", func(r chi.Router) {
		r.Use(jwt.JwtMiddleware)
		r.Get("/", fh.GetPost)
		r.Post("/", fh.CreatePost)
	})
}
