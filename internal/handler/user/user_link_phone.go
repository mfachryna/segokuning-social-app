package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	"github.com/shafaalafghany/segokuning-social-app/internal/common/response"
	"github.com/shafaalafghany/segokuning-social-app/internal/common/utils/validation"
	dto "github.com/shafaalafghany/segokuning-social-app/internal/domain/dto/user"
	"github.com/shafaalafghany/segokuning-social-app/internal/entity"
)

func (uh *UserHandler) LinkPhone(w http.ResponseWriter, r *http.Request) {
	var (
		userId string
		data   dto.UserPhone
	)

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		(&response.Response{
			HttpStatus: http.StatusBadRequest,
			Message:    "required fields are missing or invalid",
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

	if err := validation.PhoneValidation(data.Phone); err != nil {
		(&response.Response{
			HttpStatus: http.StatusBadRequest,
			Message:    err.Error(),
		}).GenerateResponse(w)
		return
	}

	ctx := r.Context()
	userId = ctx.Value("user_id").(string)
	user := entity.User{}
	user.Phone = data.Phone

	resultPhone, err := uh.ur.FindByPhone(ctx, user.Phone)
	if err != nil {
		if err != pgx.ErrNoRows {
			(&response.Response{
				HttpStatus: http.StatusInternalServerError,
				Message:    err.Error(),
			}).GenerateResponse(w)
			return
		}
	}

	if resultPhone != nil {
		(&response.Response{
			HttpStatus: http.StatusConflict,
			Message:    "phone number already existed",
		}).GenerateResponse(w)
		return
	}

	resUser, err := uh.ur.FindById(ctx, userId)
	if err != nil {
		if err == pgx.ErrNoRows {
			(&response.Response{
				HttpStatus: http.StatusNotFound,
				Message:    "User not found",
			}).GenerateResponse(w)
			return
		}

		(&response.Response{
			HttpStatus: http.StatusInternalServerError,
			Message:    err.Error(),
		}).GenerateResponse(w)
		return
	}

	resUser.Phone = data.Phone

	if err := uh.ur.Update(ctx, *resUser); err != nil {
		(&response.Response{
			HttpStatus: http.StatusInternalServerError,
			Message:    err.Error(),
		}).GenerateResponse(w)
		return
	}

	(&response.Response{
		HttpStatus: http.StatusOK,
		Message:    "successfully link your phone to phone number",
	}).GenerateResponse(w)
}
