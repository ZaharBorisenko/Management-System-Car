package db

import (
	"database/sql"
	"fmt"
	"github.com/ZaharBorisenko/Management-System-Car/internal/models"
	"log"
)

var db *sql.DB

func GetDB() *sql.DB {
	return db
}

func CloseDB() {
	if err := db.Close(); err != nil {
		log.Fatalf("Error closing to database : %v", err)
	}
}

func InitDB(c *models.Config) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.Name,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error opening to database : %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connection to database : %v", err)
	}

	fmt.Println("Successfully connected!")
}
