package entity

import (
	"github.com/Amierza/warasin-mobile/backend/helpers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Psycholog struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"psy_id"`
	Name        string    `json:"psy_name"`
	STRNumber   string    `json:"psy_str_number"`
	Email       string    `gorm:"unique; not null" json:"psy_email"`
	Password    string    `json:"psy_password"`
	WorkYear    string    `gorm:"type:varchar(4)" json:"psy_work_year"`
	Description string    `json:"psy_description"`
	PhoneNumber string    `json:"psy_phone_number,omitempty"`
	Image       string    `json:"psy_image,omitempty"`

	CityID *uuid.UUID `gorm:"type:uuid" json:"city_id"`
	City   City       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	RoleID *uuid.UUID `gorm:"type:uuid" json:"role_id"`
	Role   Role       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Consuls                  []Consulation             `gorm:"foreignKey:PsychologID"`
	PsychologLanguages       []PsychologLanguage       `gorm:"foreignKey:PsychologID"`
	PsychologSpecializations []PsychologSpecialization `gorm:"foreignKey:PsychologID"`
	Educations               []Education               `gorm:"foreignKey:PsychologID"`
	Practices                []Practice                `gorm:"foreignKey:PsychologID"`

	TimeStamp
}

func (p *Psycholog) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	var err error
	p.Password, err = helpers.HashPassword(p.Password)
	if err != nil {
		return err
	}

	return nil
}
