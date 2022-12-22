//go:generate mockgen -source interface.go -destination mock/interface.go -package mock
package event

import (
	"context"
)

// MessageProcessor processor methods must implement kafka.Worker func method interface
type MessageProcessor interface {
	ProcessMessages(ctx context.Context, workerID int)
}

type Producer interface {
	PublishMessage(ctx context.Context, data []byte, channel string) error
	Close() error
}

type Subscriber interface {
	ConsumeTopic(ctx context.Context, cancel context.CancelFunc, groupID, topic string, poolSize int)
	GetNewKafkaReader(kafkaURL []string, topic, groupID string)
	GetNewKafkaWriter(topic string)
}
