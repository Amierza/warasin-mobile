package entity

import (
	"github.com/google/uuid"
)

type PsychologSpecialization struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"psy_spe_id"`

	PsychologID      *uuid.UUID     `gorm:"type:uuid" json:"psy_id"`
	Psycholog        Psycholog      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SpecializationID *uuid.UUID     `gorm:"type:uuid" json:"spe_id"`
	Specialization   Specialization `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	TimeStamp
}
