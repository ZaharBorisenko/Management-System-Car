package main

import (
	"fmt"
	"github.com/ZaharBorisenko/Management-System-Car/internal/database/db"
	carHandler "github.com/ZaharBorisenko/Management-System-Car/internal/handler/car"
	engineHandler "github.com/ZaharBorisenko/Management-System-Car/internal/handler/engine"
	"github.com/ZaharBorisenko/Management-System-Car/internal/models"
	carService "github.com/ZaharBorisenko/Management-System-Car/internal/service/car"
	engineService "github.com/ZaharBorisenko/Management-System-Car/internal/service/engine"
	carStore "github.com/ZaharBorisenko/Management-System-Car/internal/store/car"
	engineStore "github.com/ZaharBorisenko/Management-System-Car/internal/store/engine"
	"github.com/ZaharBorisenko/Management-System-Car/pkg/logger"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	//init DB
	err := godotenv.Load(".env")
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
	db.InitDB(&cfg)
	defer db.CloseDB()
	db := db.GetDB()

	//init logger
	logger := logger.InitLogger()
	fmt.Println("slog initialization!", logger)

	//Init store
	carStore := carStore.NewCarStore(db, logger)
	engineStore := engineStore.NewEngineStore(db, logger)

	//Init service
	carService := carService.NewCarService(carStore, logger)
	engineService := engineService.NewEngineService(engineStore, logger)

	//Init handlers
	carHandler := carHandler.NewCarHandler(carService, logger)
	engineHandler := engineHandler.NewEngineHandler(engineService, logger)

	//router
	mux := http.NewServeMux()

	mux.HandleFunc("GET /cars/{id}", carHandler.GetCarById)
	mux.HandleFunc("GET /cars/{brand}", carHandler.GetCarByBrand)
	mux.HandleFunc("POST /cars", carHandler.CreateCar)

	mux.HandleFunc("GET /engine/{id}", engineHandler.GetEngineById)
	mux.HandleFunc("POST /engine", engineHandler.CreateEngine)

	// run server
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = ":8080"
	}

	err = http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatalf("Server not listening %v", err)
	}

}
