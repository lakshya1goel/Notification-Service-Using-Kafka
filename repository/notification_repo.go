package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/lakshya1goel/Notification-Service-Using-Kafka/domain/dto"
	"github.com/lakshya1goel/Notification-Service-Using-Kafka/domain/model"
	"github.com/lakshya1goel/Notification-Service-Using-Kafka/kafka"
)

type NotificationRepository interface {
	PublishNotification(ctx context.Context, req dto.NotificationRequestDto) error
}

type notificationRepository struct {
	producer kafka.KafkaProducer
}

func NewNotificationRepository() NotificationRepository {
	return &notificationRepository{
		producer: kafka.NewKafkaProducer(kafka.KafkaProducerConfig{
			Broker: "localhost:9092", //TODO: replace this with actual kafka broker
			Topic:  "notifications",  //TODO: replace this with actual kafka topic
		}),
	}
}

func (r *notificationRepository) PublishNotification(ctx context.Context, req dto.NotificationRequestDto) error {
	notification := model.Notification{
		ID:      uuid.New().String(),
		Title:   req.Title,
		Message: req.Message,
	}
	return r.producer.Produce(ctx, notification)
}
