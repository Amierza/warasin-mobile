package handler

import "github.com/Amierza/warasin-mobile/backend/service"

type (
	IUserHandler interface {
	}

	UserHandler struct {
		userService service.IUserService
	}
)

func NewUserHandler(userService service.IUserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}
