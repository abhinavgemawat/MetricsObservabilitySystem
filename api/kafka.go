package api

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

// KafkaMessage represents the message structure in Kafka
type KafkaMessage struct {
	Key   string
	Value string
	Time  time.Time
}

// ProduceMessage sends a message to the Kafka topic
func ProduceMessage(topic string, message KafkaMessage) error {
	w := kafka.Writer{
		Addr:     kafka.TCP("localhost:21250"),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(message.Key),
			Value: []byte(message.Value),
			Time:  message.Time,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to write message: %v", err)
	}

	return nil
}

// ConsumeMessages reads messages from the Kafka topic
func ConsumeMessages(topic string, handler func(KafkaMessage)) error {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   topic,
		GroupID: "metrics-group",
	})

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			return fmt.Errorf("failed to read message: %v", err)
		}

		kafkaMsg := KafkaMessage{
			Key:   string(msg.Key),
			Value: string(msg.Value),
			Time:  msg.Time,
		}

		handler(kafkaMsg)
	}
}
