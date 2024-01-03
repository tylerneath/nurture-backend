package store

import (
	models "github.com/tylerneath/nuture-backend/internal/model"
	"gorm.io/gorm"
)

type MessageStore struct {
	db *gorm.DB
}

func NewMessageStore(db *gorm.DB) *MessageStore {
	return &MessageStore{
		db: db,
	}
}

func CreateMessage(m *models.Message) error {
	panic("not implemented")
}

func GetMessage(id string) (*models.Message, error) {
	panic("not implemented")
}

func GetUserMessages(id string, userId string) (*models.Message, error) {
	panic("not implemented")
}
