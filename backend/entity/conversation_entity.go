package entity

import "github.com/google/uuid"

type Conversation struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key"`

	UserID *uuid.UUID `gorm:"type:uuid" json:"user_id"`
	User   User       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Messages []Message `gorm:"foreignKey:ConversationID"`

	TimeStamp
}
