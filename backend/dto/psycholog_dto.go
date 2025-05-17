package dto

import "errors"

const (
	// failed
	MESSAGE_FAILED_LOGIN_PSYCHOLOG = "failed login psycholog"

	// success
	MESSAGE_SUCCESS_LOGIN_PSYCHOLOG = "success login psycholog"
)

var (
	ErrGetPsychologIDFromToken = errors.New("failed get psycholog id from token")
)

type (
	PsychologLoginRequest struct {
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}

	PsychologLoginResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
)
