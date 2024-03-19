package entity

import "time"

type User struct {
	ID          string    `json:"userId"`
	Email       string    `json:"-"`
	Phone       string    `json:"-"`
	Name        string    `json:"name"`
	ImageUrl    string    `json:"imageUrl"`
	Password    string    `json:"-"`
	FriendCount int64     `json:"friendCount"`
	CreatedAt   time.Time `json:"createdAt"`
}

type UserLoginData struct {
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}
