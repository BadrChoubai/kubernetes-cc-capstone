package service

import (
	_ "github.com/lib/pq"

	"github.com/badrchoubai/services/internal/database"
	"github.com/badrchoubai/services/internal/observability/logging"
)

type Option interface {
	apply(*Service)
}

type optionFunc func(*Service)

func (f optionFunc) apply(service *Service) {
	f(service)
}

func WithName(name string) Option {
	return optionFunc(func(s *Service) {
		s.name = name
	})
}

func WithDatabase(database *database.Database) Option {
	return optionFunc(func(s *Service) {
		s.database = database
	})
}

func WithLogger(logger *logging.Logger) Option {
	return optionFunc(func(s *Service) {
		s.logger = logger
	})
}

func WithURL(url string) Option {
	return optionFunc(func(s *Service) {
		s.url = url
	})
}

// WithOptions clones the current Service, applies the supplied Options, and
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
