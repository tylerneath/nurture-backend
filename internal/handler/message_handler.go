package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tylerneath/nuture-backend/internal/common"
	"github.com/tylerneath/nuture-backend/internal/service"
)

type MessageHandler struct {
	messageService service.MessageService
}

func (v *MessageHandler) CreateMessage(c *gin.Context) {
	var newMessage common.CreateMessageRequest
	if err := c.BindJSON(&newMessage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}
     