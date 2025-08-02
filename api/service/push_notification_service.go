package services

import (
	"context"

	"firebase.google.com/go/v4/messaging"
	"github.com/lakshya1goel/Notification-Service-Using-Kafka/api/util"
	"github.com/lakshya1goel/Notification-Service-Using-Kafka/domain/model"
)

type PushNotificationService interface {
	SendPushNotification(ctx context.Context, notif model.Notification) error
}

type pushNotificationService struct {
	client *messaging.Client
}

func NewPushNotificationService(ctx context.Context) (PushNotificationService, error) {
	client, err := util.FirebaseApp.Messaging(ctx)
	if err != nil {
		return nil, err
	}
	return &pushNotificationService{client: client}, nil
}

func (p *pushNotificationService) SendPushNotification(ctx context.Context, notif model.Notification) error {
	fcmToken := "fcmToken"
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: notif.Title,
			Body:  notif.Message,
		},
		Token: fcmToken,
	}

	_, err := p.client.Send(ctx, message)
	if err != nil {
		return err
	}
	return nil
}
