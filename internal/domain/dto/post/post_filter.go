package dto

type PostFilter struct {
	Limit     *int64   `json:"limit" validate:"required,numeric" schema:"limit"`
	Offset    *int64   `json:"offset" validate:"required,numeric" schema:"offset"`
	Search    string   `json:"search" validate:"omitempty,min=3" schema:"search"`
	SearchTag []string `json:"searchTag" validate:"omitempty,min=0,dive,min=0" schema:"searchTag"`
}
