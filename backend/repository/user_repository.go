package repository

import (
	"context"

	"github.com/Amierza/warasin-mobile/backend/dto"
	"github.com/Amierza/warasin-mobile/backend/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	IUserRepository interface {
		CheckEmail(ctx context.Context, tx *gorm.DB, email string) (entity.User, bool, error)
		RegisterUser(ctx context.Context, tx *gorm.DB, user entity.User) (entity.User, error)
		UpdateUser(ctx context.Context, tx *gorm.DB, user entity.User) (entity.User, error)
		GetUserByPassword(ctx context.Context, tx *gorm.DB, password string) (entity.User, error)
		GetUserByID(ctx context.Context, tx *gorm.DB, userID string) (entity.User, error)
		GetRoleByName(ctx context.Context, tx *gorm.DB, roleName string) (entity.Role, error)
		GetPermissionsByRoleID(ctx context.Context, tx *gorm.DB, roleID string) ([]string, error)
		GetRoleByID(ctx context.Context, tx *gorm.DB, roleID string) (entity.Role, error)
		GetCityByID(ctx context.Context, tx *gorm.DB, cityID string) (entity.City, error)
		GetAllProvince(ctx context.Context, tx *gorm.DB) (dto.AllProvinceRepositoryResponse, error)
		GetAllCity(ctx context.Context, tx *gorm.DB, req dto.CityQueryRequest) (dto.AllCityRepositoryResponse, error)
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

func (ur *UserRepository) CheckEmail(ctx context.Context, tx *gorm.DB, email string) (entity.User, bool, error) {
	if tx == nil {
		tx = ur.db
	}

	var user entity.User
	if err := tx.WithContext(ctx).Preload("Role").Preload("City").Where("email = ?", email).Take(&user).Error; err != nil {
		return entity.User{}, false, err
	}

	return user, true, nil
}

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

func (ur *UserRepository) UpdateUser(ctx context.Context, tx *gorm.DB, user entity.User) (entity.User, error) {
	if tx == nil {
		tx = ur.db
	}

	if err := tx.WithContext(ctx).Updates(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (ur *UserRepository) GetUserByPassword(ctx context.Context, tx *gorm.DB, password string) (entity.User, error) {
	if tx == nil {
		tx = ur.db
	}

	var user entity.User
	if err := tx.WithContext(ctx).Where("password = ?", password).Take(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (ur *UserRepository) GetUserByID(ctx context.Context, tx *gorm.DB, userID string) (entity.User, error) {
	if tx == nil {
		tx = ur.db
	}

	var user entity.User
	if err := tx.WithContext(ctx).Preload("City.Province").Preload("Role").Where("id = ?", userID).Take(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (ur *UserRepository) GetRoleByName(ctx context.Context, tx *gorm.DB, roleName string) (entity.Role, error) {
	if tx == nil {
		tx = ur.db
	}

	var role entity.Role
	if err := tx.WithContext(ctx).Where("name = ?", roleName).Take(&role).Error; err != nil {
		return entity.Role{}, err
	}

	return role, nil
}

func (ur *UserRepository) GetRoleByID(ctx context.Context, tx *gorm.DB, roleID string) (entity.Role, error) {
	if tx == nil {
		tx = ur.db
	}

	var role entity.Role
	if err := tx.WithContext(ctx).Where("id = ?", roleID).Take(&role).Error; err != nil {
		return entity.Role{}, err
	}

	return role, nil
}

func (ur *UserRepository) GetPermissionsByRoleID(ctx context.Context, tx *gorm.DB, roleID string) ([]string, error) {
	if tx == nil {
		tx = ur.db
	}

	var endpoints []string
	if err := tx.WithContext(ctx).Table("permissions").Where("role_id = ?", roleID).Pluck("endpoint", &endpoints).Error; err != nil {
		return []string{}, err
	}

	return endpoints, nil
}

func (ar *UserRepository) GetCityByID(ctx context.Context, tx *gorm.DB, cityID string) (entity.City, error) {
	if tx == nil {
		tx = ar.db
	}

	var city entity.City
	if err := tx.WithContext(ctx).Preload("Province").Where("id = ?", cityID).Take(&city).Error; err != nil {
		return entity.City{}, err
	}

	return city, nil
}

func (ar *UserRepository) GetAllProvince(ctx context.Context, tx *gorm.DB) (dto.AllProvinceRepositoryResponse, error) {
	if tx == nil {
		tx = ar.db
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

func (ar *UserRepository) GetAllCity(ctx context.Context, tx *gorm.DB, req dto.CityQueryRequest) (dto.AllCityRepositoryResponse, error) {
	if tx == nil {
		tx = ar.db
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
