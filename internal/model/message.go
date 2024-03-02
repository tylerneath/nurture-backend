package models

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/tylerneath/nuture-backend/internal/common"
	"gorm.io/gorm"
)

type MessageType string

type MessageAction string

var (
	Log      MessageType = "log"
	Reminder MessageType = "reminder"
)

var (
	Reflect MessageAction = "reflect"
	Advise  MessageAction = "advise"
)

type Message struct {
	Base
	Text          string         `gorm:"column:text;type:varchar;size:255" json:"text"`
	MessageType   *MessageType   `gorm:"column:message_type;type:text" json:"message_type"`
	MessageAction *MessageAction `gorm:"column:message_action;type:text" json:"message_action"`
	UserID        uuid.UUID
}

func CreateMessage(db *gorm.DB) error {
	panic("Implement Me")

	return nil
}

func (m *MessageType) Valid() bool {
	if m == nil {
		return false
	}

	switch *m {
	case Log, Reminder:
		return true
	default:
		return false
	}
}

func (m *MessageAction) Valid() bool {
	if m == nil {
		return false
	}

	switch *m {
	case Reflect, Advise:
		return true
	default:
		return false
	}

}

func (m *Message) BeforeSave(tx *gorm.DB) error {

	isMessageType := m.MessageAction.Valid()
	isActionType := m.MessageType.Valid()

	if valid := common.Xor(isMessageType, isActionType); !valid {
		return fmt.Errorf("invalid message action/type for message")
	}

	return nil
}
