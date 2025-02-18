package migrations

import (
	"github.com/Amierza/warasin-mobile/backend/migrations/seed"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	if err := seed.ListUserSeeder(db); err != nil {
		return err
	}

	return nil
}
