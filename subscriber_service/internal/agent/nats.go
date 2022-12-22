package agent

import (
	"context"
	"github.com/rezaAmiri123/nov-test/pkg/event/natsimpl"
	"github.com/rezaAmiri123/nov-test/subscriber_service/internal/port/subscriber"
)

func (a *Agent) setupNats() error {
	//ctx, cancel := context.WithCancel(context.Background())
	//a.closers = append(a.closers, closer{cancel: cancel})
	//kafkaCfg := kafka.Config{
	//	Kafka: kafkaClient.Config{
	//		Brokers:    a.KafkaBrokers,
	//		GroupID:    a.KafkaGroupID,
	//		InitTopics: a.KafkaInitTopics,
	//	},
	//	KafkaTopics: kafka.KafkaTopics{
	//		UserCreate: kafkaClient.TopicConfig{
	//			TopicName: kafkaClient.CreateUserTopic,
	//		},
	//	},
	//}
	nc, err := natsimpl.NewClientConn(context.Background(), a.logger)
	if err != nil {
		return err
	}

	messageMessageProcessor := subscriber.NewMessageProcessor(a.logger, a.Application, nc)
	go func() {
		sub := messageMessageProcessor.ProcessMessage("subject", "qgroup", "durable")
		a.closers = append(a.closers, sub)
	}()

	return nil
}
