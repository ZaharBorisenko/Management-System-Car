package models

import (
	"time"

	"github.com/google/uuid"
)

type Car struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	Year        int64     `json:"year"`
	//Brand        string    `json:"brand"`
	Model        string    `json:"model"`
	FuelType     string    `json:"fuel_type"` // allowed: "petrol", "diesel", "electric", "hybrid", "hydrogen", "gas"
	Engine       Engine    `json:"engine"`
	Brand        Brand     `json:"brand"`
	Price        float64   `json:"price"`
	VIN          string    `json:"vin"`
	Mileage      int64     `json:"mileage"`
	Transmission string    `json:"transmission"` // allowed: "automatic", "manual", "robotic", "cvt"
	Color        string    `json:"color"`        // allowed: "black", "white", "gray", "silver", "blue", "red", "green", "yellow", "brown", "beige"
	BodyType     string    `json:"body_type"`    // allowed: "sedan", "hatchback", "coupe", "cabriolet", "suv", "crossover", "pickup", "wagon", "van", "micro"
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CarRequestDTO struct {
	Description string `json:"description" validate:"required,min=10,max=2000"`
	Year        int64  `json:"year" validate:"required,gte=1886,lte=2100"`
	//Brand        string  `json:"brand" validate:"required,min=2,max=50"`
	Model        string  `json:"model" validate:"required,min=1,max=100"`
	FuelType     string  `json:"fuel_type" validate:"required,oneof=petrol diesel hybrid electric gas"`
	Engine       Engine  `json:"engine" validate:"required"`
	Brand        Brand   `json:"brand"`
	Price        float64 `json:"price" validate:"required,gt=0"`
	VIN          string  `json:"vin" validate:"required,len=17,alphanum"`
	Mileage      int64   `json:"mileage" validate:"required,gte=0,lte=3000000"`
	Transmission string  `json:"transmission" validate:"required,oneof=manual automatic robot cvt"`
	Color        string  `json:"color" validate:"required,min=2,max=30"`
	BodyType     string  `json:"body_type" validate:"required,oneof=sedan hatchback wagon suv coupe convertible pickup van"`
}

type CarUpdateDTO struct {
	Description  *string  `json:"description" validate:"omitempty,min=10,max=2000"`
	Year         *int64   `json:"year" validate:"omitempty,gte=1886,lte=2100"`
	Brand        *string  `json:"brand" validate:"omitempty,min=2,max=50"`
	Model        *string  `json:"model" validate:"omitempty,min=1,max=100"`
	FuelType     *string  `json:"fuel_type" validate:"omitempty,oneof=petrol diesel hybrid electric gas"`
	EngineID     *string  `json:"engine_id" validate:"omitempty,uuid4"`
	BrandID      *string  `json:"brand_id" validate:"omitempty,uuid4"`
	Price        *float64 `json:"price" validate:"omitempty,gt=0"`
	VIN          *string  `json:"vin" validate:"omitempty,len=17,alphanum"`
	Mileage      *int64   `json:"mileage" validate:"omitempty,gte=0,lte=3000000"`
	Transmission *string  `json:"transmission" validate:"omitempty,oneof=manual automatic robot cvt"`
	Color        *string  `json:"color" validate:"omitempty,min=2,max=30"`
	BodyType     *string  `json:"body_type" validate:"omitempty,oneof=sedan hatchback wagon suv coupe convertible pickup van"`
}
