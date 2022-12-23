package commands_test

import (
	"context"
	"fmt"
	"github.com/rezaAmiri123/microservice/pkg/db/postgres"
	"github.com/rezaAmiri123/nov-test/pkg/logger/applogger"
	"github.com/rezaAmiri123/nov-test/subscriber_service/internal/adapters/pg"
	"github.com/rezaAmiri123/nov-test/subscriber_service/internal/app/commands"
	"github.com/rezaAmiri123/nov-test/subscriber_service/internal/domain"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCreateSensorHandler_Handle(t *testing.T) {
	logger := applogger.NewAppLogger(applogger.Config{})
	dbConn, err := postgres.NewPsqlDB(postgres.Config{
		PGDriver:   "pgx",
		PGHost:     "nov_postgesql",
		PGPort:     "5432",
		PGUser:     "postgres",
		PGDBName:   "microservice",
		PGPassword: "postgres",
	})
	require.NoError(t, err)
	repo := pg.NewPGSensorRepository(dbConn, logger)
	createSensor := commands.NewCreateSensorHandler(repo, logger)
	var arg []*domain.Sensor
	for i := 0; i < 5; i++ {
		arg = append(arg, &domain.Sensor{
			Name:      fmt.Sprintf("sensor %d", i),
			Value:     float64(4564.9383) + float64(i),
			Timestamp: time.Now().Add(time.Second),
		})
	}
	err = createSensor.Handle(context.Background(), arg)
	require.NoError(t, err)

}