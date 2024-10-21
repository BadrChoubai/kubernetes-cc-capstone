package users

import (
	"sync"

	"github.com/badrchoubai/services/internal/services"
)

// Service implements Service
type Service struct {
	Service *services.Service
}

var _ services.ServiceInterface = (*Service)(nil)

func NewUsersService(opts ...services.Option) *Service {
	options := &services.Options{
		Name:         "Service",
		ServiceMutex: &sync.Mutex{},
	}

	for _, opt := range opts {
		opt.Apply(options)
	}

	return &Service{
		Service: &services.Service{
			Name: options.Name,
		},
	}
}
