package handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
	"github.com/shafaalafghany/segokuning-social-app/internal/common/response"
	"github.com/shafaalafghany/segokuning-social-app/internal/common/utils/validation"
	metadto "github.com/shafaalafghany/segokuning-social-app/internal/domain/dto/meta"
	postdto "github.com/shafaalafghany/segokuning-social-app/internal/domain/dto/post"
	"go.uber.org/zap"
)

func (uh *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	var (
		filter postdto.PostFilter
	)

	if err := r.ParseForm(); err != nil {
		uh.log.Error("failed to parse form", zap.Error(err))
		(&response.Response{
			HttpStatus: http.StatusInternalServerError,
			Message:    err.Error(),
		}).GenerateResponse(w)
		return
	}

	if err := validation.ValidateParams(r, filter); err != nil {
		uh.log.Error("failed to validate params", zap.Error(err))
		(&response.Response{
			HttpStatus: http.StatusBadRequest,
			Message:    err.Error(),
		}).GenerateResponse(w)
		return
	}

	if err := schema.NewDecoder().Decode(&filter, r.Form); err != nil {
		uh.log.Error("required fields are missing or invalid", zap.Error(err))
		(&response.Response{
			HttpStatus: http.StatusBadRequest,
			Message:    err.Error(),
		}).GenerateResponse(w)
		return
	}

	if err := uh.val.Struct(filter); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, e := range validationErrors {
			uh.log.Error(validation.CustomError(e), zap.Error(err))
			(&response.Response{
				HttpStatus: http.StatusBadRequest,
				Message:    validation.CustomError(e),
			}).GenerateResponse(w)
			return
		}
	}

	ctx := r.Context()
	userId := ctx.Value("user_id").(string)

	if filter.Limit == 0 {
		filter.Limit = 5
	}
	filter.Offset = filter.Limit * filter.Offset

	data, count, err := uh.pr.GetPostWithFilter(ctx, filter, userId)
	if err != nil {
		uh.log.Error("failed to get post with filter", zap.Error(err))
		(&response.Response{
			HttpStatus: http.StatusInternalServerError,
			Message:    err.Error(),
		}).GenerateResponse(w)
		return
	}

	(&response.ResponseWithMeta{
		HttpStatus: http.StatusOK,
		Data:       data,
		Meta: metadto.Meta{
			Limit:  filter.Limit,
			Offset: filter.Offset,
			Total:  count,
		},
	}).GenerateResponseMeta(w)
}
