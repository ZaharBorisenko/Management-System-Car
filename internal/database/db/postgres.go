package db

import (
	"database/sql"
	"fmt"
	"github.com/ZaharBorisenko/Management-System-Car/internal/models"
	_ "github.com/lib/pq"
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

func InitDB(c *models.Config) error {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.Name,
	)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Successfully connected!")
	return nil
}
