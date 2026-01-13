package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"os"
	"time"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer() *Producer {
	brokers := os.Getenv("KAFKA_BROKERS") // "kafka:9092" из .env
	if brokers == "" {
		brokers = "kafka:9092"
	}

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      []string{brokers},
		Topic:        "car-events",
		RequiredAcks: -1, // At least once
		MaxAttempts:  10,
		BatchSize:    100,
		WriteTimeout: 10 * time.Second,
		Balancer:     &kafka.RoundRobin{},
	})

	return &Producer{writer: writer}
}

func (p *Producer) SendMessage(ctx context.Context, value []byte) error {
	err := p.writer.WriteMessages(ctx, kafka.Message{
		Value: value,
	})
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	return nil
}

func (p *Producer) Close() {
	p.writer.Close()
}
