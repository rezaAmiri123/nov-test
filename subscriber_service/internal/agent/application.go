package agent

import (
	"fmt"

	"github.com/rezaAmiri123/microservice/pkg/db/postgres"
	"github.com/rezaAmiri123/nov-test/subscriber_service/internal/adapters/pg"
	"github.com/rezaAmiri123/nov-test/subscriber_service/internal/app"
	"github.com/rezaAmiri123/nov-test/subscriber_service/internal/app/commands"
)

func (a *Agent) setupApplication() error {
	dbConn, err := postgres.NewPsqlDB(postgres.Config{
		PGDriver:   a.PGDriver,
		PGHost:     a.PGHost,
		PGPort:     a.PGPort,
		PGUser:     a.PGUser,
		PGDBName:   a.PGDBName,
		PGPassword: a.PGPassword,
	})
	if err != nil {
		return fmt.Errorf("cannot load db: %w", err)
	}

	//repo, err := adapters.NewGORMArticleRepository(a.DBConfig)
	repo := pg.NewPGSensorRepository(dbConn, a.logger)

	application := &app.Application{
		Commands: app.Commands{
			CreateSensor: commands.NewCreateSensorHandler(repo, a.logger),
		},
		Queries: app.Queries{},
	}
	a.Application = application
	return nil
}
