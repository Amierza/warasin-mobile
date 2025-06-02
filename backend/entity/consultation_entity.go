package entity

import (
	"github.com/google/uuid"
)

type Consultation struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey" json:"consul_id"`
	Date    string    `json:"consul_date"`
	Rate    int       `json:"consul_rate"`
	Comment string    `json:"consul_comment"`
	Status  int       `json:"consul_status"` // 0: upcoming 1: canceled 2: done

	UserID          *uuid.UUID    `gorm:"type:uuid" json:"user_id"`
	User            User          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PracticeID      *uuid.UUID    `gorm:"type:uuid" json:"prac_id"`
	Practice        Practice      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	AvailableSlotID *uuid.UUID    `gorm:"type:uuid" json:"slot_id"`
	AvailableSlot   AvailableSlot `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	TimeStamp
}
