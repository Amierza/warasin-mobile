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
	IUserRepository interface {
		// Get
		GetUserByEmail(ctx context.Context, tx *gorm.DB, email string) (entity.User, bool, error)
		GetUserByPassword(ctx context.Context, tx *gorm.DB, password string) (entity.User, bool, error)
		GetUserByID(ctx context.Context, tx *gorm.DB, userID string) (entity.User, bool, error)
		GetRoleByName(ctx context.Context, tx *gorm.DB, roleName string) (entity.Role, bool, error)
		GetPermissionsByRoleID(ctx context.Context, tx *gorm.DB, roleID string) ([]string, bool, error)
		GetRoleByID(ctx context.Context, tx *gorm.DB, roleID string) (entity.Role, bool, error)
		GetAllNewsWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.AllNewsRepositoryResponse, error)
		GetNewsByID(ctx context.Context, tx *gorm.DB, newsID string) (entity.News, bool, error)
		GetAllMotivationWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.AllMotivationRepositoryResponse, error)
		GetMotivationByID(ctx context.Context, tx *gorm.DB, motivationID string) (entity.Motivation, bool, error)
		GetPracticeByID(ctx context.Context, tx *gorm.DB, pracID string) (entity.Practice, bool, error)
		GetAvailableSlotByID(ctx context.Context, tx *gorm.DB, slotID string) (entity.AvailableSlot, bool, error)
		GetAllConsultationWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest, userID string) (dto.AllConsultationRepositoryResponseForUser, error)
		GetConsultationByID(ctx context.Context, tx *gorm.DB, consulID string) (entity.Consultation, bool, error)
		GetAllPsycholog(ctx context.Context, tx *gorm.DB, filter dto.PsychologFilter) ([]entity.Psycholog, error)
		GetPsychologByID(ctx context.Context, tx *gorm.DB, psyID string) (entity.Psycholog, bool, error)
		GetAllPractice(ctx context.Context, tx *gorm.DB, psyID string) (dto.AllPracticeRepositoryResponse, error)
		GetAllAvailableSlot(ctx context.Context, tx *gorm.DB, psyID string) (dto.AllAvailableSlotRepositoryResponse, error)
		GetNewsDetailByUserAndNewsID(ctx context.Context, tx *gorm.DB, userID string, newsID string) (entity.NewsDetail, bool, error)
		GetAllNewsDetail(ctx context.Context, tx *gorm.DB, userID string) ([]entity.NewsDetail, error)

		// Create
		RegisterUser(ctx context.Context, tx *gorm.DB, user entity.User) (entity.User, error)
		CreateConsultation(ctx context.Context, tx *gorm.DB, consultation entity.Consultation) error
		CreateNewsDetail(ctx context.Context, tx *gorm.DB, newsDetail entity.NewsDetail) error

		// Update
		UpdateUser(ctx context.Context, tx *gorm.DB, user entity.User) (entity.User, error)
		UpdateStatusBookSlot(ctx context.Context, tx *gorm.DB, slotID uuid.UUID, statusBook bool) error
		UpdateConsultation(ctx context.Context, tx *gorm.DB, consultation entity.Consultation) error

		// Delete
		DeleteConsultation(ctx context.Context, tx *gorm.DB, consulID string) error
	}

	UserRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Get
func (ur *UserRepository) GetUserByEmail(ctx context.Context, tx *gorm.DB, email string) (entity.User, bool, error) {
	if tx == nil {
		tx = ur.db
	}

	var user entity.User
	if err := tx.WithContext(ctx).Preload("Role").Preload("City").Where("email = ?", email).Take(&user).Error; err != nil {
		return entity.User{}, false, err
	}

	return user, true, nil
}
func (ur *UserRepository) GetUserByPassword(ctx context.Context, tx *gorm.DB, password string) (entity.User, bool, error) {
	if tx == nil {
		tx = ur.db
	}

	var user entity.User
	if err := tx.WithContext(ctx).Where("password = ?", password).Take(&user).Error; err != nil {
		return entity.User{}, false, err
	}

	return user, true, nil
}
func (ur *UserRepository) GetUserByID(ctx context.Context, tx *gorm.DB, userID string) (entity.User, bool, error) {
	if tx == nil {
		tx = ur.db
	}

	var user entity.User
	if err := tx.WithContext(ctx).Preload("City.Province").Preload("Role").Where("id = ?", userID).Take(&user).Error; err != nil {
		return entity.User{}, false, err
	}

	return user, true, nil
}
func (ur *UserRepository) GetRoleByName(ctx context.Context, tx *gorm.DB, roleName string) (entity.Role, bool, error) {
	if tx == nil {
		tx = ur.db
	}

	var role entity.Role
	if err := tx.WithContext(ctx).Where("name = ?", roleName).Take(&role).Error; err != nil {
		return entity.Role{}, false, err
	}

	return role, true, nil
}
func (ur *UserRepository) GetRoleByID(ctx context.Context, tx *gorm.DB, roleID string) (entity.Role, bool, error) {
	if tx == nil {
		tx = ur.db
	}

	var role entity.Role
	if err := tx.WithContext(ctx).Where("id = ?", roleID).Take(&role).Error; err != nil {
		return entity.Role{}, false, err
	}

	return role, true, nil
}
func (ur *UserRepository) GetPermissionsByRoleID(ctx context.Context, tx *gorm.DB, roleID string) ([]string, bool, error) {
	if tx == nil {
		tx = ur.db
	}

	var endpoints []string
	if err := tx.WithContext(ctx).Table("permissions").Where("role_id = ?", roleID).Pluck("endpoint", &endpoints).Error; err != nil {
		return []string{}, false, err
	}

	return endpoints, true, nil
}
func (ur *UserRepository) GetAllNewsWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.AllNewsRepositoryResponse, error) {
	if tx == nil {
		tx = ur.db
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
		query = query.Where("LOWER(title) LIKE ? OR LOWER(body) LIKE ?", searchValue, searchValue)
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
func (ur *UserRepository) GetNewsByID(ctx context.Context, tx *gorm.DB, newsID string) (entity.News, bool, error) {
	if tx == nil {
		tx = ur.db
	}

	var news entity.News
	if err := tx.WithContext(ctx).Where("id = ?", newsID).Take(&news).Error; err != nil {
		return entity.News{}, true, err
	}

	return news, true, nil
}
func (ur *UserRepository) GetAllMotivationWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.AllMotivationRepositoryResponse, error) {
	if tx == nil {
		tx = ur.db
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
		query = query.Where("LOWER(author) LIKE ? OR LOWER(content) LIKE ?", searchValue, searchValue)
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
func (ur *UserRepository) GetMotivationByID(ctx context.Context, tx *gorm.DB, motivationID string) (entity.Motivation, bool, error) {
	if tx == nil {
		tx = ur.db
	}

	var motivation entity.Motivation
	if err := tx.WithContext(ctx).Preload("MotivationCategory").Where("id = ?", motivationID).Take(&motivation).Error; err != nil {
		return entity.Motivation{}, true, err
	}

	return motivation, true, nil
}
func (ur *UserRepository) GetPracticeByID(ctx context.Context, tx *gorm.DB, pracID string) (entity.Practice, bool, error) {
	if tx == nil {
		tx = ur.db
	}

	query := tx.WithContext(ctx).Model(&entity.Practice{}).
		Preload("Psycholog.Role").
		Preload("Psycholog.City.Province").
		Preload("PracticeSchedules")

	var practice entity.Practice
	if err := query.Where("id = ?", pracID).Take(&practice).Error; err != nil {
		return entity.Practice{}, false, err
	}

	return practice, true, nil
}
func (ur *UserRepository) GetAvailableSlotByID(ctx context.Context, tx *gorm.DB, slotID string) (entity.AvailableSlot, bool, error) {
	if tx == nil {
		tx = ur.db
	}

	query := tx.WithContext(ctx).Model(&entity.AvailableSlot{}).
		Preload("Psycholog.Role").
		Preload("Psycholog.City.Province")

	var slot entity.AvailableSlot
	if err := query.Where("id = ?", slotID).Take(&slot).Error; err != nil {
		return entity.AvailableSlot{}, false, err
	}

	return slot, true, nil
}
func (ur *UserRepository) GetAllConsultationWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest, userID string) (dto.AllConsultationRepositoryResponseForUser, error) {
	if tx == nil {
		tx = ur.db
	}

	var consultations []entity.Consultation
	var err error
	var count int64

	if req.PerPage == 0 {
		req.PerPage = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}

	query := tx.WithContext(ctx).Model(&entity.Consultation{}).Where("user_id = ?", &userID).
		Preload("User.Role").
		Preload("User.City.Province").
		Preload("AvailableSlot.Psycholog.Role").
		Preload("AvailableSlot.Psycholog.City.Province").
		Preload("AvailableSlot.Psycholog.PsychologLanguages.LanguageMaster").
		Preload("AvailableSlot.Psycholog.PsychologSpecializations.Specialization").
		Preload("AvailableSlot.Psycholog.Educations").
		Preload("Practice.PracticeSchedules")

	if err := query.Count(&count).Error; err != nil {
		return dto.AllConsultationRepositoryResponseForUser{}, err
	}

	if err := query.Order("created_at DESC").Scopes(Paginate(req.Page, req.PerPage)).Find(&consultations).Error; err != nil {
		return dto.AllConsultationRepositoryResponseForUser{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(req.PerPage)))

	return dto.AllConsultationRepositoryResponseForUser{
		Consultations: consultations,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			MaxPage: totalPage,
			Count:   count,
		},
	}, err
}
func (ur *UserRepository) GetConsultationByID(ctx context.Context, tx *gorm.DB, consulID string) (entity.Consultation, bool, error) {
	if tx == nil {
		tx = ur.db
	}

	query := tx.WithContext(ctx).Model(&entity.Consultation{}).Where("id = ?", consulID).
		Preload("User.Role").
		Preload("User.City.Province").
		Preload("AvailableSlot.Psycholog.Role").
		Preload("AvailableSlot.Psycholog.City.Province").
		Preload("AvailableSlot.Psycholog.PsychologLanguages.LanguageMaster").
		Preload("AvailableSlot.Psycholog.PsychologSpecializations.Specialization").
		Preload("AvailableSlot.Psycholog.Educations").
		Preload("Practice.PracticeSchedules")

	var consultation entity.Consultation
	if err := query.Order("created_at DESC").Take(&consultation).Error; err != nil {
		return entity.Consultation{}, false, err
	}

	return consultation, true, nil
}
func (ur *UserRepository) GetAllPsycholog(ctx context.Context, tx *gorm.DB, filter dto.PsychologFilter) ([]entity.Psycholog, error) {
	if tx == nil {
		tx = ur.db
	}

	var (
		psychologs []entity.Psycholog
		err        error
	)

	query := tx.WithContext(ctx).Model(&entity.Psycholog{}).
		Preload("Role").
		Preload("City.Province").
		Preload("PsychologLanguages.LanguageMaster").
		Preload("PsychologSpecializations.Specialization").
		Preload("Educations")

	if filter.Name != "" {
		query = query.Where("name ILIKE ?", "%"+filter.Name+"%")
	}

	if filter.City != "" {
		query = query.Joins("JOIN cities ON cities.id = psychologs.city_id").
			Where("cities.name ILIKE ?", "%"+filter.City+"%")
	}

	if filter.Province != "" {
		query = query.Joins("JOIN cities ON id = psychologs.city_id").
			Joins("JOIN provinces ON provinces.id = cities.province_id").
			Where("provinces.name ILIKE ?", "%"+filter.Province+"%")
	}

	if filter.Specialization != "" {
		query = query.Joins("JOIN psycholog_specializations ON psycholog_specializations.psycholog_id = psychologs.id").
			Joins("JOIN specializations ON specializations.id = psycholog_specializations.specialization_id").
			Where("specializations.name ILIKE ?", "%"+filter.Specialization+"%")
	}

	if err := query.Order("created_at DESC").Find(&psychologs).Error; err != nil {
		return []entity.Psycholog{}, err
	}

	return psychologs, err
}
func (ur *UserRepository) GetPsychologByID(ctx context.Context, tx *gorm.DB, psyID string) (entity.Psycholog, bool, error) {
	if tx == nil {
		tx = ur.db
	}

	query := tx.WithContext(ctx).Model(&entity.Psycholog{}).
		Preload("Role").
		Preload("City.Province").
		Preload("PsychologLanguages.LanguageMaster").
		Preload("PsychologSpecializations.Specialization").
		Preload("Educations")

	var psy entity.Psycholog
	if err := query.Where("id = ?", psyID).Take(&psy).Error; err != nil {
		return entity.Psycholog{}, false, err
	}

	return psy, true, nil
}
func (ur *UserRepository) GetAllPractice(ctx context.Context, tx *gorm.DB, psyID string) (dto.AllPracticeRepositoryResponse, error) {
	if tx == nil {
		tx = ur.db
	}

	var (
		practices []entity.Practice
		err       error
	)

	query := tx.WithContext(ctx).Model(&entity.Practice{}).Where("psycholog_id = ?", psyID).
		Preload("Psycholog.Role").
		Preload("Psycholog.City.Province").
		Preload("PracticeSchedules")

	if err := query.Order("created_at DESC").Find(&practices).Error; err != nil {
		return dto.AllPracticeRepositoryResponse{}, err
	}

	return dto.AllPracticeRepositoryResponse{
		Practices: practices,
	}, err
}
func (ur *UserRepository) GetAllAvailableSlot(ctx context.Context, tx *gorm.DB, psyID string) (dto.AllAvailableSlotRepositoryResponse, error) {
	if tx == nil {
		tx = ur.db
	}

	var (
		availableSlots []entity.AvailableSlot
		err            error
	)

	query := tx.WithContext(ctx).Model(&entity.AvailableSlot{}).Where("psycholog_id = ?", psyID).
		Preload("Psycholog.Role").
		Preload("Psycholog.City.Province")

	if err := query.Order("created_at DESC").Find(&availableSlots).Error; err != nil {
		return dto.AllAvailableSlotRepositoryResponse{}, err
	}

	return dto.AllAvailableSlotRepositoryResponse{
		AvailableSlots: availableSlots,
	}, err
}
func (ur *UserRepository) GetNewsDetailByUserAndNewsID(ctx context.Context, tx *gorm.DB, userID string, newsID string) (entity.NewsDetail, bool, error) {
	if tx == nil {
		tx = ur.db
	}

	query := tx.WithContext(ctx).Model(&entity.NewsDetail{}).
		Preload("User").
		Preload("User.City.Province").
		Preload("News")

	var newsDetail entity.NewsDetail
	if err := query.Where("user_id = ? AND news_id = ?", userID, newsID).Take(&newsDetail).Error; err != nil {
		return entity.NewsDetail{}, false, err
	}

	return newsDetail, true, nil
}
func (ur *UserRepository) GetAllNewsDetail(ctx context.Context, tx *gorm.DB, userID string) ([]entity.NewsDetail, error) {
	if tx == nil {
		tx = ur.db
	}

	query := tx.WithContext(ctx).Model(&entity.NewsDetail{}).
		Preload("User").
		Preload("User.City.Province").
		Preload("News")

	var newsDetails []entity.NewsDetail
	if err := query.Where("user_id = ?", userID).Find(&newsDetails).Error; err != nil {
		return []entity.NewsDetail{}, err
	}

	return newsDetails, nil
}

// Create
func (ur *UserRepository) RegisterUser(ctx context.Context, tx *gorm.DB, user entity.User) (entity.User, error) {
	if tx == nil {
		tx = ur.db
	}

	user.ID = uuid.New()
	if err := tx.WithContext(ctx).Create(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}
func (ur *UserRepository) CreateConsultation(ctx context.Context, tx *gorm.DB, consultation entity.Consultation) error {
	if tx == nil {
		tx = ur.db
	}

	return tx.WithContext(ctx).Create(&consultation).Error
}
func (ur *UserRepository) CreateNewsDetail(ctx context.Context, tx *gorm.DB, newsDetail entity.NewsDetail) error {
	if tx == nil {
		tx = ur.db
	}

	return tx.WithContext(ctx).Create(&newsDetail).Error
}

// Update
func (ur *UserRepository) UpdateUser(ctx context.Context, tx *gorm.DB, user entity.User) (entity.User, error) {
	if tx == nil {
		tx = ur.db
	}

	if err := tx.WithContext(ctx).Updates(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}
func (ur *UserRepository) UpdateStatusBookSlot(ctx context.Context, tx *gorm.DB, slotID uuid.UUID, statusBook bool) error {
	if tx == nil {
		tx = ur.db
	}

	return tx.WithContext(ctx).
		Model(&entity.AvailableSlot{}).
		Where("id = ?", slotID).
		Update("is_booked", statusBook).Error
}
func (ur *UserRepository) UpdateConsultation(ctx context.Context, tx *gorm.DB, consultation entity.Consultation) error {
	if tx == nil {
		tx = ur.db
	}

	return tx.WithContext(ctx).Where("id = ?", consultation.ID).Updates(&consultation).Error
}

// Delete
func (ur *UserRepository) DeleteConsultation(ctx context.Context, tx *gorm.DB, consulID string) error {
	if tx == nil {
		tx = ur.db
	}

	return tx.WithContext(ctx).Where("id = ?", consulID).Delete(&entity.Consultation{}).Error
}
