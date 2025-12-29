package entity

import (
	"github.com/google/uuid"
	"time"
)

type Brand struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
