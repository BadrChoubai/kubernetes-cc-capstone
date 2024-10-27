package service

import "net/http"

func NewService(name string, options ...Option) *Service {
	mux := http.NewServeMux()
	svc := &Service{
		name:    name,
		mux:     mux,
		handler: mux,
	}

	svc = svc.WithOptions(options...)
	return svc
}

func (svc *Service) RegisterRoute(path string, handler http.Handler) {
	svc.mux.Handle(path, handler)
}

func (svc *Service) Name() string {
	return svc.name
}

func (svc *Service) Handler() http.Handler {
	return svc.handler
}
