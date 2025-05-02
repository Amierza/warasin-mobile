package service

import (
	"context"

	"github.com/Amierza/warasin-mobile/backend/dto"
	"github.com/Amierza/warasin-mobile/backend/helpers"
	"github.com/Amierza/warasin-mobile/backend/repository"
)

type (
	IAdminService interface {
		Login(ctx context.Context, req dto.AdminLoginRequest) (dto.AdminLoginResponse, error)
		RefreshToken(ctx context.Context, req dto.RefreshTokenRequest) (dto.RefreshTokenResponse, error)
	}

	AdminService struct {
		adminRepo  repository.IAdminRepository
		jwtService IJWTService
	}
)

func NewAdminService(adminRepo repository.IAdminRepository, jwtService IJWTService) *AdminService {
	return &AdminService{
		adminRepo:  adminRepo,
		jwtService: jwtService,
	}
}

func (as *AdminService) Login(ctx context.Context, req dto.AdminLoginRequest) (dto.AdminLoginResponse, error) {
	user, flag, err := as.adminRepo.CheckEmail(ctx, nil, req.Email)
	if !flag || err != nil {
		return dto.AdminLoginResponse{}, dto.ErrEmailNotFound
	}

	checkPassword, err := helpers.CheckPassword(user.Password, []byte(req.Password))
	if err != nil || !checkPassword {
		return dto.AdminLoginResponse{}, dto.ErrPasswordNotMatch
	}

	if user.Role.Name != "admin" {
		return dto.AdminLoginResponse{}, dto.ErrDeniedAccess
	}

	endpoints, err := as.adminRepo.GetPermissionsByRoleID(ctx, nil, user.RoleID.String())
	if err != nil {
		return dto.AdminLoginResponse{}, dto.ErrGetPermissionsByRoleID
	}

	accessToken, refreshToken, err := as.jwtService.GenerateToken(user.ID.String(), user.RoleID.String(), endpoints)
	if err != nil {
		return dto.AdminLoginResponse{}, dto.ErrGenerateToken
	}

	return dto.AdminLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (as *AdminService) RefreshToken(ctx context.Context, req dto.RefreshTokenRequest) (dto.RefreshTokenResponse, error) {
	_, err := as.jwtService.ValidateToken(req.RefreshToken)
	if err != nil {
		return dto.RefreshTokenResponse{}, dto.ErrValidateToken
	}

	userID, err := as.jwtService.GetUserIDByToken(req.RefreshToken)
	if err != nil {
		return dto.RefreshTokenResponse{}, dto.ErrGetUserIDFromToken
	}

	roleID, err := as.jwtService.GetRoleIDByToken(req.RefreshToken)
	if err != nil {
		return dto.RefreshTokenResponse{}, dto.ErrGetRoleFromToken
	}

	role, err := as.adminRepo.GetRoleByID(ctx, nil, roleID)
	if err != nil {
		return dto.RefreshTokenResponse{}, dto.ErrGetRoleFromID
	}

	if role.Name != "admin" {
		return dto.RefreshTokenResponse{}, dto.ErrDeniedAccess
	}

	permissions, err := as.adminRepo.GetPermissionsByRoleID(ctx, nil, roleID)
	if err != nil {
		return dto.RefreshTokenResponse{}, dto.ErrGetPermissionsByRoleID
	}

	accessToken, _, err := as.jwtService.GenerateToken(userID, roleID, permissions)
	if err != nil {
		return dto.RefreshTokenResponse{}, dto.ErrGenerateAccessToken
	}

	return dto.RefreshTokenResponse{AccessToken: accessToken}, nil
}
