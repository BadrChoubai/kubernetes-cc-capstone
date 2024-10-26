package services

import "github.com/badrchoubai/services/internal/service"

var _ IAuthService = (*AuthService)(nil)

type AuthService struct {
	service  *service.Service
	handlers []service.Handler
}

type IAuthService interface {
	Service() *service.Service
}

func NewAuthService(name string, opts ...service.Option) *AuthService {
	mainSvc := service.NewService(name, opts...)

	svc := &AuthService{
		service: mainSvc,
		handlers: []service.Handler{
			{
				Path:    "/",
				Handler: PingHandler,
			},
		},
	}

	svc.Service().RegisterRoutes(svc.handlers)
	return svc
}

func (p *AuthService) Service() *service.Service {
	return p.service
}
