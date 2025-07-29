package repository

import (
	"context"

	"github.com/lakshya1goel/Notification-Service-Using-Kafka.git/domain/dto"
	"github.com/lakshya1goel/Notification-Service-Using-Kafka.git/domain/model"
	"github.com/lakshya1goel/Notification-Service-Using-Kafka.git/kafka"
	"github.com/google/uuid"
)

type NotificationRepository interface {
	PublishNotification(ctx context.Context, req dto.NotificationRequestDto) error
}

type notificationRepository struct {
	producer kafka.KafkaProducer
}

func NewNotificationRepository() NotificationRepository {
	return &notificationRepository{
		producer: kafka.NewKafkaProducer("localhost:9092", "notifications"),
	}
}

func (r *notificationRepository) PublishNotification(ctx context.Context, req dto.NotificationRequestDto) error {
	notification := model.Notification{
		ID:        uuid.New().String(),
		Title:     req.Title,
		Message:   req.Message,
	}
	return r.producer.Produce(ctx, notification)
}
