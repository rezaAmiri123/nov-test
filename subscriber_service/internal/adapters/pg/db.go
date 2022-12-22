package pg

import (
	"github.com/jmoiron/sqlx"
	"github.com/rezaAmiri123/nov-test/pkg/logger"
)

func NewPGSensorRepository(db *sqlx.DB, log logger.Logger) *PGSensorRepository {
	return &PGSensorRepository{DB: db, Logger: log}
}

// News Repository
type PGSensorRepository struct {
	DB     *sqlx.DB
	Logger logger.Logger
}
