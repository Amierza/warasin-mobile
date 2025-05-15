package dto

import (
	"errors"

	"github.com/Amierza/warasin-mobile/backend/entity"
	"github.com/google/uuid"
)

const (
	// Failed
	// Admin
	MESSAGE_FAILED_LOGIN_ADMIN = "failed login admin"

	// Role
	MESSAGE_FAILED_GET_LIST_ROLE = "failed get all role"

	// User
	MESSAGE_FAILED_CREATE_USER     = "failed create user"
	MESSAGE_FAILED_GET_DETAIL_USER = "failed get detail user"
	MESSAGE_FAILED_GET_LIST_USER   = "failed get list user"
	MESSAGE_FAILED_UPDATE_USER     = "failed update user"
	MESSAGE_FAILED_DELETE_USER     = "failed delete user"

	// News
	MESSAGE_FAILED_CREATE_NEWS     = "failed create news"
	MESSAGE_FAILED_GET_LIST_NEWS   = "failed get list news"
	MESSAGE_FAILED_GET_DETAIL_NEWS = "failed get detail news"
	MESSAGE_FAILED_UPDATE_NEWS     = "failed update news"
	MESSAGE_FAILED_DELETE_NEWS     = "failed delete news"

	// Motivation Category
	MESSAGE_FAILED_CREATE_MOTIVATION_CATEGORY     = "failed create motivation category"
	MESSAGE_FAILED_GET_LIST_MOTIVATION_CATEGORY   = "failed get list motivation category"
	MESSAGE_FAILED_GET_DETAIL_MOTIVATION_CATEGORY = "failed get detail motivation category"
	MESSAGE_FAILED_UPDATE_MOTIVATION_CATEGORY     = "failed update motivation category"
	MESSAGE_FAILED_DELETE_MOTIVATION_CATEGORY     = "failed delete motivation category"

	// Motivation
	MESSAGE_FAILED_CREATE_MOTIVATION     = "failed create motivation"
	MESSAGE_FAILED_GET_LIST_MOTIVATION   = "failed get all motivation"
	MESSAGE_FAILED_GET_DETAIL_MOTIVATION = "failed get detail motivation"
	MESSAGE_FAILED_UPDATE_MOTIVATION     = "failed update motivation"
	MESSAGE_FAILED_DELETE_MOTIVATION     = "failed delete motivation"

	// Psycholog
	MESSAGE_FAILED_CREATE_PSYCHOLOG     = "failed create psycholog"
	MESSAGE_FAILED_GET_LIST_PSYCHOLOG   = "failed get all psycholog"
	MESSAGE_FAILED_GET_DETAIL_PSYCHOLOG = "failed get detail psycholog"
	MESSAGE_FAILED_UPDATE_PSYCHOLOG     = "failed update psycholog"
	MESSAGE_FAILED_DELETE_PSYCHOLOG     = "failed delete psycholog"

	// Success
	// Admin
	MESSAGE_SUCCESS_LOGIN_ADMIN = "success login admin"

	// Role
	MESSAGE_SUCCESS_GET_LIST_ROLE = "success get all role"

	// User
	MESSAGE_SUCCESS_CREATE_USER     = "success create user"
	MESSAGE_SUCCESS_GET_DETAIL_USER = "success get detail user"
	MESSAGE_SUCCESS_GET_LIST_USER   = "success get list user"
	MESSAGE_SUCCESS_UPDATE_USER     = "success update user"
	MESSAGE_SUCCESS_DELETE_USER     = "success delete user"

	// News
	MESSAGE_SUCCESS_CREATE_NEWS     = "success create news"
	MESSAGE_SUCCESS_GET_LIST_NEWS   = "success get list news"
	MESSAGE_SUCCESS_GET_DETAIL_NEWS = "success get detail news"
	MESSAGE_SUCCESS_UPDATE_NEWS     = "success update news"
	MESSAGE_SUCCESS_DELETE_NEWS     = "success delete news"

	// Motivation Category
	MESSAGE_SUCCESS_CREATE_MOTIVATION_CATEGORY     = "success create motivation category"
	MESSAGE_SUCCESS_GET_LIST_MOTIVATION_CATEGORY   = "success get list motivation category"
	MESSAGE_SUCCESS_GET_DETAIL_MOTIVATION_CATEGORY = "success get detail motivation category"
	MESSAGE_SUCCESS_UPDATE_MOTIVATION_CATEGORY     = "success update motivation category"
	MESSAGE_SUCCESS_DELETE_MOTIVATION_CATEGORY     = "success delete motivation category"

	// Motivation
	MESSAGE_SUCCESS_CREATE_MOTIVATION     = "success create motivation"
	MESSAGE_SUCCESS_GET_LIST_MOTIVATION   = "success get all motivation"
	MESSAGE_SUCCESS_GET_DETAIL_MOTIVATION = "success get detail motivation"
	MESSAGE_SUCCESS_UPDATE_MOTIVATION     = "success update motivation"
	MESSAGE_SUCCESS_DELETE_MOTIVATION     = "success delete motivation"

	// Psycholog
	MESSAGE_SUCCESS_CREATE_PSYCHOLOG     = "success create psycholog"
	MESSAGE_SUCCESS_GET_LIST_PSYCHOLOG   = "success get all psycholog"
	MESSAGE_SUCCESS_GET_DETAIL_PSYCHOLOG = "success get detail psycholog"
	MESSAGE_SUCCESS_UPDATE_PSYCHOLOG     = "success update psycholog"
	MESSAGE_SUCCESS_DELETE_PSYCHOLOG     = "success delete psycholog"
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
	ErrCreateMotivationCategory               = errors.New("failed create motivation category")
	ErrGetAllMotivationCategoryWithPagination = errors.New("failed get list motivation category with pagination")
	ErrGetMotivationCategoryFromID            = errors.New("failed get motivation category data from id")
	ErrUpdateMotivationCategory               = errors.New("failed update motivation category")
	ErrDeleteMotivationCategory               = errors.New("failed delete motivation category")
	ErrCreateMotivation                       = errors.New("failed create motivation")
	ErrGetAllMotivationWithPagination         = errors.New("failed get list motivation with pagination")
	ErrGetMotivationFromID                    = errors.New("failed get motivation data from id")
	ErrDeleteMotivation                       = errors.New("failed delete motivation")
	ErrUpdateMotivation                       = errors.New("failed update motivation")
	ErrMotivationNotFound                     = errors.New("failed motivation not found")
	ErrRegisterPsycholog                      = errors.New("failed to register psycholog")
	ErrInvalidSTRNumber                       = errors.New("failed invalid STR Number")
	ErrInvalidWorkYear                        = errors.New("failed invalid work year")
	ErrGetAllPsychologWithPagination          = errors.New("failed get list psycholog with pagination")
	ErrPsychologNotFound                      = errors.New("failed psycholog not found")
	ErrGetDataPsychologFromID                 = errors.New("failed get data psycholog from id")
	ErrUpdatePsycholog                        = errors.New("failed update psycholog")
	ErrDeletePsycholog                        = errors.New("failed delete psycholog")
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

	UpdateUserRequest struct {
		ID          string     `json:"-"`
		Name        string     `json:"name,omitempty"`
		Email       string     `json:"email,omitempty"`
		Image       string     `json:"image,omitempty"`
		Gender      *bool      `json:"gender,omitempty"`
		Birthdate   string     `json:"birth_date,omitempty"`
		PhoneNumber string     `json:"phone_number,omitempty"`
		CityID      *uuid.UUID `gorm:"type:uuid" json:"city_id,omitempty"`
		RoleID      *uuid.UUID `gorm:"type:uuid" json:"role_id,omitempty"`
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
		NewsID string `json:"-"`
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
		Name string `json:"name,omitempty"`
	}

	DeleteMotivationCategoryRequest struct {
		MotivationCategoryID string `json:"-"`
	}

	CreateMotivationRequest struct {
		Author               string     `json:"author"`
		Content              string     `json:"content"`
		MotivationCategoryID *uuid.UUID `json:"motivation_category_id"`
	}

	MotivationResponse struct {
		ID                 *uuid.UUID                 `json:"motivation_id"`
		Author             string                     `json:"motivation_author"`
		Content            string                     `json:"motivation_content"`
		MotivationCategory MotivationCategoryResponse `json:"motivation_category"`
	}

	MotivationPaginationResponse struct {
		PaginationResponse
		Data []MotivationResponse `json:"data"`
	}

	AllMotivationRepositoryResponse struct {
		PaginationResponse
		Motivations []entity.Motivation
	}

	UpdateMotivationRequest struct {
		ID                   string `json:"-"`
		Author               string `json:"author,omitempty"`
		Content              string `json:"content,omitempty"`
		MotivationCategoryID string `json:"motivation_category_id,omitempty"`
	}

	DeleteMotivationRequest struct {
		ID string `json:"-"`
	}

	RolePaginationResponse struct {
		Data []RoleResponse `json:"data"`
	}

	AllRoleRepositoryResponse struct {
		Roles []entity.Role
	}

	CreatePsychologRequest struct {
		Name        string     `json:"name"`
		STRNumber   string     `json:"str_number"`
		Email       string     `gorm:"unique" json:"email"`
		Password    string     `json:"password"`
		WorkYear    string     `json:"work_year"`
		Description string     `json:"description"`
		PhoneNumber string     `json:"phone_number"`
		Image       string     `json:"image"`
		CityID      *uuid.UUID `gorm:"type:uuid" json:"city_id"`
		RoleID      *uuid.UUID `gorm:"type:uuid" json:"role_id"`
	}

	PsychologResponse struct {
		ID          uuid.UUID    `json:"psy_id"`
		Name        string       `json:"psy_name"`
		STRNumber   string       `json:"psy_str_number"`
		Email       string       `json:"psy_email"`
		Password    string       `json:"psy_password"`
		WorkYear    string       `json:"psy_work_year"`
		Description string       `json:"psy_description"`
		PhoneNumber string       `json:"psy_phone_number"`
		Image       string       `json:"psy_image"`
		City        CityResponse `json:"city"`
		Role        RoleResponse `json:"role"`
	}

	PsychologPaginationResponse struct {
		PaginationResponse
		Data []PsychologResponse `json:"data"`
	}

	AllPsychologRepositoryResponse struct {
		PaginationResponse
		Psychologs []entity.Psycholog
	}

	UpdatePsychologRequest struct {
		ID          string `json:"-"`
		Name        string `json:"name,omitempty"`
		STRNumber   string `json:"str_number,omitempty"`
		Email       string `json:"email,omitempty"`
		WorkYear    string `json:"work_year,omitempty"`
		Description string `json:"description,omitempty"`
		PhoneNumber string `json:"phone_number,omitempty"`
		Image       string `json:"image,omitempty"`
		CityID      string `json:"city_id,omitempty"`
	}

	DeletePsychologRequest struct {
		ID string `json:"-"`
	}
)
