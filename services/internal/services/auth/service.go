package auth

import (
	"github.com/badrchoubai/services/internal/services"
	"sync"
)

// AuthService implements Service
type AuthService struct {
	Service *services.Service
}

var _ services.ServiceInterface = (*AuthService)(nil)

func NewAuthService(opts ...services.Option) *AuthService {
	options := &services.Options{
		Name:         "AuthService",
		ServiceMutex: &sync.Mutex{},
	}

	for _, opt := range opts {
		opt.Apply(options)
	}

	return &AuthService{
		Service: &services.Service{
			Name: options.Name,
		},
	}
}
