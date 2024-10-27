package service

import (
	"fmt"
	"net/http"
)

func (svc *Service) Index() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%s Index", svc.name)
		})
}

func (svc *Service) Health() http.Handler {
	apiClient := &http.Client{}

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			resp, err := apiClient.Get("http://localhost:8080/health")
			if err != nil {
				svc.logger.Error("reaching healthcheck", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode == http.StatusOK {
				svc.logger.Info("health check passed")
			} else {
				svc.logger.Info("health check failed")
			}
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%d", resp.StatusCode)
		})
}
