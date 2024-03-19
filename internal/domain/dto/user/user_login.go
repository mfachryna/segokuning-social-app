package dto

type UserLogin struct {
	CredentialType  string `json:"credentialType" validate:"required,eq=email|eq=phone"`
	CredentialValue string `json:"credentialValue" validate:"required"`
	Password        string `json:"password" validate:"required,min=5,max=15"`
}
