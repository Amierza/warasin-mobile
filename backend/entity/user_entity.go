package entity

import (
	"time"

	"github.com/Amierza/warasin-mobile/backend/helpers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey" json:"user_id"`
	Name        string     `json:"user_name"`
	Email       string     `gorm:"unique; not null" json:"user_email"`
	Password    string     `json:"user_password"`
	Gender      bool       `json:"user_gender,omitempty"`
	Birthdate   *time.Time `gorm:"type:date" json:"user_birth_date,omitempty"`
	PhoneNumber string     `json:"user_phone_number,omitempty"`
	Data01      int        `json:"user_data01,omitempty"`
	Data02      int        `json:"user_data02,omitempty"`
	Data03      int        `json:"user_data03,omitempty"`
	IsVerified  bool       `json:"user_is_verified"`

	CityID *uuid.UUID `gorm:"type:uuid" json:"city_id"`
	City   City       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	RoleID *uuid.UUID `gorm:"type:uuid" json:"role_id"`
	Role   Role       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Consuls     []Consulation `gorm:"foreignKey:UserID"`
	NewsDetails []NewsDetail  `gorm:"foreignKey:UserID"`

	TimeStamp
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	var err error
	u.Password, err = helpers.HashPassword(u.Password)
	if err != nil {
		return err
	}

	return nil
}
