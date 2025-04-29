package entity

import "github.com/google/uuid"

type Province struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"province_id"`
	Name string    `json:"province_name"`

	Cities []City `gorm:"foreignKey:ProvinceID"`
}
