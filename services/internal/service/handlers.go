package service

import (
	"fmt"
	"net/http"
)

func (svc *Service) Index() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf("%s Index", svc.name)))
		})
}
