package agent

import (
	"context"
	messagerabbitmq "github.com/rezaAmiri123/nov-test/subscriber_service/internal/port/rabbitmq"
	"github.com/rezaAmiri123/test-microservice/pkg/rabbitmq"
)

func (a *Agent) setupRabbitMQ() error {
	config := rabbitmq.Config{
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
	conn, err := rabbitmq.NewRabbitMQConn(config)
	if err != nil {
		return err
	}
	consumer := messagerabbitmq.NewMessageConsumer(conn, a.logger, a.Application, a.metric)
	worker := consumer.CreateSensorWorker()
	err = consumer.StartConsumer(
		context.Background(),
		a.RabbitMQWorkerPoolSize,
		a.RabbitMQExchange,
		a.RabbitMQQueue,
		a.RabbitMQRoutingKey,
		a.RabbitMQConsumerTag,
		worker,
	)
	return err
}
