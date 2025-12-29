package dto

import (
	"time"

	"github.com/google/uuid"
)

type EngineCreateRequest struct {
	Description   string `json:"description"`
	Displacement  int64  `json:"displacement" validate:"required,gte=0,lte=10000"`
	NoOfCylinders int64  `json:"no_of_cylinders" validate:"required,oneof=0 2 3 4 6 8 10 12 16"`
	CarRange      int64  `json:"car_range" validate:"gte=0,lte=2000"`
	HorsePower    int64  `json:"horse_power" validate:"required,gte=1,lte=2000"`
	Torque        int64  `json:"torque" validate:"required,gte=1,lte=5000"`
	EngineType    string `json:"engine_type" validate:"required,oneof=Turbine Hybrid petrol diesel electric rotary"`
	EmissionClass string `json:"emission_class" validate:"required,oneof=Euro 3 Euro 4 Euro 5 Euro 6 Euro 6d"`
}

type EngineUpdateRequest struct {
	Description   *string `json:"description" validate:"omitempty"`
	Displacement  *int64  `json:"displacement" validate:"omitempty,gte=0,lte=10000"`
	NoOfCylinders *int64  `json:"no_of_cylinders" validate:"omitempty,oneof=0 2 3 4 6 8 10 12 16"`
	CarRange      *int64  `json:"car_range" validate:"omitempty,gte=0,lte=2000"`
	HorsePower    *int64  `json:"horse_power" validate:"omitempty,gte=1,lte=2000"`
	Torque        *int64  `json:"torque" validate:"omitempty,gte=1,lte=5000"`
	EngineType    *string `json:"engine_type" validate:"omitempty,oneof=Turbine Hybrid petrol diesel electric rotary"`
	EmissionClass *string `json:"emission_class" validate:"omitempty,oneof=Euro 3 Euro 4 Euro 5 Euro 6 Euro 6d"`
}

type EngineResponse struct {
	ID            uuid.UUID `json:"id"`
	Description   string    `json:"description"`
	Displacement  int64     `json:"displacement"`
	NoOfCylinders int64     `json:"no_of_cylinders"`
	CarRange      int64     `json:"car_range"`
	HorsePower    int64     `json:"horse_power"`
	Torque        int64     `json:"torque"`
	EngineType    string    `json:"engine_type"`
	EmissionClass string    `json:"emission_class"`
	CreatedAt     time.Time `json:"created_at"`
}
