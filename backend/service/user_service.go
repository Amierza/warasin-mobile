package service

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"os"
	"strings"
	"time"

	"github.com/Amierza/warasin-mobile/backend/dto"
	"github.com/Amierza/warasin-mobile/backend/entity"
	"github.com/Amierza/warasin-mobile/backend/helpers"
	"github.com/Amierza/warasin-mobile/backend/repository"
	"github.com/Amierza/warasin-mobile/backend/utils"
)

type (
	IUserService interface {
		Register(ctx context.Context, req dto.UserRegisterRequest) (dto.AllUserResponse, error)
		Login(ctx context.Context, req dto.UserLoginRequest) (dto.UserLoginResponse, error)
		RefreshToken(ctx context.Context, req dto.RefreshTokenRequest) (dto.RefreshTokenResponse, error)
		SendForgotPasswordEmail(ctx context.Context, req dto.SendForgotPasswordEmailRequest) error
		ForgotPassword(ctx context.Context, req dto.ForgotPasswordRequest) (dto.ForgotPasswordResponse, error)
		UpdatePassword(ctx context.Context, req dto.UpdatePasswordRequest) (dto.UpdatePasswordResponse, error)
		SendVerificationEmail(ctx context.Context, req dto.SendVerificationEmailRequest) error
		VerifyEmail(ctx context.Context, req dto.VerifyEmailRequest) (dto.VerifyEmailResponse, error)
		GetDetailUser(ctx context.Context) (dto.AllUserResponse, error)
		UpdateUser(ctx context.Context, req dto.UpdateUserRequest) (dto.AllUserResponse, error)
		GetAllProvince(ctx context.Context) (dto.ProvincesResponse, error)
		GetAllCity(ctx context.Context, req dto.CityQueryRequest) (dto.CitiesResponse, error)
		GetAllNewsWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.NewsPaginationResponse, error)
	}

	UserService struct {
		userRepo   repository.IUserRepository
		jwtService IJWTService
	}
)

func NewUserService(userRepo repository.IUserRepository, jwtService IJWTService) *UserService {
	return &UserService{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

func (us *UserService) Register(ctx context.Context, req dto.UserRegisterRequest) (dto.AllUserResponse, error) {
	if len(req.Name) < 5 {
		return dto.AllUserResponse{}, dto.ErrInvalidName
	}

	if !helpers.IsValidEmail(req.Email) {
		return dto.AllUserResponse{}, dto.ErrInvalidEmail
	}

	if len(req.Password) < 8 {
		return dto.AllUserResponse{}, dto.ErrInvalidPassword
	}

	role, err := us.userRepo.GetRoleByName(ctx, nil, "user")
	if err != nil {
		return dto.AllUserResponse{}, dto.ErrGetRoleFromName
	}

	_, flag, err := us.userRepo.CheckEmail(ctx, nil, req.Email)
	if flag || err == nil {
		return dto.AllUserResponse{}, dto.ErrEmailAlreadyExists
	}

	user := entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     role,
	}

	userReg, err := us.userRepo.RegisterUser(ctx, nil, user)
	if err != nil {
		return dto.AllUserResponse{}, dto.ErrRegisterUser
	}

	return dto.AllUserResponse{
		ID:       userReg.ID,
		Name:     userReg.Name,
		Email:    userReg.Email,
		Password: userReg.Password,
		Role: dto.RoleResponse{
			ID:   userReg.RoleID,
			Name: user.Role.Name,
		},
	}, nil
}

func (us *UserService) Login(ctx context.Context, req dto.UserLoginRequest) (dto.UserLoginResponse, error) {
	if !helpers.IsValidEmail(req.Email) {
		return dto.UserLoginResponse{}, dto.ErrInvalidEmail
	}

	if len(req.Password) < 8 {
		return dto.UserLoginResponse{}, dto.ErrInvalidPassword
	}

	user, flag, err := us.userRepo.CheckEmail(ctx, nil, req.Email)
	if !flag || err != nil {
		return dto.UserLoginResponse{}, dto.ErrEmailNotFound
	}

	checkPassword, err := helpers.CheckPassword(user.Password, []byte(req.Password))
	if err != nil || !checkPassword {
		return dto.UserLoginResponse{}, dto.ErrPasswordNotMatch
	}

	if user.Role.Name != "user" {
		return dto.UserLoginResponse{}, dto.ErrDeniedAccess
	}

	permissions, err := us.userRepo.GetPermissionsByRoleID(ctx, nil, user.RoleID.String())
	if err != nil {
		return dto.UserLoginResponse{}, dto.ErrGetPermissionsByRoleID
	}

	accessToken, refreshToken, err := us.jwtService.GenerateToken(user.ID.String(), user.RoleID.String(), permissions)
	if err != nil {
		return dto.UserLoginResponse{}, err
	}

	return dto.UserLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (us *UserService) RefreshToken(ctx context.Context, req dto.RefreshTokenRequest) (dto.RefreshTokenResponse, error) {
	_, err := us.jwtService.ValidateToken(req.RefreshToken)

	if err != nil {
		return dto.RefreshTokenResponse{}, dto.ErrValidateToken
	}

	userID, err := us.jwtService.GetUserIDByToken(req.RefreshToken)
	if err != nil {
		return dto.RefreshTokenResponse{}, dto.ErrGetUserIDFromToken
	}

	roleID, err := us.jwtService.GetRoleIDByToken(req.RefreshToken)
	if err != nil {
		return dto.RefreshTokenResponse{}, dto.ErrGetRoleFromToken
	}

	role, err := us.userRepo.GetRoleByID(ctx, nil, roleID)
	if err != nil {
		return dto.RefreshTokenResponse{}, dto.ErrGetRoleFromID
	}

	if role.Name != "user" {
		return dto.RefreshTokenResponse{}, dto.ErrDeniedAccess
	}

	endpoints, err := us.userRepo.GetPermissionsByRoleID(ctx, nil, roleID)
	if err != nil {
		return dto.RefreshTokenResponse{}, dto.ErrGetPermissionsByRoleID
	}

	accessToken, _, err := us.jwtService.GenerateToken(userID, roleID, endpoints)
	if err != nil {
		return dto.RefreshTokenResponse{}, dto.ErrGenerateAccessToken
	}

	return dto.RefreshTokenResponse{AccessToken: accessToken}, nil
}

func makeVerificationEmail(receiverEmail string) (map[string]string, error) {
	expired := time.Now().Add(time.Hour * 24).Format("2006-01-02 15:04:05")
	plainText := fmt.Sprintf("%s_%s", receiverEmail, expired)
	token, err := utils.AESEncrypt(plainText)
	if err != nil {
		return nil, err
	}

	baseURL := os.Getenv("BASE_URL")
	verifyEmailRoute := "verify-email"
	if baseURL == "" {
		baseURL = "http://127.0.0.1:8000/api/v1/user"
	}

	verifyLink := baseURL + "/" + verifyEmailRoute + "?token=" + token

	readHTML, err := os.ReadFile("utils/email_template/verification_mail.html")
	if err != nil {
		return nil, err
	}

	data := struct {
		Email  string
		Verify string
	}{
		Email:  receiverEmail,
		Verify: verifyLink,
	}

	tmpl, err := template.New("custom").Parse(string(readHTML))
	if err != nil {
		return nil, err
	}

	var strMail bytes.Buffer
	if err := tmpl.Execute(&strMail, data); err != nil {
		return nil, err
	}

	draftEmail := map[string]string{
		"subject": "warasin",
		"body":    strMail.String(),
	}

	return draftEmail, nil
}

func makeForgotPasswordEmail(receiverEmail string) (map[string]string, error) {
	expired := time.Now().Add(time.Hour * 24).Format("2006-01-02 15:04:05")
	plainText := fmt.Sprintf("%s_%s", receiverEmail, expired)
	token, err := utils.AESEncrypt(plainText)
	if err != nil {
		return nil, err
	}

	baseURL := os.Getenv("BASE_URL")
	forgotPasswordEmailRoute := "forgot-password"
	if baseURL == "" {
		baseURL = "http://127.0.0.1:8000/api/v1/user"
	}

	forgotPasswordLink := baseURL + "/" + forgotPasswordEmailRoute + "?token=" + token

	readHTML, err := os.ReadFile("utils/email_template/forgot_password_mail.html")
	if err != nil {
		return nil, err
	}

	data := struct {
		Email          string
		ForgotPassword string
	}{
		Email:          receiverEmail,
		ForgotPassword: forgotPasswordLink,
	}

	tmpl, err := template.New("custom").Parse(string(readHTML))
	if err != nil {
		return nil, err
	}

	var strMail bytes.Buffer
	if err := tmpl.Execute(&strMail, data); err != nil {
		return nil, err
	}

	draftEmail := map[string]string{
		"subject": "warasin",
		"body":    strMail.String(),
	}

	return draftEmail, nil
}

func (us *UserService) SendForgotPasswordEmail(ctx context.Context, req dto.SendForgotPasswordEmailRequest) error {
	user, flag, err := us.userRepo.CheckEmail(ctx, nil, req.Email)
	if err != nil || !flag {
		return dto.ErrEmailNotFound
	}

	draftEmail, err := makeForgotPasswordEmail(user.Email)
	if err != nil {
		return dto.ErrMakeVerificationEmail
	}

	if err := utils.SendEmail(user.Email, draftEmail["subject"], draftEmail["body"]); err != nil {
		return dto.ErrSendEmail
	}

	return nil
}

func (us *UserService) ForgotPassword(ctx context.Context, req dto.ForgotPasswordRequest) (dto.ForgotPasswordResponse, error) {
	decryptedToken, err := utils.AESDecrypt(req.Token)
	if err != nil {
		return dto.ForgotPasswordResponse{}, dto.ErrDecryptToken
	}

	if !strings.Contains(decryptedToken, "_") {
		return dto.ForgotPasswordResponse{}, dto.ErrTokenInvalid
	}

	decryptedTokenSplit := strings.Split(decryptedToken, "_")
	if len(decryptedTokenSplit) != 2 {
		return dto.ForgotPasswordResponse{}, dto.ErrTokenInvalid
	}

	email := decryptedTokenSplit[0]
	expired := decryptedTokenSplit[1]

	now := time.Now()
	expiredTime, err := time.Parse("2006-01-02 15:04:05", expired)
	if err != nil {
		return dto.ForgotPasswordResponse{}, dto.ErrParsingExpiredTime
	}

	if expiredTime.Sub(now) < 0 {
		return dto.ForgotPasswordResponse{}, dto.ErrTokenExpired
	}

	return dto.ForgotPasswordResponse{
		Email: email,
	}, nil
}

func (us *UserService) UpdatePassword(ctx context.Context, req dto.UpdatePasswordRequest) (dto.UpdatePasswordResponse, error) {
	user, flag, err := us.userRepo.CheckEmail(ctx, nil, req.Email)
	if err != nil || !flag {
		return dto.UpdatePasswordResponse{}, dto.ErrUserNotFound
	}

	oldPassword := user.Password

	newPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		return dto.UpdatePasswordResponse{}, dto.ErrHashPassword
	}

	user.Password = newPassword

	_, err = us.userRepo.UpdateUser(ctx, nil, user)
	if err != nil {
		return dto.UpdatePasswordResponse{}, dto.ErrUpdateUser
	}

	return dto.UpdatePasswordResponse{
		OldPassword: oldPassword,
		NewPassword: newPassword,
	}, nil
}

func (us *UserService) SendVerificationEmail(ctx context.Context, req dto.SendVerificationEmailRequest) error {
	user, flag, err := us.userRepo.CheckEmail(ctx, nil, req.Email)
	if err != nil || !flag {
		return dto.ErrEmailNotFound
	}

	draftEmail, err := makeVerificationEmail(user.Email)
	if err != nil {
		return dto.ErrMakeVerificationEmail
	}

	if err := utils.SendEmail(user.Email, draftEmail["subject"], draftEmail["body"]); err != nil {
		return dto.ErrSendEmail
	}

	return nil
}

func (us *UserService) VerifyEmail(ctx context.Context, req dto.VerifyEmailRequest) (dto.VerifyEmailResponse, error) {
	decryptedToken, err := utils.AESDecrypt(req.Token)
	if err != nil {
		return dto.VerifyEmailResponse{}, dto.ErrDecryptToken
	}

	if !strings.Contains(decryptedToken, "_") {
		return dto.VerifyEmailResponse{}, dto.ErrTokenInvalid
	}

	decryptedTokenSplit := strings.Split(decryptedToken, "_")
	if len(decryptedTokenSplit) != 2 {
		return dto.VerifyEmailResponse{}, dto.ErrTokenInvalid
	}

	email := decryptedTokenSplit[0]
	expired := decryptedTokenSplit[1]

	now := time.Now()
	expiredTime, err := time.Parse("2006-01-02 15:04:05", expired)
	if err != nil {
		return dto.VerifyEmailResponse{}, dto.ErrParsingExpiredTime
	}

	if expiredTime.Sub(now) < 0 {
		return dto.VerifyEmailResponse{
			Email:      email,
			IsVerified: false,
		}, dto.ErrTokenExpired
	}

	user, flag, err := us.userRepo.CheckEmail(ctx, nil, email)
	if !flag || err != nil {
		return dto.VerifyEmailResponse{}, dto.ErrUserNotFound
	}

	if user.IsVerified {
		return dto.VerifyEmailResponse{}, dto.ErrEmailALreadyVerified
	}

	updatedUser, err := us.userRepo.UpdateUser(ctx, nil, entity.User{
		ID:         user.ID,
		IsVerified: true,
	})
	if err != nil {
		return dto.VerifyEmailResponse{}, dto.ErrUpdateUser
	}

	return dto.VerifyEmailResponse{
		Email:      email,
		IsVerified: updatedUser.IsVerified,
	}, nil
}

func (us *UserService) GetDetailUser(ctx context.Context) (dto.AllUserResponse, error) {
	token := ctx.Value("Authorization").(string)

	userId, err := us.jwtService.GetUserIDByToken(token)
	if err != nil {
		return dto.AllUserResponse{}, dto.ErrGetUserIDFromToken
	}

	user, err := us.userRepo.GetUserByID(ctx, nil, userId)
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

func (us *UserService) UpdateUser(ctx context.Context, req dto.UpdateUserRequest) (dto.AllUserResponse, error) {
	token := ctx.Value("Authorization").(string)
	userID, err := us.jwtService.GetUserIDByToken(token)

	user, err := us.userRepo.GetUserByID(ctx, nil, userID)
	if err != nil {
		return dto.AllUserResponse{}, dto.ErrGetDataUserFromID
	}

	if req.CityID != nil {
		city, err := us.userRepo.GetCityByID(ctx, nil, req.CityID.String())
		if err != nil {
			return dto.AllUserResponse{}, dto.ErrGetCityByID
		}

		user.City = city
	}

	if req.RoleID != nil {
		role, err := us.userRepo.GetRoleByID(ctx, nil, req.RoleID.String())
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

		_, flag, err := us.userRepo.CheckEmail(ctx, nil, req.Email)
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

	updatedUser, err := us.userRepo.UpdateUser(ctx, nil, user)
	if err != nil {
		return dto.AllUserResponse{}, dto.ErrUpdateUser
	}

	res := dto.AllUserResponse{
		ID:          updatedUser.ID,
		Name:        updatedUser.Name,
		Email:       updatedUser.Email,
		Password:    updatedUser.Password,
		Image:       updatedUser.Image,
		Gender:      updatedUser.Gender,
		Birthdate:   updatedUser.Birthdate.String(),
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

func (as *UserService) GetAllProvince(ctx context.Context) (dto.ProvincesResponse, error) {
	data, err := as.userRepo.GetAllProvince(ctx, nil)
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

func (as *UserService) GetAllCity(ctx context.Context, req dto.CityQueryRequest) (dto.CitiesResponse, error) {
	data, err := as.userRepo.GetAllCity(ctx, nil, req)
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

func (us *UserService) GetAllNewsWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.NewsPaginationResponse, error) {
	dataWithPaginate, err := us.userRepo.GetAllNewsWithPagination(ctx, nil, req)
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
