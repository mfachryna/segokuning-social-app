package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/shafaalafghany/segokuning-social-app/internal/interfaces"
)

type UserHandler struct {
	ur interfaces.UserRepository
}

func NewUserHandler(r chi.Router, ur interfaces.UserRepository) {
	uh := &UserHandler{ur}

	r.Post("/", uh.Register)
}
