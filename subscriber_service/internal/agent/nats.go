package agent

import (
	"context"
	nats "github.com/rezaAmiri123/nov-test/pkg/event"
	"github.com/rezaAmiri123/nov-test/pkg/event/natsimpl"
	"github.com/rezaAmiri123/nov-test/subscriber_service/internal/port/subscriber"
)

func (a *Agent) setupNats() error {
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

	messageProcessor := subscriber.NewMessageProcessor(a.logger, a.Application, nc)
	sub, err := messageProcessor.ProcessMessage(nats.CreateSnsorTopic, "", "")
	if err != nil {
		return err
	}
	a.closers = append(a.closers, sub)

	return nil
}
