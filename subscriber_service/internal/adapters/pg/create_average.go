package pg

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/nov-test/subscriber_service/internal/domain"
)

const createAverage = `INSERT INTO averages (average) VALUES ($1) RETURNING *`

func (r *PGSensorRepository) CreateAverage(ctx context.Context, average float64) (domain.Average, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PGSensorRepository.CreateAverage")
	defer span.Finish()

	var a domain.Average
	if err := r.DB.QueryRowxContext(
		ctx,
		createAverage,
		average,
	).StructScan(&a); err != nil {
		return domain.Average{}, fmt.Errorf("postgres connot create user: %w", err)
	}

	return a, nil
}
