package main

import (
	"context"
	"fmt"
	"github.com/ZaharBorisenko/Management-System-Car/internal/broker/kafka"
	"github.com/ZaharBorisenko/Management-System-Car/internal/database/db"
	carHandler "github.com/ZaharBorisenko/Management-System-Car/internal/handler/car"
	engineHandler "github.com/ZaharBorisenko/Management-System-Car/internal/handler/engine"
	"github.com/ZaharBorisenko/Management-System-Car/internal/models"
	"github.com/ZaharBorisenko/Management-System-Car/internal/routes"
	carService "github.com/ZaharBorisenko/Management-System-Car/internal/service/car"
	engineService "github.com/ZaharBorisenko/Management-System-Car/internal/service/engine"
	carStore "github.com/ZaharBorisenko/Management-System-Car/internal/store/car"
	engineStore "github.com/ZaharBorisenko/Management-System-Car/internal/store/engine"
	"github.com/ZaharBorisenko/Management-System-Car/pkg/logger"
	"github.com/ZaharBorisenko/Management-System-Car/pkg/rateLimiter"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	//init DB
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	cfg := models.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}
	err = db.InitDB(&cfg)
	if err != nil {
		log.Fatalf("Failed to init DB: %v", err)
	}
	defer db.CloseDB()
	db := db.GetDB()

	//init logger
	logger := logger.InitLogger()
	fmt.Println("slog initialization!", logger)

	// Init store
	engineStore := engineStore.NewEngineStore(db)
	carStore := carStore.NewCarStore(db, engineStore, logger)

	// Init service
	carService := carService.NewCarService(carStore, logger)
	engineService := engineService.NewEngineService(engineStore)

	// Init handlers
	carHandler := carHandler.NewCarHandler(carService, logger)
	engineHandler := engineHandler.NewEngineHandler(engineService, logger)

	// init rate limiter
	limiter := rateLimiter.NewLimiter(10, time.Minute)

	//router
	routing := routes.RegisterRoutes(carHandler, engineHandler)
	handlerRout := limiter.Middleware(routing)

	//==========TEST KAFKA====================
	producer := kafka.NewProducer()
	defer producer.Close()

	consumer := kafka.NewConsumer()
	defer consumer.Close()

	// Тест producer:
	ctx := context.Background()
	err = producer.SendMessage(ctx, []byte("Test message from Car Manager!"))
	if err != nil {
		log.Printf("Kafka producer error: %v", err)
	} else {
		log.Println("Kafka message sent successfully")
	}

	// Тест consumer:
	go func() {
		for {
			msg, err := consumer.ReadMessage(ctx)
			if err != nil {
				log.Printf("Kafka consumer error: %v", err)
				time.Sleep(1 * time.Second)
				continue
			}
			log.Printf("Received Kafka message: %s", msg)
		}
	}()

	//==========TEST KAFKA END====================

	// run server
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = ":8080"
	}

	err = http.ListenAndServe(port, handlerRout)
	if err != nil {
		log.Fatalf("Server not listening %v", err)
	}

}
