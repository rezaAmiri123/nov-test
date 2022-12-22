package agent

import (
	"github.com/rezaAmiri123/nov-test/pkg/logger/applogger"
)

func (a *Agent) setupLogger() error {
	appLogger := applogger.NewAppLogger(applogger.Config{
		LogLevel:   a.LogLevel,
		LogDevMode: a.LogDevMode,
		LogEncoder: a.LogEncoder,
	})
	appLogger.InitLogger()
	appLogger.WithName("ServiceMessage")
	a.logger = appLogger
	return nil
}
