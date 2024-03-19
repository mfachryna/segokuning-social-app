package dto

import (
	dto "github.com/shafaalafghany/segokuning-social-app/internal/domain/dto/meta"
	"github.com/shafaalafghany/segokuning-social-app/internal/entity"
)

type Users struct {
	Message string        `json:"message"`
	Data    []entity.User `json:"data"`
	Meta    dto.Meta      `json:"meta"`
}
