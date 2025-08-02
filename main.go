package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lakshya1goel/Notification-Service-Using-Kafka/api/controller"
	routes "github.com/lakshya1goel/Notification-Service-Using-Kafka/api/router"
	services "github.com/lakshya1goel/Notification-Service-Using-Kafka/api/service"
	"github.com/lakshya1goel/Notification-Service-Using-Kafka/api/util"
	"github.com/lakshya1goel/Notification-Service-Using-Kafka/kafka"
)

func main() {

	ctx := context.Background()

	if err := util.InitFirebase(ctx); err != nil {
		log.Fatalf("Firebase init failed: %v", err)
	}

	pushNotificationSender, err := services.NewPushNotificationService(ctx)
	if err != nil {
		log.Fatalf("Push Notification Sender initialization failed: %v", err)
	}

	consumer := kafka.NewKafkaConsumer(kafka.KafkaConsumerConfig{
		Brokers:    []string{"localhost:9092"},    //TODO: replace this with actual kafka brokers
		Topic:      "notifications",               //TODO: replace this with actual kafka topic
		GroupID:    "notification-consumer-group", //TODO: replace this with actual kafka group id
		PushSender: pushNotificationSender,
	})
	consumer.Start(ctx)

	router := gin.Default()

	apiRouter := router.Group("/api")
	{
		routes.NotificationRoutes(apiRouter, controller.NewNotificationController())
	}

	router.Run(":8000")
}
