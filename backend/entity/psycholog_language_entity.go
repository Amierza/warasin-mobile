package entity

import (
	"github.com/google/uuid"
)

type PsychologLanguage struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"psy_lang_id"`

	PsychologID      *uuid.UUID     `gorm:"uuid" json:"psy_id"`
	Psycholog        Psycholog      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	LanguageMasterID *uuid.UUID     `gorm:"type:uuid" json:"lang_id"`
	LanguageMaster   LanguageMaster `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	TimeStamp
}
