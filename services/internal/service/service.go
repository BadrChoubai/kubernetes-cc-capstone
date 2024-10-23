package service

import (
	"sync"

	"github.com/badrchoubai/services/internal/database"
	"github.com/badrchoubai/services/internal/encoding"
	logging "github.com/badrchoubai/services/internal/observability/logging/zap"
)

var _ IService = (*Service)(nil)

type Service struct {
	Name           string
	ServiceMutex   *sync.Mutex
	Logger         *logging.Logger
	EncoderDecoder *encoding.ServerEncoderDecoder
	DB             *database.Database
}

// IService interface
type IService interface {
	WithOptions(opts ...Option) *Service
}

func NewService(name string, opts ...Option) *Service {
	svc := &Service{
		Name:         name,
		ServiceMutex: &sync.Mutex{},
	}

	return svc.WithOptions(opts...)
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
