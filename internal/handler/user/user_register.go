package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/shafaalafghany/segokuning-social-app/internal/common/response"
	"github.com/shafaalafghany/segokuning-social-app/internal/common/utils/validation"
	dto "github.com/shafaalafghany/segokuning-social-app/internal/domain/dto/user"
	"github.com/shafaalafghany/segokuning-social-app/internal/entity"
	"github.com/shafaalafghany/segokuning-social-app/pkg/jwt"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func (uh *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var (
		data    dto.UserCreate
		resData interface{}
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

	ctx := r.Context()
	uuid := uuid.NewString()
	credType := data.CredentialType
	user := entity.User{
		ID:   uuid,
		Name: data.Name,
	}

	tokenString, err := jwt.SignedToken(jwt.Claim{UserId: uuid})
	if err != nil {
		uh.log.Error("failed to sign token", zap.Error(err))
		(&response.Response{
			HttpStatus: http.StatusInternalServerError,
			Message:    err.Error(),
		}).GenerateResponse(w)
		return
	}

	if credType == "phone" {
		if err := validation.PhoneValidation(data.CredentialValue); err != nil {
			uh.log.Error("failed to validate phone credential", zap.Error(err))
			(&response.Response{
				HttpStatus: http.StatusBadRequest,
				Message:    err.Error(),
			}).GenerateResponse(w)
			return
		}

		count, err := uh.ur.PhoneCheck(ctx, data.CredentialValue)
		if err != nil {
			uh.log.Error("failed to get phone", zap.Error(err))
			(&response.Response{
				HttpStatus: http.StatusInternalServerError,
				Message:    err.Error(),
			}).GenerateResponse(w)
			return
		}

		if count > 0 {
			uh.log.Error("phone is already used")
			(&response.Response{
				HttpStatus: http.StatusConflict,
				Message:    "phone is already used",
			}).GenerateResponse(w)
			return
		}

		user.Phone = data.CredentialValue

		resData = dto.PhoneData{
			Phone:       user.Phone,
			Name:        user.Name,
			AccessToken: tokenString,
		}
	} else {
		if err := validation.EmailValidation(data.CredentialValue); err != nil {
			uh.log.Error("failed to validate email credential", zap.Error(err))
			(&response.Response{
				HttpStatus: http.StatusBadRequest,
				Message:    err.Error(),
			}).GenerateResponse(w)
			return
		}

		count, err := uh.ur.EmailCheck(ctx, data.CredentialValue)
		if err != nil {
			uh.log.Error("failed to get email", zap.Error(err))
			(&response.Response{
				HttpStatus: http.StatusInternalServerError,
				Message:    err.Error(),
			}).GenerateResponse(w)
			return
		}

		if count > 0 {
			uh.log.Error("email is already used")
			(&response.Response{
				HttpStatus: http.StatusConflict,
				Message:    "email is already used",
			}).GenerateResponse(w)
			return
		}

		user.Email = data.CredentialValue
		resData = dto.EmailData{
			Email:       user.Email,
			Name:        user.Name,
			AccessToken: tokenString,
		}
	}

	salt, err := strconv.Atoi(uh.cfg.App.BcryptSalt)
	if err != nil {
		uh.log.Error("failed to convert string salt", zap.Error(err))
		(&response.Response{
			HttpStatus: http.StatusInternalServerError,
			Message:    err.Error(),
		}).GenerateResponse(w)
		return
	}

	hashedPasswordChan := make(chan string)
	go func() {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), salt)
		if err != nil {
			uh.log.Error("failed to hash password", zap.Error(err))
			(&response.Response{
				HttpStatus: http.StatusConflict,
				Message:    "email is already used",
			}).GenerateResponse(w)
			return
		}
		hashedPasswordChan <- string(hashedPassword)
	}()

	user.Password = <-hashedPasswordChan

	if err := uh.ur.Insert(ctx, user, credType); err != nil {
		uh.log.Error("failed to insert data", zap.Error(err))
		(&response.Response{
			HttpStatus: http.StatusInternalServerError,
			Message:    err.Error(),
		}).GenerateResponse(w)
		return
	}

	(&response.Response{
		HttpStatus: http.StatusCreated,
		Message:    "User registered successfully",
		Data:       resData,
	}).GenerateResponse(w)
}
