package auth

import (
	"context"

	"github.com/badrchoubai/services/internal/config"
	"github.com/badrchoubai/services/internal/database"
	"github.com/badrchoubai/services/internal/observability/logging"
	"github.com/badrchoubai/services/internal/service"
)

// NewAuthService creates a new service for handling Authentication and Authorization
func NewAuthService(ctx context.Context, cfg *config.AppConfig, logger *logging.Logger) (*service.Service, error) {
	db, err := database.NewDatabase(cfg)
	if err != nil {
		logger.Error("establishing database connection", err)
		return nil, err
	}

	svc, err := service.NewService(
		ctx,
		"auth-service-v1",
		service.WithLogger(logger),
		service.WithDatabase(db),
	)

	if svc != nil {
		addRoutes(svc)

		return svc, nil
	}

	return nil, err
}
