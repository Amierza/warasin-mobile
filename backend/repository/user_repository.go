package repository

import (
	"context"

	"github.com/Amierza/warasin-mobile/backend/entity"
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
	if err := tx.WithContext(ctx).Where("email = ?", email).Take(&user).Error; err != nil {
		return entity.User{}, false, err
	}

	return user, true, nil
}

func (ur *UserRepository) RegisterUser(ctx context.Context, tx *gorm.DB, user entity.User) (entity.User, error) {
	if tx == nil {
		tx = ur.db
	}

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
	if err := tx.WithContext(ctx).Where("id = ?", userID).Take(&user).Error; err != nil {
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
