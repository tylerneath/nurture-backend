package repo

import (
	"github.com/google/uuid"
	models "github.com/tylerneath/nuture-backend/internal/model"
	"gorm.io/gorm"
)

type MessageRepository interface {
	Create(uuid.UUID, models.Message) error
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &messageRepository{db: db}
}

func (m *messageRepository) Create(userId uuid.UUID, message models.Message) error {
	return m.db.Transaction(func(tx *gorm.DB) error {
		user := models.User{}
		if err := tx.First(&user, userId).Error; err != nil {
			return err
		}

		message.UserID = user.ID
		if err := tx.Create(&message).Error; err != nil {
			return err
		}

		return nil
	})
}
