package service

import (
	"context"
	"fmt"

	"github.com/Amierza/warasin-mobile/backend/dto"
	"github.com/Amierza/warasin-mobile/backend/entity"
	"github.com/Amierza/warasin-mobile/backend/helpers"
	"github.com/Amierza/warasin-mobile/backend/repository"
	"github.com/go-playground/validator/v10"
)

type (
	IUserService interface {
		Register(ctx context.Context, req dto.UserRegisterRequest) (dto.UserResponse, error)
		Login(ctx context.Context, req dto.UserLoginRequest) (dto.UserLoginResponse, error)
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
