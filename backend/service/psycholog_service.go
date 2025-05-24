package service

import (
	"context"
	"log"

	"github.com/Amierza/warasin-mobile/backend/dto"
	"github.com/Amierza/warasin-mobile/backend/helpers"
	"github.com/Amierza/warasin-mobile/backend/repository"
)

type (
	IPsychologService interface {
		// Authentication
		Login(ctx context.Context, req dto.PsychologLoginRequest) (dto.PsychologLoginResponse, error)
		RefreshToken(ctx context.Context, req dto.RefreshTokenRequest) (dto.RefreshTokenResponse, error)

		// Psycholog
		GetDetailPsycholog(ctx context.Context) (dto.PsychologResponse, error)
	}

	PsychologService struct {
		psychologRepo repository.IPsychologRepository
		masterRepo    repository.IMasterRepository
		jwtService    IJWTService
	}
)

func NewPsychologService(psychologRepo repository.IPsychologRepository, masterRepo repository.IMasterRepository, jwtService IJWTService) *PsychologService {
	return &PsychologService{
		psychologRepo: psychologRepo,
		masterRepo:    masterRepo,
		jwtService:    jwtService,
	}
}

// Authentication
func (ps *PsychologService) Login(ctx context.Context, req dto.PsychologLoginRequest) (dto.PsychologLoginResponse, error) {
	if !helpers.IsValidEmail(req.Email) {
		return dto.PsychologLoginResponse{}, dto.ErrInvalidEmail
	}

	if len(req.Password) < 8 {
		return dto.PsychologLoginResponse{}, dto.ErrInvalidPassword
	}

	psycholog, flag, err := ps.psychologRepo.CheckEmail(ctx, nil, req.Email)
	if !flag || err != nil {
		return dto.PsychologLoginResponse{}, dto.ErrEmailNotFound
	}

	if psycholog.Role.Name != "psycholog" {
		return dto.PsychologLoginResponse{}, dto.ErrDeniedAccess
	}

	checkPassword, err := helpers.CheckPassword(psycholog.Password, []byte(req.Password))
	if err != nil || !checkPassword {
		return dto.PsychologLoginResponse{}, dto.ErrPasswordNotMatch
	}

	permissions, err := ps.psychologRepo.GetPermissionsByRoleID(ctx, nil, psycholog.RoleID.String())
	if err != nil {
		return dto.PsychologLoginResponse{}, dto.ErrGetPermissionsByRoleID
	}

	accessToken, refreshToken, err := ps.jwtService.GenerateToken(psycholog.ID.String(), psycholog.RoleID.String(), permissions)
	if err != nil {
		return dto.PsychologLoginResponse{}, err
	}

	return dto.PsychologLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
func (ps *PsychologService) RefreshToken(ctx context.Context, req dto.RefreshTokenRequest) (dto.RefreshTokenResponse, error) {
	_, err := ps.jwtService.ValidateToken(req.RefreshToken)

	if err != nil {
		return dto.RefreshTokenResponse{}, dto.ErrValidateToken
	}

	userID, err := ps.jwtService.GetUserIDByToken(req.RefreshToken)
	if err != nil {
		return dto.RefreshTokenResponse{}, dto.ErrGetUserIDFromToken
	}

	roleID, err := ps.jwtService.GetRoleIDByToken(req.RefreshToken)
	if err != nil {
		return dto.RefreshTokenResponse{}, dto.ErrGetRoleFromToken
	}

	role, err := ps.psychologRepo.GetRoleByID(ctx, nil, roleID)
	if err != nil {
		return dto.RefreshTokenResponse{}, dto.ErrGetRoleFromID
	}

	log.Println(role.Name)

	if role.Name != "psycholog" {
		return dto.RefreshTokenResponse{}, dto.ErrDeniedAccess
	}

	endpoints, err := ps.psychologRepo.GetPermissionsByRoleID(ctx, nil, roleID)
	if err != nil {
		return dto.RefreshTokenResponse{}, dto.ErrGetPermissionsByRoleID
	}

	accessToken, _, err := ps.jwtService.GenerateToken(userID, roleID, endpoints)
	if err != nil {
		return dto.RefreshTokenResponse{}, dto.ErrGenerateAccessToken
	}

	return dto.RefreshTokenResponse{AccessToken: accessToken}, nil
}

// Psycholog
func (ps *PsychologService) GetDetailPsycholog(ctx context.Context) (dto.PsychologResponse, error) {
	token := ctx.Value("Authorization").(string)

	psychologId, err := ps.jwtService.GetUserIDByToken(token)
	if err != nil {
		return dto.PsychologResponse{}, dto.ErrGetPsychologIDFromToken
	}

	psycholog, err := ps.psychologRepo.GetPsychologByID(ctx, nil, psychologId)
	if err != nil {
		return dto.PsychologResponse{}, dto.ErrUserNotFound
	}

	return dto.PsychologResponse{
		ID:          psycholog.ID,
		Name:        psycholog.Name,
		STRNumber:   psycholog.STRNumber,
		Email:       psycholog.Email,
		Password:    psycholog.Password,
		WorkYear:    psycholog.WorkYear,
		Description: psycholog.Description,
		PhoneNumber: psycholog.PhoneNumber,
		Image:       psycholog.Image,
		City: dto.CityResponse{
			ID:   psycholog.CityID,
			Name: psycholog.City.Name,
			Type: psycholog.City.Type,
			Province: dto.ProvinceResponse{
				ID:   psycholog.City.ProvinceID,
				Name: psycholog.City.Province.Name,
			},
		},
		Role: dto.RoleResponse{
			ID:   psycholog.RoleID,
			Name: psycholog.Role.Name,
		},
	}, nil
}
