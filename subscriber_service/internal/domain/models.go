package domain

import (
	"github.com/google/uuid"
	"time"
)

type Sensor struct {
	Name      string    `json:"name" db:"name" validate:"required,min=6,max=30"`
	Timestamp time.Time `json:"timestamp" db:"timestamp"`
	Value     float64   `json:"value" db:"value" validate:"required,min=8,max=15"`
	//CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Average struct {
	AverageID uuid.UUID `json:"average_id" db:"average_id"`
	Average   float64   `json:"average" db:"average"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
