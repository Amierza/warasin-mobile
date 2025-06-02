package entity

import (
	"github.com/google/uuid"
)

type AvailableSlot struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"slot_id"`
	Start    string    `json:"slot_start"`
	End      string    `json:"slot_end"`
	IsBooked bool      `json:"slot_is_booked"`

	PsychologID *uuid.UUID `gorm:"type:uuid" json:"psy_id"`
	Psycholog   Psycholog  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Consuls []Consultation `gorm:"foreignKey:AvailableSlotID"`

	TimeStamp
}
