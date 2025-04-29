package entity

import (
	"github.com/google/uuid"
)

type Education struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey" json:"edu_id"`
	Degree         string    `json:"edu_degree"`
	Major          string    `json:"edu_major"`
	Institution    string    `json:"edu_institution"`
	GraduationYear string    `gorm:"type:varchar(4)" json:"edu_graduation_year"`

	PsychologID *uuid.UUID `gorm:"type:uuid" json:"psy_id"`
	Psycholog   Psycholog  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	TimeStamp
}
