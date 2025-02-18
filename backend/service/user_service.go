package service

import "github.com/Amierza/warasin-mobile/backend/repository"

type (
	IUserService interface {
	}

	UserService struct {
		userRepo   repository.IUserRepository
		jwtService IJWTService
	}
)

func NewUserService(userRepo repository.IUserRepository, jwtService IJWTService) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}
