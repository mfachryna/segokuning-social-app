package dto

import "time"

type UserCreate struct {
	CredentialType  string    `json:"credentialType" validate:"required,eq=email|eq=phone"`
	CredentialValue string    `json:"credentialValue" validate:"required"`
	Name            string    `json:"name" validate:"required,min=5,max=50"`
	Password        string    `json:"password" validate:"required,min=5,max=15"`
	CreatedAt       time.Time `json:"createdAt"`
}

type EmailData struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}

type PhoneData struct {
	Phone       string `json:"phone"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}
