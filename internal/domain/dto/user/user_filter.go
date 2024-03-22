package dto

type UserFilter struct {
	Limit      int64  `json:"limit" validate:"omitempty,numeric,min=0" schema:"limit"`
	Offset     int64  `json:"offset" validate:"omitempty,numeric,min=0" schema:"offset"`
	OnlyFriend bool   `json:"onlyFriend" validate:"omitempty,boolean" schema:"onlyFriend"`
	SortBy     string `json:"sortBy" validate:"omitempty,eq=friendCount|eq=createdAt" schema:"sortBy"`
	OrderBy    string `json:"orderBy" validate:"omitempty,eq=asc|eq=desc" schema:"orderBy"`
	Search     string `json:"search" validate:"omitempty,min=1" schema:"search"`
}
