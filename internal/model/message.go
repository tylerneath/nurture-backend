package models

import (
	"errors"

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
	Text          string        `gorm:"column:text;type:varchar;size:255" json:"text"`
	MessageType   MessageType   `gorm:"column:message_type;type:text" json:"message_type"`
	MessageAction MessageAction `gorm:"column:message_action;type:text" json:"message_action"`
	UserID        uuid.UUID
}

func isValidMessageType(messageType MessageType) bool {
	return messageType == Log || messageType == Reminder
}

func isValidMessageAction(messageAction MessageAction) bool {
	return messageAction == Reflect || messageAction == Advise
}

func (m *Message) BeforeSave(tx *gorm.DB) error {
	println(isValidMessageType(m.MessageType), isValidMessageAction(m.MessageAction))
	if !common.Xor(isValidMessageType(m.MessageType), isValidMessageAction(m.MessageAction)) {
		return errors.New("invalid message type or action. either one should be set")
	}
	return nil
}
