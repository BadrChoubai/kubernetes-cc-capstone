package users

import (
	"github.com/badrchoubai/services/pkg/service"
	"sync"
)

var _ service.Service = (*UserService)(nil)

type UserService struct {
	ServiceMutex sync.Mutex
}

func NewUsersService() service.Service {
	ser := &UserService{
		ServiceMutex: sync.Mutex{},
	}

	return ser
}
