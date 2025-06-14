package entity

import "github.com/google/uuid"

type Message struct {
	ID      uuid.UUID `gorm:"type:uuid;primary_key"`
	Sender  string    // "user" or "assistant"
	Content string

	ConversationID *uuid.UUID   `gorm:"type:uuid" json:"conversation_id"`
	Conversation   Conversation `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	TimeStamp
}
