package auth

import (
	"github.com/badrchoubai/services/pkg/service"
	"sync"
)

var _ service.Service = (*AuthService)(nil)

type AuthService struct {
	ServiceMutex sync.Mutex
}

func NewAuthService() service.Service {
	ser := &AuthService{
		ServiceMutex: sync.Mutex{},
	}

	return ser
}
