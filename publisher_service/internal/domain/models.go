package domain

import "time"

type Sensor struct {
	Name      string    `json:"name" db:"name" validate:"required,min=6,max=30"`
	Timestamp time.Time `json:"timestamp" db:"timestamp"`
	Value     float64   `json:"value" db:"value" validate:"required,min=8,max=15"`
	//CreatedAt time.Time `json:"created_at" db:"created_at"`
}
