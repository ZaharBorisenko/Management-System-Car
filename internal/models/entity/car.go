package entity

import (
	"github.com/google/uuid"
	"time"
)

type Car struct {
	ID           uuid.UUID
	Description  string
	Year         int64
	Model        string
	FuelType     string
	EngineID     uuid.UUID
	BrandID      uuid.UUID
	Price        float64
	VIN          string
	Mileage      int64
	Transmission string
	Color        string
	BodyType     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
