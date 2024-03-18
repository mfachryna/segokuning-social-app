package dto

type UserCreate struct {
	CredentialType  string `json:"credentialType" validate:"required,eq=email|eq=phone"`
	CredentialValue string `json:"credentialValue" validate:"required"`
	Name            string `json:"name" validate:"required,min=5,max=50"`
	password        string `json:"password" validate:"required,min=5,max=15"`
}
