// Package auth provides services and HTTP handlers for handling authentication and authorization functionality.
// This package includes the core logic for user login, registration, token management, and permissions checks,
// leveraging configuration, database, and logging services.
package auth

import (
	"context"
	"go.uber.org/zap"

	"github.com/badrchoubai/services/internal/config"
	"github.com/badrchoubai/services/internal/database"
	"github.com/badrchoubai/services/internal/service"
)

// NewAuthService creates a new service for handling Authentication and Authorization
func NewAuthService(ctx context.Context, cfg config.AppConfig, logger *zap.Logger) (*service.Service, error) {
	db, err := database.NewDatabase(ctx, cfg)
	if err != nil {
		logger.Error(
			"establishing database connection",
			zap.Error(err),
		)

		return nil, err
	}

	svc, err := service.NewService(
		ctx,
		"auth-v1",
		service.WithLogger(logger),
		service.WithDatabase(db),
	)

	if svc != nil {
		addRoutes(svc)

		return svc, nil
	}

	return nil, err
}
