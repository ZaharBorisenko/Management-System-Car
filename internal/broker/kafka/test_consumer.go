package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"os"
)

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer() *Consumer {
	brokers := os.Getenv("KAFKA_BROKERS") // "kafka:9092"
	if brokers == "" {
		brokers = "kafka:9092"
	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokers},
		Topic:   "car-events",
		GroupID: "car-group",
	})

	return &Consumer{reader: reader}
}

func (c *Consumer) ReadMessage(ctx context.Context) (string, error) {
	msg, err := c.reader.ReadMessage(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to read message: %w", err)
	}
	return string(msg.Value), nil
}

func (c *Consumer) Close() {
	c.reader.Close()
}
