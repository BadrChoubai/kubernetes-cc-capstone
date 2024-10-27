package service

import (
	"net/http"
)

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

// RegisterRoute adds new http.Handler to Service mux
// routes are accessible by "/svc.URL()/path"
func (svc *Service) RegisterRoute(path string, handler http.Handler) {
	if path == "" {
		path = "/"
	}

	svc.mux.Handle(svc.URL()+path, handler)
}

func (svc *Service) Name() string {
	return svc.name
}

func (svc *Service) URL() string {
	return svc.url
}

func (svc *Service) Handler() http.Handler {
	return svc.handler
}

func (svc *Service) Mux() *http.ServeMux {
	return svc.mux
}
