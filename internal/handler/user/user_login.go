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
	"github.com/shafaalafghany/segokuning-social-app/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var (
		data     dto.UserLogin
		userData entity.User
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
	credType := data.CredentialType
	user := entity.User{}

	if credType == "phone" {
		if err := validation.PhoneValidation(data.CredentialValue); err != nil {
			(&response.Response{
				HttpStatus: http.StatusBadRequest,
				Message:    err.Error(),
			}).GenerateResponse(w)
			return
		}

		user.Phone = data.CredentialValue

		result, err := uh.ur.FindByPhone(ctx, user.Phone)
		if err != nil {
			if err == pgx.ErrNoRows {
				(&response.Response{
					HttpStatus: http.StatusNotFound,
					Message:    "phone not found",
				}).GenerateResponse(w)
				return
			}

			(&response.Response{
				HttpStatus: http.StatusInternalServerError,
				Message:    err.Error(),
			}).GenerateResponse(w)
			return
		}

		userData = *result
	} else {
		if err := validation.EmailValidation(data.CredentialValue); err != nil {
			(&response.Response{
				HttpStatus: http.StatusBadRequest,
				Message:    err.Error(),
			}).GenerateResponse(w)
			return
		}

		user.Email = data.CredentialValue

		result, err := uh.ur.FindByEmail(ctx, data.CredentialValue)
		if err != nil {
			if err == pgx.ErrNoRows {
				(&response.Response{
					HttpStatus: http.StatusNotFound,
					Message:    "phone not found",
				}).GenerateResponse(w)
				return
			}

			(&response.Response{
				HttpStatus: http.StatusInternalServerError,
				Message:    err.Error(),
			}).GenerateResponse(w)
			return
		}

		userData = *result
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(data.Password)); err != nil {
		(&response.Response{
			HttpStatus: http.StatusBadRequest,
			Message:    "password mismatched",
		}).GenerateResponse(w)
		return
	}

	tokenString, err := jwt.SignedToken(jwt.Claim{UserId: userData.ID})
	if err != nil {
		(&response.Response{
			HttpStatus: http.StatusInternalServerError,
			Message:    err.Error(),
		}).GenerateResponse(w)
		return
	}

	res := &entity.UserLoginData{
		Email:       userData.Email,
		Phone:       userData.Phone,
		Name:        userData.Name,
		AccessToken: tokenString,
	}

	(&response.Response{
		HttpStatus: http.StatusOK,
		Message:    "User logged successfully",
		Data:       res,
	}).GenerateResponse(w)
}
