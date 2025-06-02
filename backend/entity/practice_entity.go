package entity

import (
	"github.com/google/uuid"
)

type Practice struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"prac_id"`
	Type        string    `json:"prac_type"` // Online / Offline
	Name        string    `json:"prac_name"`
	Address     string    `json:"prac_address"`
	PhoneNumber string    `json:"prac_phone_number"`

	PsychologID *uuid.UUID `gorm:"type:uuid" json:"psy_id"`
	Psycholog   Psycholog  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	PracticeSchedules []PracticeSchedule `gorm:"foreignKey:PracticeID"`
	Consuls           []Consultation     `gorm:"foreignKey:PracticeID"`

	TimeStamp
}
