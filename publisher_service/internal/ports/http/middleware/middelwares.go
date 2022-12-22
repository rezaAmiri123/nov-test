package middleware

import (
	"github.com/rezaAmiri123/nov-test/pkg/logger"
	"github.com/rezaAmiri123/nov-test/publisher_service/internal/app"
)

// Middleware manager
type MiddlewareManager struct {
	logger logger.Logger
	app    *app.Application
	//origins    []string
}

// Middleware manager constructor
func NewMiddlewareManager(logger logger.Logger, app *app.Application) *MiddlewareManager {
	return &MiddlewareManager{logger: logger, app: app}
}
