package service

import (
	"context"
	"time"

	"github.com/Amierza/warasin-mobile/backend/dto"
	"github.com/Amierza/warasin-mobile/backend/entity"
	"github.com/Amierza/warasin-mobile/backend/helpers"
	"github.com/Amierza/warasin-mobile/backend/repository"
	"github.com/google/uuid"
)

type (
	IAdminService interface {
		// Authentication
		Login(ctx context.Context, req dto.AdminLoginRequest) (dto.AdminLoginResponse, error)
		RefreshToken(ctx context.Context, req dto.RefreshTokenRequest) (dto.RefreshTokenResponse, error)

		// Get Role
		GetAllRole(ctx context.Context) (dto.RolePaginationResponse, error)

		// User
		CreateUser(ctx context.Context, req dto.CreateUserRequest) (dto.AllUserResponse, error)
		GetAllUserWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.UserPaginationResponse, error)
		GetDetailUser(ctx context.Context, userID string) (dto.AllUserResponse, error)
		UpdateUser(ctx context.Context, req dto.UpdateUserRequest) (dto.AllUserResponse, error)
		DeleteUser(ctx context.Context, req dto.DeleteUserRequest) (dto.AllUserResponse, error)

		// News
		CreateNews(ctx context.Context, req dto.CreateNewsRequest) (dto.NewsResponse, error)
		GetAllNewsWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.NewsPaginationResponse, error)
		GetDetailNews(ctx context.Context, newsID string) (dto.NewsResponse, error)
		UpdateNews(ctx context.Context, req dto.UpdateNewsRequest) (dto.NewsResponse, error)
		DeleteNews(ctx context.Context, req dto.DeleteNewsRequest) (dto.NewsResponse, error)

		// Motivation Category
		CreateMotivationCategory(ctx context.Context, req dto.CreateMotivationCategoryRequest) (dto.MotivationCategoryResponse, error)
		GetAllMotivationCategoryWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.MotivationCategoryPaginationResponse, error)
		GetDetailMotivationCategory(ctx context.Context, motivationCategoryID string) (dto.MotivationCategoryResponse, error)
		UpdateMotivationCategory(ctx context.Context, req dto.UpdateMotivationCategoryRequest) (dto.MotivationCategoryResponse, error)
		DeleteMotivationCategory(ctx context.Context, req dto.DeleteMotivationCategoryRequest) (dto.MotivationCategoryResponse, error)

		// Motivation
		CreateMotivation(ctx context.Context, req dto.CreateMotivationRequest) (dto.MotivationResponse, error)
		GetAllMotivationWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.MotivationPaginationResponse, error)
		GetDetailMotivation(ctx context.Context, motivationID string) (dto.MotivationResponse, error)
		UpdateMotivation(ctx context.Context, req dto.UpdateMotivationRequest) (dto.MotivationResponse, error)
		DeleteMotivation(ctx context.Context, req dto.DeleteMotivationRequest) (dto.MotivationResponse, error)

		// Psycholog
		CreatePsycholog(ctx context.Context, req dto.CreatePsychologRequest) (dto.PsychologResponse, error)
		GetAllPsychologWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.PsychologPaginationResponse, error)
		UpdatePsycholog(ctx context.Context, req dto.UpdatePsychologRequest) (dto.PsychologResponse, error)
		DeletePsycholog(ctx context.Context, req dto.DeletePsychologRequest) (dto.PsychologResponse, error)

		// User Motivation
		GetAllUserMotivationWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.UserMotivationPaginationResponse, error)

		// User News
		GetAllUserNewsWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.UserNewsPaginationResponse, error)

		// Language Master
		GetAllLanguageMaster(ctx context.Context) (dto.AllLanguageMasterResponse, error)

		// Specialization
		GetAllSpecialization(ctx context.Context) (dto.AllSpecializationResponse, error)

		// Consultation
		GetAllConsultationWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.ConsultationPaginationResponse, error)
	}

	AdminService struct {
		adminRepo  repository.IAdminRepository
		masterRepo repository.IMasterRepository
		jwtService IJWTService
	}
)

func NewAdminService(adminRepo repository.IAdminRepository, masterRepo repository.IMasterRepository, jwtService IJWTService) *AdminService {
	return &AdminService{
		adminRepo:  adminRepo,
		masterRepo: masterRepo,
		jwtService: jwtService,
	}
}

// Authentication
func (as *AdminService) Login(ctx context.Context, req dto.AdminLoginRequest) (dto.AdminLoginResponse, error) {
	user, flag, err := as.adminRepo.GetUserByEmail(ctx, nil, req.Email)
	if !flag || err != nil {
		return dto.AdminLoginResponse{}, dto.ErrEmailNotFound
	}

	if user.Role.Name != "admin" {
		return dto.AdminLoginResponse{}, dto.ErrDeniedAccess
	}

	checkPassword, err := helpers.CheckPassword(user.Password, []byte(req.Password))
	if err != nil || !checkPassword {
		return dto.AdminLoginResponse{}, dto.ErrPasswordNotMatch
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

// Role
func (as *AdminService) GetAllRole(ctx context.Context) (dto.RolePaginationResponse, error) {
	dataWithPaginate, err := as.adminRepo.GetAllRole(ctx, nil)
	if err != nil {
		return dto.RolePaginationResponse{}, dto.ErrGetAllNewsWithPagination
	}

	var datas []dto.RoleResponse
	for _, role := range dataWithPaginate.Roles {
		data := dto.RoleResponse{
			ID:   &role.ID,
			Name: role.Name,
		}

		datas = append(datas, data)
	}

	return dto.RolePaginationResponse{
		Data: datas,
	}, nil
}

// User
func (as *AdminService) CreateUser(ctx context.Context, req dto.CreateUserRequest) (dto.AllUserResponse, error) {
	if len(req.Name) < 5 {
		return dto.AllUserResponse{}, dto.ErrInvalidName
	}

	if !helpers.IsValidEmail(req.Email) {
		return dto.AllUserResponse{}, dto.ErrInvalidEmail
	}

	_, flag, err := as.adminRepo.GetUserByEmail(ctx, nil, req.Email)
	if flag || err == nil {
		return dto.AllUserResponse{}, dto.ErrEmailAlreadyExists
	}

	if len(req.Password) < 8 {
		return dto.AllUserResponse{}, dto.ErrInvalidPassword
	}

	birthdateFormatted, err := helpers.ParseBirthdate(req.Birthdate)
	if err != nil {
		return dto.AllUserResponse{}, dto.ErrFormatBirthdate
	}

	phoneNumberFormatted, err := helpers.StandardizePhoneNumber(req.PhoneNumber)
	if err != nil {
		return dto.AllUserResponse{}, dto.ErrFormatPhoneNumber
	}

	city, err := as.masterRepo.GetCityByID(ctx, nil, req.CityID.String())
	if err != nil {
		return dto.AllUserResponse{}, dto.ErrGetCityByID
	}

	role, err := as.adminRepo.GetRoleByID(ctx, nil, req.RoleID.String())
	if err != nil {
		return dto.AllUserResponse{}, dto.ErrGetRoleFromID
	}

	user := entity.User{
		ID:          uuid.New(),
		Name:        req.Name,
		Email:       req.Email,
		Password:    req.Password,
		Birthdate:   birthdateFormatted,
		PhoneNumber: phoneNumberFormatted,
		City:        city,
		Role:        role,
	}

	err = as.adminRepo.CreateUser(ctx, nil, user)
	if err != nil {
		return dto.AllUserResponse{}, dto.ErrRegisterUser
	}

	res := dto.AllUserResponse{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		Password:    user.Password,
		Birthdate:   user.Birthdate.String(),
		PhoneNumber: user.PhoneNumber,
		Data01:      user.Data01,
		Data02:      user.Data02,
		Data03:      user.Data03,
		IsVerified:  user.IsVerified,
		City: dto.CityResponse{
			ID:   &city.ID,
			Name: user.City.Name,
			Type: user.City.Type,
			Province: dto.ProvinceResponse{
				ID:   user.City.ProvinceID,
				Name: user.City.Province.Name,
			},
		},
		Role: dto.RoleResponse{
			ID:   &role.ID,
			Name: user.Role.Name,
		},
	}

	return res, nil
}
func (as *AdminService) GetAllUserWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.UserPaginationResponse, error) {
	dataWithPaginate, err := as.adminRepo.GetAllUserWithPagination(ctx, nil, req)
	if err != nil {
		return dto.UserPaginationResponse{}, dto.ErrGetAllUserWithPagination
	}

	var datas []dto.AllUserResponse
	for _, user := range dataWithPaginate.Users {
		data := dto.AllUserResponse{
			ID:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
			Image:    user.Image,
			Gender:   user.Gender,
			Birthdate: func(t *time.Time) string {
				if t != nil {
					return t.Format("2006-01-02")
				}
				return ""
			}(user.Birthdate),
			PhoneNumber: user.PhoneNumber,
			Data01:      user.Data01,
			Data02:      user.Data02,
			Data03:      user.Data03,
			IsVerified:  user.IsVerified,
			City: dto.CityResponse{
				ID:   user.CityID,
				Name: user.City.Name,
				Type: user.City.Type,
				Province: dto.ProvinceResponse{
					ID:   user.City.ProvinceID,
					Name: user.City.Province.Name,
				},
			},
			Role: dto.RoleResponse{
				ID:   user.RoleID,
				Name: user.Role.Name,
			},
		}

		datas = append(datas, data)
	}

	return dto.UserPaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}
func (as *AdminService) GetDetailUser(ctx context.Context, userID string) (dto.AllUserResponse, error) {
	user, err := as.adminRepo.GetUserByID(ctx, nil, userID)
	if err != nil {
		return dto.AllUserResponse{}, dto.ErrUserNotFound
	}

	return dto.AllUserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Image:    user.Image,
		Gender:   user.Gender,
		Birthdate: func(t *time.Time) string {
			if t != nil {
				return t.Format("2006-01-02")
			}
			return ""
		}(user.Birthdate),
		PhoneNumber: user.PhoneNumber,
		Data01:      user.Data01,
		Data02:      user.Data02,
		Data03:      user.Data03,
		IsVerified:  user.IsVerified,
		City: dto.CityResponse{
			ID:   user.CityID,
			Name: user.City.Name,
			Type: user.City.Type,
			Province: dto.ProvinceResponse{
				ID:   user.City.ProvinceID,
				Name: user.City.Province.Name,
			},
		},
		Role: dto.RoleResponse{
			ID:   user.RoleID,
			Name: user.Role.Name,
		},
	}, nil
}
func (as *AdminService) UpdateUser(ctx context.Context, req dto.UpdateUserRequest) (dto.AllUserResponse, error) {
	user, err := as.adminRepo.GetUserByID(ctx, nil, req.ID)
	if err != nil {
		return dto.AllUserResponse{}, dto.ErrGetUserFromID
	}

	if req.CityID != nil {
		city, err := as.masterRepo.GetCityByID(ctx, nil, req.CityID.String())
		if err != nil {
			return dto.AllUserResponse{}, dto.ErrGetCityByID
		}

		user.City = city
	}

	if req.RoleID != nil {
		role, err := as.adminRepo.GetRoleByID(ctx, nil, req.RoleID.String())
		if err != nil {
			return dto.AllUserResponse{}, dto.ErrGetRoleFromID
		}

		user.Role = role
	}

	if req.Name != "" {
		if len(req.Name) < 5 {
			return dto.AllUserResponse{}, dto.ErrInvalidName
		}

		user.Name = req.Name
	}

	if req.Email != "" {
		if !helpers.IsValidEmail(req.Email) {
			return dto.AllUserResponse{}, dto.ErrInvalidEmail
		}

		_, flag, err := as.adminRepo.GetUserByEmail(ctx, nil, req.Email)
		if flag || err == nil {
			return dto.AllUserResponse{}, dto.ErrEmailAlreadyExists
		}

		user.Email = req.Email
	}

	if req.Image != "" {
		user.Image = req.Image
	}

	if req.Gender != nil {
		user.Gender = req.Gender
	}

	if req.Birthdate != "" {
		t, err := helpers.ParseBirthdate(req.Birthdate)
		if err != nil {
			return dto.AllUserResponse{}, dto.ErrFormatBirthdate
		}
		user.Birthdate = t
	}

	if req.PhoneNumber != "" {
		phoneNumberFormatted, err := helpers.StandardizePhoneNumber(req.PhoneNumber)
		if err != nil {
			return dto.AllUserResponse{}, dto.ErrFormatPhoneNumber
		}

		user.PhoneNumber = phoneNumberFormatted
	}

	err = as.adminRepo.UpdateUser(ctx, nil, user)
	if err != nil {
		return dto.AllUserResponse{}, dto.ErrUpdateUser
	}

	res := dto.AllUserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Image:    user.Image,
		Gender:   user.Gender,
		Birthdate: func() string {
			if user.Birthdate != nil {
				return user.Birthdate.Format("2006-01-02")
			}
			return ""
		}(),
		PhoneNumber: user.PhoneNumber,
		Data01:      user.Data01,
		Data02:      user.Data02,
		Data03:      user.Data03,
		IsVerified:  user.IsVerified,
		City: dto.CityResponse{
			ID:   user.CityID,
			Name: user.City.Name,
			Type: user.City.Type,
			Province: dto.ProvinceResponse{
				ID:   user.City.ProvinceID,
				Name: user.City.Province.Name,
			},
		},
		Role: dto.RoleResponse{
			ID:   user.RoleID,
			Name: user.Role.Name,
		},
	}

	return res, nil
}
func (as *AdminService) DeleteUser(ctx context.Context, req dto.DeleteUserRequest) (dto.AllUserResponse, error) {
	deletedUser, err := as.adminRepo.GetUserByID(ctx, nil, req.UserID)
	if err != nil {
		return dto.AllUserResponse{}, dto.ErrGetUserFromID
	}

	err = as.adminRepo.DeleteUserByID(ctx, nil, req.UserID)
	if err != nil {
		return dto.AllUserResponse{}, dto.ErrDeleteUserByID
	}

	res := dto.AllUserResponse{
		ID:          deletedUser.ID,
		Name:        deletedUser.Name,
		Email:       deletedUser.Email,
		Password:    deletedUser.Password,
		Birthdate:   deletedUser.Birthdate.String(),
		PhoneNumber: deletedUser.PhoneNumber,
		Data01:      deletedUser.Data01,
		Data02:      deletedUser.Data02,
		Data03:      deletedUser.Data03,
		IsVerified:  deletedUser.IsVerified,
		City: dto.CityResponse{
			ID:   deletedUser.CityID,
			Name: deletedUser.City.Name,
			Type: deletedUser.City.Type,
			Province: dto.ProvinceResponse{
				ID:   deletedUser.City.ProvinceID,
				Name: deletedUser.City.Province.Name,
			},
		},
		Role: dto.RoleResponse{
			ID:   deletedUser.RoleID,
			Name: deletedUser.Role.Name,
		},
	}

	return res, nil
}

// News
func (as *AdminService) CreateNews(ctx context.Context, req dto.CreateNewsRequest) (dto.NewsResponse, error) {
	flag, _, err := as.adminRepo.GetNewsByTitle(ctx, nil, req.Title)
	if flag {
		return dto.NewsResponse{}, dto.ErrNewsTitleAlreadyExists
	}

	news := entity.News{
		ID:    uuid.New(),
		Image: req.Image,
		Title: req.Title,
		Body:  req.Body,
		Date:  req.Date,
	}

	err = as.adminRepo.CreateNews(ctx, nil, news)
	if err != nil {
		return dto.NewsResponse{}, dto.ErrCreateNews
	}

	res := dto.NewsResponse{
		ID:    &news.ID,
		Image: news.Image,
		Title: news.Title,
		Body:  news.Body,
		Date:  news.Date,
	}

	return res, nil
}
func (as *AdminService) GetAllNewsWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.NewsPaginationResponse, error) {
	dataWithPaginate, err := as.adminRepo.GetAllNewsWithPagination(ctx, nil, req)
	if err != nil {
		return dto.NewsPaginationResponse{}, dto.ErrGetAllNewsWithPagination
	}

	var datas []dto.NewsResponse
	for _, news := range dataWithPaginate.News {
		data := dto.NewsResponse{
			ID:    &news.ID,
			Image: news.Image,
			Title: news.Title,
			Body:  news.Body,
			Date:  news.Date,
		}

		datas = append(datas, data)
	}

	return dto.NewsPaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}
func (as *AdminService) GetDetailNews(ctx context.Context, newsID string) (dto.NewsResponse, error) {
	news, err := as.adminRepo.GetNewsByID(ctx, nil, newsID)
	if err != nil {
		return dto.NewsResponse{}, dto.ErrGetNewsFromID
	}

	return dto.NewsResponse{
		ID:    &news.ID,
		Image: news.Image,
		Title: news.Title,
		Body:  news.Body,
		Date:  news.Date,
	}, nil
}
func (as *AdminService) UpdateNews(ctx context.Context, req dto.UpdateNewsRequest) (dto.NewsResponse, error) {
	news, err := as.adminRepo.GetNewsByID(ctx, nil, req.ID)
	if err != nil {
		return dto.NewsResponse{}, dto.ErrGetNewsFromID
	}

	if req.Image != "" {
		news.Image = req.Image
	}
	if req.Title != "" {
		if req.Title == news.Title {
			return dto.NewsResponse{}, dto.ErrNewsTitleAlreadyExists
		}
		news.Title = req.Title
	}
	if req.Body != "" {
		news.Body = req.Body
	}
	if req.Date != "" {
		news.Date = req.Date
	}

	err = as.adminRepo.UpdateNews(ctx, nil, news)
	if err != nil {
		return dto.NewsResponse{}, dto.ErrUpdateNews
	}

	res := dto.NewsResponse{
		ID:    &news.ID,
		Image: news.Image,
		Title: news.Title,
		Body:  news.Body,
		Date:  news.Date,
	}

	return res, nil
}
func (as *AdminService) DeleteNews(ctx context.Context, req dto.DeleteNewsRequest) (dto.NewsResponse, error) {
	deletedNews, err := as.adminRepo.GetNewsByID(ctx, nil, req.NewsID)
	if err != nil {
		return dto.NewsResponse{}, dto.ErrGetNewsFromID
	}

	err = as.adminRepo.DeleteNewsByID(ctx, nil, req.NewsID)
	if err != nil {
		return dto.NewsResponse{}, dto.ErrDeleteNews
	}

	res := dto.NewsResponse{
		ID:    &deletedNews.ID,
		Image: deletedNews.Image,
		Title: deletedNews.Title,
		Body:  deletedNews.Body,
		Date:  deletedNews.Date,
	}

	return res, nil
}

// Motivation Category
func (as *AdminService) CreateMotivationCategory(ctx context.Context, req dto.CreateMotivationCategoryRequest) (dto.MotivationCategoryResponse, error) {
	flag, _, err := as.adminRepo.GetMotivationCategoryByName(ctx, nil, req.Name)
	if flag {
		return dto.MotivationCategoryResponse{}, dto.ErrMotivationCategoryNameAlreadyExists
	}

	motivationCategory := entity.MotivationCategory{
		ID:   uuid.New(),
		Name: req.Name,
	}

	err = as.adminRepo.CreateMotivationCategory(ctx, nil, motivationCategory)
	if err != nil {
		return dto.MotivationCategoryResponse{}, dto.ErrCreateMotivationCategory
	}

	res := dto.MotivationCategoryResponse{
		ID:   &motivationCategory.ID,
		Name: motivationCategory.Name,
	}

	return res, nil
}
func (as *AdminService) GetAllMotivationCategoryWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.MotivationCategoryPaginationResponse, error) {
	dataWithPaginate, err := as.adminRepo.GetAllMotivationCategoryWithPagination(ctx, nil, req)
	if err != nil {
		return dto.MotivationCategoryPaginationResponse{}, dto.ErrGetAllMotivationCategoryWithPagination
	}

	var datas []dto.MotivationCategoryResponse
	for _, motivationCategory := range dataWithPaginate.MotivationCategories {
		data := dto.MotivationCategoryResponse{
			ID:   &motivationCategory.ID,
			Name: motivationCategory.Name,
		}

		datas = append(datas, data)
	}

	return dto.MotivationCategoryPaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}
func (as *AdminService) GetDetailMotivationCategory(ctx context.Context, motivationCategoryID string) (dto.MotivationCategoryResponse, error) {
	motivationCategory, err := as.adminRepo.GetMotivationCategoryByID(ctx, nil, motivationCategoryID)
	if err != nil {
		return dto.MotivationCategoryResponse{}, dto.ErrGetMotivationCategoryFromID
	}

	return dto.MotivationCategoryResponse{
		ID:   &motivationCategory.ID,
		Name: motivationCategory.Name,
	}, nil
}
func (as *AdminService) UpdateMotivationCategory(ctx context.Context, req dto.UpdateMotivationCategoryRequest) (dto.MotivationCategoryResponse, error) {
	motivationCategory, err := as.adminRepo.GetMotivationCategoryByID(ctx, nil, req.ID)
	if err != nil {
		return dto.MotivationCategoryResponse{}, dto.ErrGetMotivationCategoryFromID
	}

	if req.Name != "" {
		if req.Name == motivationCategory.Name {
			return dto.MotivationCategoryResponse{}, dto.ErrMotivationCategoryNameAlreadyExists
		}
		motivationCategory.Name = req.Name
	}

	err = as.adminRepo.UpdateMotivationCategory(ctx, nil, motivationCategory)
	if err != nil {
		return dto.MotivationCategoryResponse{}, dto.ErrUpdateMotivationCategory
	}

	res := dto.MotivationCategoryResponse{
		ID:   &motivationCategory.ID,
		Name: motivationCategory.Name,
	}

	return res, nil
}
func (as *AdminService) DeleteMotivationCategory(ctx context.Context, req dto.DeleteMotivationCategoryRequest) (dto.MotivationCategoryResponse, error) {
	deletedMotivationCategory, err := as.adminRepo.GetMotivationCategoryByID(ctx, nil, req.MotivationCategoryID)
	if err != nil {
		return dto.MotivationCategoryResponse{}, dto.ErrGetMotivationCategoryFromID
	}

	err = as.adminRepo.DeleteMotivationCategoryByID(ctx, nil, req.MotivationCategoryID)
	if err != nil {
		return dto.MotivationCategoryResponse{}, dto.ErrDeleteMotivationCategory
	}

	res := dto.MotivationCategoryResponse{
		ID:   &deletedMotivationCategory.ID,
		Name: deletedMotivationCategory.Name,
	}

	return res, nil
}

// Motivation
func (as *AdminService) CreateMotivation(ctx context.Context, req dto.CreateMotivationRequest) (dto.MotivationResponse, error) {
	motivationCategory, err := as.adminRepo.GetMotivationCategoryByID(ctx, nil, req.MotivationCategoryID.String())
	if err != nil {
		return dto.MotivationResponse{}, dto.ErrGetMotivationCategoryFromID
	}

	flag, _, err := as.adminRepo.GetMotivationByContent(ctx, nil, req.Content)
	if flag {
		return dto.MotivationResponse{}, dto.ErrMotivationContentAlreadyExists
	}

	motivation := entity.Motivation{
		ID:                   uuid.New(),
		Author:               req.Author,
		Content:              req.Content,
		MotivationCategoryID: req.MotivationCategoryID,
	}

	err = as.adminRepo.CreateMotivation(ctx, nil, motivation)
	if err != nil {
		return dto.MotivationResponse{}, dto.ErrCreateMotivation
	}

	res := dto.MotivationResponse{
		ID:      &motivation.ID,
		Author:  motivation.Author,
		Content: motivation.Content,
		MotivationCategory: dto.MotivationCategoryResponse{
			ID:   &motivationCategory.ID,
			Name: motivationCategory.Name,
		},
	}

	return res, nil
}
func (as *AdminService) GetAllMotivationWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.MotivationPaginationResponse, error) {
	dataWithPaginate, err := as.adminRepo.GetAllMotivationWithPagination(ctx, nil, req)
	if err != nil {
		return dto.MotivationPaginationResponse{}, dto.ErrGetAllMotivationWithPagination
	}

	var datas []dto.MotivationResponse
	for _, motivation := range dataWithPaginate.Motivations {
		data := dto.MotivationResponse{
			ID:      &motivation.ID,
			Author:  motivation.Author,
			Content: motivation.Content,
			MotivationCategory: dto.MotivationCategoryResponse{
				ID:   &motivation.MotivationCategory.ID,
				Name: motivation.MotivationCategory.Name,
			},
		}

		datas = append(datas, data)
	}

	return dto.MotivationPaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}
func (as *AdminService) GetDetailMotivation(ctx context.Context, motivationID string) (dto.MotivationResponse, error) {
	motivation, err := as.adminRepo.GetMotivationByID(ctx, nil, motivationID)
	if err != nil {
		return dto.MotivationResponse{}, dto.ErrMotivationNotFound
	}

	return dto.MotivationResponse{
		ID:      &motivation.ID,
		Author:  motivation.Author,
		Content: motivation.Content,
		MotivationCategory: dto.MotivationCategoryResponse{
			ID:   &motivation.MotivationCategory.ID,
			Name: motivation.MotivationCategory.Name,
		},
	}, nil
}
func (as *AdminService) UpdateMotivation(ctx context.Context, req dto.UpdateMotivationRequest) (dto.MotivationResponse, error) {
	motivation, err := as.adminRepo.GetMotivationByID(ctx, nil, req.ID)
	if err != nil {
		return dto.MotivationResponse{}, dto.ErrGetMotivationFromID
	}

	if req.Author != "" {
		motivation.Author = req.Author
	}

	if req.Content != "" {
		if motivation.Content == req.Content {
			return dto.MotivationResponse{}, dto.ErrMotivationContentAlreadyExists
		}
		motivation.Content = req.Content
	}

	if req.Author != "" {
		motivation.Author = req.Author
	}

	if req.MotivationCategoryID != "" {
		motivationCategory, err := as.adminRepo.GetMotivationCategoryByID(ctx, nil, req.MotivationCategoryID)
		if err != nil {
			return dto.MotivationResponse{}, dto.ErrGetMotivationCategoryFromID
		}
		motivation.MotivationCategoryID = &motivationCategory.ID
		motivation.MotivationCategory.Name = motivationCategory.Name
	}

	err = as.adminRepo.UpdateMotivation(ctx, nil, motivation)
	if err != nil {
		return dto.MotivationResponse{}, dto.ErrUpdateMotivation
	}

	res := dto.MotivationResponse{
		ID:      &motivation.ID,
		Author:  motivation.Author,
		Content: motivation.Content,
		MotivationCategory: dto.MotivationCategoryResponse{
			ID:   motivation.MotivationCategoryID,
			Name: motivation.MotivationCategory.Name,
		},
	}

	return res, nil
}
func (as *AdminService) DeleteMotivation(ctx context.Context, req dto.DeleteMotivationRequest) (dto.MotivationResponse, error) {
	deletedMotivation, err := as.adminRepo.GetMotivationByID(ctx, nil, req.ID)
	if err != nil {
		return dto.MotivationResponse{}, dto.ErrGetMotivationFromID
	}

	err = as.adminRepo.DeleteMotivationByID(ctx, nil, req.ID)
	if err != nil {
		return dto.MotivationResponse{}, dto.ErrDeleteMotivation
	}

	res := dto.MotivationResponse{
		ID:      &deletedMotivation.ID,
		Author:  deletedMotivation.Author,
		Content: deletedMotivation.Content,
		MotivationCategory: dto.MotivationCategoryResponse{
			ID:   &deletedMotivation.MotivationCategory.ID,
			Name: deletedMotivation.MotivationCategory.Name,
		},
	}

	return res, nil
}

// Psycholog
func (as *AdminService) CreatePsycholog(ctx context.Context, req dto.CreatePsychologRequest) (dto.PsychologResponse, error) {
	if len(req.Name) < 5 {
		return dto.PsychologResponse{}, dto.ErrInvalidName
	}

	if !helpers.IsValidSTRNumber(req.STRNumber) {
		return dto.PsychologResponse{}, dto.ErrInvalidSTRNumber
	}

	if len(req.WorkYear) < 4 {
		return dto.PsychologResponse{}, dto.ErrInvalidWorkYear
	}

	if !helpers.IsValidEmail(req.Email) {
		return dto.PsychologResponse{}, dto.ErrInvalidEmail
	}

	_, flag, err := as.masterRepo.GetPsychologByEmail(ctx, nil, req.Email)
	if flag || err == nil {
		return dto.PsychologResponse{}, dto.ErrEmailAlreadyExists
	}

	if len(req.Password) < 8 {
		return dto.PsychologResponse{}, dto.ErrInvalidPassword
	}

	phoneNumberFormatted, err := helpers.StandardizePhoneNumber(req.PhoneNumber)
	if err != nil {
		return dto.PsychologResponse{}, dto.ErrFormatPhoneNumber
	}

	city, err := as.masterRepo.GetCityByID(ctx, nil, req.CityID.String())
	if err != nil {
		return dto.PsychologResponse{}, dto.ErrGetCityByID
	}

	role, err := as.adminRepo.GetRoleByID(ctx, nil, req.RoleID.String())
	if err != nil {
		return dto.PsychologResponse{}, dto.ErrGetRoleFromID
	}

	psycholog := entity.Psycholog{
		ID:          uuid.New(),
		Name:        req.Name,
		STRNumber:   req.STRNumber,
		Email:       req.Email,
		Password:    req.Password,
		WorkYear:    req.WorkYear,
		Description: req.Description,
		PhoneNumber: phoneNumberFormatted,
		Image:       req.Image,
		City:        city,
		Role:        role,
	}

	err = as.adminRepo.CreatePsycholog(ctx, nil, psycholog)
	if err != nil {
		return dto.PsychologResponse{}, dto.ErrRegisterPsycholog
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
			ID:   &psycholog.City.ID,
			Name: psycholog.City.Name,
			Type: psycholog.City.Type,
			Province: dto.ProvinceResponse{
				ID:   &psycholog.City.Province.ID,
				Name: psycholog.City.Province.Name,
			},
		},
		Role: dto.RoleResponse{
			ID:   &psycholog.Role.ID,
			Name: psycholog.Role.Name,
		},
	}

	return res, nil
}
func (as *AdminService) GetAllPsychologWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.PsychologPaginationResponse, error) {
	dataWithPaginate, err := as.adminRepo.GetAllPsychologWithPagination(ctx, nil, req)
	if err != nil {
		return dto.PsychologPaginationResponse{}, dto.ErrGetAllPsychologWithPagination
	}

	var datas []dto.PsychologResponse

	for _, psycholog := range dataWithPaginate.Psychologs {
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

		datas = append(datas, data)
	}

	return dto.PsychologPaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}
func (as *AdminService) UpdatePsycholog(ctx context.Context, req dto.UpdatePsychologRequest) (dto.PsychologResponse, error) {
	psycholog, flag, err := as.masterRepo.GetPsychologByID(ctx, nil, req.ID)
	if err != nil || !flag {
		return dto.PsychologResponse{}, dto.ErrPsychologNotFound
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

		_, flag, err := as.masterRepo.GetPsychologByEmail(ctx, nil, req.Email)
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

	if req.CityID != "" {
		city, err := as.masterRepo.GetCityByID(ctx, nil, req.CityID)
		if err != nil {
			return dto.PsychologResponse{}, dto.ErrGetCityByID
		}

		psycholog.City = city
	}

	if len(req.LanguageMasterIDs) > 0 {
		err := as.adminRepo.DeletePsychologLanguageByPsychologID(ctx, nil, psycholog.ID.String())
		if err != nil {
			return dto.PsychologResponse{}, dto.ErrDeletePsychologLanguageByPsychologID
		}

		var newLangs []entity.PsychologLanguage

		for _, langID := range req.LanguageMasterIDs {
			languageMaster, found, err := as.adminRepo.GetLanguageMasterByID(ctx, nil, langID)
			if err != nil || !found {
				return dto.PsychologResponse{}, dto.ErrLanguageMasterNotFound
			}

			newLangs = append(newLangs, entity.PsychologLanguage{
				ID:               uuid.New(),
				PsychologID:      &psycholog.ID,
				LanguageMasterID: &languageMaster.ID,
			})
		}

		err = as.adminRepo.CreatePsychologLanguages(ctx, nil, newLangs)
		if err != nil {
			return dto.PsychologResponse{}, dto.ErrCreatePsychologLanguages
		}

		psycholog.PsychologLanguages = newLangs
	}

	if len(req.SpecializationIDs) > 0 {
		err := as.adminRepo.DeletePsychologSpecializationByPsychologID(ctx, nil, psycholog.ID.String())
		if err != nil {
			return dto.PsychologResponse{}, dto.ErrDeletePsychologSpecializationByPsychologID
		}

		var newSpecializations []entity.PsychologSpecialization

		for _, speID := range req.SpecializationIDs {
			specialization, found, err := as.adminRepo.GetSpecializationByID(ctx, nil, speID)
			if err != nil || !found {
				return dto.PsychologResponse{}, dto.ErrSpecializationNotFound
			}

			newSpecializations = append(newSpecializations, entity.PsychologSpecialization{
				ID:               uuid.New(),
				PsychologID:      &psycholog.ID,
				SpecializationID: &specialization.ID,
			})
		}

		err = as.adminRepo.CreatePsychologSpecializations(ctx, nil, newSpecializations)
		if err != nil {
			return dto.PsychologResponse{}, dto.ErrCreatePsychologSpecializations
		}

		psycholog.PsychologSpecializations = newSpecializations
	}

	if len(req.Educations) > 0 {
		err := as.adminRepo.DeleteEducationByPsychologID(ctx, nil, psycholog.ID.String())
		if err != nil {
			return dto.PsychologResponse{}, dto.ErrDeleteEducationByPsychologID
		}

		var newEducations []entity.Education
		for _, newEducation := range req.Educations {
			newEducations = append(newEducations, entity.Education{
				ID:             uuid.New(),
				Degree:         newEducation.Degree,
				Major:          newEducation.Major,
				Institution:    newEducation.Institution,
				GraduationYear: newEducation.GraduationYear,
			})
		}

		err = as.adminRepo.CreateEducations(ctx, nil, newEducations)
		if err != nil {
			return dto.PsychologResponse{}, dto.ErrCreateEducations
		}

		psycholog.Educations = newEducations
	}

	err = as.adminRepo.UpdatePsycholog(ctx, nil, psycholog)
	if err != nil {
		return dto.PsychologResponse{}, dto.ErrUpdatePsycholog
	}

	psycholog, found, err := as.masterRepo.GetPsychologByID(ctx, nil, req.ID)
	if err != nil || !found {
		return dto.PsychologResponse{}, dto.ErrPsychologNotFound
	}

	var languageMasters []dto.LanguageMasterResponse
	for _, lang := range psycholog.PsychologLanguages {
		languageMasters = append(languageMasters, dto.LanguageMasterResponse{
			ID:   &lang.LanguageMaster.ID,
			Name: lang.LanguageMaster.Name,
		})
	}

	var specializations []dto.SpecializationResponse
	for _, spe := range psycholog.PsychologSpecializations {
		specializations = append(specializations, dto.SpecializationResponse{
			ID:          &spe.Specialization.ID,
			Name:        spe.Specialization.Name,
			Description: spe.Specialization.Description,
		})
	}

	var educations []dto.EducationResponse
	for _, edu := range psycholog.Educations {
		educations = append(educations, dto.EducationResponse{
			ID:             &edu.ID,
			Degree:         edu.Degree,
			Major:          edu.Major,
			Institution:    edu.Institution,
			GraduationYear: edu.GraduationYear,
		})
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
		LanguageMasters: languageMasters,
		Specializations: specializations,
		Educations:      educations,
	}

	return res, nil
}
func (as *AdminService) DeletePsycholog(ctx context.Context, req dto.DeletePsychologRequest) (dto.PsychologResponse, error) {
	deletedPsycholog, flag, err := as.masterRepo.GetPsychologByID(ctx, nil, req.ID)
	if err != nil || !flag {
		return dto.PsychologResponse{}, dto.ErrGetPsychologFromID
	}

	err = as.adminRepo.DeletePsychologByID(ctx, nil, req.ID)
	if err != nil {
		return dto.PsychologResponse{}, dto.ErrDeletePsycholog
	}

	res := dto.PsychologResponse{
		ID:          deletedPsycholog.ID,
		Name:        deletedPsycholog.Name,
		STRNumber:   deletedPsycholog.STRNumber,
		Email:       deletedPsycholog.Email,
		Password:    deletedPsycholog.Password,
		WorkYear:    deletedPsycholog.WorkYear,
		Description: deletedPsycholog.Description,
		PhoneNumber: deletedPsycholog.PhoneNumber,
		Image:       deletedPsycholog.Image,
		City: dto.CityResponse{
			ID:   deletedPsycholog.CityID,
			Name: deletedPsycholog.City.Name,
			Type: deletedPsycholog.City.Type,
			Province: dto.ProvinceResponse{
				ID:   deletedPsycholog.City.ProvinceID,
				Name: deletedPsycholog.City.Province.Name,
			},
		},
		Role: dto.RoleResponse{
			ID:   deletedPsycholog.RoleID,
			Name: deletedPsycholog.Role.Name,
		},
	}

	return res, nil
}

// User Motivation
func (as *AdminService) GetAllUserMotivationWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.UserMotivationPaginationResponse, error) {
	dataWithPaginate, err := as.adminRepo.GetAllUserMotivationWithPagination(ctx, nil, req)
	if err != nil {
		return dto.UserMotivationPaginationResponse{}, dto.ErrGetAllPsychologWithPagination
	}

	var datas []dto.UserMotivationResponse

	for _, userMotivation := range dataWithPaginate.UserMotivations {
		data := dto.UserMotivationResponse{
			ID: &userMotivation.ID,
			User: dto.AllUserResponse{
				ID:       userMotivation.User.ID,
				Name:     userMotivation.User.Name,
				Email:    userMotivation.User.Email,
				Password: userMotivation.User.Password,
				Image:    userMotivation.User.Image,
				Gender:   userMotivation.User.Gender,
				Birthdate: func(t *time.Time) string {
					if t != nil {
						return t.Format("2006-01-02")
					}
					return ""
				}(userMotivation.User.Birthdate),
				PhoneNumber: userMotivation.User.PhoneNumber,
				Data01:      userMotivation.User.Data01,
				Data02:      userMotivation.User.Data02,
				Data03:      userMotivation.User.Data03,
				IsVerified:  userMotivation.User.IsVerified,
				City: dto.CityResponse{
					ID:   userMotivation.User.CityID,
					Name: userMotivation.User.City.Name,
					Type: userMotivation.User.City.Type,
					Province: dto.ProvinceResponse{
						ID:   userMotivation.User.City.ProvinceID,
						Name: userMotivation.User.City.Province.Name,
					},
				},
				Role: dto.RoleResponse{
					ID:   userMotivation.User.RoleID,
					Name: userMotivation.User.Role.Name,
				},
			},
			Motivation: dto.MotivationResponse{
				ID:      &userMotivation.Motivation.ID,
				Author:  userMotivation.Motivation.Author,
				Content: userMotivation.Motivation.Content,
				MotivationCategory: dto.MotivationCategoryResponse{
					ID:   &userMotivation.Motivation.MotivationCategory.ID,
					Name: userMotivation.Motivation.MotivationCategory.Name,
				},
			},
		}

		datas = append(datas, data)
	}

	return dto.UserMotivationPaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}

// User News
func (as *AdminService) GetAllUserNewsWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.UserNewsPaginationResponse, error) {
	dataWithPaginate, err := as.adminRepo.GetAllUserNewsWithPagination(ctx, nil, req)
	if err != nil {
		return dto.UserNewsPaginationResponse{}, dto.ErrGetAllNewsWithPagination
	}

	var datas []dto.UserNewsResponse

	for _, userNews := range dataWithPaginate.UserNews {
		data := dto.UserNewsResponse{
			ID:   &userNews.ID,
			Date: userNews.Date,
			User: dto.AllUserResponse{
				ID:       userNews.User.ID,
				Name:     userNews.User.Name,
				Email:    userNews.User.Email,
				Password: userNews.User.Password,
				Image:    userNews.User.Image,
				Gender:   userNews.User.Gender,
				Birthdate: func(t *time.Time) string {
					if t != nil {
						return t.Format("2006-01-02")
					}
					return ""
				}(userNews.User.Birthdate),
				PhoneNumber: userNews.User.PhoneNumber,
				Data01:      userNews.User.Data01,
				Data02:      userNews.User.Data02,
				Data03:      userNews.User.Data03,
				IsVerified:  userNews.User.IsVerified,
				City: dto.CityResponse{
					ID:   userNews.User.CityID,
					Name: userNews.User.City.Name,
					Type: userNews.User.City.Type,
					Province: dto.ProvinceResponse{
						ID:   userNews.User.City.ProvinceID,
						Name: userNews.User.City.Province.Name,
					},
				},
				Role: dto.RoleResponse{
					ID:   userNews.User.RoleID,
					Name: userNews.User.Role.Name,
				},
			},
			News: dto.NewsResponse{
				ID:    &userNews.News.ID,
				Image: userNews.News.Image,
				Title: userNews.News.Title,
				Body:  userNews.News.Body,
				Date:  userNews.News.Date,
			},
		}

		datas = append(datas, data)
	}

	return dto.UserNewsPaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}

// Language Master
func (as *AdminService) GetAllLanguageMaster(ctx context.Context) (dto.AllLanguageMasterResponse, error) {
	data, err := as.adminRepo.GetAllLanguageMaster(ctx, nil)
	if err != nil {
		return dto.AllLanguageMasterResponse{}, dto.ErrGetAllLanguageMaster
	}

	var datas []dto.LanguageMasterResponse
	for _, languageMaster := range data.LanguageMasters {
		datas = append(datas, dto.LanguageMasterResponse{
			ID:   &languageMaster.ID,
			Name: languageMaster.Name,
		})
	}

	return dto.AllLanguageMasterResponse{
		LanguageMasters: datas,
	}, nil
}

// Specialization
func (as *AdminService) GetAllSpecialization(ctx context.Context) (dto.AllSpecializationResponse, error) {
	data, err := as.adminRepo.GetAllSpecialization(ctx, nil)
	if err != nil {
		return dto.AllSpecializationResponse{}, dto.ErrGetAllLanguageMaster
	}

	var datas []dto.SpecializationResponse
	for _, specialization := range data.Specializations {
		datas = append(datas, dto.SpecializationResponse{
			ID:          &specialization.ID,
			Name:        specialization.Name,
			Description: specialization.Description,
		})
	}

	return dto.AllSpecializationResponse{
		Specializations: datas,
	}, nil
}

// Consultation
func (as *AdminService) GetAllConsultationWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.ConsultationPaginationResponse, error) {
	dataWithPaginate, err := as.adminRepo.GetAllConsultationWithPagination(ctx, nil, req)
	if err != nil {
		return dto.ConsultationPaginationResponse{}, dto.ErrGetAllConsultationWithPagination
	}

	var datas []dto.ConsultationResponse
	for _, consultation := range dataWithPaginate.Consultations {
		data := dto.ConsultationResponse{
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
			Psycholog: dto.PsychologResponse{
				ID:          consultation.Psycholog.ID,
				Name:        consultation.Psycholog.Name,
				STRNumber:   consultation.Psycholog.STRNumber,
				Email:       consultation.Psycholog.Email,
				Password:    consultation.Psycholog.Password,
				WorkYear:    consultation.Psycholog.WorkYear,
				Description: consultation.Psycholog.Description,
				PhoneNumber: consultation.Psycholog.PhoneNumber,
				Image:       consultation.Psycholog.Image,
				City: dto.CityResponse{
					ID:   &consultation.Psycholog.City.ID,
					Name: consultation.Psycholog.City.Name,
					Type: consultation.Psycholog.City.Type,
					Province: dto.ProvinceResponse{
						ID:   &consultation.Psycholog.City.Province.ID,
						Name: consultation.Psycholog.City.Province.Name,
					},
				},
				Role: dto.RoleResponse{
					ID:   &consultation.Psycholog.Role.ID,
					Name: consultation.Psycholog.Role.Name,
				},
			},
		}

		datas = append(datas, data)
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
