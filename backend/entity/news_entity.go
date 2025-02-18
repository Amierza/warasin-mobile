package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type News struct {
	ID          uuid.UUID    `gorm:"type:uuid;primaryKey" json:"news_id"`
	NewsDetails []NewsDetail `gorm:"foreignKey:NewsID"`
	Image       string       `json:"news_image"`
	Title       string       `json:"news_title"`
	Body        string       `json:"news_body"`
	Date        string       `gorm:"type:date" json:"news_date"`

	TimeStamp
}

func (n *News) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	n.ID = uuid.New()

	return nil
}
