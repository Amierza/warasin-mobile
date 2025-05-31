package repository

import (
	"context"

	"github.com/Amierza/warasin-mobile/backend/dto"
	"github.com/Amierza/warasin-mobile/backend/entity"
	"gorm.io/gorm"
)

type (
	IMasterRepository interface {
		// Get / Read
		// Role
		GetRoleByID(ctx context.Context, tx *gorm.DB, roleID string) (entity.Role, error)
		// City & Province
		GetCityByID(ctx context.Context, tx *gorm.DB, cityID string) (entity.City, error)
		GetAllProvince(ctx context.Context, tx *gorm.DB) (dto.AllProvinceRepositoryResponse, error)
		GetAllCity(ctx context.Context, tx *gorm.DB, req dto.CityQueryRequest) (dto.AllCityRepositoryResponse, error)
		// Psycholog
		GetPsychologByID(ctx context.Context, tx *gorm.DB, psychologID string) (entity.Psycholog, bool, error)
		GetPsychologByEmail(ctx context.Context, tx *gorm.DB, email string) (entity.Psycholog, bool, error)
	}

	MasterRepository struct {
		db *gorm.DB
	}
)

func NewMasterRepository(db *gorm.DB) *MasterRepository {
	return &MasterRepository{
		db: db,
	}
}

// Get / read
// Role
func (mr *MasterRepository) GetRoleByID(ctx context.Context, tx *gorm.DB, roleID string) (entity.Role, error) {
	if tx == nil {
		tx = mr.db
	}

	var role entity.Role
	if err := tx.WithContext(ctx).Where("id = ?", roleID).Take(&role).Error; err != nil {
		return entity.Role{}, err
	}

	return role, nil
}

// City & Province
func (mr *MasterRepository) GetCityByID(ctx context.Context, tx *gorm.DB, cityID string) (entity.City, error) {
	if tx == nil {
		tx = mr.db
	}

	var city entity.City
	if err := tx.WithContext(ctx).Preload("Province").Where("id = ?", cityID).Take(&city).Error; err != nil {
		return entity.City{}, err
	}

	return city, nil
}
func (mr *MasterRepository) GetAllProvince(ctx context.Context, tx *gorm.DB) (dto.AllProvinceRepositoryResponse, error) {
	if tx == nil {
		tx = mr.db
	}

	var provinces []entity.Province
	var err error

	if err := tx.WithContext(ctx).Model(&entity.Province{}).Find(&provinces).Error; err != nil {
		return dto.AllProvinceRepositoryResponse{}, err
	}

	return dto.AllProvinceRepositoryResponse{
		Provinces: provinces,
	}, err
}
func (mr *MasterRepository) GetAllCity(ctx context.Context, tx *gorm.DB, req dto.CityQueryRequest) (dto.AllCityRepositoryResponse, error) {
	if tx == nil {
		tx = mr.db
	}

	var cities []entity.City
	var err error

	query := tx.WithContext(ctx).Model(&entity.City{})

	if req.ProvinceID != "" {
		query = query.Where("province_id = ?", req.ProvinceID)
	}

	if err := query.Find(&cities).Error; err != nil {
		return dto.AllCityRepositoryResponse{}, err
	}

	return dto.AllCityRepositoryResponse{
		Cities: cities,
	}, err
}

// Psycholog
func (mr *MasterRepository) GetPsychologByID(ctx context.Context, tx *gorm.DB, psychologID string) (entity.Psycholog, bool, error) {
	if tx == nil {
		tx = mr.db
	}

	var psycholog entity.Psycholog
	query := tx.WithContext(ctx).
		Preload("Role").
		Preload("City.Province").
		Preload("PsychologLanguages.LanguageMaster").
		Preload("PsychologSpecializations.Specialization").
		Preload("Educations")

	if err := query.Where("id = ?", psychologID).Take(&psycholog).Error; err != nil {
		return entity.Psycholog{}, false, err
	}

	return psycholog, true, nil
}
func (mr *MasterRepository) GetPsychologByEmail(ctx context.Context, tx *gorm.DB, email string) (entity.Psycholog, bool, error) {
	if tx == nil {
		tx = mr.db
	}

	var psycholog entity.Psycholog
	query := tx.WithContext(ctx).
		Preload("Role").
		Preload("City.Province").
		Preload("PsychologLanguages.LanguageMaster").
		Preload("PsychologSpecializations.Specialization").
		Preload("Educations")

	if err := query.Where("email = ?", email).Take(&psycholog).Error; err != nil {
		return entity.Psycholog{}, false, err
	}

	return psycholog, true, nil
}
