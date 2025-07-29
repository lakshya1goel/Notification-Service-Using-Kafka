package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lakshya1goel/Notification-Service-Using-Kafka.git/api/controller"
	routes "github.com/lakshya1goel/Notification-Service-Using-Kafka.git/api/router"
	services "github.com/lakshya1goel/Notification-Service-Using-Kafka.git/api/service"
	"github.com/lakshya1goel/Notification-Service-Using-Kafka.git/api/util"
	"github.com/lakshya1goel/Notification-Service-Using-Kafka.git/kafka"
)

func main() {

	ctx := context.Background()

	if err := util.InitFirebase(ctx); err != nil {
		log.Fatalf("Firebase init failed: %v", err)
	}

	pushSender, err := services.NewPushNotificationService(ctx)
	if err != nil {
		log.Fatalf("Push Sender init failed: %v", err)
	}

	consumer := kafka.NewKafkaConsumer(
		[]string{"localhost:9092"},
		"notifications",
		"notification-consumer-group",
		pushSender,
	)
	consumer.Start(ctx)

	router := gin.Default()

	apiRouter := router.Group("/api")
	{
		routes.NotificationRoutes(apiRouter, controller.NewNotificationController())
	}

	router.Run(":8000")
}
