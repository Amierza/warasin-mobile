package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NewsDetail struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey" json:"news_detail_id"`
	UserID uuid.UUID `gorm:"type:uuid" json:"user_id"`
	User   User      `gorm:"foreignKey:UserID"`
	NewsID uuid.UUID `gorm:"type:uuid" json:"news_id"`
	News   User      `gorm:"foreignKey:NewsID"`
	Date   string    `gorm:"type:date" json:"news_detail_date"`

	TimeStamp
}

func (n *NewsDetail) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	n.ID = uuid.New()

	return nil
}
