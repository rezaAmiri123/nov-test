package commands

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/rezaAmiri123/nov-test/pkg/logger"
	"github.com/rezaAmiri123/nov-test/subscriber_service/internal/domain"
)

type CreateSensorHandler struct {
	logger logger.Logger
	repo   domain.Repository
}

func NewCreateSensorHandler(repo domain.Repository, logger logger.Logger) *CreateSensorHandler {
	if repo == nil {
		panic("repo is nil")
	}
	return &CreateSensorHandler{repo: repo, logger: logger}
}

func (h CreateSensorHandler) Handle(ctx context.Context, arg []*domain.Sensor) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateSensorHandler.Handle")
	defer span.Finish()

	var average float64
	for _, val := range arg {
		average += val.Value
	}
	average = average / float64(len(arg))
	avg, err := h.repo.CreateAverage(ctx, average)
	if err != nil {
		return err
	}
	err = h.repo.CreateSensor(ctx, arg, avg)
	return err
}
