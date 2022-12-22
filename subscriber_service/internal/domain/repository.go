package domain

import "context"

type Repository interface {
	CreateSensor(ctx context.Context, arg []*Sensor, average Average) error
	CreateAverage(ctx context.Context, average float64) (Average, error)
}
