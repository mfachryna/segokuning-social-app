package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/shafaalafghany/segokuning-social-app/internal/common/response"
	"github.com/shafaalafghany/segokuning-social-app/internal/common/utils/validation"
	dto "github.com/shafaalafghany/segokuning-social-app/internal/domain/dto/comment"
	"github.com/shafaalafghany/segokuning-social-app/internal/entity"
)

func (uh *CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	var (
		userId string
		data   dto.CommentCreate
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

	if err := validation.UuidValidation(data.PostId); err != nil {
		(&response.Response{
			HttpStatus: http.StatusBadRequest,
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
			(&response.Response{
				HttpStatus: http.StatusNotFound,
				Message:    "Post not found",
			}).GenerateResponse(w)
			return
		}

		(&response.Response{
			HttpStatus: http.StatusInternalServerError,
			Message:    err.Error(),
		}).GenerateResponse(w)
		return
	}

	fmt.Println(userId, post.UserId)

	count, err := uh.fr.FindByRelation(ctx, userId, post.UserId)
	if err != nil {
		(&response.Response{
			HttpStatus: http.StatusInternalServerError,
			Message:    err.Error(),
		}).GenerateResponse(w)
		return
	}

	if count <= 0 {
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
		fmt.Println(err.Error())
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
