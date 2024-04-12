package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/tylerneath/nuture-backend/internal/common"
	models "github.com/tylerneath/nuture-backend/internal/model"
	"github.com/tylerneath/nuture-backend/internal/repo"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type MessageService interface {
	CreateMessage(req common.CreateMessageRequest) (string, error)
	GetByID(id string) (common.MessageResponse, error)
}

type messageService struct {
	messageRepository repo.MessageRepository
	log               *zap.Logger
}

func (s *messageService) CreateMessage(req common.CreateMessageRequest) (string, error) {
	// validate the request
	if req.Text == "" || (req.MessageType == "" && req.MessageAction == "") {
		return "", ErrMissingField
	}
	message := models.Message{
		Text:          req.Text,
		MessageType:   req.MessageType,
		MessageAction: req.MessageAction,
	}

	id, err := s.messageRepository.Create(&message)
	if err != nil {
		return "", NewInternalServerError(err)
	}

	return id.String(), nil
}

func (s *messageService) GetByID(id string) (common.MessageResponse, error) {
	var message common.MessageResponse
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return message, ErrInvalidID
	}
	msg, err := s.messageRepository.Get(parsedID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return message, ErrMessageNotFound
		} else {
			return message, NewInternalServerError(err)
		}
	}
	message = common.MessageResponse{
		ID:        msg.ID,
		UserID:    msg.UserID,
		Message:   msg.Text,
		CreatedAt: msg.CreatedAt,
	}

	return message, nil
}

func NewMessageService(ctx context.Context, messageRepo repo.MessageRepository, log *zap.Logger) MessageService {
	return &messageService{
		messageRepository: messageRepo,
		log:               log,
	}
}
