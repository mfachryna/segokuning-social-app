package entity

import "time"

type User struct {
	ID          string
	Email       string
	Phone       string
	Name        string `json:"name"`
	ImageUrl    string `json:"imageUrl"`
	Password    string
	FriendCount int64     `json:"friendCount"`
	CreatedAt   time.Time `json:"createdAt"`
}

type UserLoginData struct {
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}
