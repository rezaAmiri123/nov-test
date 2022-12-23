package rabbitmq

import (
	"context"
	"github.com/opentracing/opentracing-go"
	eventMessages "github.com/rezaAmiri123/nov-test/publisher_service/proto/event"
	"github.com/rezaAmiri123/nov-test/subscriber_service/internal/domain"
	"github.com/rezaAmiri123/test-microservice/pkg/rabbitmq"
	"github.com/streadway/amqp"
	"google.golang.org/protobuf/proto"
)

func (c *MessageConsumer) CreateSensorWorker() rabbitmq.Worker {
	return func(ctx context.Context, messages <-chan amqp.Delivery) {
		for delivery := range messages {
			span, ctx := opentracing.StartSpanFromContext(ctx, "MessageConsumer.worker")
			c.metric.CreateSensorAsyncRequests.Inc()
			c.logger.Infof("processDeliveries deliveryTag% v", delivery.DeliveryTag)

			var m eventMessages.CreateSensor
			if err := proto.Unmarshal(delivery.Body, &m); err != nil {
				//s.log.WarnMsg("proto.Unmarshal", err)
				//s.commitErrMessage(ctx, r, m)
				return
			}

			var sensors []*domain.Sensor
			for _, val := range m.Sensors {
				sensors = append(sensors, &domain.Sensor{
					Name:      val.GetName(),
					Timestamp: val.GetTimestamp().AsTime(),
					Value:     val.GetValue(),
				})
			}

			err := c.app.Commands.CreateSensor.Handle(ctx, sensors)
			if err != nil {
				c.logger.Errorf("error create email consumer", err)
			}

			if err != nil {
				if err := delivery.Reject(false); err != nil {
					c.logger.Errorf("Err delivery.Reject: %v", err)
				}
				c.logger.Errorf("Failed to process delivery: %v", err)
				c.metric.ErrorAsyncRequests.Inc()
			} else {
				err = delivery.Ack(false)
				if err != nil {
					c.logger.Errorf("Failed to acknowledge delivery: %v", err)
				}
				//c.logger.Info("email created")
				c.metric.SuccessAsyncRequests.Inc()
			}
			span.Finish()
		}

		c.logger.Info("Deliveries channel closed")
	}
}
