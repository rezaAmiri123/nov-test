package agent

import (
	"context"
	"github.com/rezaAmiri123/nov-test/pkg/event/natsimpl"
	"github.com/rezaAmiri123/nov-test/pkg/rabbitmq"
	"github.com/rezaAmiri123/nov-test/pkg/rabbitmq/publisher"
	"github.com/rezaAmiri123/nov-test/publisher_service/internal/app"
	"github.com/rezaAmiri123/nov-test/publisher_service/internal/app/commands"
)

func (a *Agent) setupApplication() error {
	config := natsimpl.Config{
		Url:          a.NatsUrl,
		ClusterID:    a.NatsClusterID,
		ClientID:     a.NatsClientID,
		PingInterval: a.NatsPingInterval,
		PingMaxOut:   a.NatsPingMaxOut,
	}
	nc, err := natsimpl.NewClientConn(context.Background(), a.logger, config)
	if err != nil {
		return err
	}
	nats := natsimpl.NewNats(nc, a.logger)
	a.closers = append(a.closers, nats)
	//nats := &natsimpl.Nats{}

	rabbitConfig := rabbitmq.Config{
		User:           a.RabbitMQUser,
		Password:       a.RabbitMQPassword,
		Host:           a.RabbitMQHost,
		Port:           a.RabbitMQPort,
		Exchange:       a.RabbitMQExchange,
		Queue:          a.RabbitMQQueue,
		RoutingKey:     a.RabbitMQRoutingKey,
		ConsumerTag:    a.RabbitMQConsumerTag,
		WorkerPoolSize: a.RabbitMQWorkerPoolSize,
	}
	rabbitPublisher, err := publisher.NewPublisher(rabbitConfig, a.logger)

	if err != nil {
		return err
	}
	application := &app.Application{
		Commands: app.Commands{
			CreateSensor: commands.NewCreateSensorHandler(a.logger, nats, rabbitPublisher),
		},
	}

	a.Application = application
	return nil
}
