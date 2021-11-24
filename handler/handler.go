package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/helyus1412/messaging-api/helper"
	"github.com/helyus1412/messaging-api/services"
)

type Handler struct {
	service services.Services
}

func NewHandler(services services.Services) *Handler {
	return &Handler{services}
}

func (h *Handler) GetAllMessages(c *gin.Context) {
	allMessages, err := h.service.AllMessages()

	if err != nil {
		resErr := helper.APIFailure(500, "internal server error", err.Error())
		c.JSON(500, resErr)
		return
	}

	res := helper.APIResponse(200, "success get all messages", allMessages)

	c.JSON(200, res)
}
