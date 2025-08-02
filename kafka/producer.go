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

type KafkaProducerConfig struct {
	Broker string
	Topic  string
}

func NewKafkaProducer(config KafkaProducerConfig) KafkaProducer {
	return &kafkaProducer{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(config.Broker),
			Topic:    config.Topic,
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
