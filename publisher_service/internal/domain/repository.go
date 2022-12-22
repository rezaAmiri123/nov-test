package domain

import "context"

type Repository interface {
	CreateSensor(ctx context.Context, arg *[]Sensor) error
}
