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

		// Consultation
		GetAllConsultationWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.ConsultationPaginationResponseForPsycholog, error)
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

	psycholog, flag, err := ps.masterRepo.GetPsychologByEmail(ctx, nil, req.Email)
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

// Consultation
func (ps *PsychologService) GetAllConsultationWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.ConsultationPaginationResponseForPsycholog, error) {
	token := ctx.Value("Authorization").(string)

	psyID, err := ps.jwtService.GetUserIDByToken(token)
	if err != nil {
		return dto.ConsultationPaginationResponseForPsycholog{}, dto.ErrGetPsychologFromID
	}

	dataWithPaginate, err := ps.psychologRepo.GetAllConsultationWithPagination(ctx, nil, req, psyID)
	if err != nil {
		return dto.ConsultationPaginationResponseForPsycholog{}, dto.ErrGetAllConsultationWithPagination
	}

	var (
		psycholog     dto.PsychologResponse
		consultations []dto.ConsultationResponseForPsycholog
	)

	psycholog = dto.PsychologResponse{
		ID:          dataWithPaginate.Consultations[0].Psycholog.ID,
		Name:        dataWithPaginate.Consultations[0].Psycholog.Name,
		STRNumber:   dataWithPaginate.Consultations[0].Psycholog.STRNumber,
		Email:       dataWithPaginate.Consultations[0].Psycholog.Email,
		Password:    dataWithPaginate.Consultations[0].Psycholog.Password,
		WorkYear:    dataWithPaginate.Consultations[0].Psycholog.WorkYear,
		Description: dataWithPaginate.Consultations[0].Psycholog.Description,
		PhoneNumber: dataWithPaginate.Consultations[0].Psycholog.PhoneNumber,
		Image:       dataWithPaginate.Consultations[0].Psycholog.Image,
		City: dto.CityResponse{
			ID:   dataWithPaginate.Consultations[0].Psycholog.CityID,
			Name: dataWithPaginate.Consultations[0].Psycholog.City.Name,
			Type: dataWithPaginate.Consultations[0].Psycholog.City.Type,
			Province: dto.ProvinceResponse{
				ID:   dataWithPaginate.Consultations[0].Psycholog.City.ProvinceID,
				Name: dataWithPaginate.Consultations[0].Psycholog.City.Province.Name,
			},
		},
		Role: dto.RoleResponse{
			ID:   dataWithPaginate.Consultations[0].Psycholog.RoleID,
			Name: dataWithPaginate.Consultations[0].Psycholog.Role.Name,
		},
	}

	for _, consultation := range dataWithPaginate.Consultations {
		consultations = append(consultations, dto.ConsultationResponseForPsycholog{
			ID:      consultation.ID,
			Date:    consultation.Date.String(),
			Rate:    consultation.Rate,
			Comment: consultation.Comment,
			User: dto.AllUserResponse{
				ID:          consultation.User.ID,
				Name:        consultation.User.Name,
				Email:       consultation.User.Email,
				Password:    consultation.User.Password,
				Birthdate:   consultation.User.Birthdate.String(),
				PhoneNumber: consultation.User.PhoneNumber,
				Data01:      consultation.User.Data01,
				Data02:      consultation.User.Data02,
				Data03:      consultation.User.Data03,
				IsVerified:  consultation.User.IsVerified,
				City: dto.CityResponse{
					ID:   &consultation.User.City.ID,
					Name: consultation.User.City.Name,
					Type: consultation.User.City.Type,
					Province: dto.ProvinceResponse{
						ID:   consultation.User.City.ProvinceID,
						Name: consultation.User.City.Province.Name,
					},
				},
				Role: dto.RoleResponse{
					ID:   &consultation.User.Role.ID,
					Name: consultation.User.Role.Name,
				},
			},
		})
	}

	datas := dto.AllConsultationResponseForPsycholog{
		Psycholog:    psycholog,
		Consultation: consultations,
	}

	return dto.ConsultationPaginationResponseForPsycholog{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}
