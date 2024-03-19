package dto

type FriendData struct {
	UserId string `json:"userId" validate:"required"`
}
