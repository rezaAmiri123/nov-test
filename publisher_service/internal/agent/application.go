package agent

import (
	"context"
	"github.com/rezaAmiri123/nov-test/pkg/event/natsimpl"
	"github.com/rezaAmiri123/nov-test/publisher_service/internal/app"
	"github.com/rezaAmiri123/nov-test/publisher_service/internal/app/commands"
)

func (a *Agent) setupApplication() error {
	nc, err := natsimpl.NewClientConn(context.Background(), a.logger)
	if err != nil {
		return err
	}
	nats := natsimpl.NewNats(nc, a.logger)
	a.closers = append(a.closers, nats)

	application := &app.Application{
		Commands: app.Commands{
			CreateSensor: commands.NewCreateSensorHandler(a.logger, nats),
		},
	}

	a.Application = application
	return nil
}
