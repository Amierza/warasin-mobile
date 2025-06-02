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
		GetPracticeByID(ctx context.Context, tx *gorm.DB, pracID string) (entity.Practice, bool, error)
		GetAvailableSlotByID(ctx context.Context, tx *gorm.DB, slotID string) (entity.AvailableSlot, bool, error)
		GetAllConsultationWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest, userID string) (dto.AllConsultationRepositoryResponseForUser, error)

		// Create
		RegisterUser(ctx context.Context, tx *gorm.DB, user entity.User) (entity.User, error)
		CreateConsultation(ctx context.Context, tx *gorm.DB, consultation entity.Consultation) error

		// Update
		UpdateUser(ctx context.Context, tx *gorm.DB, user entity.User) (entity.User, error)
		UpdateStatusBookSlot(ctx context.Context, tx *gorm.DB, slotID uuid.UUID, statusBook bool) error
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
