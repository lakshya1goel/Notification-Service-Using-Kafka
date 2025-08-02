package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lakshya1goel/Notification-Service-Using-Kafka/domain/dto"
	"github.com/lakshya1goel/Notification-Service-Using-Kafka/usecase"
)

type NotificationController struct {
	notificationUsecase usecase.NotificationUseCase
}

func NewNotificationController() *NotificationController {
	return &NotificationController{
		notificationUsecase: usecase.NewNotificationUseCase(),
	}
}

func (ctrl *NotificationController) SendNotification(c *gin.Context) {
	var req dto.NotificationRequestDto
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}

	if err := ctrl.notificationUsecase.SendNotification(c, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Notification published to Kafka"})
}
