package repository

import (
	"context"
	"math"
	"strings"

	"github.com/Amierza/warasin-mobile/backend/dto"
	"github.com/Amierza/warasin-mobile/backend/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	IAdminRepository interface {
		// Get
		GetUserByEmail(ctx context.Context, tx *gorm.DB, email string) (entity.User, bool, error)
		GetUserByID(ctx context.Context, tx *gorm.DB, userID string) (entity.User, error)
		GetRoleByID(ctx context.Context, tx *gorm.DB, roleID string) (entity.Role, error)
		GetPermissionsByRoleID(ctx context.Context, tx *gorm.DB, roleID string) ([]string, error)
		GetAllUserWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.AllUserRepositoryResponse, error)
		GetAllNewsWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.AllNewsRepositoryResponse, error)
		GetNewsByID(ctx context.Context, tx *gorm.DB, newsID string) (entity.News, error)
		GetNewsByTitle(ctx context.Context, tx *gorm.DB, title string) (bool, entity.News, error)
		GetAllMotivationCategoryWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.AllMotivationCategoryRepositoryResponse, error)
		GetMotivationCategoryByID(ctx context.Context, tx *gorm.DB, motivationCategoryID string) (entity.MotivationCategory, error)
		GetMotivationCategoryByName(ctx context.Context, tx *gorm.DB, motivationCategoryName string) (bool, entity.MotivationCategory, error)
		GetAllMotivationWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.AllMotivationRepositoryResponse, error)
		GetMotivationByID(ctx context.Context, tx *gorm.DB, motivationID string) (entity.Motivation, error)
		GetMotivationByContent(ctx context.Context, tx *gorm.DB, content string) (bool, entity.Motivation, error)
		GetAllRole(ctx context.Context, tx *gorm.DB) (dto.AllRoleRepositoryResponse, error)
		GetAllPsychologWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.AllPsychologRepositoryResponse, error)
		GetAllUserMotivationWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.AllUserMotivationRepositoryResponse, error)
		GetAllUserNewsWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.AllUserNewsRepositoryResponse, error)
		GetAllConsultationWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.AllConsultationRepositoryResponse, error)
		GetLanguageMasterByID(ctx context.Context, tx *gorm.DB, languageMasterID string) (entity.LanguageMaster, bool, error)
		GetSpecializationByID(ctx context.Context, tx *gorm.DB, specializationID string) (entity.Specialization, bool, error)
		GetEducationByID(ctx context.Context, tx *gorm.DB, eduID string) (entity.Education, bool, error)

		// Create
		CreateUser(ctx context.Context, tx *gorm.DB, user entity.User) error
		CreateNews(ctx context.Context, tx *gorm.DB, news entity.News) error
		CreateMotivationCategory(ctx context.Context, tx *gorm.DB, motivationCategory entity.MotivationCategory) error
		CreateMotivation(ctx context.Context, tx *gorm.DB, motivation entity.Motivation) error
		CreatePsycholog(ctx context.Context, tx *gorm.DB, psycholog entity.Psycholog) error
		CreatePsychologLanguages(ctx context.Context, tx *gorm.DB, psychologLanguages []entity.PsychologLanguage) error
		CreatePsychologSpecializations(ctx context.Context, tx *gorm.DB, psychologSpecializations []entity.PsychologSpecialization) error
		CreateEducations(ctx context.Context, tx *gorm.DB, educations []entity.Education) error

		// Update
		UpdateUser(ctx context.Context, tx *gorm.DB, user entity.User) error
		UpdateNews(ctx context.Context, tx *gorm.DB, user entity.News) error
		UpdateMotivationCategory(ctx context.Context, tx *gorm.DB, motivationCategory entity.MotivationCategory) error
		UpdateMotivation(ctx context.Context, tx *gorm.DB, motivation entity.Motivation) error
		UpdatePsycholog(ctx context.Context, tx *gorm.DB, psycholog entity.Psycholog) error

		// Delete
		DeleteUserByID(ctx context.Context, tx *gorm.DB, userID string) error
		DeleteNewsByID(ctx context.Context, tx *gorm.DB, newsID string) error
		DeleteMotivationCategoryByID(ctx context.Context, tx *gorm.DB, motivationCategoryID string) error
		DeleteMotivationByID(ctx context.Context, tx *gorm.DB, motivationID string) error
		DeletePsychologByID(ctx context.Context, tx *gorm.DB, psychologID string) error
		DeletePsychologLanguageByPsychologID(ctx context.Context, tx *gorm.DB, psychologID string) error
		DeletePsychologSpecializationByPsychologID(ctx context.Context, tx *gorm.DB, psychologID string) error
		DeleteEducationByPsychologID(ctx context.Context, tx *gorm.DB, psychologID string) error
	}

	AdminRepository struct {
		db *gorm.DB
	}
)

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{
		db: db,
	}
}

// Get
func (ar *AdminRepository) GetUserByEmail(ctx context.Context, tx *gorm.DB, email string) (entity.User, bool, error) {
	if tx == nil {
		tx = ar.db
	}

	var user entity.User
	if err := tx.WithContext(ctx).Preload("Role").Preload("City").Where("email = ?", email).Take(&user).Error; err != nil {
		return entity.User{}, false, err
	}

	return user, true, nil
}
func (ar *AdminRepository) GetUserByID(ctx context.Context, tx *gorm.DB, userID string) (entity.User, error) {
	if tx == nil {
		tx = ar.db
	}

	var user entity.User
	if err := tx.WithContext(ctx).Preload("City.Province").Preload("Role").Where("id = ?", userID).Take(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}
func (ar *AdminRepository) GetRoleByID(ctx context.Context, tx *gorm.DB, roleID string) (entity.Role, error) {
	if tx == nil {
		tx = ar.db
	}

	var role entity.Role
	if err := tx.WithContext(ctx).Where("id = ?", roleID).Take(&role).Error; err != nil {
		return entity.Role{}, err
	}

	return role, nil
}
func (ar *AdminRepository) GetPermissionsByRoleID(ctx context.Context, tx *gorm.DB, roleID string) ([]string, error) {
	if tx == nil {
		tx = ar.db
	}

	var endpoints []string
	if err := tx.WithContext(ctx).Table("permissions").Where("role_id = ?", roleID).Pluck("endpoint", &endpoints).Error; err != nil {
		return []string{}, err
	}

	return endpoints, nil
}
func (ar *AdminRepository) GetAllUserWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.AllUserRepositoryResponse, error) {
	if tx == nil {
		tx = ar.db
	}

	var users []entity.User
	var err error
	var count int64

	if req.PerPage == 0 {
		req.PerPage = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	var adminIDs []uuid.UUID
	if err := tx.WithContext(ctx).Model(&entity.Role{}).Where("name != ?", "admin").Pluck("id", &adminIDs).Error; err != nil {
		return dto.AllUserRepositoryResponse{}, err
	}

	query := tx.WithContext(ctx).Model(&entity.User{}).Where("role_id IN (?)", adminIDs)

	if req.Search != "" {
		searchValue := "%" + strings.ToLower(req.Search) + "%"
		query = query.Where("LOWER(name) LIKE ? OR LOWER(email) LIKE ? OR LOWER(phone_number) LIKE ?",
			searchValue, searchValue, searchValue)
	}

	query = query.Preload("City.Province").Preload("Role")

	if err := query.Count(&count).Error; err != nil {
		return dto.AllUserRepositoryResponse{}, err
	}

	if err := query.Order("created_at DESC").Scopes(Paginate(req.Page, req.PerPage)).Find(&users).Error; err != nil {
		return dto.AllUserRepositoryResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(req.PerPage)))

	return dto.AllUserRepositoryResponse{
		Users: users,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			MaxPage: totalPage,
			Count:   count,
		},
	}, err
}
func (ar *AdminRepository) GetAllNewsWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.AllNewsRepositoryResponse, error) {
	if tx == nil {
		tx = ar.db
	}

	var news []entity.News
	var err error
	var count int64

	if req.PerPage == 0 {
		req.PerPage = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	query := tx.WithContext(ctx).Model(&entity.News{})

	if req.Search != "" {
		searchValue := "%" + strings.ToLower(req.Search) + "%"
		query = query.Where("LOWER(image) LIKE ? OR LOWER(title) LIKE ? OR LOWER(body) LIKE ?",
			searchValue, searchValue, searchValue)
	}

	if err := query.Count(&count).Error; err != nil {
		return dto.AllNewsRepositoryResponse{}, err
	}

	if err := query.Order("created_at DESC").Scopes(Paginate(req.Page, req.PerPage)).Find(&news).Error; err != nil {
		return dto.AllNewsRepositoryResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(req.PerPage)))

	return dto.AllNewsRepositoryResponse{
		News: news,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			MaxPage: totalPage,
			Count:   count,
		},
	}, err
}
func (ar *AdminRepository) GetNewsByID(ctx context.Context, tx *gorm.DB, newsID string) (entity.News, error) {
	if tx == nil {
		tx = ar.db
	}

	var news entity.News
	if err := tx.WithContext(ctx).Where("id = ?", newsID).Take(&news).Error; err != nil {
		return entity.News{}, err
	}

	return news, nil
}
func (ar *AdminRepository) GetNewsByTitle(ctx context.Context, tx *gorm.DB, title string) (bool, entity.News, error) {
	if tx == nil {
		tx = ar.db
	}

	var news entity.News
	if err := tx.WithContext(ctx).Where("title = ?", title).Take(&news).Error; err != nil {
		return false, entity.News{}, err
	}

	return true, news, nil
}
func (ar *AdminRepository) GetAllMotivationCategoryWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.AllMotivationCategoryRepositoryResponse, error) {
	if tx == nil {
		tx = ar.db
	}

	var motivationCategories []entity.MotivationCategory
	var err error
	var count int64

	if req.PerPage == 0 {
		req.PerPage = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	query := tx.WithContext(ctx).Model(&entity.MotivationCategory{})

	if req.Search != "" {
		searchValue := "%" + strings.ToLower(req.Search) + "%"
		query = query.Where("LOWER(name) LIKE ?", searchValue)
	}

	if err := query.Count(&count).Error; err != nil {
		return dto.AllMotivationCategoryRepositoryResponse{}, err
	}

	if err := query.Order("created_at DESC").Scopes(Paginate(req.Page, req.PerPage)).Find(&motivationCategories).Error; err != nil {
		return dto.AllMotivationCategoryRepositoryResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(req.PerPage)))

	return dto.AllMotivationCategoryRepositoryResponse{
		MotivationCategories: motivationCategories,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			MaxPage: totalPage,
			Count:   count,
		},
	}, err
}
func (ar *AdminRepository) GetMotivationCategoryByID(ctx context.Context, tx *gorm.DB, motivationCategoryID string) (entity.MotivationCategory, error) {
	if tx == nil {
		tx = ar.db
	}

	var motivationCategory entity.MotivationCategory
	if err := tx.WithContext(ctx).Where("id = ?", motivationCategoryID).Take(&motivationCategory).Error; err != nil {
		return entity.MotivationCategory{}, err
	}

	return motivationCategory, nil
}
func (ar *AdminRepository) GetMotivationCategoryByName(ctx context.Context, tx *gorm.DB, motivationCategoryName string) (bool, entity.MotivationCategory, error) {
	if tx == nil {
		tx = ar.db
	}

	var motivationCategory entity.MotivationCategory
	if err := tx.WithContext(ctx).Where("name = ?", motivationCategoryName).Take(&motivationCategory).Error; err != nil {
		return false, entity.MotivationCategory{}, err
	}

	return true, motivationCategory, nil
}
func (ar *AdminRepository) GetAllMotivationWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.AllMotivationRepositoryResponse, error) {
	if tx == nil {
		tx = ar.db
	}

	var motivations []entity.Motivation
	var err error
	var count int64

	if req.PerPage == 0 {
		req.PerPage = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	query := tx.WithContext(ctx).Model(&entity.Motivation{}).Preload("MotivationCategory")

	if req.Search != "" {
		searchValue := "%" + strings.ToLower(req.Search) + "%"
		query = query.Where("LOWER(name) LIKE ?", searchValue)
	}

	if err := query.Count(&count).Error; err != nil {
		return dto.AllMotivationRepositoryResponse{}, err
	}

	if err := query.Order("created_at DESC").Scopes(Paginate(req.Page, req.PerPage)).Find(&motivations).Error; err != nil {
		return dto.AllMotivationRepositoryResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(req.PerPage)))

	return dto.AllMotivationRepositoryResponse{
		Motivations: motivations,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			MaxPage: totalPage,
			Count:   count,
		},
	}, err
}
func (ar *AdminRepository) GetMotivationByID(ctx context.Context, tx *gorm.DB, motivationID string) (entity.Motivation, error) {
	if tx == nil {
		tx = ar.db
	}

	var motivation entity.Motivation
	if err := tx.WithContext(ctx).Preload("MotivationCategory").Where("id = ?", motivationID).Take(&motivation).Error; err != nil {
		return entity.Motivation{}, err
	}

	return motivation, nil
}
func (ar *AdminRepository) GetMotivationByContent(ctx context.Context, tx *gorm.DB, content string) (bool, entity.Motivation, error) {
	if tx == nil {
		tx = ar.db
	}

	var motivation entity.Motivation
	if err := tx.WithContext(ctx).Preload("MotivationCategory").Where("content = ?", content).Take(&motivation).Error; err != nil {
		return false, entity.Motivation{}, err
	}

	return true, motivation, nil
}
func (ar *AdminRepository) GetAllRole(ctx context.Context, tx *gorm.DB) (dto.AllRoleRepositoryResponse, error) {
	if tx == nil {
		tx = ar.db
	}

	var (
		roles []entity.Role
		err   error
	)

	if err := tx.WithContext(ctx).Model(&entity.Role{}).Find(&roles).Error; err != nil {
		return dto.AllRoleRepositoryResponse{}, err
	}

	return dto.AllRoleRepositoryResponse{
		Roles: roles,
	}, err
}
func (ar *AdminRepository) GetAllPsychologWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.AllPsychologRepositoryResponse, error) {
	if tx == nil {
		tx = ar.db
	}

	var psychologs []entity.Psycholog
	var err error
	var count int64

	if req.PerPage == 0 {
		req.PerPage = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	query := tx.WithContext(ctx).Model(&entity.Psycholog{}).
		Preload("Role").
		Preload("City.Province").
		Preload("PsychologLanguages.LanguageMaster").
		Preload("PsychologSpecializations.Specialization").
		Preload("Educations")

	if req.Search != "" {
		searchValue := "%" + strings.ToLower(req.Search) + "%"
		query = query.Where("LOWER(name) LIKE ? OR LOWER(email) LIKE ?", searchValue, searchValue)
	}

	if err := query.Count(&count).Error; err != nil {
		return dto.AllPsychologRepositoryResponse{}, err
	}

	if err := query.Order("created_at DESC").Scopes(Paginate(req.Page, req.PerPage)).Find(&psychologs).Error; err != nil {
		return dto.AllPsychologRepositoryResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(req.PerPage)))

	return dto.AllPsychologRepositoryResponse{
		Psychologs: psychologs,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			MaxPage: totalPage,
			Count:   count,
		},
	}, err
}
func (ar *AdminRepository) GetAllUserMotivationWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.AllUserMotivationRepositoryResponse, error) {
	if tx == nil {
		tx = ar.db
	}

	var (
		userMotivations []entity.UserMotivation
		err             error
		count           int64
	)

	if req.PerPage == 0 {
		req.PerPage = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	query := tx.WithContext(ctx).Model(&entity.UserMotivation{}).
		Preload("User.Role").
		Preload("User.City.Province").
		Preload("Motivation.MotivationCategory")

	// if req.Search != "" {
	// 	searchValue := "%" + strings.ToLower(req.Search) + "%"
	// 	query = query.Where("LOWER(name) LIKE ?", searchValue)
	// }

	if err := query.Count(&count).Error; err != nil {
		return dto.AllUserMotivationRepositoryResponse{}, err
	}

	if err := query.Order("created_at DESC").Scopes(Paginate(req.Page, req.PerPage)).Find(&userMotivations).Error; err != nil {
		return dto.AllUserMotivationRepositoryResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(req.PerPage)))

	return dto.AllUserMotivationRepositoryResponse{
		UserMotivations: userMotivations,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			MaxPage: totalPage,
			Count:   count,
		},
	}, err
}
func (ar *AdminRepository) GetAllUserNewsWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.AllUserNewsRepositoryResponse, error) {
	if tx == nil {
		tx = ar.db
	}

	var (
		userNews []entity.NewsDetail
		err      error
		count    int64
	)

	if req.PerPage == 0 {
		req.PerPage = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	query := tx.WithContext(ctx).Model(&entity.NewsDetail{}).
		Preload("User.Role").
		Preload("User.City.Province").
		Preload("News")

	if err := query.Count(&count).Error; err != nil {
		return dto.AllUserNewsRepositoryResponse{}, err
	}

	if err := query.Order("created_at DESC").Scopes(Paginate(req.Page, req.PerPage)).Find(&userNews).Error; err != nil {
		return dto.AllUserNewsRepositoryResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(req.PerPage)))

	return dto.AllUserNewsRepositoryResponse{
		UserNews: userNews,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			MaxPage: totalPage,
			Count:   count,
		},
	}, err
}
func (ar *AdminRepository) GetAllConsultationWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.AllConsultationRepositoryResponse, error) {
	if tx == nil {
		tx = ar.db
	}

	var consultations []entity.Consulation
	var err error
	var count int64

	if req.PerPage == 0 {
		req.PerPage = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	query := tx.WithContext(ctx).Model(&entity.Consulation{}).
		Preload("User.Role").
		Preload("User.City.Province").
		Preload("Psycholog.Role").
		Preload("Psycholog.City.Province")

	// if req.Search != "" {
	// 	query = query.Where("rate = ?", req.Search)
	// }

	if err := query.Count(&count).Error; err != nil {
		return dto.AllConsultationRepositoryResponse{}, err
	}

	if err := query.Order("created_at DESC").Scopes(Paginate(req.Page, req.PerPage)).Find(&consultations).Error; err != nil {
		return dto.AllConsultationRepositoryResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(req.PerPage)))

	return dto.AllConsultationRepositoryResponse{
		Consultations: consultations,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			MaxPage: totalPage,
			Count:   count,
		},
	}, err
}
func (ar *AdminRepository) GetLanguageMasterByID(ctx context.Context, tx *gorm.DB, languageMasterID string) (entity.LanguageMaster, bool, error) {
	if tx == nil {
		tx = ar.db
	}

	var languageMaster entity.LanguageMaster
	if err := tx.WithContext(ctx).Where("id = ?", languageMasterID).Take(&languageMaster).Error; err != nil {
		return entity.LanguageMaster{}, false, err
	}

	return languageMaster, true, nil
}
func (ar *AdminRepository) GetSpecializationByID(ctx context.Context, tx *gorm.DB, specializationID string) (entity.Specialization, bool, error) {
	if tx == nil {
		tx = ar.db
	}

	var specialization entity.Specialization
	if err := tx.WithContext(ctx).Where("id = ?", specializationID).Take(&specialization).Error; err != nil {
		return entity.Specialization{}, false, err
	}

	return specialization, true, nil
}
func (ar *AdminRepository) GetEducationByID(ctx context.Context, tx *gorm.DB, educationID string) (entity.Education, bool, error) {
	if tx == nil {
		tx = ar.db
	}

	var education entity.Education
	if err := tx.WithContext(ctx).Where("id = ?", educationID).Take(&education).Error; err != nil {
		return entity.Education{}, false, err
	}

	return education, true, nil
}

// Create
func (ar *AdminRepository) CreateUser(ctx context.Context, tx *gorm.DB, user entity.User) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Create(&user).Error
}
func (ar *AdminRepository) CreateNews(ctx context.Context, tx *gorm.DB, news entity.News) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Create(&news).Error
}
func (ar *AdminRepository) CreateMotivationCategory(ctx context.Context, tx *gorm.DB, motivationCategory entity.MotivationCategory) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Create(&motivationCategory).Error
}
func (ar *AdminRepository) CreateMotivation(ctx context.Context, tx *gorm.DB, motivation entity.Motivation) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Create(&motivation).Error
}
func (ar *AdminRepository) CreatePsycholog(ctx context.Context, tx *gorm.DB, psycholog entity.Psycholog) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Create(&psycholog).Error
}
func (ar *AdminRepository) CreatePsychologLanguages(ctx context.Context, tx *gorm.DB, psychologLanguages []entity.PsychologLanguage) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Create(&psychologLanguages).Error
}
func (ar *AdminRepository) CreatePsychologSpecializations(ctx context.Context, tx *gorm.DB, psychologSpecializations []entity.PsychologSpecialization) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Create(&psychologSpecializations).Error
}
func (ar *AdminRepository) CreateEducations(ctx context.Context, tx *gorm.DB, educations []entity.Education) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Create(educations).Error
}

// Update
func (ar *AdminRepository) UpdateUser(ctx context.Context, tx *gorm.DB, user entity.User) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Where("id = ?", user.ID).Updates(&user).Error
}
func (ar *AdminRepository) UpdateNews(ctx context.Context, tx *gorm.DB, news entity.News) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Where("id = ?", news.ID).Updates(&news).Error
}
func (ar *AdminRepository) UpdateMotivationCategory(ctx context.Context, tx *gorm.DB, motivationCategory entity.MotivationCategory) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Where("id = ?", motivationCategory.ID).Updates(&motivationCategory).Error
}
func (ar *AdminRepository) UpdateMotivation(ctx context.Context, tx *gorm.DB, motivation entity.Motivation) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Select("author", "content", "motivation_category_id").Where("id = ?", motivation.ID).Updates(&motivation).Error
}
func (ar *AdminRepository) UpdatePsycholog(ctx context.Context, tx *gorm.DB, psycholog entity.Psycholog) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Where("id = ?", psycholog.ID).Updates(&psycholog).Error
}

// Delete
func (ar *AdminRepository) DeleteUserByID(ctx context.Context, tx *gorm.DB, userID string) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Where("id = ?", userID).Delete(&entity.User{}).Error
}
func (ar *AdminRepository) DeleteNewsByID(ctx context.Context, tx *gorm.DB, newsID string) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Where("id = ?", newsID).Delete(&entity.News{}).Error
}
func (ar *AdminRepository) DeleteMotivationCategoryByID(ctx context.Context, tx *gorm.DB, motivationCategoryID string) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Where("id = ?", motivationCategoryID).Delete(&entity.MotivationCategory{}).Error
}
func (ar *AdminRepository) DeleteMotivationByID(ctx context.Context, tx *gorm.DB, motivationID string) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Where("id = ?", motivationID).Delete(&entity.Motivation{}).Error
}
func (ar *AdminRepository) DeletePsychologByID(ctx context.Context, tx *gorm.DB, psychologID string) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Where("id = ?", psychologID).Delete(&entity.Psycholog{}).Error
}
func (ar *AdminRepository) DeletePsychologLanguageByPsychologID(ctx context.Context, tx *gorm.DB, psychologID string) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Where("psycholog_id = ?", psychologID).Delete(&entity.PsychologLanguage{}).Error
}
func (ar *AdminRepository) DeletePsychologSpecializationByPsychologID(ctx context.Context, tx *gorm.DB, psychologID string) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Where("psycholog_id = ?", psychologID).Delete(&entity.PsychologSpecialization{}).Error
}
func (ar *AdminRepository) DeleteEducationByPsychologID(ctx context.Context, tx *gorm.DB, psychologID string) error {
	if tx == nil {
		tx = ar.db
	}

	return tx.WithContext(ctx).Where("psycholog_id = ?", psychologID).Delete(&entity.Education{}).Error
}
