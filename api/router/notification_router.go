package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lakshya1goel/Notification-Service-Using-Kafka/api/controller"
)

func NotificationRoutes(router *gin.RouterGroup, notificationController *controller.NotificationController) {
	notificationRouter := router.Group("/notification")
	{
		notificationRouter.POST("/", notificationController.SendNotification)
	}
}
