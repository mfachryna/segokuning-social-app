package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/shafaalafghany/segokuning-social-app/internal/common/response"
	"github.com/shafaalafghany/segokuning-social-app/internal/common/utils/validation"
	dto "github.com/shafaalafghany/segokuning-social-app/internal/domain/dto/post"
	"github.com/shafaalafghany/segokuning-social-app/internal/entity"
)

func (uh *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var (
		userId string
		data   dto.PostCreate
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

	ctx := r.Context()
	userId = ctx.Value("user_id").(string)
	postId := uuid.NewString()

	postEntity := entity.Post{
		ID:         postId,
		PostInHtml: data.PostInHtml,
		Tags:       data.Tags,
	}
	if err := uh.pr.Insert(ctx, postEntity, userId); err != nil {
		fmt.Println(err.Error())
		(&response.Response{
			HttpStatus: http.StatusInternalServerError,
			Message:    err.Error(),
		}).GenerateResponse(w)
		return
	}

	(&response.Response{
		HttpStatus: http.StatusOK,
		Message:    "Add post success",
	}).GenerateResponse(w)

}
