package app

import (
	"github.com/rezaAmiri123/nov-test/publisher_service/internal/app/commands"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Queries struct {
}

type Commands struct {
	CreateSensor *commands.CreateSensorHandler
}
