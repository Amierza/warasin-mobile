package repository

import (
	"context"

	"github.com/Amierza/warasin-mobile/backend/entity"
	"gorm.io/gorm"
)

type (
	IPsychologRepository interface {
		// Get
		GetPermissionsByRoleID(ctx context.Context, tx *gorm.DB, roleID string) ([]string, error)
		GetRoleByID(ctx context.Context, tx *gorm.DB, roleID string) (entity.Role, error)
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
