package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	interfaces "github.com/shafaalafghany/segokuning-social-app/internal/interfaces"
)

type UserHandler struct {
	ur        interfaces.UserRepository
	validator *validator.Validate
}

func NewUserHandler(r chi.Router, ur interfaces.UserRepository, val *validator.Validate) {
	uh := &UserHandler{ur, val}

	r.Post("/", uh.Register)
}
