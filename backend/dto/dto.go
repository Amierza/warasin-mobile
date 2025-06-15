package dto

import (
	"errors"
	"mime/multipart"

	"github.com/Amierza/warasin-mobile/backend/entity"
	"github.com/google/uuid"
)

const (
	// ====================================== Failed ======================================
	MESSAGE_FAILED_GET_DATA_FROM_BODY = "failed get data from body"
	// File
	MESSAGE_FAILED_READ_PHOTO = "failed read photo"
	MESSAGE_FAILED_OPEN_PHOTO = "failed open photo"
	// Middleware
	MESSAGE_FAILED_PROSES_REQUEST             = "failed proses request"
	MESSAGE_FAILED_ACCESS_DENIED              = "failed access denied"
	MESSAGE_FAILED_TOKEN_NOT_FOUND            = "failed token not found"
	MESSAGE_FAILED_TOKEN_NOT_VALID            = "failed token not valid"
	MESSAGE_FAILED_TOKEN_DENIED_ACCESS        = "failed token denied access"
	MESSAGE_FAILED_INAVLID_ENPOINTS_TOKEN     = "failed invalid endpoints in token"
	MESSAGE_FAILED_INAVLID_ROUTE_FORMAT_TOKEN = "failed invalid route format in token"
	// Authentication
	MESSAGE_FAILED_REGISTER_USER = "failed register user"
	MESSAGE_FAILED_LOGIN_USER    = "failed login user"
	MESSAGE_FAILED_REFRESH_TOKEN = "failed refresh token"
	MESSAGE_FAILED_LOGIN_ADMIN   = "failed login admin"
	// Send Email
	MESSAGE_FAILED_SEND_VERIFICATION_EMAIL     = "failed to send verification email"
	MESSAGE_FAILED_VERIFY_EMAIL                = "failed to verify email"
	MESSAGE_FAILED_SEND_FORGOT_PASSWORD_EMAIL  = "failed to send forgot password email"
	MESSAGE_FAILED_CHECK_FORGOT_PASSWORD_TOKEN = "failed to check forgot password token"
	// City & Province
	MESSAGE_FAILED_GET_LIST_CITY     = "failed get list city"
	MESSAGE_FAILED_GET_LIST_PROVINCE = "failed get list province"
	// Role
	MESSAGE_FAILED_GET_LIST_ROLE = "failed get all role"
	// User
	MESSAGE_FAILED_CREATE_USER     = "failed create user"
	MESSAGE_FAILED_GET_DETAIL_USER = "failed get detail user"
	MESSAGE_FAILED_GET_LIST_USER   = "failed get list user"
	MESSAGE_FAILED_UPDATE_USER     = "failed update user"
	MESSAGE_FAILED_DELETE_USER     = "failed delete user"
	MESSAGE_FAILED_UPDATE_PASSWORD = "failed to update password"
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
	MESSAGE_FAILED_LOGIN_PSYCHOLOG      = "failed login psycholog"
	MESSAGE_FAILED_CREATE_PSYCHOLOG     = "failed create psycholog"
	MESSAGE_FAILED_GET_LIST_PSYCHOLOG   = "failed get all psycholog"
	MESSAGE_FAILED_GET_DETAIL_PSYCHOLOG = "failed get detail psycholog"
	MESSAGE_FAILED_UPDATE_PSYCHOLOG     = "failed update psycholog"
	MESSAGE_FAILED_DELETE_PSYCHOLOG     = "failed delete psycholog"
	// User Motivation
	MESSAGE_FAILED_GET_PSYCHOLOG_LIST_USER_MOTIVATION = "failed get all user motivation"
	MESSAGE_FAILED_CREATE_USER_MOTIVATION             = "success create user motivation"
	MESSAGE_FAILED_GET_LIST_USER_MOTIVATION           = "success get all user motivation"
	// News Detail
	MESSAGE_FAILED_GET_LIST_NEWS_DETAIL = "failed get all news detail"
	MESSAGE_FAILED_CREATE_NEWS_DETAIL   = "failed create news detail"
	// Consultation
	MESSAGE_FAILED_CREATE_CONSULTATION     = "failed create consultation"
	MESSAGE_FAILED_GET_LIST_CONSULTATION   = "failed get all consultation"
	MESSAGE_FAILED_GET_DETAIL_CONSULTATION = "failed get detail consultation"
	MESSAGE_FAILED_UPDATE_CONSULTATION     = "failed update consultation"
	MESSAGE_FAILED_DELETE_CONSULTATION     = "failed delete consultation"
	// Language Master
	MESSAGE_FAILED_GET_ALL_LANGUAGE_MASTER = "failed get all language master"
	// Specialization
	MESSAGE_FAILED_GET_ALL_SPECIALIZATION = "failed get all specialization"
	// Practice
	MESSAGE_FAILED_CREATE_PRACTICE   = "failed create practice"
	MESSAGE_FAILED_GET_LIST_PRACTICE = "failed get all practice"
	MESSAGE_FAILED_UPDATE_PRACTICE   = "failed update practice"
	MESSAGE_FAILED_DELETE_PRACTICE   = "failed delete practice"
	// Available Slot
	MESSAGE_FAILED_GET_LIST_AVAILABLE_SLOT = "failed get all available slot"
	// Chat
	MESSAGE_FAILED_HANDLE_CHAT = "chat failed"

	// ====================================== Success ======================================
	// Authentication
	MESSAGE_SUCCESS_REGISTER_USER = "success register user"
	MESSAGE_SUCCESS_LOGIN_USER    = "success login user"
	MESSAGE_SUCCESS_LOGIN_ADMIN   = "success login admin"
	MESSAGE_SUCCESS_REFRESH_TOKEN = "success refresh token"
	// Send Email
	MESSAGE_SUCCESS_SEND_VERIFICATION_EMAIL     = "success to send verification email"
	MESSAGE_SUCCESS_VERIFY_EMAIL                = "success to verify email"
	MESSAGE_SUCCESS_SEND_FORGOT_PASSWORD_EMAIL  = "success to send forgot password email"
	MESSAGE_SUCCESS_CHECK_FORGOT_PASSWORD_TOKEN = "success to check forgot password token"
	// City & Province
	MESSAGE_SUCCESS_GET_LIST_CITY     = "success get list city"
	MESSAGE_SUCCESS_GET_LIST_PROVINCE = "success get list province"
	// Role
	MESSAGE_SUCCESS_GET_LIST_ROLE = "success get all role"
	// User
	MESSAGE_SUCCESS_CREATE_USER     = "success create user"
	MESSAGE_SUCCESS_GET_DETAIL_USER = "success get detail user"
	MESSAGE_SUCCESS_GET_LIST_USER   = "success get list user"
	MESSAGE_SUCCESS_UPDATE_USER     = "success update user"
	MESSAGE_SUCCESS_DELETE_USER     = "success delete user"
	MESSAGE_SUCCESS_UPDATE_PASSWORD = "success to update password"
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
	MESSAGE_SUCCESS_LOGIN_PSYCHOLOG      = "success login psycholog"
	MESSAGE_SUCCESS_CREATE_PSYCHOLOG     = "success create psycholog"
	MESSAGE_SUCCESS_GET_LIST_PSYCHOLOG   = "success get all psycholog"
	MESSAGE_SUCCESS_GET_DETAIL_PSYCHOLOG = "success get detail psycholog"
	MESSAGE_SUCCESS_UPDATE_PSYCHOLOG     = "success update psycholog"
	MESSAGE_SUCCESS_DELETE_PSYCHOLOG     = "success delete psycholog"
	// User Motivation
	MESSAGE_SUCCESS_GET_PSYCHOLOG_LIST_USER_MOTIVATION = "success get all user motivation"
	MESSAGE_SUCCESS_CREATE_USER_MOTIVATION             = "success create user motivation"
	MESSAGE_SUCCESS_GET_LIST_USER_MOTIVATION           = "success get all user motivation"
	// News Detail
	MESSAGE_SUCCESS_GET_LIST_NEWS_DETAIL = "success get all news detail"
	MESSAGE_SUCCESS_CREATE_NEWS_DETAIL   = "success create news detail"
	// Consultation
	MESSAGE_SUCCESS_CREATE_CONSULTATION     = "success create consultation"
	MESSAGE_SUCCESS_GET_LIST_CONSULTATION   = "success get all consultation"
	MESSAGE_SUCCESS_GET_DETAIL_CONSULTATION = "success get detail consultation"
	MESSAGE_SUCCESS_UPDATE_CONSULTATION     = "success update consultation"
	MESSAGE_SUCCESS_DELETE_CONSULTATION     = "success delete consultation"
	// Language Master
	MESSAGE_SUCCESS_GET_ALL_LANGUAGE_MASTER = "success get all language master"
	// Specialization
	MESSAGE_SUCCESS_GET_ALL_SPECIALIZATION = "success get all specialization"
	// Practice
	MESSAGE_SUCCESS_CREATE_PRACTICE   = "success create practice"
	MESSAGE_SUCCESS_GET_LIST_PRACTICE = "success get all practice"
	MESSAGE_SUCCESS_UPDATE_PRACTICE   = "success update practice"
	MESSAGE_SUCCESS_DELETE_PRACTICE   = "success delete practice"
	// Available Slot
	MESSAGE_SUCCESS_GET_LIST_AVAILABLE_SLOT = "success get all available slot"
	// Chat
	MESSAGE_SUCCESS_HANDLE_CHAT = "chat success"
)

var (
	// Parse
	ErrParseUUID             = errors.New("failed parse uuid")
	ErrParseConsultationDate = errors.New("failed parse consultation date")
	ErrParseDate             = errors.New("failed parse date")
	// Middleware
	ErrDeniedAccess           = errors.New("denied access")
	ErrGetPermissionsByRoleID = errors.New("failed get all permission by role id")
	// Input Validation
	ErrFieldEmpty                 = errors.New("failed there are empty fields")
	ErrFormatBirthdate            = errors.New("failed parse birthdate input")
	ErrInvalidSTRNumber           = errors.New("failed invalid STR Number")
	ErrInvalidWorkYear            = errors.New("failed invalid work year")
	ErrInvalidName                = errors.New("failed invalid name")
	ErrInvalidEmail               = errors.New("failed invalid email")
	ErrInvalidPassword            = errors.New("failed invalid password")
	ErrFormatPhoneNumber          = errors.New("failed standarize phone number input")
	ErrInvalidPracticeName        = errors.New("failed invalid practice name")
	ErrInvalidPracticeType        = errors.New("failed invalid practice type")
	ErrInvalidRateConsultation    = errors.New("failed invalid rate consultation")
	ErrConsultationCommentToShort = errors.New("failed consultation comment to short")
	ErrInvalidStatusConsultation  = errors.New("failed invalid status consultation")
	ErrInvalidStatusInput         = errors.New("failed invalid status input")
	ErrInvalidPsychologSchedule   = errors.New("failed invalid psycholog schedule")
	// Authentication
	ErrRegisterUser = errors.New("failed to register user")
	// Email
	ErrEmailAlreadyExists      = errors.New("email already exists")
	ErrEmailNotFound           = errors.New("email not found")
	ErrMakeVerificationEmail   = errors.New("failed to make verification email")
	ErrMakeForgotPasswordEmail = errors.New("failed to make forgot password email")
	ErrSendEmail               = errors.New("failed to send email")
	ErrEmailAlreadyVerified    = errors.New("email is already verfied")
	// Password
	ErrPasswordNotMatch = errors.New("password not match")
	ErrHashPassword     = errors.New("failed to hash password")
	// Token
	ErrGenerateToken           = errors.New("failed to generate token")
	ErrGenerateAccessToken     = errors.New("failed to generate access token")
	ErrGenerateRefreshToken    = errors.New("failed to generate refresh token")
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
	ErrDecryptToken            = errors.New("failed to decrypt token")
	ErrTokenInvalid            = errors.New("token invalid")
	ErrValidateToken           = errors.New("failed to validate token")
	ErrParsingExpiredTime      = errors.New("failed to parsing expired time")
	ErrTokenExpired            = errors.New("token expired")
	ErrInvalidToken            = errors.New("token invalid expired")
	// City & Province
	ErrGetCityByID    = errors.New("failed get city by id")
	ErrGetAllProvince = errors.New("failed get list province")
	// Role
	ErrGetRoleIDFromToken = errors.New("failed get role id from token")
	ErrGetRoleFromToken   = errors.New("failed get role from token")
	ErrGetRoleFromName    = errors.New("failed get role by role name")
	ErrGetRoleFromID      = errors.New("failed get role by role id")
	// Psycholog
	ErrGetPsychologIDFromToken       = errors.New("failed get psycholog id from token")
	ErrRegisterPsycholog             = errors.New("failed to register psycholog")
	ErrGetAllPsychologWithPagination = errors.New("failed get list psycholog with pagination")
	ErrGetAllPsycholog               = errors.New("failed get list psycholog")
	ErrPsychologNotFound             = errors.New("failed psycholog not found")
	ErrGetPsychologFromID            = errors.New("failed get` psycholog from id")
	ErrUpdatePsycholog               = errors.New("failed update psycholog")
	ErrDeletePsycholog               = errors.New("failed delete psycholog")
	// User
	ErrUserNotFound             = errors.New("user not found")
	ErrGetAllUserWithPagination = errors.New("failed get list user with pagination")
	ErrUpdateUser               = errors.New("failed to update user")
	ErrDeleteUserByID           = errors.New("failed delete user by id")
	ErrGetUserByPassword        = errors.New("failed to get user by password")
	ErrGetUserIDFromToken       = errors.New("failed get user id from token")
	ErrGetUserFromID            = errors.New("failed get user by id")
	// News
	ErrCreateNews               = errors.New("failed create news")
	ErrGetAllNewsWithPagination = errors.New("failed get list news with pagination")
	ErrGetNewsFromID            = errors.New("failed to get news data from id")
	ErrGetNewsFromTitle         = errors.New("failed to get news data from title")
	ErrNewsTitleAlreadyExists   = errors.New("failed news title already exists")
	ErrUpdateNews               = errors.New("failed update news")
	ErrDeleteNews               = errors.New("failed delete news")
	ErrNewsNotFound             = errors.New("failed news not found")
	// Motivation Category
	ErrCreateMotivationCategory               = errors.New("failed create motivation category")
	ErrGetAllMotivationCategoryWithPagination = errors.New("failed get list motivation category with pagination")
	ErrGetMotivationCategoryFromID            = errors.New("failed get motivation category data from id")
	ErrGetMotivationCategoryFromName          = errors.New("failed get motivation category data from name")
	ErrMotivationCategoryNameAlreadyExists    = errors.New("failed motivation category name is exists")
	ErrUpdateMotivationCategory               = errors.New("failed update motivation category")
	ErrDeleteMotivationCategory               = errors.New("failed delete motivation category")
	// Motivation
	ErrCreateMotivation               = errors.New("failed create motivation")
	ErrGetAllMotivationWithPagination = errors.New("failed get list motivation with pagination")
	ErrGetMotivationFromID            = errors.New("failed get motivation data from id")
	ErrGetMotivationFromContent       = errors.New("failed to get motivation data from content")
	ErrMotivationContentAlreadyExists = errors.New("failed motivation content already exists")
	ErrDeleteMotivation               = errors.New("failed delete motivation")
	ErrUpdateMotivation               = errors.New("failed update motivation")
	ErrMotivationNotFound             = errors.New("failed motivation not found")
	// Consultation
	ErrGetAllConsultationWithPagination = errors.New("failed get list consultation with pagination")
	ErrConsultationAlreadyBooked        = errors.New("failed consultation already booked")
	ErrConsultationNotFound             = errors.New("failed consultation not found")
	ErrUpdateConsultation               = errors.New("failed update consultation")
	ErrCreateConsultation               = errors.New("failed create consultation")
	ErrDeleteConsultation               = errors.New("failed delete consultation")
	// User motivation
	ErrGetAllUserMotivation        = errors.New("failed all user motivation")
	ErrUserMotivationAlreadyExists = errors.New("failed user motivation already exists")
	ErrCreateUserMotivation        = errors.New("failed create user motivation")
	ErrUserMotivationNotFound      = errors.New("failed user motivation not found")
	// News Detail
	ErrGetAllNewsDetail        = errors.New("failed all news detail")
	ErrCreateNewsDetail        = errors.New("failed create news detail")
	ErrNewsDetailAlreadyExists = errors.New("failed news detail already exists")
	ErrNewsDetailNotFound      = errors.New("failed news detail not found")
	// Language Master
	ErrLanguageMasterNotFound = errors.New("failed language master not found")
	ErrGetAllLanguageMaster   = errors.New("failed get all language master")
	// Psycholog Language
	ErrDeletePsychologLanguageByPsychologID = errors.New("failed delete psycholog language by psycholog id")
	ErrCreatePsychologLanguages             = errors.New("failed create psycholog languages")
	// Specialization
	ErrSpecializationNotFound = errors.New("failed specialization not found")
	// Psycholog Specialization
	ErrDeletePsychologSpecializationByPsychologID = errors.New("failed delete psycholog specialization by psycholog id")
	ErrCreatePsychologSpecializations             = errors.New("failed create psycholog specializations")
	// Education
	ErrEducationIsExists            = errors.New("failed education is exists")
	ErrDeleteEducationByPsychologID = errors.New("failed delete education by psycholog id")
	ErrCreateEducations             = errors.New("failed create educations")
	// Practice
	ErrPracticeAlreadyExists = errors.New("failed practice already exists")
	ErrPracticeNotFound      = errors.New("failed practice not found")
	ErrCreatePractice        = errors.New("failed create practice")
	ErrGetAllPractice        = errors.New("failed get all practice")
	ErrUpdatePractice        = errors.New("failed update practice")
	ErrDeletePractice        = errors.New("failed delete practice")
	// Practice Schedule
	ErrAddPracticeSchedule     = errors.New("failed add practice schedule")
	ErrCreatePracticeSchedule  = errors.New("failed create practice schedule")
	ErrDeletePracticeSchedules = errors.New("failed delete practice schedules")
	// Available Slot
	ErrAvailableSlotAlreadyExists = errors.New("failed available slot already exists")
	ErrGetAllAvailableSlot        = errors.New("failed get all available slot")
	ErrAvailableSlotNotFound      = errors.New("failed available slot not found")
	ErrUpdateStatusBookSlot       = errors.New("failed update book status slot")
	ErrCreateAvailableSlots       = errors.New("failed create available slots")
	// Chat
	ErrCreateConversation = errors.New("failed create conversation")
	ErrSaveMessage        = errors.New("failed save message")
	ErrGetChatGPTResponse = errors.New("failed get chat gpt response")
	ErrGetMessages        = errors.New("failed get messages")
	// File
	ErrInvalidExtensionPhoto = errors.New("only jpg/jpeg/png allowed")
	ErrCreateFile            = errors.New("failed create file")
	ErrSaveFile              = errors.New("failed save file")
)

type (
	// Role
	RolePaginationResponse struct {
		Data []RoleResponse `json:"data"`
	}
	AllRoleRepositoryResponse struct {
		Roles []entity.Role
	}
	RoleResponse struct {
		ID   *uuid.UUID `json:"role_id"`
		Name string     `json:"role_name"`
	}
	// Email
	SendForgotPasswordEmailRequest struct {
		Email string `json:"email" form:"email" binding:"required"`
	}
	ForgotPasswordRequest struct {
		Token string `json:"token" form:"token" binding:"required"`
	}
	ForgotPasswordResponse struct {
		Email string `json:"email" form:"email" binding:"required"`
	}
	SendVerificationEmailRequest struct {
		Email string `json:"email" form:"email" binding:"required"`
	}
	VerifyEmailRequest struct {
		Token string `json:"token" form:"token" binding:"required"`
	}
	VerifyEmailResponse struct {
		Email      string `json:"email"`
		IsVerified bool   `json:"is_verified"`
	}
	// City & Province
	ProvinceResponse struct {
		ID   *uuid.UUID `json:"province_id"`
		Name string     `json:"province_name"`
	}
	CityResponse struct {
		ID       *uuid.UUID       `json:"city_id"`
		Name     string           `json:"city_name"`
		Type     string           `json:"city_type"`
		Province ProvinceResponse `json:"province"`
	}
	ProvincesResponse struct {
		Data []ProvinceResponse `json:"data"`
	}
	CityResponseCustom struct {
		ID   *uuid.UUID `json:"city_id"`
		Name string     `json:"city_name"`
		Type string     `json:"city_type"`
	}
	CityQueryRequest struct {
		ProvinceID string `json:"province_id" form:"province_id"`
	}
	CitiesResponse struct {
		Data []CityResponseCustom `json:"data"`
	}
	AllProvinceRepositoryResponse struct {
		Provinces []entity.Province
	}
	AllCityRepositoryResponse struct {
		Cities []entity.City
	}
	// Authentication
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
	AdminLoginRequest struct {
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}
	AdminLoginResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
	PsychologLoginRequest struct {
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}
	PsychologLoginResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
	RefreshTokenRequest struct {
		RefreshToken string `json:"refresh_token"`
	}

	RefreshTokenResponse struct {
		AccessToken string `json:"access_token"`
	}
	// User
	AllUserResponse struct {
		ID          uuid.UUID    `json:"user_id"`
		Name        string       `json:"user_name"`
		Email       string       `json:"user_email"`
		Password    string       `json:"user_password"`
		Image       string       `json:"user_image"`
		Gender      *bool        `json:"user_gender"`
		Birthdate   string       `json:"user_birth_date"`
		PhoneNumber string       `json:"user_phone_number"`
		IsVerified  *bool        `json:"is_verified"`
		Data01      int          `json:"user_data01"`
		Data02      int          `json:"user_data02"`
		Data03      int          `json:"user_data03"`
		City        CityResponse `json:"city"`
		Role        RoleResponse `json:"role"`
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
		Name        string                `json:"name"`
		Email       string                `json:"email"`
		Password    string                `json:"password"`
		Image       string                `json:"image,omitempty"`
		Gender      *bool                 `json:"gender"`
		Birthdate   string                `json:"birth_date"`
		PhoneNumber string                `json:"phone_number"`
		CityID      *uuid.UUID            `json:"city_id"`
		RoleID      *uuid.UUID            `json:"role_id"`
		FileHeader  *multipart.FileHeader `json:"fileheader,omitempty"`
		FileReader  multipart.File        `json:"filereader,omitempty"`
	}
	UpdateUserRequest struct {
		ID          string                `json:"-"`
		Name        string                `json:"name,omitempty"`
		Email       string                `json:"email,omitempty"`
		Image       string                `json:"image,omitempty"`
		Gender      *bool                 `json:"gender,omitempty"`
		Birthdate   string                `json:"birth_date,omitempty"`
		PhoneNumber string                `json:"phone_number,omitempty"`
		CityID      *uuid.UUID            `json:"city_id,omitempty"`
		RoleID      *uuid.UUID            `json:"role_id,omitempty"`
		FileHeader  *multipart.FileHeader `json:"fileheader,omitempty"`
		FileReader  multipart.File        `json:"filereader,omitempty"`
	}
	DeleteUserRequest struct {
		UserID string `json:"-"`
	}
	UpdatePasswordRequest struct {
		Email    string `json:"email" form:"email" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}
	UpdatePasswordResponse struct {
		OldPassword string `json:"old_password" form:"old_password" binding:"required"`
		NewPassword string `json:"new_password" form:"new_password" binding:"required"`
	}
	// News
	CreateNewsRequest struct {
		Image      string                `json:"image,omitempty"`
		Title      string                `json:"title"`
		Body       string                `json:"body"`
		Date       string                `json:"date"`
		FileHeader *multipart.FileHeader `json:"fileheader,omitempty"`
		FileReader multipart.File        `json:"filereader,omitempty"`
	}
	NewsResponse struct {
		ID    *uuid.UUID `json:"news_id"`
		Image string     `json:"news_image"`
		Title string     `json:"news_title"`
		Body  string     `json:"news_body"`
		Date  string     `json:"news_date"`
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
		ID         string                `json:"-"`
		Image      string                `json:"image,omitempty"`
		Title      string                `json:"title,omitempty"`
		Body       string                `json:"body,omitempty"`
		Date       string                `json:"date,omitempty"`
		FileHeader *multipart.FileHeader `json:"fileheader,omitempty"`
		FileReader multipart.File        `json:"filereader,omitempty"`
	}
	DeleteNewsRequest struct {
		NewsID string `json:"-"`
	}
	// Motivation Category
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
	// Motivation
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
	// Psycholog
	CreatePsychologRequest struct {
		Name              string                `json:"name"`
		STRNumber         string                `json:"str_number"`
		Email             string                `gorm:"unique" json:"email"`
		Password          string                `json:"password"`
		WorkYear          string                `json:"work_year"`
		Description       string                `json:"description"`
		PhoneNumber       string                `json:"phone_number"`
		Image             string                `json:"image,omitempty"`
		CityID            *uuid.UUID            `json:"city_id"`
		RoleID            *uuid.UUID            `json:"role_id"`
		LanguageMasterIDs []string              `json:"language_master,omitempty"`
		SpecializationIDs []string              `json:"specialization,omitempty"`
		Educations        []EducationResponse   `json:"education,omitempty"`
		FileHeader        *multipart.FileHeader `json:"fileheader,omitempty"`
		FileReader        multipart.File        `json:"filereader,omitempty"`
	}
	PsychologResponse struct {
		ID              uuid.UUID                `json:"psy_id"`
		Name            string                   `json:"psy_name"`
		STRNumber       string                   `json:"psy_str_number"`
		Email           string                   `json:"psy_email"`
		Password        string                   `json:"psy_password"`
		WorkYear        string                   `json:"psy_work_year"`
		Description     string                   `json:"psy_description"`
		PhoneNumber     string                   `json:"psy_phone_number"`
		Image           string                   `json:"psy_image"`
		City            CityResponse             `json:"city"`
		Role            RoleResponse             `json:"role"`
		LanguageMasters []LanguageMasterResponse `json:"language"`
		Specializations []SpecializationResponse `json:"specialization"`
		Educations      []EducationResponse      `json:"education"`
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
		ID                string                `json:"-"`
		Name              string                `json:"name,omitempty"`
		STRNumber         string                `json:"str_number,omitempty"`
		Email             string                `json:"email,omitempty"`
		WorkYear          string                `json:"work_year,omitempty"`
		Description       string                `json:"description,omitempty"`
		PhoneNumber       string                `json:"phone_number,omitempty"`
		Image             string                `json:"image,omitempty"`
		CityID            *uuid.UUID            `json:"city_id,omitempty"`
		LanguageMasterIDs []string              `json:"language_master,omitempty"`
		SpecializationIDs []string              `json:"specialization,omitempty"`
		Educations        []EducationResponse   `json:"education,omitempty"`
		FileHeader        *multipart.FileHeader `json:"fileheader,omitempty"`
		FileReader        multipart.File        `json:"filereader,omitempty"`
	}
	DeletePsychologRequest struct {
		ID string `json:"-"`
	}
	// Language Master
	LanguageMasterResponse struct {
		ID   *uuid.UUID `json:"lang_id"`
		Name string     `json:"lang_name"`
	}
	AllLanguageMasterRepositoryResponse struct {
		LanguageMasters []entity.LanguageMaster
	}
	AllLanguageMasterResponse struct {
		LanguageMasters []LanguageMasterResponse `json:"language_master"`
	}
	// Specialization
	SpecializationResponse struct {
		ID          *uuid.UUID `json:"spe_id"`
		Name        string     `json:"spe_name"`
		Description string     `json:"spe_desc"`
	}
	AllSpecializationRepositoryResponse struct {
		Specializations []entity.Specialization
	}
	AllSpecializationResponse struct {
		Specializations []SpecializationResponse `json:"specialization"`
	}
	// Education
	EducationResponse struct {
		ID             *uuid.UUID `json:"edu_id"`
		Degree         string     `json:"edu_degree"`
		Major          string     `json:"edu_major"`
		Institution    string     `json:"edu_institution"`
		GraduationYear string     `json:"edu_graduation_year"`
	}
	// User Motivation
	CreateUserMotivationRequest struct {
		DisplayDate  string `json:"user_mot_display_date"`
		Reaction     int    `json:"user_mot_reaction"`
		MotivationID string `json:"mot_id"`
	}
	UserMotivationResponse struct {
		ID          *uuid.UUID         `json:"user_mot_id"`
		DisplayDate string             `json:"user_mot_date"`
		Reaction    int                `json:"user_mot_reaction"`
		User        AllUserResponse    `json:"user"`
		Motivation  MotivationResponse `json:"motivation"`
	}
	UserMotivationResponseCustom struct {
		ID          uuid.UUID          `json:"user_mot_id"`
		DisplayDate string             `json:"user_mot_date"`
		Reaction    int                `json:"user_mot_reaction"`
		Motivation  MotivationResponse `json:"motivation"`
	}
	AllUserMotivationRepositoryResponse struct {
		PaginationResponse
		UserMotivations []entity.UserMotivation
	}
	UserMotivationPaginationResponse struct {
		PaginationResponse
		Data []UserMotivationResponse `json:"data"`
	}
	// News Detail
	UserNewsResponse struct {
		ID   *uuid.UUID      `json:"news_detail_id"`
		Date string          `json:"news_detail_date"`
		User AllUserResponse `json:"user"`
		News NewsResponse    `json:"news"`
	}
	NewsDetailResponse struct {
		ID   *uuid.UUID   `json:"news_detail_id"`
		Date string       `json:"news_detail_date"`
		News NewsResponse `json:"news"`
	}
	AllUserNewsRepositoryResponse struct {
		PaginationResponse
		UserNews []entity.NewsDetail
	}
	UserNewsPaginationResponse struct {
		PaginationResponse
		Data []UserNewsResponse `json:"data"`
	}
	CreateNewsDetailRequest struct {
		Date   string `json:"news_detail_date"`
		NewsID string `json:"news_id"`
	}
	// Practice
	CreatePracticeRequest struct {
		Type        string `json:"prac_type"`
		Name        string `json:"prac_name"`
		Address     string `json:"prac_address"`
		PhoneNumber string `json:"prac_phone_number"`
	}
	UpdatePracticeRequest struct {
		ID          string `json:"-"`
		Type        string `json:"prac_type"`
		Name        string `json:"prac_name"`
		Address     string `json:"prac_address"`
		PhoneNumber string `json:"prac_phone_number"`
	}
	PracticeScheduleResponse struct {
		ID    uuid.UUID `json:"prac_sched_id"`
		Day   string    `json:"prac_sched_day"`
		Open  string    `json:"prac_sched_open"`
		Close string    `json:"prac_sched_close"`
	}
	PracticeResponse struct {
		ID                uuid.UUID                  `json:"prac_id"`
		Type              string                     `json:"prac_type"`
		Name              string                     `json:"prac_name"`
		Address           string                     `json:"prac_address"`
		PhoneNumber       string                     `json:"prac_phone_number"`
		PracticeSchedules []PracticeScheduleResponse `json:"practice_schedule"`
	}
	AllPracticeRepositoryResponse struct {
		Practices []entity.Practice
	}
	AllPracticeResponse struct {
		Psycholog PsychologResponse  `json:"psycholog"`
		Practices []PracticeResponse `json:"practice"`
	}
	// Available Slot
	AvailableSlotResponse struct {
		ID       uuid.UUID `json:"slot_id"`
		Start    string    `json:"slot_start"`
		End      string    `json:"slot_end"`
		IsBooked bool      `json:"slot_is_booked"`
	}
	AllAvailableSlotRepositoryResponse struct {
		AvailableSlots []entity.AvailableSlot
	}
	AllAvailableSlotResponse struct {
		Psycholog      PsychologResponse       `json:"psycholog"`
		AvailableSlots []AvailableSlotResponse `json:"available_slot"`
	}
	// Consultation
	ConsultationResponse struct {
		ID            uuid.UUID             `json:"consul_id"`
		Date          string                `json:"consul_date"`
		Rate          int                   `json:"consul_rate"`
		Comment       string                `json:"consul_comment"`
		Status        int                   `json:"consul_status"`
		User          AllUserResponse       `json:"user"`
		AvailableSlot AvailableSlotResponse `json:"available_slot"`
		Practice      PracticeResponse      `json:"practice"`
	}
	AllConsultationRepositoryResponse struct {
		PaginationResponse
		Consultations []entity.Consultation
	}
	AllConsultationResponse struct {
		Psycholog    PsychologResponse      `json:"psycholog"`
		Consultation []ConsultationResponse `json:"consultation"`
	}
	ConsultationPaginationResponse struct {
		PaginationResponse
		Data AllConsultationResponse `json:"data"`
	}
	UpdateConsultationRequest struct {
		Status *int `json:"status"`
	}
	CreateConsultationRequest struct {
		Date            string `json:"consul_date"`
		AvailableSlotID string `json:"slot_id"`
		PracticeID      string `json:"prac_id"`
	}
	ConsultationResponseForUser struct {
		ID            uuid.UUID             `json:"consul_id"`
		Date          string                `json:"consul_date"`
		Rate          int                   `json:"consul_rate"`
		Comment       string                `json:"consul_comment"`
		Status        int                   `json:"consul_status"`
		Psycholog     PsychologResponse     `json:"psycholog"`
		AvailableSlot AvailableSlotResponse `json:"available_slot"`
		Practice      PracticeResponse      `json:"practice"`
	}
	AllConsultationRepositoryResponseForUser struct {
		PaginationResponse
		Consultations []entity.Consultation
	}
	AllConsultationResponseForUser struct {
		User         AllUserResponse               `json:"user"`
		Consultation []ConsultationResponseForUser `json:"consultation"`
	}
	ConsultationPaginationResponseForUser struct {
		PaginationResponse
		Data AllConsultationResponseForUser `json:"data"`
	}
	UpdateConsultationRequestForUser struct {
		ID              string `json:"-"`
		Date            string `json:"consul_date,omitempty"`
		Rate            *int   `json:"consul_rate,omitempty"`
		Status          *int   `json:"consul_status,omitempty"`
		Comment         string `json:"consul_comment,omitempty"`
		AvailableSlotID string `json:"slot_id,omitempty"`
		PracticeID      string `json:"prac_id,omitempty"`
	}
	PsychologFilter struct {
		Name           string
		City           string
		Province       string
		Specialization string
	}
	// chat
	ChatRequest struct {
		Message        string    `json:"message" binding:"required"`
		ConversationID uuid.UUID `json:"conversation_id,omitempty"`
		UserID         uuid.UUID `json:"user_id"`
	}
	ChatResponse struct {
		Response       string    `json:"response"`
		ConversationID uuid.UUID `json:"conversation_id"`
	}
)
