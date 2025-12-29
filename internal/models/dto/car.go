package dto

import (
	"github.com/google/uuid"
	"time"
)

type CarCreateRequest struct {
	Description  string    `json:"description" validate:"required,min=10,max=2000"`
	Year         int64     `json:"year" validate:"required,gte=1886,lte=2100"`
	Model        string    `json:"model" validate:"required,min=1,max=100"`
	FuelType     string    `json:"fuel_type" validate:"required,oneof=petrol diesel hybrid electric gas"`
	EngineID     uuid.UUID `json:"engine_id" validate:"required,uuid4"`
	BrandID      uuid.UUID `json:"brand_id" validate:"required,uuid4"`
	Price        float64   `json:"price" validate:"required,gt=0"`
	VIN          string    `json:"vin" validate:"required,len=17,alphanum"`
	Mileage      int64     `json:"mileage" validate:"gte=0,lte=3000000"`
	Transmission string    `json:"transmission" validate:"required,oneof=manual automatic robot cvt"`
	Color        string    `json:"color" validate:"required,min=2,max=30"`
	BodyType     string    `json:"body_type" validate:"required,oneof=sedan hatchback wagon suv coupe convertible pickup van"`
}

type CarResponse struct {
	ID           uuid.UUID   `json:"id"`
	Description  string      `json:"description"`
	Year         int64       `json:"year"`
	Model        string      `json:"model"`
	FuelType     string      `json:"fuel_type"`
	Price        float64     `json:"price"`
	VIN          string      `json:"vin"`
	Mileage      int64       `json:"mileage"`
	Transmission string      `json:"transmission"`
	Color        string      `json:"color"`
	BodyType     string      `json:"body_type"`
	Brand        BrandShort  `json:"brand"`
	Engine       EngineShort `json:"engine"`
	CreatedAt    time.Time   `json:"created_at"`
}

type BrandShort struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type EngineShort struct {
	ID         uuid.UUID `json:"id"`
	HorsePower int64     `json:"horse_power"`
	EngineType string    `json:"engine_type"`
}

type CarUpdateRequest struct {
	Description  *string    `json:"description" validate:"omitempty,min=10,max=2000"`
	Year         *int64     `json:"year" validate:"omitempty,gte=1886,lte=2100"`
	Model        *string    `json:"model" validate:"omitempty"`
	FuelType     *string    `json:"fuel_type" validate:"omitempty,oneof=petrol diesel hybrid electric gas"`
	EngineID     *uuid.UUID `json:"engine_id" validate:"omitempty,uuid4"`
	BrandID      *uuid.UUID `json:"brand_id" validate:"omitempty,uuid4"`
	Price        *float64   `json:"price" validate:"omitempty,gt=0"`
	VIN          *string    `json:"vin" validate:"omitempty,len=17"`
	Mileage      *int64     `json:"mileage" validate:"omitempty,gte=0"`
	Transmission *string    `json:"transmission" validate:"omitempty,oneof=manual automatic robot cvt"`
	Color        *string    `json:"color" validate:"omitempty"`
	BodyType     *string    `json:"body_type" validate:"omitempty"`
}
