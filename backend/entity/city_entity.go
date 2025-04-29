package entity

import (
	"github.com/google/uuid"
)

type City struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"city_id"`
	Name string    `json:"city_name"`
	Type string    `json:"city_type"`

	ProvinceID *uuid.UUID `gorm:"type:uuid" json:"province_id"`
	Province   Province   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Users      []User      `gorm:"foreignKey:CityID"`
	Psychologs []Psycholog `gorm:"foreignKey:CityID"`

	TimeStamp
}
