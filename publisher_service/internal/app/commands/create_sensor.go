package commands

import (
	"context"
	"github.com/opentracing/opentracing-go"
	messageClient "github.com/rezaAmiri123/nov-test/pkg/event"
	"github.com/rezaAmiri123/nov-test/pkg/logger"
	"github.com/rezaAmiri123/nov-test/publisher_service/internal/domain"
	"github.com/rezaAmiri123/nov-test/publisher_service/proto/event"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CreateSensorHandler struct {
	logger logger.Logger
	//repo            domain.Repository
	messageProducer messageClient.Producer
}

func NewCreateSensorHandler(logger logger.Logger, producer messageClient.Producer) *CreateSensorHandler {
	//if repo == nil {
	//	panic("repo is nil")
	//}
	return &CreateSensorHandler{logger: logger, messageProducer: producer}
}

func (h CreateSensorHandler) Handle(ctx context.Context, arg []domain.Sensor) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateSensorHandler.Handle")
	defer span.Finish()

	var sensors []*event.Sensor
	for _, val := range arg {
		sensors = append(sensors, &event.Sensor{
			Name:      val.Name,
			Timestamp: timestamppb.New(val.Timestamp),
			Value:     val.Value,
		})
	}
	req := &event.CreateSensor{
		Sensors: sensors,
	}

	message, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	//// TODO we need to change topic name
	err = h.messageProducer.PublishMessage(ctx, message, messageClient.CreateSnsorTopic)
	return err
}
