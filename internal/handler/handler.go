package handler

import (
	"context"

	"github.com/tylerneath/nuture-backend/internal/service"
	"go.uber.org/zap"
)

type Handler struct {
	UserHandler    UserHandler
	MessageHandler MessageHandler
}

func New(ctx context.Context, messageService service.MessageService, userService service.UserService, log *zap.Logger) (*UserHandler, *MessageHandler) {
	return &UserHandler{
			userService: userService,
			log:         log,
		}, &MessageHandler{
			messageService: messageService,
		}
}
