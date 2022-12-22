package impl

import (
	"context"
	"github.com/nats-io/stan.go"
	"github.com/rezaAmiri123/nov-test/pkg/logger"
)

type Nats struct {
	Sc     stan.Conn
	Logger logger.Logger
}

func (n *Nats) PublishMessage(ctx context.Context, data []byte, channel string) error {
	ach := func(s string, err2 error) {}
	_, err := n.Sc.PublishAsync(channel, data, ach)
	if err != nil {
		n.Logger.Errorf("Error during async publish: %v\n", err)
	}
	return err
}

func (n *Nats) Close() error {
	return n.Sc.Close()
}
