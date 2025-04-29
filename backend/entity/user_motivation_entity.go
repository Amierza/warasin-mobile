package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserMotivation struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"user_mot_id"`
	DisplayDate time.Time `json:"user_mot_display_date"`
	Reaction    int       `json:"user_mot_reaction"`

	UserID       *uuid.UUID `gorm:"type:uuid" json:"user_id"`
	User         User       `gorm:"constraint:onUpdate:CASCADE,OnDelete:SET NULL;"`
	MotivationID *uuid.UUID `gorm:"type:uuid" json:"mot_id"`
	Motivation   Motivation `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	TimeStamp
}
