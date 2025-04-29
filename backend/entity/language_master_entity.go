package entity

import (
	"github.com/google/uuid"
)

type LanguageMaster struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"lang_id"`
	Name string    `json:"lang_name"`

	PsychologLanguages []PsychologLanguage `gorm:"foreignKey:LanguageMasterID"`

	TimeStamp
}
