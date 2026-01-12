package producer

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func main() {
	ctx := context.Background()

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      []string{"localhost:9092"},
		Topic:        "my-topic",
		RequiredAcks: -1,                  // Подтверждение от всех реплик
		MaxAttempts:  10,                  //кол-во попыток доставки(по умолчанию всегда 10)
		BatchSize:    100,                 // Ограничение на количество сообщений(по дефолту 100)
		WriteTimeout: 10 * time.Second,    //время ожидания для записи(по умолчанию 10сек)
		Balancer:     &kafka.RoundRobin{}, //балансировщик.
	})
	defer writer.Close()

	err := writer.WriteMessages(ctx, kafka.Message{
		Value: []byte("Hello, Kafka!"),
	})

	if err != nil {
		log.Fatal("Ошибка при отправке:", err)
	}
	fmt.Println("Сообщение отправлено с гарантией At least once.")
}
