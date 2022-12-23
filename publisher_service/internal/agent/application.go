package agent

import (
	"context"
	"github.com/rezaAmiri123/nov-test/pkg/event/natsimpl"
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
	application := &app.Application{
		Commands: app.Commands{
			CreateSensor: commands.NewCreateSensorHandler(a.logger, nats),
		},
	}

	a.Application = application
	return nil
}
