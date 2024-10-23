package users

import "github.com/badrchoubai/services/internal/service"

func NewUsersService(name string, opts ...service.Option) *service.Service {
	return service.NewService(name, opts...)
}
