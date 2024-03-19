package dto

import "github.com/shafaalafghany/segokuning-social-app/internal/entity"

type Friend struct {
	Data []entity.User `json:"data"`
}

type FriendData struct {
	UserId string `json:"userId" validate:"required"`
}
