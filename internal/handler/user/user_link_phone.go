package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	"github.com/shafaalafghany/segokuning-social-app/internal/common/response"
	"github.com/shafaalafghany/segokuning-social-app/internal/common/utils/validation"
	dto "github.com/shafaalafghany/segokuning-social-app/internal/domain/dto/user"
	"go.uber.org/zap"
)

func (uh *UserHandler) LinkPhone(w http.ResponseWriter, r *http.Request) {
	var (
		userId string
		data   dto.UserPhone
	)

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		uh.log.Info("required fields are missing or invalid", zap.Error(err))
		(&response.Response{
			HttpStatus: http.StatusBadRequest,
			Message:    "required fields are missing or invalid",
		}).GenerateResponse(w)
		return
	}

	if err := uh.val.Struct(data); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, e := range validationErrors {
			uh.log.Info(validation.CustomError(e), zap.Error(err))
			(&response.Response{
				HttpStatus: http.StatusBadRequest,
				Message:    validation.CustomError(e),
			}).GenerateResponse(w)
			return
		}
	}

	if err := validation.PhoneValidation(data.Phone); err != nil {
		uh.log.Info("failed to validate phone credential", zap.Error(err))
		(&response.Response{
			HttpStatus: http.StatusBadRequest,
			Message:    err.Error(),
		}).GenerateResponse(w)
		return
	}

	ctx := r.Context()
	userId = ctx.Value("user_id").(string)

	resultPhone, err := uh.ur.FindByPhone(ctx, data.Phone)
	if err != nil && err != pgx.ErrNoRows {
		uh.log.Info("failed to get phone", zap.Error(err))
		(&response.Response{
			HttpStatus: http.StatusInternalServerError,
			Message:    err.Error(),
		}).GenerateResponse(w)
		return
	}

	if resultPhone != nil {
		if resultPhone.ID == userId && resultPhone.Phone != "" {
			uh.log.Info("you already have a phone number")
			(&response.Response{
				HttpStatus: http.StatusBadRequest,
				Message:    "You already have a phone number",
			}).GenerateResponse(w)
			return
		}

		uh.log.Info("phone number already existed")
		(&response.Response{
			HttpStatus: http.StatusConflict,
			Message:    "phone number already existed",
		}).GenerateResponse(w)
		return
	}

	resUser, err := uh.ur.FindById(ctx, userId)
	if err != nil {
		if err == pgx.ErrNoRows {
			uh.log.Info("user is not found", zap.Error(err))
			(&response.Response{
				HttpStatus: http.StatusNotFound,
				Message:    "User not found",
			}).GenerateResponse(w)
			return
		}

		uh.log.Info("failed to get user", zap.Error(err))
		(&response.Response{
			HttpStatus: http.StatusInternalServerError,
			Message:    err.Error(),
		}).GenerateResponse(w)
		return
	}

	if resUser.Phone != "" {
		uh.log.Info("cannot change phone number if you already have one")
		(&response.Response{
			HttpStatus: http.StatusBadRequest,
			Message:    "cannot change phone number if you already have one",
		}).GenerateResponse(w)
		return
	}

	resUser.Phone = data.Phone

	if err := uh.ur.Update(ctx, *resUser); err != nil {
		uh.log.Info("failed to update user", zap.Error(err))
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
