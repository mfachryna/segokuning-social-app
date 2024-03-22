package dto

type PostFilter struct {
	Limit     int64    `json:"limit" validate:"omitempty,numeric,min=0" schema:"limit"`
	Offset    int64    `json:"offset" validate:"omitempty,numeric,min=0" schema:"offset"`
	Search    string   `json:"search" validate:"omitempty,min=1" schema:"search"`
	SearchTag []string `json:"searchTag" validate:"omitempty,min=0,dive,min=1" schema:"searchTag"`
}
