package dto

import (
	"github.com/shafaalafghany/segokuning-social-app/internal/entity"
)

type Comment struct {
	Comment   string      `json:"comment"`
	Creator   entity.User `json:"creator"`
	CreatedAt string      `json:"createdAt"`
}
