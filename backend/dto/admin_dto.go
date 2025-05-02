package dto

import (
	"errors"

	"github.com/Amierza/warasin-mobile/backend/entity"
	"github.com/google/uuid"
)

const (
// failed

// success

)

var (
	ErrDeleteUserByID  = errors.New("failed delete user by id")
	ErrFormatBirthdate = errors.New("failed parse birthdate input")
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
		ID          uuid.UUID  `json:"id"`
		Name        string     `json:"name"`
		Email       string     `json:"email"`
		Password    string     `json:"password"`
		Gender      bool       `json:"gender"`
		Birthdate   string     `json:"birth_date"`
		PhoneNumber string     `json:"phone_number"`
		CityID      *uuid.UUID `gorm:"type:uuid" json:"city_id"`
		RoleID      *uuid.UUID `gorm:"type:uuid" json:"role_id"`
	}
)
