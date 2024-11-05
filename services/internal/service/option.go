package service

import (
	_ "github.com/lib/pq" // Register Postgres driver for database access
	"go.uber.org/zap"

	"github.com/badrchoubai/services/internal/database"
)

// Option represents a configuration option for a Service.
// Each Option modifies a Service instance when applied.
type Option interface {
	apply(*Service)
}

type optionFunc func(*Service)

func (f optionFunc) apply(service *Service) {
	f(service)
}

// WithDatabase returns an Option that sets the database for a Service instance.
// It allows customization of the Service's database during initialization.
func WithDatabase(db *database.Database) Option {
	return optionFunc(func(s *Service) {
		s.database = db
	})
}

// WithLogger returns an Option that sets the logger for a Service instance.
// It allows customization of the Service's logging behavior during initialization.
func WithLogger(logger *zap.Logger) Option {
	return optionFunc(func(s *Service) {
		s.logger = logger
	})
}

// WithOptions clones the current Service, applies the supplied list of Option, and
// returns the resulting Service. It's safe to use concurrently.
func (svc *Service) WithOptions(opts ...Option) *Service {
	s := svc.clone()
	for _, opt := range opts {
		opt.apply(s)
	}
	return s
}

func (svc *Service) clone() *Service {
	clone := *svc
	return &clone
}
