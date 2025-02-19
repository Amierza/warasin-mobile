package entity

import (
	"github.com/google/uuid"
)

type City struct {
	ID         uuid.UUID   `gorm:"type:uuid;primaryKey" json:"city_id"`
	Name       string      `json:"city_name"`
	Users      []User      `gorm:"foreignKey:CityID"`
	Psychologs []Psycholog `gorm:"foreignKey:CityID"`

	TimeStamp
}
