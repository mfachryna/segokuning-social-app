package entity

import "time"

type User struct {
	ID        string
	Email     string
	Phone     string
	Name      string
	Password  string
	CreatedAt time.Time
}

type UserLoginData struct {
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}
