package image

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/shafaalafghany/segokuning-social-app/config"
	"github.com/shafaalafghany/segokuning-social-app/pkg/jwt"
)

type ImageHandler struct {
	val *validator.Validate
	cfg config.Configuration
}

func NewImageHandler(r chi.Router, val validator.Validate, cfg config.Configuration) {
	ih := &ImageHandler{
		val: &val,
		cfg: cfg,
	}

	r.Route("/image", func(r chi.Router) {
		r.Use(jwt.JwtMiddleware)
		r.Post("/", ih.Store)
	})
}
