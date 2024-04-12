package repo

import (
	"github.com/google/uuid"
	models "github.com/tylerneath/nuture-backend/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type MessageRepository interface {
	Create(*models.Message) (uuid.UUID, error)
	Get(uuid.UUID) (models.Message, error)
}

type messageRepository struct {
	log *zap.Logger
	db  *gorm.DB
}

func NewMessageRepository(db *gorm.DB, log *zap.Logger) MessageRepository {
	return &messageRepository{db: db, log: log}
}

func (m *messageRepository) Create(message *models.Message) (uuid.UUID, error) {
	err := m.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(message).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return uuid.UUID{}, err
	}
	return message.ID, nil
}

func (m *messageRepository) Get(id uuid.UUID) (models.Message, error) {
	message := models.Message{}
	if err := m.db.First(&message, id).Error; err != nil {
		return message, err
	}
	return message, nil
}

func (m *messageRepository) GetByUserID(userId uuid.UUID) ([]models.Message, error) {
	messages := []models.Message{}
	if err := m.db.Where("user_id = ?", userId).Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}
