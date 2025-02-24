package dto

import (
	"errors"

	"github.com/google/uuid"
)

const (
	// failed
	MESSAGE_FAILED_GET_DATA_FROM_BODY  = "failed get data from body"
	MESSAGE_FAILED_REGISTER_USER       = "failed register user"
	MESSAGE_FAILED_LOGIN_USER          = "failed login user"
	MESSAGE_FAILED_PROSES_REQUEST      = "failed proses request"
	MESSAGE_FAILED_TOKEN_NOT_FOUND     = "failed token not found"
	MESSAGE_FAILED_TOKEN_NOT_VALID     = "failed token not valid"
	MESSAGE_FAILED_TOKEN_DENIED_ACCESS = "failed token denied access"

	// success
	MESSAGE_SUCCESS_GET_DATA_FROM_BODY = "success get data from body"
	MESSAGE_SUCCESS_REGISTER_USER      = "success register user"
	MESSAGE_SUCCESS_LOGIN_USER         = "success login user"
)

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrRegisterUser       = errors.New("failed to register user")
	ErrEmailNotFound      = errors.New("email not found")
	ErrPasswordNotMatch   = errors.New("password not match")
)

type (
	UserResponse struct {
		ID       uuid.UUID `json:"user_id"`
		Name     string    `json:"name"`
		Email    string    `json:"email"`
		Password string    `json:"password"`
	}

	UserRegisterRequest struct {
		Name     string `json:"name" form:"name" validate:"required,min=5"`
		Email    string `json:"email" form:"email" validate:"required,email"`
		Password string `json:"password" form:"password" validate:"required,min=8"`
	}

	UserLoginRequest struct {
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}

	UserLoginResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
)
