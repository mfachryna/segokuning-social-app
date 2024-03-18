package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/shafaalafghany/segokuning-social-app/internal/common/response"
	"github.com/shafaalafghany/segokuning-social-app/internal/common/utils/validation"
	"golang.org/x/crypto/bcrypt"
)

func (uh *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var registerData domain.UserRegister
	if err := json.NewDecoder(r.Body).Decode(&registerData); err != nil {
		fmt.Println(err.Error())
		(&response.Response{
			HttpStatus: http.StatusBadRequest,
			Message:    response.ValidationAndParseBodyError,
		}).GenerateResponse(w)
		return
	}

	if err := uh.validator.Struct(registerData); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, e := range validationErrors {
			fmt.Println(err.Error())
			(&response.Response{
				HttpStatus: http.StatusBadRequest,
				Message:    validation.CustomError(e),
			}).GenerateResponse(w)
		}
	}

	// registerData.Username = strings.ToLower(registerData.Username)

	var count int
	err := uh.db.QueryRow("SELECT COUNT(*) FROM users WHERE username = $1", registerData.Username).Scan(&count)

	if err != nil {
		fmt.Println(err.Error())
		response.Error(w, apierror.CustomServerError(err.Error()))
		return
	}

	if count > 0 {
		err := apierror.ClientAlreadyExists()
		fmt.Println(err.Message)
		response.Error(w, err)
		return
	}

	date := time.Now()

	hashedPasswordChan := make(chan string)
	go func() {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerData.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println(err.Error())
			response.Error(w, apierror.CustomServerError(err.Error()))
			return
		}
		hashedPasswordChan <- string(hashedPassword)
	}()

	var id string
	uuid := uuid.New()
	if err := uh.db.QueryRow(`INSERT INTO users (id,username,name,password,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id`, uuid, registerData.Username, registerData.Name, <-hashedPasswordChan, date, date).Scan(&id); err != nil {
		fmt.Println(err.Error())
		response.Error(w, apierror.CustomServerError(err.Error()))
		return
	}

	tokenString, err := jwt.SignedToken(jwt.Claim{
		UserId: id,
	})
	if err != nil {
		fmt.Println(err.Error())
		response.Error(w, apierror.CustomServerError("Failed to generate access token"))
		return
	}

	res := &domain.UserAuthResponse{
		Name:        registerData.Name,
		Username:    registerData.Username,
		AccessToken: tokenString,
	}

	response.Success(w, apisuccess.RegisterResponse(res))
}
