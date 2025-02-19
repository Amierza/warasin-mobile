package dto

import (
	"errors"

	"github.com/google/uuid"
)

const (
	// failed
	MESSAGE_FAILED_GET_DATA_FROM_BODY = "failed get data from body"
	MESSAGE_FAILED_REGISTER_USER      = "failed register user"

	// success
	MESSAGE_SUCCESS_GET_DATA_FROM_BODY = "success get data from body"
	MESSAGE_SUCCESS_REGISTER_USER      = "success register user"
)

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrRegisterUser       = errors.New("failed to register user")
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
)
