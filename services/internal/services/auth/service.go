package auth

import (
	"github.com/badrchoubai/services/internal/services"
	"sync"
)

var _ services.Service = (*AuthService)(nil)

// AuthService implements Service
type AuthService struct {
	ServiceMutex sync.Mutex
}

func NewAuthService() services.Service {
	ser := &AuthService{
		ServiceMutex: sync.Mutex{},
	}

	return ser
}
