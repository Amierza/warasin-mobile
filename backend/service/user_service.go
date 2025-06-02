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
	"github.com/google/uuid"
)

type (
	IUserService interface {
		// Authentication
		Register(ctx context.Context, req dto.UserRegisterRequest) (dto.AllUserResponse, error)
		Login(ctx context.Context, req dto.UserLoginRequest) (dto.UserLoginResponse, error)
		RefreshToken(ctx context.Context, req dto.RefreshTokenRequest) (dto.RefreshTokenResponse, error)

		// Forgot Password
		SendForgotPasswordEmail(ctx context.Context, req dto.SendForgotPasswordEmailRequest) error
		ForgotPassword(ctx context.Context, req dto.ForgotPasswordRequest) (dto.ForgotPasswordResponse, error)
		UpdatePassword(ctx context.Context, req dto.UpdatePasswordRequest) (dto.UpdatePasswordResponse, error)

		// Verification Email
		SendVerificationEmail(ctx context.Context, req dto.SendVerificationEmailRequest) error
		VerifyEmail(ctx context.Context, req dto.VerifyEmailRequest) (dto.VerifyEmailResponse, error)

		// User
		GetDetailUser(ctx context.Context) (dto.AllUserResponse, error)
		UpdateUser(ctx context.Context, req dto.UpdateUserRequest) (dto.AllUserResponse, error)

		// News
		GetAllNewsWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.NewsPaginationResponse, error)
		GetDetailNews(ctx context.Context, newsID string) (dto.NewsResponse, error)

		// Consultation
		CreateConsultation(ctx context.Context, req dto.CreateConsultationRequest) (dto.ConsultationResponse, error)
		GetAllConsultationWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.ConsultationPaginationResponseForUser, error)
		GetDetailConsultation(ctx context.Context, consulID string) (dto.ConsultationResponseForUser, error)
		UpdateConsultation(ctx context.Context, req dto.UpdateConsultationRequestForUser, consulID string) (dto.ConsultationResponseForUser, error)
	}

	UserService struct {
		userRepo   repository.IUserRepository
		masterRepo repository.IMasterRepository
		jwtService IJWTService
	}
)

func NewUserService(userRepo repository.IUserRepository, masterRepo repository.IMasterRepository, jwtService IJWTService) *UserService {
	return &UserService{
		userRepo:   userRepo,
		masterRepo: masterRepo,
		jwtService: jwtService,
	}
}

// Authentication
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

	role, _, err := us.userRepo.GetRoleByName(ctx, nil, "user")
	if err != nil {
		return dto.AllUserResponse{}, dto.ErrGetRoleFromName
	}

	_, flag, err := us.userRepo.GetUserByEmail(ctx, nil, req.Email)
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

	user, flag, err := us.userRepo.GetUserByEmail(ctx, nil, req.Email)
	if !flag || err != nil {
		return dto.UserLoginResponse{}, dto.ErrEmailNotFound
	}

	if user.Role.Name != "user" {
		return dto.UserLoginResponse{}, dto.ErrDeniedAccess
	}

	checkPassword, err := helpers.CheckPassword(user.Password, []byte(req.Password))
	if err != nil || !checkPassword {
		return dto.UserLoginResponse{}, dto.ErrPasswordNotMatch
	}

	permissions, _, err := us.userRepo.GetPermissionsByRoleID(ctx, nil, user.RoleID.String())
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

	role, _, err := us.userRepo.GetRoleByID(ctx, nil, roleID)
	if err != nil {
		return dto.RefreshTokenResponse{}, dto.ErrGetRoleFromID
	}

	if role.Name != "user" {
		return dto.RefreshTokenResponse{}, dto.ErrDeniedAccess
	}

	endpoints, _, err := us.userRepo.GetPermissionsByRoleID(ctx, nil, roleID)
	if err != nil {
		return dto.RefreshTokenResponse{}, dto.ErrGetPermissionsByRoleID
	}

	accessToken, _, err := us.jwtService.GenerateToken(userID, roleID, endpoints)
	if err != nil {
		return dto.RefreshTokenResponse{}, dto.ErrGenerateAccessToken
	}

	return dto.RefreshTokenResponse{AccessToken: accessToken}, nil
}

// Forgot Password
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
	user, flag, err := us.userRepo.GetUserByEmail(ctx, nil, req.Email)
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
	user, flag, err := us.userRepo.GetUserByEmail(ctx, nil, req.Email)
	if err != nil || !flag {
		return dto.UpdatePasswordResponse{}, dto.ErrEmailNotFound
	}

	if len(req.Password) < 8 {
		return dto.UpdatePasswordResponse{}, dto.ErrInvalidPassword
	}

	oldPassword := user.Password

	newPassword, err := helpers.HashPassword(req.Password)
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

// Verification Email
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
func (us *UserService) SendVerificationEmail(ctx context.Context, req dto.SendVerificationEmailRequest) error {
	user, flag, err := us.userRepo.GetUserByEmail(ctx, nil, req.Email)
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

	user, flag, err := us.userRepo.GetUserByEmail(ctx, nil, email)
	if !flag || err != nil {
		return dto.VerifyEmailResponse{}, dto.ErrUserNotFound
	}

	if *user.IsVerified {
		return dto.VerifyEmailResponse{}, dto.ErrEmailAlreadyVerified
	}

	trueValue := true
	updatedUser, err := us.userRepo.UpdateUser(ctx, nil, entity.User{
		ID:         user.ID,
		IsVerified: &trueValue,
	})
	if err != nil {
		return dto.VerifyEmailResponse{}, dto.ErrUpdateUser
	}

	return dto.VerifyEmailResponse{
		Email:      email,
		IsVerified: *updatedUser.IsVerified,
	}, nil
}

// User
func (us *UserService) GetDetailUser(ctx context.Context) (dto.AllUserResponse, error) {
	token := ctx.Value("Authorization").(string)

	userId, err := us.jwtService.GetUserIDByToken(token)
	if err != nil {
		return dto.AllUserResponse{}, dto.ErrGetUserIDFromToken
	}

	user, flag, err := us.userRepo.GetUserByID(ctx, nil, userId)
	if err != nil || !flag {
		return dto.AllUserResponse{}, dto.ErrUserNotFound
	}

	return dto.AllUserResponse{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		Password:    user.Password,
		Image:       user.Image,
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
	}, nil
}
func (us *UserService) UpdateUser(ctx context.Context, req dto.UpdateUserRequest) (dto.AllUserResponse, error) {
	token := ctx.Value("Authorization").(string)
	userID, err := us.jwtService.GetUserIDByToken(token)

	user, flag, err := us.userRepo.GetUserByID(ctx, nil, userID)
	if err != nil || !flag {
		return dto.AllUserResponse{}, dto.ErrGetUserFromID
	}

	if req.CityID != nil {
		city, err := us.masterRepo.GetCityByID(ctx, nil, req.CityID.String())
		if err != nil {
			return dto.AllUserResponse{}, dto.ErrGetCityByID
		}

		user.City = city
	}

	if req.RoleID != nil {
		role, _, err := us.userRepo.GetRoleByID(ctx, nil, req.RoleID.String())
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

		_, flag, err := us.userRepo.GetUserByEmail(ctx, nil, req.Email)
		if flag || err == nil {
			return dto.AllUserResponse{}, dto.ErrEmailAlreadyExists
		}

		user.Email = req.Email
		falseValue := false
		user.IsVerified = &falseValue
	}

	if req.Image != "" {
		user.Image = req.Image
	}

	if req.Gender != nil {
		user.Gender = req.Gender
	}

	if req.Birthdate != "" {
		t, err := helpers.ValidateAndNormalizeDateString(req.Birthdate)
		if err != nil {
			return dto.AllUserResponse{}, dto.ErrFormatBirthdate
		}
		user.Birthdate = t
	}

	if req.PhoneNumber != "" {
		phoneNumberFormatted, err := helpers.StandardizePhoneNumber(req.PhoneNumber, true)
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

// News
func (us *UserService) GetAllNewsWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.NewsPaginationResponse, error) {
	dataWithPaginate, err := us.userRepo.GetAllNewsWithPagination(ctx, nil, req)
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
func (us *UserService) GetDetailNews(ctx context.Context, newsID string) (dto.NewsResponse, error) {
	news, _, err := us.userRepo.GetNewsByID(ctx, nil, newsID)
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

// Consultation
func (us *UserService) CreateConsultation(ctx context.Context, req dto.CreateConsultationRequest) (dto.ConsultationResponse, error) {
	token := ctx.Value("Authorization").(string)

	userID, err := us.jwtService.GetUserIDByToken(token)
	if err != nil {
		return dto.ConsultationResponse{}, dto.ErrGetUserIDFromToken
	}

	u, flag, err := us.userRepo.GetUserByID(ctx, nil, userID)
	if err != nil || !flag {
		return dto.ConsultationResponse{}, dto.ErrUserNotFound
	}

	a, flag, err := us.userRepo.GetAvailableSlotByID(ctx, nil, req.AvailableSlotID)
	if err != nil || !flag {
		return dto.ConsultationResponse{}, dto.ErrAvailableSlotNotFound
	}

	p, flag, err := us.userRepo.GetPracticeByID(ctx, nil, req.PracticeID)
	if err != nil || !flag {
		return dto.ConsultationResponse{}, dto.ErrPracticeNotFound
	}

	consultation := entity.Consultation{
		ID:              uuid.New(),
		Date:            req.Date,
		Rate:            0,
		Comment:         "",
		Status:          0,
		UserID:          &u.ID,
		PracticeID:      &p.ID,
		AvailableSlotID: &a.ID,
	}

	err = us.userRepo.CreateConsultation(ctx, nil, consultation)
	if err != nil {
		return dto.ConsultationResponse{}, dto.ErrCreateConsultation
	}

	a.IsBooked = true
	err = us.userRepo.UpdateStatusBookSlot(ctx, nil, a.ID, a.IsBooked)
	if err != nil {
		return dto.ConsultationResponse{}, dto.ErrUpdateStatusBookSlot
	}

	user := dto.AllUserResponse{
		ID:          u.ID,
		Name:        u.Name,
		Email:       u.Email,
		Password:    u.Password,
		Birthdate:   u.Birthdate,
		PhoneNumber: u.PhoneNumber,
		Data01:      u.Data01,
		Data02:      u.Data02,
		Data03:      u.Data03,
		IsVerified:  u.IsVerified,
		City: dto.CityResponse{
			ID:   &u.City.ID,
			Name: u.City.Name,
			Type: u.City.Type,
			Province: dto.ProvinceResponse{
				ID:   u.City.ProvinceID,
				Name: u.City.Province.Name,
			},
		},
		Role: dto.RoleResponse{
			ID:   &u.Role.ID,
			Name: u.Role.Name,
		},
	}

	availableSlot := dto.AvailableSlotResponse{
		ID:       a.ID,
		Start:    a.Start,
		End:      a.End,
		IsBooked: a.IsBooked,
	}

	practice := dto.PracticeResponse{
		ID:          p.ID,
		Type:        p.Type,
		Name:        p.Name,
		Address:     p.Address,
		PhoneNumber: p.PhoneNumber,
	}

	dayName, err := helpers.GetDayName(consultation.Date)
	if err != nil {
		return dto.ConsultationResponse{}, dto.ErrParseConsultationDate
	}

	var matchedSchedules []dto.PracticeScheduleResponse
	for _, prac := range p.PracticeSchedules {
		if dayName == prac.Day {
			matchedSchedules = append(matchedSchedules, dto.PracticeScheduleResponse{
				ID:    prac.ID,
				Day:   prac.Day,
				Open:  prac.Open,
				Close: prac.Close,
			})
		}
	}
	practice.PracticeSchedules = matchedSchedules

	return dto.ConsultationResponse{
		ID:            consultation.ID,
		Date:          consultation.Date,
		Rate:          consultation.Rate,
		Comment:       consultation.Comment,
		Status:        consultation.Status,
		User:          user,
		AvailableSlot: availableSlot,
		Practice:      practice,
	}, nil
}
func (us *UserService) GetAllConsultationWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.ConsultationPaginationResponseForUser, error) {
	token := ctx.Value("Authorization").(string)

	userID, err := us.jwtService.GetUserIDByToken(token)
	if err != nil {
		return dto.ConsultationPaginationResponseForUser{}, dto.ErrGetUserIDFromToken
	}

	dataWithPaginate, err := us.userRepo.GetAllConsultationWithPagination(ctx, nil, req, userID)
	if err != nil {
		return dto.ConsultationPaginationResponseForUser{}, dto.ErrGetAllConsultationWithPagination
	}

	var (
		user          dto.AllUserResponse
		consultations []dto.ConsultationResponseForUser
	)

	user = dto.AllUserResponse{
		ID:          dataWithPaginate.Consultations[0].User.ID,
		Name:        dataWithPaginate.Consultations[0].User.Name,
		Email:       dataWithPaginate.Consultations[0].User.Email,
		Password:    dataWithPaginate.Consultations[0].User.Password,
		Birthdate:   dataWithPaginate.Consultations[0].User.Birthdate,
		PhoneNumber: dataWithPaginate.Consultations[0].User.PhoneNumber,
		Data01:      dataWithPaginate.Consultations[0].User.Data01,
		Data02:      dataWithPaginate.Consultations[0].User.Data02,
		Data03:      dataWithPaginate.Consultations[0].User.Data03,
		IsVerified:  dataWithPaginate.Consultations[0].User.IsVerified,
		City: dto.CityResponse{
			ID:   &dataWithPaginate.Consultations[0].User.City.ID,
			Name: dataWithPaginate.Consultations[0].User.City.Name,
			Type: dataWithPaginate.Consultations[0].User.City.Type,
			Province: dto.ProvinceResponse{
				ID:   dataWithPaginate.Consultations[0].User.City.ProvinceID,
				Name: dataWithPaginate.Consultations[0].User.City.Province.Name,
			},
		},
		Role: dto.RoleResponse{
			ID:   &dataWithPaginate.Consultations[0].User.Role.ID,
			Name: dataWithPaginate.Consultations[0].User.Role.Name,
		},
	}

	for _, consultation := range dataWithPaginate.Consultations {
		dayName, err := helpers.GetDayName(consultation.Date)
		if err != nil {
			return dto.ConsultationPaginationResponseForUser{}, dto.ErrParseConsultationDate
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

		data := dto.ConsultationResponseForUser{
			ID:      consultation.ID,
			Date:    consultation.Date,
			Rate:    consultation.Rate,
			Comment: consultation.Comment,
			Status:  consultation.Status,
			Psycholog: dto.PsychologResponse{
				ID:          consultation.AvailableSlot.Psycholog.ID,
				Name:        consultation.AvailableSlot.Psycholog.Name,
				STRNumber:   consultation.AvailableSlot.Psycholog.STRNumber,
				Email:       consultation.AvailableSlot.Psycholog.Email,
				Password:    consultation.AvailableSlot.Psycholog.Password,
				WorkYear:    consultation.AvailableSlot.Psycholog.WorkYear,
				Description: consultation.AvailableSlot.Psycholog.Description,
				PhoneNumber: consultation.AvailableSlot.Psycholog.PhoneNumber,
				Image:       consultation.AvailableSlot.Psycholog.Image,
				City: dto.CityResponse{
					ID:   consultation.AvailableSlot.Psycholog.CityID,
					Name: consultation.AvailableSlot.Psycholog.City.Name,
					Type: consultation.AvailableSlot.Psycholog.City.Type,
					Province: dto.ProvinceResponse{
						ID:   consultation.AvailableSlot.Psycholog.City.ProvinceID,
						Name: consultation.AvailableSlot.Psycholog.City.Province.Name,
					},
				},
				Role: dto.RoleResponse{
					ID:   consultation.AvailableSlot.Psycholog.RoleID,
					Name: consultation.AvailableSlot.Psycholog.Role.Name,
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

		// LanguageMasters
		for _, lang := range consultation.AvailableSlot.Psycholog.PsychologLanguages {
			data.Psycholog.LanguageMasters = append(data.Psycholog.LanguageMasters, dto.LanguageMasterResponse{
				ID:   &lang.LanguageMaster.ID,
				Name: lang.LanguageMaster.Name,
			})
		}

		// Specializations
		for _, spec := range consultation.AvailableSlot.Psycholog.PsychologSpecializations {
			data.Psycholog.Specializations = append(data.Psycholog.Specializations, dto.SpecializationResponse{
				ID:          &spec.Specialization.ID,
				Name:        spec.Specialization.Name,
				Description: spec.Specialization.Description,
			})
		}

		// Educations
		for _, edu := range consultation.AvailableSlot.Psycholog.Educations {
			data.Psycholog.Educations = append(data.Psycholog.Educations, dto.EducationResponse{
				ID:             &edu.ID,
				Degree:         edu.Degree,
				Major:          edu.Major,
				Institution:    edu.Institution,
				GraduationYear: edu.GraduationYear,
			})
		}

		consultations = append(consultations, data)
	}

	datas := dto.AllConsultationResponseForUser{
		User:         user,
		Consultation: consultations,
	}

	return dto.ConsultationPaginationResponseForUser{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}
func (us *UserService) GetDetailConsultation(ctx context.Context, consulID string) (dto.ConsultationResponseForUser, error) {
	consultation, _, err := us.userRepo.GetConsultationByID(ctx, nil, consulID)
	if err != nil {
		return dto.ConsultationResponseForUser{}, dto.ErrConsultationNotFound
	}

	dayName, err := helpers.GetDayName(consultation.Date)
	if err != nil {
		return dto.ConsultationResponseForUser{}, dto.ErrParseConsultationDate
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

	data := dto.ConsultationResponseForUser{
		ID:      consultation.ID,
		Date:    consultation.Date,
		Rate:    consultation.Rate,
		Comment: consultation.Comment,
		Status:  consultation.Status,
		Psycholog: dto.PsychologResponse{
			ID:          consultation.AvailableSlot.Psycholog.ID,
			Name:        consultation.AvailableSlot.Psycholog.Name,
			STRNumber:   consultation.AvailableSlot.Psycholog.STRNumber,
			Email:       consultation.AvailableSlot.Psycholog.Email,
			Password:    consultation.AvailableSlot.Psycholog.Password,
			WorkYear:    consultation.AvailableSlot.Psycholog.WorkYear,
			Description: consultation.AvailableSlot.Psycholog.Description,
			PhoneNumber: consultation.AvailableSlot.Psycholog.PhoneNumber,
			Image:       consultation.AvailableSlot.Psycholog.Image,
			City: dto.CityResponse{
				ID:   consultation.AvailableSlot.Psycholog.CityID,
				Name: consultation.AvailableSlot.Psycholog.City.Name,
				Type: consultation.AvailableSlot.Psycholog.City.Type,
				Province: dto.ProvinceResponse{
					ID:   consultation.AvailableSlot.Psycholog.City.ProvinceID,
					Name: consultation.AvailableSlot.Psycholog.City.Province.Name,
				},
			},
			Role: dto.RoleResponse{
				ID:   consultation.AvailableSlot.Psycholog.RoleID,
				Name: consultation.AvailableSlot.Psycholog.Role.Name,
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

	// LanguageMasters
	for _, lang := range consultation.AvailableSlot.Psycholog.PsychologLanguages {
		data.Psycholog.LanguageMasters = append(data.Psycholog.LanguageMasters, dto.LanguageMasterResponse{
			ID:   &lang.LanguageMaster.ID,
			Name: lang.LanguageMaster.Name,
		})
	}

	// Specializations
	for _, spec := range consultation.AvailableSlot.Psycholog.PsychologSpecializations {
		data.Psycholog.Specializations = append(data.Psycholog.Specializations, dto.SpecializationResponse{
			ID:          &spec.Specialization.ID,
			Name:        spec.Specialization.Name,
			Description: spec.Specialization.Description,
		})
	}

	// Educations
	for _, edu := range consultation.AvailableSlot.Psycholog.Educations {
		data.Psycholog.Educations = append(data.Psycholog.Educations, dto.EducationResponse{
			ID:             &edu.ID,
			Degree:         edu.Degree,
			Major:          edu.Major,
			Institution:    edu.Institution,
			GraduationYear: edu.GraduationYear,
		})
	}

	return data, nil
}
func (us *UserService) UpdateConsultation(ctx context.Context, req dto.UpdateConsultationRequestForUser, consulID string) (dto.ConsultationResponseForUser, error) {
	consul, flag, err := us.userRepo.GetConsultationByID(ctx, nil, consulID)
	if err != nil || !flag {
		return dto.ConsultationResponseForUser{}, dto.ErrConsultationNotFound
	}

	if req.Date != "" {
		date, err := helpers.ValidateAndNormalizeDateString(req.Date)
		if err != nil {
			return dto.ConsultationResponseForUser{}, dto.ErrParseConsultationDate
		}

		consul.Date = date
	}

	if req.Rate != nil {
		valid := false
		switch *req.Rate {
		case 1, 2, 3, 4, 5:
			valid = true
		}
		if !valid {
			return dto.ConsultationResponseForUser{}, dto.ErrInvalidRateConsultation
		}

		consul.Rate = *req.Rate
	}

	if req.Comment != "" {
		if len(req.Comment) < 5 {
			return dto.ConsultationResponseForUser{}, dto.ErrConsultationCommentToShort
		}

		consul.Comment = req.Comment
	}

	if req.AvailableSlotID != "" {
		slot, flag, err := us.userRepo.GetAvailableSlotByID(ctx, nil, req.AvailableSlotID)
		if err != nil || !flag {
			return dto.ConsultationResponseForUser{}, dto.ErrAvailableSlotNotFound
		}

		err = us.userRepo.UpdateStatusBookSlot(ctx, nil, *consul.AvailableSlotID, false)
		if err != nil {
			return dto.ConsultationResponseForUser{}, dto.ErrUpdateStatusBookSlot
		}

		slotID, err := uuid.Parse(req.AvailableSlotID)
		if err != nil {
			return dto.ConsultationResponseForUser{}, dto.ErrParseUUID
		}

		err = us.userRepo.UpdateStatusBookSlot(ctx, nil, slotID, true)
		if err != nil {
			return dto.ConsultationResponseForUser{}, dto.ErrUpdateStatusBookSlot
		}

		slot.IsBooked = true
		consul.AvailableSlot = slot
		consul.AvailableSlotID = &slotID
	}

	if req.PracticeID != "" {
		prac, flag, err := us.userRepo.GetPracticeByID(ctx, nil, req.PracticeID)
		if err != nil || !flag {
			return dto.ConsultationResponseForUser{}, dto.ErrPracticeNotFound
		}

		pracID, err := uuid.Parse(req.PracticeID)
		if err != nil {
			return dto.ConsultationResponseForUser{}, dto.ErrParseUUID
		}

		consul.Practice = prac
		consul.PracticeID = &pracID
	}

	dayName, err := helpers.GetDayName(consul.Date)
	if err != nil {
		return dto.ConsultationResponseForUser{}, dto.ErrParseConsultationDate
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

	data := dto.ConsultationResponseForUser{
		ID:      consul.ID,
		Date:    consul.Date,
		Rate:    consul.Rate,
		Comment: consul.Comment,
		Status:  consul.Status,
		Psycholog: dto.PsychologResponse{
			ID:          consul.AvailableSlot.Psycholog.ID,
			Name:        consul.AvailableSlot.Psycholog.Name,
			STRNumber:   consul.AvailableSlot.Psycholog.STRNumber,
			Email:       consul.AvailableSlot.Psycholog.Email,
			Password:    consul.AvailableSlot.Psycholog.Password,
			WorkYear:    consul.AvailableSlot.Psycholog.WorkYear,
			Description: consul.AvailableSlot.Psycholog.Description,
			PhoneNumber: consul.AvailableSlot.Psycholog.PhoneNumber,
			Image:       consul.AvailableSlot.Psycholog.Image,
			City: dto.CityResponse{
				ID:   consul.AvailableSlot.Psycholog.CityID,
				Name: consul.AvailableSlot.Psycholog.City.Name,
				Type: consul.AvailableSlot.Psycholog.City.Type,
				Province: dto.ProvinceResponse{
					ID:   consul.AvailableSlot.Psycholog.City.ProvinceID,
					Name: consul.AvailableSlot.Psycholog.City.Province.Name,
				},
			},
			Role: dto.RoleResponse{
				ID:   consul.AvailableSlot.Psycholog.RoleID,
				Name: consul.AvailableSlot.Psycholog.Role.Name,
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

	// LanguageMasters
	for _, lang := range consul.AvailableSlot.Psycholog.PsychologLanguages {
		data.Psycholog.LanguageMasters = append(data.Psycholog.LanguageMasters, dto.LanguageMasterResponse{
			ID:   &lang.LanguageMaster.ID,
			Name: lang.LanguageMaster.Name,
		})
	}

	// Specializations
	for _, spec := range consul.AvailableSlot.Psycholog.PsychologSpecializations {
		data.Psycholog.Specializations = append(data.Psycholog.Specializations, dto.SpecializationResponse{
			ID:          &spec.Specialization.ID,
			Name:        spec.Specialization.Name,
			Description: spec.Specialization.Description,
		})
	}

	// Educations
	for _, edu := range consul.AvailableSlot.Psycholog.Educations {
		data.Psycholog.Educations = append(data.Psycholog.Educations, dto.EducationResponse{
			ID:             &edu.ID,
			Degree:         edu.Degree,
			Major:          edu.Major,
			Institution:    edu.Institution,
			GraduationYear: edu.GraduationYear,
		})
	}

	err = us.userRepo.UpdateConsultation(ctx, nil, consul)
	if err != nil {
		return dto.ConsultationResponseForUser{}, dto.ErrUpdateConsultation
	}

	return data, nil
}
