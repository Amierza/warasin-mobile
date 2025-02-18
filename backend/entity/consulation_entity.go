package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Consulation struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey" json:"consul_id"`
	UserID      uuid.UUID  `gorm:"type:uuid" json:"user_id"`
	User        User       `gorm:"foreignKey:UserID"`
	PsychologID uuid.UUID  `gorm:"type:uuid" json:"psy_id"`
	Psycholog   Psycholog  `gorm:"foreignKey:PsychologID"`
	Date        *time.Time `gorm:"type:date" json:"consul_date"`
	Rate        int        `json:"consul_rate"`
	Comment     string     `json:"consul_comment"`

	TimeStamp
}

func (c *Consulation) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	c.ID = uuid.New()

	return nil
}
