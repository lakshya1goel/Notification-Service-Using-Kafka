package kafka

import (
	"context"
	"encoding/json"
	"log"

	services "github.com/lakshya1goel/Notification-Service-Using-Kafka.git/api/service"
	"github.com/lakshya1goel/Notification-Service-Using-Kafka.git/domain/model"
	"github.com/segmentio/kafka-go"
)

type KafkaConsumer interface {
	Start(ctx context.Context)
}

type kafkaConsumer struct {
	reader *kafka.Reader
	pushSender services.PushNotificationService
}

func NewKafkaConsumer(brokers []string, topic, groupID string, pushSender services.PushNotificationService) KafkaConsumer {
	return &kafkaConsumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:  brokers,
			Topic:    topic,
			GroupID:  groupID,
			MinBytes: 10e3,
			MaxBytes: 10e6,
		}),
		pushSender: pushSender,
	}
}

func (c *kafkaConsumer) Start(ctx context.Context) {
	go func() {
		for {
			m, err := c.reader.ReadMessage(ctx)
			if err != nil {
				log.Println("Kafka read error:", err)
				continue
			}

			var notif model.Notification
			if err := json.Unmarshal(m.Value, &notif); err != nil {
				log.Println("Unmarshal error:", err)
				continue
			}

			log.Println("Received notification:", notif)
		}
	}()
}
