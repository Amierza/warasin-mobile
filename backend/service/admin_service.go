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
	IAdminService interface {
		Login(ctx context.Context, req dto.AdminLoginRequest) (dto.AdminLoginResponse, error)
		RefreshToken(ctx context.Context, req dto.RefreshTokenRequest) (dto.RefreshTokenResponse, error)
		CreateUser(ctx context.Context, req dto.CreateUserRequest) (dto.AllUserResponse, error)
		GetAllUserWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.UserPaginationResponse, error)
		UpdateUser(ctx context.Context, req dto.UpdateUserRequest) (dto.AllUserResponse, error)
		DeleteUser(ctx context.Context, req dto.DeleteUserRequest) (dto.AllUserResponse, error)
		CreateNews(ctx context.Context, req dto.CreateNewsRequest) (dto.NewsResponse, error)
		GetAllNewsWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.NewsPaginationResponse, error)
		UpdateNews(ctx context.Context, req dto.UpdateNewsRequest) (dto.NewsResponse, error)
		DeleteNews(ctx context.Context, req dto.DeleteNewsRequest) (dto.NewsResponse, error)
		GetAllMotivationCategoryWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.MotivationCategoryPaginationResponse, error)
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

func (as *AdminService) CreateUser(ctx context.Context, req dto.CreateUserRequest) (dto.AllUserResponse, error) {
	if len(req.Name) < 5 {
		return dto.AllUserResponse{}, dto.ErrInvalidName
	}

	if !helpers.IsValidEmail(req.Email) {
		return dto.AllUserResponse{}, dto.ErrInvalidEmail
	}

	_, flag, err := as.adminRepo.CheckEmail(ctx, nil, req.Email)
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

	city, err := as.adminRepo.GetCityByID(ctx, nil, req.CityID.String())
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
		Birthdate:   user.Birthdate,
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
			ID:          user.ID,
			Name:        user.Name,
			Email:       user.Email,
			Password:    user.Password,
			Gender:      user.Gender,
			Birthdate:   user.Birthdate,
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

func (as *AdminService) UpdateUser(ctx context.Context, req dto.UpdateUserRequest) (dto.AllUserResponse, error) {
	user, err := as.adminRepo.GetUserByID(ctx, nil, req.ID)
	if err != nil {
		return dto.AllUserResponse{}, dto.ErrGetDataUserFromID
	}

	if req.CityID != nil {
		city, err := as.adminRepo.GetCityByID(ctx, nil, req.CityID.String())
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

		_, flag, err := as.adminRepo.CheckEmail(ctx, nil, req.Email)
		if flag || err == nil {
			return dto.AllUserResponse{}, dto.ErrEmailAlreadyExists
		}

		user.Email = req.Email
	}

	if req.Birthdate != nil {
		user.Birthdate = req.Birthdate
	}

	if req.PhoneNumber != "" {
		phoneNumberFormatted, err := helpers.StandardizePhoneNumber(req.PhoneNumber)
		if err != nil {
			return dto.AllUserResponse{}, dto.ErrFormatPhoneNumber
		}

		user.PhoneNumber = phoneNumberFormatted
	}

	updatedUser, err := as.adminRepo.UpdateUser(ctx, nil, user)
	if err != nil {
		return dto.AllUserResponse{}, dto.ErrUpdateUser
	}

	res := dto.AllUserResponse{
		ID:          updatedUser.ID,
		Name:        updatedUser.Name,
		Email:       updatedUser.Email,
		Password:    updatedUser.Password,
		Birthdate:   updatedUser.Birthdate,
		PhoneNumber: updatedUser.PhoneNumber,
		Data01:      updatedUser.Data01,
		Data02:      updatedUser.Data02,
		Data03:      updatedUser.Data03,
		IsVerified:  updatedUser.IsVerified,
		City: dto.CityResponse{
			ID:   updatedUser.CityID,
			Name: updatedUser.City.Name,
			Type: updatedUser.City.Type,
			Province: dto.ProvinceResponse{
				ID:   updatedUser.City.ProvinceID,
				Name: updatedUser.City.Province.Name,
			},
		},
		Role: dto.RoleResponse{
			ID:   updatedUser.RoleID,
			Name: updatedUser.Role.Name,
		},
	}

	return res, nil
}

func (as *AdminService) DeleteUser(ctx context.Context, req dto.DeleteUserRequest) (dto.AllUserResponse, error) {
	deletedUser, err := as.adminRepo.GetUserByID(ctx, nil, req.UserID)
	if err != nil {
		return dto.AllUserResponse{}, dto.ErrGetDataUserFromID
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
		Birthdate:   deletedUser.Birthdate,
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

func (as *AdminService) CreateNews(ctx context.Context, req dto.CreateNewsRequest) (dto.NewsResponse, error) {
	news := entity.News{
		ID:    uuid.New(),
		Image: req.Image,
		Title: req.Title,
		Body:  req.Body,
		Date:  req.Date,
	}

	err := as.adminRepo.CreateNews(ctx, nil, news)
	if err != nil {
		return dto.NewsResponse{}, dto.ErrCreateNews
	}

	res := dto.NewsResponse{
		ID:    news.ID,
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
			ID:    news.ID,
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

func (as *AdminService) UpdateNews(ctx context.Context, req dto.UpdateNewsRequest) (dto.NewsResponse, error) {
	news, err := as.adminRepo.GetNewsByID(ctx, nil, req.ID)
	if err != nil {
		return dto.NewsResponse{}, dto.ErrGetNewsFromID
	}

	if req.Image != "" {
		news.Image = req.Image
	}
	if req.Title != "" {
		news.Title = req.Title
	}
	if req.Body != "" {
		news.Body = req.Body
	}
	if req.Date != "" {
		news.Date = req.Date
	}

	updatedNews, err := as.adminRepo.UpdateNews(ctx, nil, news)
	if err != nil {
		return dto.NewsResponse{}, dto.ErrUpdateNews
	}

	res := dto.NewsResponse{
		ID:    updatedNews.ID,
		Image: updatedNews.Image,
		Title: updatedNews.Title,
		Body:  updatedNews.Body,
		Date:  updatedNews.Date,
	}

	return res, nil
}

func (as *AdminService) DeleteNews(ctx context.Context, req dto.DeleteNewsRequest) (dto.NewsResponse, error) {
	deletedNews, err := as.adminRepo.GetNewsByID(ctx, nil, req.UserID)
	if err != nil {
		return dto.NewsResponse{}, dto.ErrGetNewsFromID
	}

	err = as.adminRepo.DeleteNewsByID(ctx, nil, req.UserID)
	if err != nil {
		return dto.NewsResponse{}, dto.ErrDeleteNews
	}

	res := dto.NewsResponse{
		ID:    deletedNews.ID,
		Image: deletedNews.Image,
		Title: deletedNews.Title,
		Body:  deletedNews.Body,
		Date:  deletedNews.Date,
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
