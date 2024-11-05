// Package database handles database
package database

import (
	"context"
	"database/sql"

	"github.com/badrchoubai/services/internal/config"
)

var _ IDatabase = (*Database)(nil)

// Database struct
type Database struct {
	db *sql.DB
}

// IDatabase interface
type IDatabase interface {
	Close() error
	DB() *sql.DB
	Ping(ctx context.Context) error
}

// DB returns pointer reference of sql.DB on Database
func (d *Database) DB() *sql.DB {
	return d.db
}

// Close calls *sql.DB Close() returning an error
func (d *Database) Close() error {
	if err := d.db.Close(); err != nil {
		return err
	}
	return nil
}

// Ping calls *sql.DB PingContext
func (d *Database) Ping(ctx context.Context) error {
	return d.db.PingContext(ctx)
}

// NewDatabase creates database connection returning a new instance of Database or error
func NewDatabase(cfg *config.AppConfig) (*Database, error) {
	db, err := connect(cfg)
	if err != nil {
		return nil, err
	}

	return &Database{
		db: db,
	}, nil
}

func connect(cfg *config.AppConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DbConnectionString())
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(cfg.MaxOpenConns())
	db.SetMaxIdleConns(cfg.MaxIdleConns())
	db.SetConnMaxIdleTime(cfg.ConnMaxIdleTime())
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime())

	return db, nil
}
