package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/lakshya1goel/Notification-Service-Using-Kafka/domain/model"
	"github.com/segmentio/kafka-go"
)

type KafkaProducer interface {
	Produce(ctx context.Context, notif model.Notification) error
}

type kafkaProducer struct {
	writer *kafka.Writer
}

func NewKafkaProducer(broker, topic string) KafkaProducer {
	return &kafkaProducer{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(broker),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		},
	}
}

func (kp *kafkaProducer) Produce(ctx context.Context, notif model.Notification) error {
	payload, err := json.Marshal(notif)
	if err != nil {
		return err
	}

	err = kp.writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(notif.ID),
		Value: payload,
	})
	if err != nil {
		log.Println("Kafka write failed:", err)
	}
	return err
}
