package entity

import (
	"time"

	"github.com/google/uuid"
)

type Consulation struct {
	ID      uuid.UUID  `gorm:"type:uuid;primaryKey" json:"consul_id"`
	Date    *time.Time `gorm:"type:date" json:"consul_date"`
	Rate    int        `json:"consul_rate"`
	Comment string     `json:"consul_comment"`

	UserID      *uuid.UUID `gorm:"type:uuid" json:"user_id"`
	User        User       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PsychologID *uuid.UUID `gorm:"type:uuid" json:"psy_id"`
	Psycholog   Psycholog  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	TimeStamp
}
