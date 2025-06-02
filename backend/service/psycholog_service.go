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

		// Practice
		GetAllPractice(ctx context.Context) (dto.AllPracticeResponse, error)

		// Available Slot
		GetAllAvailableSlot(ctx context.Context) (dto.AllAvailableSlotResponse, error)

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

// Practice
func (ps *PsychologService) GetAllPractice(ctx context.Context) (dto.AllPracticeResponse, error) {
	token := ctx.Value("Authorization").(string)

	psyID, err := ps.jwtService.GetUserIDByToken(token)
	if err != nil {
		return dto.AllPracticeResponse{}, dto.ErrGetPsychologFromID
	}

	datas, err := ps.psychologRepo.GetAllPractice(ctx, nil, psyID)
	if err != nil {
		return dto.AllPracticeResponse{}, dto.ErrGetAllPractice
	}

	psycholog := dto.PsychologResponse{
		ID:          datas.Practices[0].Psycholog.ID,
		Name:        datas.Practices[0].Psycholog.Name,
		STRNumber:   datas.Practices[0].Psycholog.STRNumber,
		Email:       datas.Practices[0].Psycholog.Email,
		Password:    datas.Practices[0].Psycholog.Password,
		WorkYear:    datas.Practices[0].Psycholog.WorkYear,
		Description: datas.Practices[0].Psycholog.Description,
		PhoneNumber: datas.Practices[0].Psycholog.PhoneNumber,
		Image:       datas.Practices[0].Psycholog.Image,
		City: dto.CityResponse{
			ID:   datas.Practices[0].Psycholog.CityID,
			Name: datas.Practices[0].Psycholog.City.Name,
			Type: datas.Practices[0].Psycholog.City.Type,
			Province: dto.ProvinceResponse{
				ID:   datas.Practices[0].Psycholog.City.ProvinceID,
				Name: datas.Practices[0].Psycholog.City.Province.Name,
			},
		},
		Role: dto.RoleResponse{
			ID:   datas.Practices[0].Psycholog.RoleID,
			Name: datas.Practices[0].Psycholog.Role.Name,
		},
	}

	var practices []dto.PracticeResponse
	for _, practice := range datas.Practices {
		data := dto.PracticeResponse{
			ID:          practice.ID,
			Type:        practice.Type,
			Name:        practice.Name,
			Address:     practice.Address,
			PhoneNumber: practice.PhoneNumber,
		}

		// PracticeSchedule
		for _, pracSche := range practice.PracticeSchedules {
			data.PracticeSchedules = append(data.PracticeSchedules, dto.PracticeScheduleResponse{
				ID:    pracSche.ID,
				Day:   pracSche.Day,
				Open:  pracSche.Open,
				Close: pracSche.Close,
			})
		}

		practices = append(practices, data)
	}

	return dto.AllPracticeResponse{
		Psycholog: psycholog,
		Practices: practices,
	}, nil
}

// Available Slot
func (ps *PsychologService) GetAllAvailableSlot(ctx context.Context) (dto.AllAvailableSlotResponse, error) {
	token := ctx.Value("Authorization").(string)

	psyID, err := ps.jwtService.GetUserIDByToken(token)
	if err != nil {
		return dto.AllAvailableSlotResponse{}, dto.ErrGetPsychologFromID
	}

	datas, err := ps.psychologRepo.GetAllAvailableSlot(ctx, nil, psyID)
	if err != nil {
		return dto.AllAvailableSlotResponse{}, dto.ErrGetAllAvailableSlot
	}

	psycholog := dto.PsychologResponse{
		ID:          datas.AvailableSlots[0].Psycholog.ID,
		Name:        datas.AvailableSlots[0].Psycholog.Name,
		STRNumber:   datas.AvailableSlots[0].Psycholog.STRNumber,
		Email:       datas.AvailableSlots[0].Psycholog.Email,
		Password:    datas.AvailableSlots[0].Psycholog.Password,
		WorkYear:    datas.AvailableSlots[0].Psycholog.WorkYear,
		Description: datas.AvailableSlots[0].Psycholog.Description,
		PhoneNumber: datas.AvailableSlots[0].Psycholog.PhoneNumber,
		Image:       datas.AvailableSlots[0].Psycholog.Image,
		City: dto.CityResponse{
			ID:   datas.AvailableSlots[0].Psycholog.CityID,
			Name: datas.AvailableSlots[0].Psycholog.City.Name,
			Type: datas.AvailableSlots[0].Psycholog.City.Type,
			Province: dto.ProvinceResponse{
				ID:   datas.AvailableSlots[0].Psycholog.City.ProvinceID,
				Name: datas.AvailableSlots[0].Psycholog.City.Province.Name,
			},
		},
		Role: dto.RoleResponse{
			ID:   datas.AvailableSlots[0].Psycholog.RoleID,
			Name: datas.AvailableSlots[0].Psycholog.Role.Name,
		},
	}

	var availableSlots []dto.AvailableSlotResponse
	for _, availableSlot := range datas.AvailableSlots {
		data := dto.AvailableSlotResponse{
			ID:       availableSlot.ID,
			Start:    availableSlot.Start,
			End:      availableSlot.End,
			IsBooked: availableSlot.IsBooked,
		}

		availableSlots = append(availableSlots, data)
	}

	return dto.AllAvailableSlotResponse{
		Psycholog:      psycholog,
		AvailableSlots: availableSlots,
	}, nil
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
		ID:          dataWithPaginate.Consultations[0].AvailableSlot.Psycholog.ID,
		Name:        dataWithPaginate.Consultations[0].AvailableSlot.Psycholog.Name,
		STRNumber:   dataWithPaginate.Consultations[0].AvailableSlot.Psycholog.STRNumber,
		Email:       dataWithPaginate.Consultations[0].AvailableSlot.Psycholog.Email,
		Password:    dataWithPaginate.Consultations[0].AvailableSlot.Psycholog.Password,
		WorkYear:    dataWithPaginate.Consultations[0].AvailableSlot.Psycholog.WorkYear,
		Description: dataWithPaginate.Consultations[0].AvailableSlot.Psycholog.Description,
		PhoneNumber: dataWithPaginate.Consultations[0].AvailableSlot.Psycholog.PhoneNumber,
		Image:       dataWithPaginate.Consultations[0].AvailableSlot.Psycholog.Image,
		City: dto.CityResponse{
			ID:   dataWithPaginate.Consultations[0].AvailableSlot.Psycholog.CityID,
			Name: dataWithPaginate.Consultations[0].AvailableSlot.Psycholog.City.Name,
			Type: dataWithPaginate.Consultations[0].AvailableSlot.Psycholog.City.Type,
			Province: dto.ProvinceResponse{
				ID:   dataWithPaginate.Consultations[0].AvailableSlot.Psycholog.City.ProvinceID,
				Name: dataWithPaginate.Consultations[0].AvailableSlot.Psycholog.City.Province.Name,
			},
		},
		Role: dto.RoleResponse{
			ID:   dataWithPaginate.Consultations[0].AvailableSlot.Psycholog.RoleID,
			Name: dataWithPaginate.Consultations[0].AvailableSlot.Psycholog.Role.Name,
		},
	}

	for _, consultation := range dataWithPaginate.Consultations {
		data := dto.ConsultationResponseForPsycholog{
			ID:      consultation.ID,
			Rate:    consultation.Rate,
			Comment: consultation.Comment,
			Status:  consultation.Status,
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
			AvailableSlot: dto.AvailableSlotResponse{
				ID:       consultation.AvailableSlot.ID,
				Start:    consultation.AvailableSlot.Start,
				End:      consultation.AvailableSlot.End,
				IsBooked: consultation.AvailableSlot.IsBooked,
			},
			Practice: dto.PracticeResponse{
				ID:          consultation.Practice.ID,
				Type:        consultation.Practice.Type,
				Name:        psycholog.Name,
				Address:     consultation.Practice.Address,
				PhoneNumber: consultation.Practice.PhoneNumber,
			},
		}

		consultations = append(consultations, data)
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
