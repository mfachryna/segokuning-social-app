package dto

type UserFilter struct {
	Limit      *int64 `json:"limit" validate:"required,numeric" schema:"limit"`
	Offset     *int64 `json:"offset" validate:"required,numeric" schema:"offset"`
	OnlyFriend bool   `json:"onlyFriend" schema:"onlyFriend"`
	SortBy     string `json:"sortBy" validate:"omitempty,eq=friendCount|eq=createdAt" schema:"sortBy"`
	OrderBy    string `json:"orderBy" validate:"omitempty,eq=asc|eq=desc" schema:"orderBy"`
	Search     string `json:"search" validate:"omitempty,min=3" schema:"search"`
}
