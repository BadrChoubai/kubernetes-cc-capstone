package users

import (
	"github.com/badrchoubai/services/internal/services"
	"sync"
)

var _ services.Service = (*UserService)(nil)

// UserService implements Service
type UserService struct {
	ServiceMutex sync.Mutex
}

func NewUsersService() services.Service {
	ser := &UserService{
		ServiceMutex: sync.Mutex{},
	}

	return ser
}
