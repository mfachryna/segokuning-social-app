package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	"github.com/shafaalafghany/segokuning-social-app/internal/common/response"
	"github.com/shafaalafghany/segokuning-social-app/internal/common/utils/validation"
	dto "github.com/shafaalafghany/segokuning-social-app/internal/domain/dto/user"
)

func (uh *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var (
		userId string
		data   dto.UserUpdate
	)

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		(&response.Response{
			HttpStatus: http.StatusBadRequest,
			Message:    "required fields are missing or invalid",
		}).GenerateResponse(w)
		return
	}
	if err := validation.UrlValidation(data.ImageUrl); err != nil {
		(&response.Response{
			HttpStatus: http.StatusBadRequest,
			Message:    "URL malformed",
		}).GenerateResponse(w)

		return
	}

	if err := uh.val.Struct(data); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, e := range validationErrors {
			(&response.Response{
				HttpStatus: http.StatusBadRequest,
				Message:    validation.CustomError(e),
			}).GenerateResponse(w)
			return
		}
	}

	if err := validation.UrlValidation(data.ImageUrl); err != nil {
		(&response.Response{
			HttpStatus: http.StatusBadRequest,
			Message:    err.Error(),
		}).GenerateResponse(w)
		return
	}

	ctx := r.Context()
	userId = ctx.Value("user_id").(string)

	result, err := uh.ur.FindById(ctx, userId)
	if err != nil {
		if err == pgx.ErrNoRows {
			(&response.Response{
				HttpStatus: http.StatusNotFound,
				Message:    "user not found",
			}).GenerateResponse(w)
			return
		}

		(&response.Response{
			HttpStatus: http.StatusInternalServerError,
			Message:    err.Error(),
		}).GenerateResponse(w)
		return
	}

	result.Name = data.Name
	result.ImageUrl = data.ImageUrl

	if err := uh.ur.Update(ctx, *result); err != nil {
		(&response.Response{
			HttpStatus: http.StatusInternalServerError,
			Message:    err.Error(),
		}).GenerateResponse(w)
		return
	}

	(&response.Response{
		HttpStatus: http.StatusOK,
		Message:    "successfully update user profile",
	}).GenerateResponse(w)
}
