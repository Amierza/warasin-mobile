package repository

import (
	"context"
	"math"

	"github.com/Amierza/warasin-mobile/backend/dto"
	"github.com/Amierza/warasin-mobile/backend/entity"
	"gorm.io/gorm"
)

type (
	IPsychologRepository interface {
		// Get
		GetPermissionsByRoleID(ctx context.Context, tx *gorm.DB, roleID string) ([]string, error)
		GetRoleByID(ctx context.Context, tx *gorm.DB, roleID string) (entity.Role, error)
		GetAllPractice(ctx context.Context, tx *gorm.DB, psyID string) (dto.AllPracticeRepositoryResponse, error)
		GetAllAvailableSlot(ctx context.Context, tx *gorm.DB, psyID string) (dto.AllAvailableSlotRepositoryResponse, error)
		GetAllConsultationWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest, psychologID string) (dto.AllConsultationRepositoryResponse, error)
	}

	PsychologRepository struct {
		db *gorm.DB
	}
)

func NewPsychologRepository(db *gorm.DB) *PsychologRepository {
	return &PsychologRepository{
		db: db,
	}
}

// Get
func (pr *PsychologRepository) GetPermissionsByRoleID(ctx context.Context, tx *gorm.DB, roleID string) ([]string, error) {
	if tx == nil {
		tx = pr.db
	}

	var endpoints []string
	if err := tx.WithContext(ctx).Table("permissions").Where("role_id = ?", roleID).Pluck("endpoint", &endpoints).Error; err != nil {
		return []string{}, err
	}

	return endpoints, nil
}
func (pr *PsychologRepository) GetRoleByID(ctx context.Context, tx *gorm.DB, roleID string) (entity.Role, error) {
	if tx == nil {
		tx = pr.db
	}

	var role entity.Role
	if err := tx.WithContext(ctx).Where("id = ?", roleID).Take(&role).Error; err != nil {
		return entity.Role{}, err
	}

	return role, nil
}
func (pr *PsychologRepository) GetAllPractice(ctx context.Context, tx *gorm.DB, psyID string) (dto.AllPracticeRepositoryResponse, error) {
	if tx == nil {
		tx = pr.db
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
func (pr *PsychologRepository) GetAllAvailableSlot(ctx context.Context, tx *gorm.DB, psyID string) (dto.AllAvailableSlotRepositoryResponse, error) {
	if tx == nil {
		tx = pr.db
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
func (pr *PsychologRepository) GetAllConsultationWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest, psychologID string) (dto.AllConsultationRepositoryResponse, error) {
	if tx == nil {
		tx = pr.db
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

	query := tx.WithContext(ctx).Model(&entity.Consultation{}).Where("available_slots.psycholog_id = ?", &psychologID).
		Preload("User.Role").
		Preload("User.City.Province").
		Preload("AvailableSlot.Psycholog.Role").
		Preload("AvailableSlot.Psycholog.City.Province").
		Preload("Practice.PracticeSchedules").
		Joins("JOIN available_slots ON consultations.available_slot_id = available_slots.id")

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
