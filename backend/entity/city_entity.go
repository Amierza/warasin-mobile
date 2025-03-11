package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type City struct {
	ID         uuid.UUID   `gorm:"type:uuid;primaryKey" json:"city_id"`
	Name       string      `json:"city_name"`
	Users      []User      `gorm:"foreignKey:CityID"`
	Psychologs []Psycholog `gorm:"foreignKey:CityID"`

	TimeStamp
}

func (c *City) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	c.ID = uuid.New()

	return nil
}
