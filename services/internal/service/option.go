package service

import "github.com/badrchoubai/services/internal/observability/logging"

type Option interface {
	apply(*Service)
}

type optionFunc func(*Service)

func (f optionFunc) apply(service *Service) {
	f(service)
}

func WithLogger(logger *logging.Logger) Option {
	return optionFunc(func(s *Service) {
		s.logger = logger
	})
}
