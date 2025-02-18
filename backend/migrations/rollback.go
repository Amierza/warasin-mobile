package migrations

import (
	"github.com/Amierza/warasin-mobile/backend/entity"
	"gorm.io/gorm"
)

func Rollback(db *gorm.DB) error {
	tables := []interface{}{
		&entity.Consulation{},
		&entity.NewsDetail{},
		&entity.User{},
		&entity.Psycholog{},
		&entity.News{},
		&entity.City{},
	}

	for _, table := range tables {
		if err := db.Migrator().DropTable(table); err != nil {
			return err
		}
	}

	return nil
}
