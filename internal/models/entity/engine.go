package entity

import (
	"github.com/google/uuid"
	"time"
)

type Engine struct {
	ID            uuid.UUID
	Description   string
	Displacement  int64
	NoOfCylinders int64
	CarRange      int64
	HorsePower    int64
	Torque        int64
	EngineType    string
	EmissionClass string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
