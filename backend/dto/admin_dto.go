package dto

import (
	"errors"

	"github.com/Amierza/warasin-mobile/backend/entity"
	"github.com/google/uuid"
)

const (
	// failed
	MESSAGE_FAILED_CREATE_USER = "failed create user"
	MESSAGE_FAILED_CREATE_NEWS = "failed create news"
	MESSAGE_FAILED_DELETE_USER = "failed delete user"
	// success
	MESSAGE_SUCCESS_CREATE_USER = "success create user"
	MESSAGE_SUCCESS_CREATE_NEWS = "success create news"
	MESSAGE_SUCCESS_DELETE_USER = "success delete user"
)

var (
	ErrDeleteUserByID  = errors.New("failed delete user by id")
	ErrFormatBirthdate = errors.New("failed parse birthdate input")
	ErrCreateNews      = errors.New("failed create news")
)

type (
	AdminLoginRequest struct {
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}

	AdminLoginResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	UserPaginationResponse struct {
		PaginationResponse
		Data []AllUserResponse `json:"data"`
	}

	AllUserRepositoryResponse struct {
		PaginationResponse
		Users []entity.User
	}

	CreateUserRequest struct {
		Name        string     `json:"name"`
		Email       string     `json:"email"`
		Password    string     `json:"password"`
		Gender      bool       `json:"gender"`
		Birthdate   string     `json:"birth_date"`
		PhoneNumber string     `json:"phone_number"`
		CityID      *uuid.UUID `gorm:"type:uuid" json:"city_id"`
		RoleID      *uuid.UUID `gorm:"type:uuid" json:"role_id"`
	}

	DeleteUserRequest struct {
		UserID uuid.UUID `json:"user_id"`
	}

	CreateNewsRequest struct {
		Image string `json:"image"`
		Title string `json:"title"`
		Body  string `json:"body"`
		Date  string `gorm:"type:date" json:"date"`
	}

	NewsResponse struct {
		ID    uuid.UUID `json:"news_id"`
		Image string    `json:"news_image"`
		Title string    `json:"news_title"`
		Body  string    `json:"news_body"`
		Date  string    `gorm:"type:date" json:"news_date"`
	}
)
