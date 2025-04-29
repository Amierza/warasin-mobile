package entity

import (
	"github.com/google/uuid"
)

type SpecializationDetail struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"spe_detail_id"`
	Name string    `json:"spe_detail_name"`

	SpecializationID *uuid.UUID     `gorm:"type:uuid" json:"spe_id"`
	Specialization   Specialization `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	TimeStamp
}
