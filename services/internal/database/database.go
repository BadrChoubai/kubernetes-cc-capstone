package database

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"

	"github.com/badrchoubai/services/internal/config"
	logging "github.com/badrchoubai/services/internal/observability/logging/zap"
)

var _ IDatabase = (*Database)(nil)

type (
	IDatabase interface {
		DB() *sql.DB
		Close() error
	}

	Database struct {
		context context.Context
		db      *sql.DB
		logger  *logging.Logger
	}
)

func NewConnection(ctx context.Context, logger *logging.Logger, cfg *config.DatabaseSettings) (*Database, error) {
	logger.Info("opening database connection")

	db, err := sql.Open("postgres", cfg.DbConnectionString)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.DbMaxOpenConns)
	db.SetMaxIdleConns(cfg.DbMaxIdleConns)
	db.SetConnMaxIdleTime(cfg.DbConnMaxLifetime)
	db.SetConnMaxLifetime(cfg.DbConnMaxLifetime)

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return &Database{
		context: ctx,
		db:      db,
	}, nil

}

func (d *Database) DB() *sql.DB {
	return d.db
}

func (d *Database) Close() error {
	if err := d.db.Close(); err != nil {
		d.logger.Error("closing database connection", err)
	}
	return nil
}
