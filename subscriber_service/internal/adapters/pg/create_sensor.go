package pg

import (
	"context"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/nov-test/subscriber_service/internal/domain"
	"time"
)

const createSensor = `INSERT INTO sensors (average_id, name, timestamp, value) VALUES ($1, $2, $3, $4) RETURNING *`

func (r *PGSensorRepository) CreateSensor(ctx context.Context, arg []*domain.Sensor, average domain.Average) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PGSensorRepository.CreateAverage")
	defer span.Finish()

	for _, val := range arg {
		if err := r.DB.QueryRowxContext(
			ctx,
			createSensor,
			&average.AverageID,
			&val.Name,
			&val.Timestamp,
			&val.Value,
		).Err(); err != nil {
			return err
		}
	}

	//var sensors []Sensor
	//for _, val := range arg {
	//	sensors = append(sensors, Sensor{
	//		AverageID: average.AverageID,
	//		Name:      val.Name,
	//		Value:     val.Value,
	//		Timestamp: val.Timestamp,
	//	})
	//}
	//_, err := r.DB.NamedExecContext(ctx, createSensor, sensors)
	return nil
}

type Sensor struct {
	AverageID uuid.UUID `json:"average_id" db:"average_id"`
	Name      string    `json:"name" db:"name" `
	Timestamp time.Time `json:"timestamp" db:"timestamp"`
	Value     float64   `json:"value" db:"value"`
}
