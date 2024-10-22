package users

import (
	"github.com/badrchoubai/services/internal/encoding"
	logging "github.com/badrchoubai/services/internal/observability/logging/zap"
	"sync"

	"github.com/badrchoubai/services/internal/services"
)

var _ services.IService = (*Service)(nil)

// Service implements ServiceInterface
type Service struct {
	Name           string
	ServiceMutex   *sync.Mutex
	Logger         *logging.Logger
	EncoderDecoder *encoding.ServerEncoderDecoder
}

func NewUsersService(opts ...services.Option) *Service {
	options := &services.Options{
		Name:         "Service",
		ServiceMutex: &sync.Mutex{},
	}

	for _, opt := range opts {
		opt.Apply(options)
	}

	return &Service{
		Name:           options.Name,
		ServiceMutex:   options.ServiceMutex,
		EncoderDecoder: options.EncoderDecoder,
		Logger:         options.Logger,
	}
}
