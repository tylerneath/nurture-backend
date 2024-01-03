package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type messageType string

const (
	reflect messageType = "reflect"
	recall  messageType = "recall"
)

type Base struct {
	ID        uuid.UUID      `gorm:"column:id;type:uuid;primary_key;" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"updated_at" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"deleted_at,index" json:"deletedAt"`
}

type User struct {
	Base
	Name     string
	Username string `gorm:"uniqueIndex;"`
	Email    string `gorm:"unique;not null"`
	Messages []Message
}

type Message struct {
	Base
	Text          string
	MessageType   MessageType   `gorm:"column:message_type;type:text" json:"message_type"`
	MessageAction MessageAction `gorm:"column:message_action;type:text" json:"message_action"`
	UserID        uuid.UUID
}

type MessageType struct {
	Base
	Type string `gorm:"unique;not null"`
}

type MessageAction struct {
	Base
	Action string `gorm:"unique;not null"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) error {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return errors.New("unable to create uuid for object")
	}
	base.ID = uuid
	return nil

}
