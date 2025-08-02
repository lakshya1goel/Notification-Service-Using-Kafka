package usecase

import (
	"context"
	"errors"

	"github.com/lakshya1goel/Notification-Service-Using-Kafka/domain/dto"
	"github.com/lakshya1goel/Notification-Service-Using-Kafka/repository"
)

type NotificationUseCase interface {
	SendNotification(ctx context.Context, req dto.NotificationRequestDto) error
}

type notificationUseCase struct {
	notificationRepo repository.NotificationRepository
}

func NewNotificationUseCase() NotificationUseCase {
	return &notificationUseCase{
		notificationRepo: repository.NewNotificationRepository(),
	}
}

func (uc *notificationUseCase) SendNotification(ctx context.Context, req dto.NotificationRequestDto) error {
	if req.Title == "" || req.Message == "" {
		return errors.New("invalid request: title and message are required")
	}

	err := uc.notificationRepo.PublishNotification(ctx, req)
	if err != nil {
		return errors.New("failed to publish notification: " + err.Error())
	}
	return nil
}
