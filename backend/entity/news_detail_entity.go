package entity

import (
	"github.com/google/uuid"
)

type NewsDetail struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"news_detail_id"`
	Date string    `gorm:"type:date" json:"news_detail_date"`

	UserID *uuid.UUID `gorm:"type:uuid" json:"user_id"`
	User   User       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	NewsID *uuid.UUID `gorm:"type:uuid" json:"news_id"`
	News   News       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	TimeStamp
}
