package entity

import (
	"github.com/google/uuid"
)

type Motivation struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey" json:"mot_id"`
	Content string    `json:"mot_content"`
	Author  string    `json:"mot_author"`

	MotivationCategoryID *uuid.UUID         `gorm:"type:uuid" json:"mot_cat_id"`
	MotivationCategory   MotivationCategory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	UserMotivations []UserMotivation `gorm:"foreignKey:MotivationID"`

	TimeStamp
}
