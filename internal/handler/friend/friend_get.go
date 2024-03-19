package handler

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
	"github.com/shafaalafghany/segokuning-social-app/internal/common/response"
	"github.com/shafaalafghany/segokuning-social-app/internal/common/utils/validation"
	metadto "github.com/shafaalafghany/segokuning-social-app/internal/domain/dto/meta"
	userdto "github.com/shafaalafghany/segokuning-social-app/internal/domain/dto/user"
)

func (uh *FriendHandler) GetFriend(w http.ResponseWriter, r *http.Request) {
	var (
		userId string
		filter userdto.UserFilter
	)
	if err := r.ParseForm(); err != nil {
		(&response.Response{
			HttpStatus: http.StatusInternalServerError,
			Message:    err.Error(),
		}).GenerateResponse(w)
		return
	}

	if err := schema.NewDecoder().Decode(&filter, r.Form); err != nil {
		fmt.Println(err.Error())
		(&response.Response{
			HttpStatus: http.StatusBadRequest,
			Message:    err.Error(),
		}).GenerateResponse(w)
		return
	}

	if err := uh.val.Struct(filter); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, e := range validationErrors {
			fmt.Println(err.Error())
			(&response.Response{
				HttpStatus: http.StatusBadRequest,
				Message:    validation.CustomError(e),
			}).GenerateResponse(w)
			return
		}
	}

	ctx := r.Context()
	userId = ctx.Value("user_id").(string)

	*filter.Offset = *filter.Limit * (*filter.Offset)
	data, count, err := uh.ur.GetUserWithFilter(ctx, userId, filter)
	if err != nil {
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
			Limit:  *filter.Limit,
			Offset: *filter.Offset,
			Total:  count,
		},
	}).GenerateResponseMeta(w)
}
