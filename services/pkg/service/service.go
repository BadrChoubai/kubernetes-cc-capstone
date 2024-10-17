package service

import "net/http"

type Service interface {
	RegisterRouter(router *http.ServeMux)
}
