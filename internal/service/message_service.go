package service

import (
	"context"

	"github.com/tylerneath/nuture-backend/internal/common"
	"github.com/tylerneath/nuture-backend/internal/repo"
)

type MessageService interface {
	CreateMessage(req common.CreateMessageRequest) error
}

type messageService struct {
	messageRepository repo.MessageRepository
	userRepository    repo.UserRepository
}

func (s *messageService) CreateMessage(req common.CreateMessageRequest) error {
	return nil
}

func NewMessageService(ctx context.Context, messageRepo repo.MessageRepository, userRepo repo.UserRepository) MessageService {
	return &messageService{
		messageRepository: messageRepo,
		userRepository:    userRepo,
	}
}
