package entity

import (
	"github.com/google/uuid"
)

type News struct {
	ID    uuid.UUID `gorm:"type:uuid;primaryKey" json:"news_id"`
	Image string    `json:"news_image"`
	Title string    `json:"news_title"`
	Body  string    `json:"news_body"`
	Date  string    `gorm:"type:date" json:"news_date"`

	NewsDetails []NewsDetail `gorm:"foreignKey:NewsID"`

	TimeStamp
}
