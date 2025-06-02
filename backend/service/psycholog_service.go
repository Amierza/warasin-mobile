package service

import (
	"context"

	"github.com/Amierza/warasin-mobile/backend/dto"
	"github.com/Amierza/warasin-mobile/backend/entity"
	"github.com/Amierza/warasin-mobile/backend/helpers"
	"github.com/Amierza/warasin-mobile/backend/repository"
	"github.com/google/uuid"
)

type (
	IPsychologService interface {
		// Authentication
		Login(ctx context.Context, req dto.PsychologLoginRequest) (dto.PsychologLoginResponse, error)
		RefreshToken(ctx context.Context, req dto.RefreshTokenRequest) (dto.RefreshTokenResponse, error)

		// Practice
		CreatePractice(ctx context.Context, req dto.CreatePracticeRequest) (dto.PracticeResponse, error)
		GetAllPractice(ctx context.Context) (dto.AllPracticeResponse, error)
		UpdatePractice(ctx context.Context, req dto.UpdatePracticeRequest, practiceID string) (dto.PracticeResponse, error)
		DeletePractice(ctx context.Context, practiceID string) (dto.PracticeResponse, error)

		// Available Slot
		GetAllAvailableSlot(ctx context.Context) (dto.AllAvailableSlotResponse, error)

		// Consultation
		GetAllConsultationWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.ConsultationPaginationResponse, error)
		UpdateConsultation(ctx context.Context, req dto.UpdateConsultationRequest, consulID string) (dto.ConsultationResponse, error)
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

	permissions, _, err := ps.psychologRepo.GetPermissionsByRoleID(ctx, nil, psycholog.RoleID.String())
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

	role, _, err := ps.psychologRepo.GetRoleByID(ctx, nil, roleID)
	if err != nil {
		return dto.RefreshTokenResponse{}, dto.ErrGetRoleFromID
	}

	if role.Name != "psycholog" {
		return dto.RefreshTokenResponse{}, dto.ErrDeniedAccess
	}

	endpoints, _, err := ps.psychologRepo.GetPermissionsByRoleID(ctx, nil, roleID)
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
func (ps *PsychologService) CreatePractice(ctx context.Context, req dto.CreatePracticeRequest) (dto.PracticeResponse, error) {
	token := ctx.Value("Authorization").(string)

	psyID, err := ps.jwtService.GetUserIDByToken(token)
	if err != nil {
		return dto.PracticeResponse{}, dto.ErrGetPsychologIDFromToken
	}

	psychologID, err := uuid.Parse(psyID)
	if err != nil {
		return dto.PracticeResponse{}, dto.ErrParseUUID
	}

	if len(req.Name) < 5 {
		return dto.PracticeResponse{}, dto.ErrInvalidPracticeName
	}

	phoneNumberFormatted, err := helpers.StandardizePhoneNumber(req.PhoneNumber, false)
	if err != nil {
		return dto.PracticeResponse{}, dto.ErrFormatPhoneNumber
	}

	practice := entity.Practice{
		ID:          uuid.New(),
		Type:        req.Type,
		Name:        req.Name,
		Address:     req.Address,
		PhoneNumber: phoneNumberFormatted,
		PsychologID: &psychologID,
	}

	err = ps.psychologRepo.CreatePractice(ctx, nil, practice)
	if err != nil {
		return dto.PracticeResponse{}, dto.ErrCreatePractice
	}

	var schedules []entity.PracticeSchedule
	switch req.Type {
	case "Konsultasi Online":
		days := []string{"Thursday", "Friday", "Saturday"}
		for _, day := range days {
			schedules = append(schedules, entity.PracticeSchedule{
				ID:         uuid.New(),
				Day:        day,
				Open:       "07:00",
				Close:      "18:00",
				PracticeID: &practice.ID,
			})
		}
	case "Praktek Klinik":
		days := []string{"Monday", "Tuesday", "Wednesday"}
		for _, day := range days {
			schedules = append(schedules, entity.PracticeSchedule{
				ID:         uuid.New(),
				Day:        day,
				Open:       "07:00",
				Close:      "18:00",
				PracticeID: &practice.ID,
			})
		}
	default:
		return dto.PracticeResponse{}, dto.ErrAddPracticeSchedule
	}

	var practiceSchedules []dto.PracticeScheduleResponse
	for _, schedule := range schedules {
		practiceSchedules = append(practiceSchedules, dto.PracticeScheduleResponse{
			ID:    schedule.ID,
			Day:   schedule.Day,
			Open:  schedule.Open,
			Close: schedule.Close,
		})
	}

	err = ps.psychologRepo.CreatePracticeSchedule(ctx, nil, schedules)
	if err != nil {
		return dto.PracticeResponse{}, dto.ErrCreatePracticeSchedule
	}

	return dto.PracticeResponse{
		ID:                practice.ID,
		Type:              practice.Type,
		Name:              practice.Name,
		Address:           practice.Address,
		PhoneNumber:       phoneNumberFormatted,
		PracticeSchedules: practiceSchedules,
	}, nil
}
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
func (ps *PsychologService) UpdatePractice(ctx context.Context, req dto.UpdatePracticeRequest, practiceID string) (dto.PracticeResponse, error) {
	prac, flag, err := ps.psychologRepo.GetPracticeByID(ctx, nil, practiceID)
	if err != nil || !flag {
		return dto.PracticeResponse{}, dto.ErrPracticeNotFound
	}

	if req.Type != "" {
		err = ps.psychologRepo.DeletePracticeSchedule(ctx, nil, practiceID)
		if err != nil {
			return dto.PracticeResponse{}, dto.ErrDeletePracticeSchedules
		}

		var schedules []entity.PracticeSchedule
		switch req.Type {
		case "Konsultasi Online":
			days := []string{"Thursday", "Friday", "Saturday"}
			for _, day := range days {
				schedules = append(schedules, entity.PracticeSchedule{
					ID:         uuid.New(),
					Day:        day,
					Open:       "07:00",
					Close:      "18:00",
					PracticeID: &prac.ID,
				})
			}
			prac.Type = req.Type
		case "Praktek Klinik":
			days := []string{"Monday", "Tuesday", "Wednesday"}
			for _, day := range days {
				schedules = append(schedules, entity.PracticeSchedule{
					ID:         uuid.New(),
					Day:        day,
					Open:       "07:00",
					Close:      "18:00",
					PracticeID: &prac.ID,
				})
			}
			prac.Type = req.Type
		default:
			return dto.PracticeResponse{}, dto.ErrAddPracticeSchedule
		}

		err = ps.psychologRepo.CreatePracticeSchedule(ctx, nil, schedules)
		if err != nil {
			return dto.PracticeResponse{}, dto.ErrCreatePracticeSchedule
		}

		prac.PracticeSchedules = schedules
	}

	if req.Name != "" {
		if len(req.Name) < 5 {
			return dto.PracticeResponse{}, dto.ErrInvalidPracticeName
		}

		prac.Name = req.Name
	}

	if req.Address != "" {
		prac.Address = req.Address
	}

	if req.PhoneNumber != "" {
		phoneNumberFormatted, err := helpers.StandardizePhoneNumber(req.PhoneNumber, false)
		if err != nil {
			return dto.PracticeResponse{}, dto.ErrFormatPhoneNumber
		}

		prac.PhoneNumber = phoneNumberFormatted
	}

	err = ps.psychologRepo.UpdatePractice(ctx, nil, prac)
	if err != nil {
		return dto.PracticeResponse{}, dto.ErrUpdatePractice
	}

	var practiceSchedules []dto.PracticeScheduleResponse
	for _, sch := range prac.PracticeSchedules {
		practiceSchedules = append(practiceSchedules, dto.PracticeScheduleResponse{
			ID:    sch.ID,
			Day:   sch.Day,
			Open:  sch.Open,
			Close: sch.Close,
		})
	}

	return dto.PracticeResponse{
		ID:                prac.ID,
		Type:              prac.Type,
		Name:              prac.Name,
		Address:           prac.Address,
		PhoneNumber:       prac.PhoneNumber,
		PracticeSchedules: practiceSchedules,
	}, nil
}
func (ps *PsychologService) DeletePractice(ctx context.Context, practiceID string) (dto.PracticeResponse, error) {
	deletedPractice, flag, err := ps.psychologRepo.GetPracticeByID(ctx, nil, practiceID)
	if err != nil || !flag {
		return dto.PracticeResponse{}, dto.ErrPracticeNotFound
	}

	err = ps.psychologRepo.DeletePracticeSchedule(ctx, nil, practiceID)
	if err != nil {
		return dto.PracticeResponse{}, dto.ErrDeletePracticeSchedules
	}

	err = ps.psychologRepo.DeletePracticeByID(ctx, nil, practiceID)
	if err != nil {
		return dto.PracticeResponse{}, dto.ErrDeletePractice
	}

	var practiceSchedules []dto.PracticeScheduleResponse
	for _, sch := range deletedPractice.PracticeSchedules {
		practiceSchedules = append(practiceSchedules, dto.PracticeScheduleResponse{
			ID:    sch.ID,
			Day:   sch.Day,
			Open:  sch.Open,
			Close: sch.Close,
		})
	}

	res := dto.PracticeResponse{
		ID:                deletedPractice.ID,
		Type:              deletedPractice.Type,
		Name:              deletedPractice.Name,
		Address:           deletedPractice.Address,
		PhoneNumber:       deletedPractice.PhoneNumber,
		PracticeSchedules: practiceSchedules,
	}

	return res, nil
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
func (ps *PsychologService) GetAllConsultationWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.ConsultationPaginationResponse, error) {
	token := ctx.Value("Authorization").(string)

	psyID, err := ps.jwtService.GetUserIDByToken(token)
	if err != nil {
		return dto.ConsultationPaginationResponse{}, dto.ErrGetPsychologFromID
	}

	dataWithPaginate, err := ps.psychologRepo.GetAllConsultationWithPagination(ctx, nil, req, psyID)
	if err != nil {
		return dto.ConsultationPaginationResponse{}, dto.ErrGetAllConsultationWithPagination
	}

	var (
		psycholog     dto.PsychologResponse
		consultations []dto.ConsultationResponse
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

	// LanguageMasters
	for _, lang := range dataWithPaginate.Consultations[0].AvailableSlot.Psycholog.PsychologLanguages {
		psycholog.LanguageMasters = append(psycholog.LanguageMasters, dto.LanguageMasterResponse{
			ID:   &lang.LanguageMaster.ID,
			Name: lang.LanguageMaster.Name,
		})
	}

	// Specializations
	for _, spec := range dataWithPaginate.Consultations[0].AvailableSlot.Psycholog.PsychologSpecializations {
		psycholog.Specializations = append(psycholog.Specializations, dto.SpecializationResponse{
			ID:          &spec.Specialization.ID,
			Name:        spec.Specialization.Name,
			Description: spec.Specialization.Description,
		})
	}

	// Educations
	for _, edu := range dataWithPaginate.Consultations[0].AvailableSlot.Psycholog.Educations {
		psycholog.Educations = append(psycholog.Educations, dto.EducationResponse{
			ID:             &edu.ID,
			Degree:         edu.Degree,
			Major:          edu.Major,
			Institution:    edu.Institution,
			GraduationYear: edu.GraduationYear,
		})
	}

	for _, consultation := range dataWithPaginate.Consultations {
		dayName, err := helpers.GetDayName(consultation.Date)
		if err != nil {
			return dto.ConsultationPaginationResponse{}, dto.ErrParseConsultationDate
		}

		var practiceSchedules []dto.PracticeScheduleResponse
		for _, pracSch := range consultation.Practice.PracticeSchedules {
			if dayName == pracSch.Day {
				practiceSchedules = append(practiceSchedules, dto.PracticeScheduleResponse{
					ID:    pracSch.ID,
					Day:   pracSch.Day,
					Open:  pracSch.Open,
					Close: pracSch.Close,
				})
			}
		}

		data := dto.ConsultationResponse{
			ID:      consultation.ID,
			Date:    consultation.Date,
			Rate:    consultation.Rate,
			Comment: consultation.Comment,
			Status:  consultation.Status,
			User: dto.AllUserResponse{
				ID:          consultation.User.ID,
				Name:        consultation.User.Name,
				Email:       consultation.User.Email,
				Password:    consultation.User.Password,
				Birthdate:   consultation.User.Birthdate,
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
				ID:                consultation.Practice.ID,
				Type:              consultation.Practice.Type,
				Name:              consultation.Practice.Name,
				Address:           consultation.Practice.Address,
				PhoneNumber:       consultation.Practice.PhoneNumber,
				PracticeSchedules: practiceSchedules,
			},
		}

		consultations = append(consultations, data)
	}

	datas := dto.AllConsultationResponse{
		Psycholog:    psycholog,
		Consultation: consultations,
	}

	return dto.ConsultationPaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}
func (ps *PsychologService) UpdateConsultation(ctx context.Context, req dto.UpdateConsultationRequest, consulID string) (dto.ConsultationResponse, error) {
	consul, flag, err := ps.psychologRepo.GetConsultationByID(ctx, nil, consulID)
	if err != nil || !flag {
		return dto.ConsultationResponse{}, dto.ErrConsultationNotFound
	}

	if req.Status != nil {
		consul.Status = *req.Status
	}

	dayName, err := helpers.GetDayName(consul.Date)
	if err != nil {
		return dto.ConsultationResponse{}, dto.ErrParseConsultationDate
	}

	var practiceSchedules []dto.PracticeScheduleResponse
	for _, pracSch := range consul.Practice.PracticeSchedules {
		if dayName == pracSch.Day {
			practiceSchedules = append(practiceSchedules, dto.PracticeScheduleResponse{
				ID:    pracSch.ID,
				Day:   pracSch.Day,
				Open:  pracSch.Open,
				Close: pracSch.Close,
			})
		}
	}

	data := dto.ConsultationResponse{
		ID:      consul.ID,
		Date:    consul.Date,
		Rate:    consul.Rate,
		Comment: consul.Comment,
		Status:  consul.Status,
		User: dto.AllUserResponse{
			ID:          consul.User.ID,
			Name:        consul.User.Name,
			Email:       consul.User.Email,
			Password:    consul.User.Password,
			Birthdate:   consul.User.Birthdate,
			PhoneNumber: consul.User.PhoneNumber,
			Data01:      consul.User.Data01,
			Data02:      consul.User.Data02,
			Data03:      consul.User.Data03,
			IsVerified:  consul.User.IsVerified,
			City: dto.CityResponse{
				ID:   &consul.User.City.ID,
				Name: consul.User.City.Name,
				Type: consul.User.City.Type,
				Province: dto.ProvinceResponse{
					ID:   consul.User.City.ProvinceID,
					Name: consul.User.City.Province.Name,
				},
			},
			Role: dto.RoleResponse{
				ID:   &consul.User.Role.ID,
				Name: consul.User.Role.Name,
			},
		},
		AvailableSlot: dto.AvailableSlotResponse{
			ID:       consul.AvailableSlot.ID,
			Start:    consul.AvailableSlot.Start,
			End:      consul.AvailableSlot.End,
			IsBooked: consul.AvailableSlot.IsBooked,
		},
		Practice: dto.PracticeResponse{
			ID:                consul.Practice.ID,
			Type:              consul.Practice.Type,
			Name:              consul.Practice.Name,
			Address:           consul.Practice.Address,
			PhoneNumber:       consul.Practice.PhoneNumber,
			PracticeSchedules: practiceSchedules,
		},
	}

	err = ps.psychologRepo.UpdateConsultation(ctx, nil, consul)
	if err != nil {
		return dto.ConsultationResponse{}, dto.ErrUpdateConsultation
	}

	return data, nil
}
