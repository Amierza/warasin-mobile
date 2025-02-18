package entity

import (
	"github.com/Amierza/warasin-mobile/backend/helpers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Psycholog struct {
	ID          uuid.UUID     `gorm:"type:uuid;primaryKey" json:"psy_id"`
	CityID      uuid.UUID     `gorm:"type:uuid" json:"city_id"`
	City        City          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Consuls     []Consulation `gorm:"foreignKey:PsychologID"`
	Name        string        `json:"psy_name"`
	Email       string        `gorm:"unique; not null" json:"psy_email"`
	Password    string        `json:"psy_password"`
	WorkYear    string        `gorm:"type:varchar(4)" json:"psy_work_year"`
	Description string        `json:"psy_description"`
	PhoneNumber string        `json:"psy_phone_number,omitempty"`
	Image       string        `json:"psy_image,omitempty"`

	TimeStamp
}

func (p *Psycholog) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	p.ID = uuid.New()

	var err error
	p.Password, err = helpers.HashPassword(p.Password)
	if err != nil {
		return err
	}

	return nil
}
