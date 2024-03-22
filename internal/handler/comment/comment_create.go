package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/shafaalafghany/segokuning-social-app/internal/common/response"
	"github.com/shafaalafghany/segokuning-social-app/internal/common/utils/validation"
	dto "github.com/shafaalafghany/segokuning-social-app/internal/domain/dto/comment"
	"github.com/shafaalafghany/segokuning-social-app/internal/entity"
	"go.uber.org/zap"
)

func (uh *CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	var (
		userId string
		data   dto.CommentCreate
	)

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		uh.log.Error("required fields are missing or invalid", zap.Error(err))
		(&response.Response{
			HttpStatus: http.StatusBadRequest,
			Message:    "required fields are missing or invalid",
		}).GenerateResponse(w)
		return
	}

	if err := uh.val.Struct(data); err != nil {
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

	if err := validation.UuidValidation(data.PostId); err != nil {
		uh.log.Error("failed to validate uuid", zap.Error(err))
		(&response.Response{
			HttpStatus: http.StatusNotFound,
			Message:    err.Error(),
		}).GenerateResponse(w)
		return
	}

	ctx := r.Context()
	userId = ctx.Value("user_id").(string)
	commentId := uuid.NewString()

	post, err := uh.pr.FindById(ctx, data.PostId)
	if err != nil {
		if err == pgx.ErrNoRows {
			uh.log.Error("user is not found", zap.Error(err))
			(&response.Response{
				HttpStatus: http.StatusNotFound,
				Message:    "Post not found",
			}).GenerateResponse(w)
			return
		}

		uh.log.Error("failed to get user", zap.Error(err))
		(&response.Response{
			HttpStatus: http.StatusInternalServerError,
			Message:    err.Error(),
		}).GenerateResponse(w)
		return
	}

	count, err := uh.fr.FindByRelation(ctx, userId, post.UserId)
	if err != nil {
		uh.log.Error("failed to get user relation", zap.Error(err))
		(&response.Response{
			HttpStatus: http.StatusInternalServerError,
			Message:    err.Error(),
		}).GenerateResponse(w)
		return
	}

	if count <= 0 {
		uh.log.Error("you cannot commented because you are not friend with this user")
		(&response.Response{
			HttpStatus: http.StatusBadRequest,
			Message:    "You cannot commented because you are not friend with this user",
		}).GenerateResponse(w)
		return
	}

	commentEntity := entity.Comment{
		ID:      commentId,
		Comment: data.Comment,
		UserId:  userId,
		PostId:  data.PostId,
	}

	if err := uh.cr.Insert(ctx, commentEntity); err != nil {
		uh.log.Error("failed to insert data", zap.Error(err))
		(&response.Response{
			HttpStatus: http.StatusInternalServerError,
			Message:    err.Error(),
		}).GenerateResponse(w)
		return
	}

	(&response.Response{
		HttpStatus: http.StatusOK,
		Message:    "Add comment success",
	}).GenerateResponse(w)

}
