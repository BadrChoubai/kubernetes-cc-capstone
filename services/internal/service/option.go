package service

import (
	"github.com/badrchoubai/services/internal/database"
	"github.com/badrchoubai/services/internal/encoding"
	logging "github.com/badrchoubai/services/internal/observability/logging/zap"
)

type Option interface {
	apply(*Service)
}

type optionFunc func(*Service)

func (f optionFunc) apply(service *Service) {
	f(service)
}

func WithDbConnection(conn *database.Database) Option {
	return optionFunc(func(service *Service) {
		service.DB = conn
	})
}

func WithEncoderDecoder(edc *encoding.ServerEncoderDecoder) Option {
	return optionFunc(func(service *Service) {
		service.EncoderDecoder = edc
	})
}

func WithLogger(logger *logging.Logger) Option {
	return optionFunc(func(service *Service) {
		service.Logger = logger
	})
}
