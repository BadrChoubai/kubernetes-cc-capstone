package users

import (
	"github.com/badrchoubai/services/internal/services"
	"sync"
)

// UserService implements Service
type UserService struct {
	Service *services.Service
}

var _ services.ServiceInterface = (*UserService)(nil)

func NewUsersService(opts ...services.Option) *UserService {
	options := &services.Options{
		Name:         "UserService",
		ServiceMutex: &sync.Mutex{},
	}

	for _, opt := range opts {
		opt.Apply(options)
	}

	return &UserService{
		Service: &services.Service{
			Name: options.Name,
		},
	}
}
