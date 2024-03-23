package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
	"github.com/shafaalafghany/segokuning-social-app/internal/common/response"
	"github.com/shafaalafghany/segokuning-social-app/internal/common/utils/validation"
	dto "github.com/shafaalafghany/segokuning-social-app/internal/domain/dto/friend"
	"go.uber.org/zap"
)

func (uh *FriendHandler) CreateFriend(w http.ResponseWriter, r *http.Request) {
	var (
		userId string
		data   dto.FriendData
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

	ctx := r.Context()
	userId = ctx.Value("user_id").(string)
	friendId := data.UserId

	if userId == friendId {
		uh.log.Info("cannot add self as friend")
		(&response.Response{
			HttpStatus: http.StatusBadRequest,
			Message:    "Cannot add self as friend",
		}).GenerateResponse(w)
		return
	}

	if err := validation.UuidValidation(friendId); err != nil {
		uh.log.Info("failed to validate uuid", zap.Error(err))
		(&response.Response{
			HttpStatus: http.StatusNotFound,
			Message:    err.Error(),
		}).GenerateResponse(w)
		return
	}

	_, err := uh.ur.FindById(ctx, friendId)
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

	count, err := uh.fr.FindByRelation(ctx, userId, friendId)
	if err != nil {
		uh.log.Info("failed to get user relation", zap.Error(err))
		(&response.Response{
			HttpStatus: http.StatusInternalServerError,
			Message:    err.Error(),
		}).GenerateResponse(w)
		return
	}

	if count > 0 {
		uh.log.Info("you already add this user as friend")
		(&response.Response{
			HttpStatus: http.StatusInternalServerError,
			Message:    "You are already add this user as friend",
		}).GenerateResponse(w)
		return
	}

	if err := uh.fr.Insert(ctx, userId, friendId); err != nil {
		uh.log.Info("failed to add user", zap.Error(err))
		(&response.Response{
			HttpStatus: http.StatusInternalServerError,
			Message:    err.Error(),
		}).GenerateResponse(w)
		return
	}

	(&response.Response{
		HttpStatus: http.StatusOK,
		Message:    "Add friend success",
	}).GenerateResponse(w)
}
