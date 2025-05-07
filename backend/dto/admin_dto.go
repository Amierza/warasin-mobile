package dto

import (
	"errors"

	"github.com/Amierza/warasin-mobile/backend/entity"
	"github.com/google/uuid"
)

const (
	// failed
	MESSAGE_FAILED_CREATE_USER                  = "failed create user"
	MESSAGE_FAILED_DELETE_USER                  = "failed delete user"
	MESSAGE_FAILED_CREATE_NEWS                  = "failed create news"
	MESSAGE_FAILED_GET_LIST_NEWS                = "failed get list news"
	MESSAGE_FAILED_UPDATE_NEWS                  = "failed update news"
	MESSAGE_FAILED_DELETE_NEWS                  = "failed delete news"
	MESSAGE_FAILED_GET_LIST_MOTIVATION_CATEGORY = "failed get list motivation category"
	MESSAGE_FAILED_CREATE_MOTIVATION_CATEGORY   = "failed create motivation category"
	MESSAGE_FAILED_UPDATE_MOTIVATION_CATEGORY   = "failed update motivation category"
	// success
	MESSAGE_SUCCESS_CREATE_USER                  = "success create user"
	MESSAGE_SUCCESS_DELETE_USER                  = "success delete user"
	MESSAGE_SUCCESS_CREATE_NEWS                  = "success create news"
	MESSAGE_SUCCESS_GET_LIST_NEWS                = "success get list news"
	MESSAGE_SUCCESS_UPDATE_NEWS                  = "success update news"
	MESSAGE_SUCCESS_DELETE_NEWS                  = "success delete news"
	MESSAGE_SUCCESS_GET_LIST_MOTIVATION_CATEGORY = "success get list motivation category"
	MESSAGE_SUCCESS_CREATE_MOTIVATION_CATEGORY   = "success create motivation category"
	MESSAGE_SUCCESS_UPDATE_MOTIVATION_CATEGORY   = "success update motivation category"
)

var (
	ErrGetAllUserWithPagination               = errors.New("failed get list user with pagination")
	ErrDeleteUserByID                         = errors.New("failed delete user by id")
	ErrFormatBirthdate                        = errors.New("failed parse birthdate input")
	ErrCreateNews                             = errors.New("failed create news")
	ErrGetAllNewsWithPagination               = errors.New("failed get list news with pagination")
	ErrGetNewsFromID                          = errors.New("failed to get news data from id")
	ErrUpdateNews                             = errors.New("failed update news")
	ErrDeleteNews                             = errors.New("failed delete news")
	ErrGetAllMotivationCategoryWithPagination = errors.New("failed get list motivation category with pagination")
	ErrCreateMotivationCategory               = errors.New("failed create motivation category")
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
		UserID string `json:"-"`
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

	NewsPaginationResponse struct {
		PaginationResponse
		Data []NewsResponse `json:"data"`
	}

	AllNewsRepositoryResponse struct {
		PaginationResponse
		News []entity.News
	}

	UpdateNewsRequest struct {
		ID    string `json:"-"`
		Image string `json:"image,omitempty"`
		Title string `json:"title,omitempty"`
		Body  string `json:"body,omitempty"`
		Date  string `gorm:"type:date" json:"date,omitempty"`
	}

	DeleteNewsRequest struct {
		UserID string `json:"-"`
	}

	CreateMotivationRequest struct {
		Author               string `json:"author"`
		Content              string `json:"content"`
		MotivationCategoryID string `json:"mot_cat_id"`
	}

	MotivationResponse struct {
		ID                   string `json:"motivation_id"`
		Author               string `json:"motivation_author"`
		Content              string `json:"motivation_content"`
		MotivationCategoryID string `json:"mot_cat_id"`
	}

	CreateMotivationCategoryRequest struct {
		Name string `json:"name"`
	}

	MotivationCategoryResponse struct {
		ID   *uuid.UUID `json:"motivation_category_id"`
		Name string     `json:"motivation_category_name"`
	}

	MotivationCategoryPaginationResponse struct {
		PaginationResponse
		Data []MotivationCategoryResponse `json:"data"`
	}

	AllMotivationCategoryRepositoryResponse struct {
		PaginationResponse
		MotivationCategories []entity.MotivationCategory
	}

	UpdateMotivationCategoryRequest struct {
		ID   string `json:"-"`
		Name string `json:"name"`
	}
)
