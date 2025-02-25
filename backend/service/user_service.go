package service

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"os"
	"strings"
	"time"

	"github.com/Amierza/warasin-mobile/backend/dto"
	"github.com/Amierza/warasin-mobile/backend/entity"
	"github.com/Amierza/warasin-mobile/backend/helpers"
	"github.com/Amierza/warasin-mobile/backend/repository"
	"github.com/Amierza/warasin-mobile/backend/utils"
	"github.com/go-playground/validator/v10"
)

type (
	IUserService interface {
		Register(ctx context.Context, req dto.UserRegisterRequest) (dto.UserResponse, error)
		Login(ctx context.Context, req dto.UserLoginRequest) (dto.UserLoginResponse, error)
		SendVerificationEmail(ctx context.Context, req dto.SendVerificationEmailRequest) error
		VerifyEmail(ctx context.Context, req dto.VerifyEmailRequest) (dto.VerifyEmailResponse, error)
	}

	UserService struct {
		userRepo   repository.IUserRepository
		jwtService IJWTService
	}
)

func NewUserService(userRepo repository.IUserRepository, jwtService IJWTService) *UserService {
	return &UserService{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

func (us *UserService) Register(ctx context.Context, req dto.UserRegisterRequest) (dto.UserResponse, error) {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		var errorMessages []string
		for _, err := range err.(validator.ValidationErrors) {
			errorMessages = append(errorMessages, err.Error())
		}

		return dto.UserResponse{}, fmt.Errorf("validation errors: %v", errorMessages)
	}

	_, flag, err := us.userRepo.CheckEmail(ctx, nil, req.Email)
	if flag || err == nil {
		return dto.UserResponse{}, dto.ErrEmailAlreadyExists
	}

	user := entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	userReg, err := us.userRepo.RegisterUser(ctx, nil, user)
	if err != nil {
		return dto.UserResponse{}, dto.ErrRegisterUser
	}

	return dto.UserResponse{
		ID:       userReg.ID,
		Name:     userReg.Name,
		Email:    userReg.Email,
		Password: userReg.Password,
	}, nil
}

func (us *UserService) Login(ctx context.Context, req dto.UserLoginRequest) (dto.UserLoginResponse, error) {
	user, flag, err := us.userRepo.CheckEmail(ctx, nil, req.Email)
	if !flag || err != nil {
		return dto.UserLoginResponse{}, dto.ErrEmailNotFound
	}

	checkPassword, err := helpers.CheckPassword(user.Password, []byte(req.Password))
	if err != nil || !checkPassword {
		return dto.UserLoginResponse{}, dto.ErrPasswordNotMatch
	}

	accessToken, refreshToken, err := us.jwtService.GenerateToken(user.ID.String())
	if err != nil {
		return dto.UserLoginResponse{}, err
	}

	return dto.UserLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func makeVerificationEmail(receiverEmail string) (map[string]string, error) {
	expired := time.Now().Add(time.Hour * 24).Format("2006-01-02 15:04:05")
	plainText := fmt.Sprintf("%s_%s", receiverEmail, expired)
	token, err := utils.AESEncrypt(plainText)
	if err != nil {
		return nil, err
	}

	baseURL := os.Getenv("BASE_URL")
	verifyEmailRoute := "verify-email"
	if baseURL == "" {
		baseURL = "http://127.0.0.1:8000/api/v1/user"
	}

	verifyLink := baseURL + "/" + verifyEmailRoute + "?token=" + token

	readHTML, err := os.ReadFile("utils/email_template/base_mail.html")
	if err != nil {
		return nil, err
	}

	data := struct {
		Email  string
		Verify string
	}{
		Email:  receiverEmail,
		Verify: verifyLink,
	}

	tmpl, err := template.New("custome").Parse(string(readHTML))
	if err != nil {
		return nil, err
	}

	var strMail bytes.Buffer
	if err := tmpl.Execute(&strMail, data); err != nil {
		return nil, err
	}

	draftEmail := map[string]string{
		"subject": "warasin",
		"body":    strMail.String(),
	}

	return draftEmail, nil
}

func (us *UserService) SendVerificationEmail(ctx context.Context, req dto.SendVerificationEmailRequest) error {
	user, flag, err := us.userRepo.CheckEmail(ctx, nil, req.Email)
	if err != nil || !flag {
		return dto.ErrEmailNotFound
	}

	draftEmail, err := makeVerificationEmail(user.Email)
	if err != nil {
		return dto.ErrMakeVerificationEmail
	}

	if err := utils.SendEmail(user.Email, draftEmail["subject"], draftEmail["body"]); err != nil {
		return dto.ErrSendEmail
	}

	return nil
}

func (us *UserService) VerifyEmail(ctx context.Context, req dto.VerifyEmailRequest) (dto.VerifyEmailResponse, error) {
	decryptedToken, err := utils.AESDecrypt(req.Token)
	if err != nil {
		return dto.VerifyEmailResponse{}, dto.ErrDecryptToken
	}

	if !strings.Contains(decryptedToken, "_") {
		return dto.VerifyEmailResponse{}, dto.ErrTokenInvalid
	}

	decryptedTokenSplit := strings.Split(decryptedToken, "_")
	if len(decryptedTokenSplit) != 2 {
		return dto.VerifyEmailResponse{}, dto.ErrTokenInvalid
	}

	email := decryptedTokenSplit[0]
	expired := decryptedTokenSplit[1]

	now := time.Now()
	expiredTime, err := time.Parse("2006-01-02 15:04:05", expired)
	if err != nil {
		return dto.VerifyEmailResponse{}, dto.ErrParsingExpiredTime
	}

	if expiredTime.Sub(now) < 0 {
		return dto.VerifyEmailResponse{
			Email:      email,
			IsVerified: false,
		}, dto.ErrTokenExpired
	}

	user, flag, err := us.userRepo.CheckEmail(ctx, nil, email)
	if !flag || err != nil {
		return dto.VerifyEmailResponse{}, dto.ErrUserNotFound
	}

	if user.IsVerified {
		return dto.VerifyEmailResponse{}, dto.ErrEmailALreadyVerified
	}

	updatedUser, err := us.userRepo.UpdateUser(ctx, nil, entity.User{
		ID:         user.ID,
		IsVerified: true,
	})
	if err != nil {
		return dto.VerifyEmailResponse{}, dto.ErrUpdateUser
	}

	return dto.VerifyEmailResponse{
		Email:      email,
		IsVerified: updatedUser.IsVerified,
	}, nil
}
