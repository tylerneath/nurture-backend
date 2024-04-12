package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tylerneath/nuture-backend/internal/common"
	"github.com/tylerneath/nuture-backend/internal/service"
	"go.uber.org/zap"
)

type MessageHandler struct {
	messageService service.MessageService
	log            *zap.Logger
}

func NewMessageHandler(messageService service.MessageService, log *zap.Logger) *MessageHandler {
	return &MessageHandler{
		log:            log,
		messageService: messageService,
	}
}

func (v *MessageHandler) CreateMessage(c *gin.Context) {
	var newMessage common.CreateMessageRequest
	if err := c.BindJSON(&newMessage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := v.messageService.CreateMessage(newMessage)
	if err != nil {
		if errors.Is(err, service.ErrMissingField) {
			c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("missing text field")})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
		}
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (v *MessageHandler) GetMessageByID(c *gin.Context) {
	id := c.Params.ByName("id")
	message, err := v.messageService.GetByID(id)
	if err != nil {
		if err == service.ErrMessageNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Message not found"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": message})

}

func (v *MessageHandler) DeleteMessage(c *gin.Context) {
	// delete message based on message id
	// must have a user to delete message 
	

	panic("implement me")
}
