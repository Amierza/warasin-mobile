package service

import (
	"context"

	"github.com/Amierza/warasin-mobile/backend/dto"
	"github.com/Amierza/warasin-mobile/backend/helpers"
	"github.com/Amierza/warasin-mobile/backend/repository"
)

type (
	IMasterService interface {
		// Province & City
		GetAllProvince(ctx context.Context) (dto.ProvincesResponse, error)
		GetAllCity(ctx context.Context, req dto.CityQueryRequest) (dto.CitiesResponse, error)

		// Psycholog
		GetDetailPsycholog(ctx context.Context, psychologID string) (dto.PsychologResponse, error)
		UpdatePsycholog(ctx context.Context, req dto.UpdatePsychologRequest) (dto.PsychologResponse, error)
	}

	MasterService struct {
		masterRepo repository.IMasterRepository
		jwtService IJWTService
	}
)

func NewMasterService(masterRepo repository.IMasterRepository, jwtService IJWTService) *MasterService {
	return &MasterService{
		masterRepo: masterRepo,
		jwtService: jwtService,
	}
}

// Get Province & City
func (ms *MasterService) GetAllProvince(ctx context.Context) (dto.ProvincesResponse, error) {
	data, err := ms.masterRepo.GetAllProvince(ctx, nil)
	if err != nil {
		return dto.ProvincesResponse{}, dto.ErrGetAllProvince
	}

	var datas []dto.ProvinceResponse
	for _, province := range data.Provinces {
		data := dto.ProvinceResponse{
			ID:   &province.ID,
			Name: province.Name,
		}

		datas = append(datas, data)
	}

	return dto.ProvincesResponse{
		Data: datas,
	}, nil
}
func (ms *MasterService) GetAllCity(ctx context.Context, req dto.CityQueryRequest) (dto.CitiesResponse, error) {
	data, err := ms.masterRepo.GetAllCity(ctx, nil, req)
	if err != nil {
		return dto.CitiesResponse{}, dto.ErrGetAllProvince
	}

	var datas []dto.CityResponseCustom
	for _, city := range data.Cities {
		data := dto.CityResponseCustom{
			ID:   &city.ID,
			Name: city.Name,
			Type: city.Type,
		}

		datas = append(datas, data)
	}

	return dto.CitiesResponse{
		Data: datas,
	}, nil
}

// Psycholog
func (ms *MasterService) GetDetailPsycholog(ctx context.Context, psychologID string) (dto.PsychologResponse, error) {
	token := ctx.Value("Authorization").(string)

	roleID, err := ms.jwtService.GetRoleIDByToken(token)
	if err != nil {
		return dto.PsychologResponse{}, dto.ErrGetRoleIDFromToken
	}

	role, err := ms.masterRepo.GetRoleByID(ctx, nil, roleID)
	if err != nil {
		return dto.PsychologResponse{}, dto.ErrGetRoleFromID
	}

	if role.Name != "admin" {
		psychologId, err := ms.jwtService.GetUserIDByToken(token)
		if err != nil {
			return dto.PsychologResponse{}, dto.ErrGetPsychologIDFromToken
		}

		psychologID = psychologId
	}

	psycholog, flag, err := ms.masterRepo.GetPsychologByID(ctx, nil, psychologID)
	if err != nil || !flag {
		return dto.PsychologResponse{}, dto.ErrPsychologNotFound
	}

	data := dto.PsychologResponse{
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
	}

	// LanguageMasters
	for _, lang := range psycholog.PsychologLanguages {
		data.LanguageMasters = append(data.LanguageMasters, dto.LanguageMasterResponse{
			ID:   &lang.LanguageMaster.ID,
			Name: lang.LanguageMaster.Name,
		})
	}

	// Specializations
	for _, spec := range psycholog.PsychologSpecializations {
		data.Specializations = append(data.Specializations, dto.SpecializationResponse{
			ID:          &spec.Specialization.ID,
			Name:        spec.Specialization.Name,
			Description: spec.Specialization.Description,
		})
	}

	// Educations
	for _, edu := range psycholog.Educations {
		data.Educations = append(data.Educations, dto.EducationResponse{
			ID:             &edu.ID,
			Degree:         edu.Degree,
			Major:          edu.Major,
			Institution:    edu.Institution,
			GraduationYear: edu.GraduationYear,
		})
	}

	return data, nil
}
func (ms *MasterService) UpdatePsycholog(ctx context.Context, req dto.UpdatePsychologRequest) (dto.PsychologResponse, error) {
	token := ctx.Value("Authorization").(string)

	roleID, err := ms.jwtService.GetRoleIDByToken(token)
	if err != nil {
		return dto.PsychologResponse{}, dto.ErrGetRoleIDFromToken
	}

	role, err := ms.masterRepo.GetRoleByID(ctx, nil, roleID)
	if err != nil {
		return dto.PsychologResponse{}, dto.ErrGetRoleFromID
	}

	if role.Name != "admin" {
		psychologId, err := ms.jwtService.GetUserIDByToken(token)
		if err != nil {
			return dto.PsychologResponse{}, dto.ErrGetPsychologIDFromToken
		}

		req.ID = psychologId
	}

	psycholog, flag, err := ms.masterRepo.GetPsychologByID(ctx, nil, req.ID)
	if err != nil || !flag {
		return dto.PsychologResponse{}, dto.ErrPsychologNotFound
	}

	if req.CityID != "" {
		city, err := ms.masterRepo.GetCityByID(ctx, nil, req.CityID)
		if err != nil {
			return dto.PsychologResponse{}, dto.ErrGetCityByID
		}

		psycholog.City = city
	}

	if req.Name != "" {
		if len(req.Name) < 5 {
			return dto.PsychologResponse{}, dto.ErrInvalidName
		}

		psycholog.Name = req.Name
	}

	if req.STRNumber != "" {
		if !helpers.IsValidSTRNumber(req.STRNumber) {
			return dto.PsychologResponse{}, dto.ErrInvalidSTRNumber
		}

		psycholog.STRNumber = req.STRNumber
	}

	if req.Email != "" {
		if !helpers.IsValidEmail(req.Email) {
			return dto.PsychologResponse{}, dto.ErrInvalidEmail
		}

		_, flag, err := ms.masterRepo.GetPsychologByEmail(ctx, nil, req.Email)
		if flag || err == nil {
			return dto.PsychologResponse{}, dto.ErrEmailAlreadyExists
		}

		psycholog.Email = req.Email
	}

	if req.WorkYear != "" {
		if len(req.WorkYear) < 4 {
			return dto.PsychologResponse{}, dto.ErrInvalidWorkYear
		}

		psycholog.WorkYear = req.WorkYear
	}

	if req.Description != "" {
		psycholog.Description = req.Description
	}

	if req.PhoneNumber != "" {
		phoneNumberFormatted, err := helpers.StandardizePhoneNumber(req.PhoneNumber)
		if err != nil {
			return dto.PsychologResponse{}, dto.ErrFormatPhoneNumber
		}

		psycholog.PhoneNumber = phoneNumberFormatted
	}

	if req.Image != "" {
		psycholog.Image = req.Image
	}

	err = ms.masterRepo.UpdatePsycholog(ctx, nil, psycholog)
	if err != nil {
		return dto.PsychologResponse{}, dto.ErrUpdatePsycholog
	}

	res := dto.PsychologResponse{
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
	}

	return res, nil
}
