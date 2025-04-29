package entity

import (
	"github.com/google/uuid"
)

type MotivationCategory struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"mot_cat_id"`
	Name string    `json:"mot_cat_name"`

	Motivations []Motivation `gorm:"foreignKey:MotivationCategoryID"`

	TimeStamp
}
