package entity

import (
	"github.com/google/uuid"
)

type Specialization struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"spe_id"`
	Name        string    `json:"spe_name"`
	Description string    `json:"spe_desc"`

	SpecializationDetails    []SpecializationDetail    `gorm:"foreignKey:SpecializationID"`
	PsychologSpecializations []PsychologSpecialization `gorm:"foreignKey:SpecializationID"`

	TimeStamp
}
