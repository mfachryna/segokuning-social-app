package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/shafaalafghany/segokuning-social-app/config"
	interfaces "github.com/shafaalafghany/segokuning-social-app/internal/interfaces"
	"github.com/shafaalafghany/segokuning-social-app/pkg/jwt"
)

type CommentHandler struct {
	fr  interfaces.FriendRepository
	cr  interfaces.CommentRepository
	pr  interfaces.PostRepository
	val *validator.Validate
	cfg config.Configuration
}

func NewCommentHandler(r chi.Router, fr interfaces.FriendRepository, cr interfaces.CommentRepository, pr interfaces.PostRepository, val *validator.Validate, cfg config.Configuration) {
	fh := &CommentHandler{
		fr:  fr,
		cr:  cr,
		pr:  pr,
		val: val,
		cfg: cfg,
	}

	r.Route("/post/comment", func(r chi.Router) {
		r.Use(jwt.JwtMiddleware)
		r.Post("/", fh.CreateComment)
	})
}
