package migrations

import (
	"github.com/Amierza/warasin-mobile/backend/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&entity.City{},
		&entity.News{},
		&entity.Psycholog{},
		&entity.User{},
		&entity.NewsDetail{},
		&entity.Consulation{},
	); err != nil {
		return err
	}

	return nil
}
