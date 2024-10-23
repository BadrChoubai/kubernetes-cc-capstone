package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/badrchoubai/services/internal/config"
	"github.com/badrchoubai/services/internal/database"
	"github.com/badrchoubai/services/internal/encoding"
	"github.com/badrchoubai/services/internal/middleware"
	"github.com/badrchoubai/services/internal/observability"
	"github.com/badrchoubai/services/internal/observability/logging/zap"
	"github.com/badrchoubai/services/internal/router"
	"github.com/badrchoubai/services/internal/server"
	"github.com/badrchoubai/services/internal/service"
	services "github.com/badrchoubai/services/internal/services/auth"
)

func run(ctx context.Context, cfg *config.AppConfig) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer cancel()

	logger, err := logging.NewLogger()
	if err != nil {
		return err
	}

	db, err := database.NewConnection(ctx, logger, cfg.DbConn())
	if err != nil {
		return err
	}

	authService := services.NewAuthService(
		"auth-service",
		service.WithLogger(logger),
		service.WithEncoderDecoder(encoding.NewEncoderDecoder(logger)),
		service.WithDbConnection(db),
	)

	httpRouter := router.NewRouter(
		fmt.Sprintf("%s-router", authService.Name),
		router.WithLogger(logger),
		router.WithService(authService),
		router.WithMiddleware(observability.RequestLoggingMiddleware(logger)),
		router.WithMiddleware(middleware.Heartbeat("/health")),
	)

	srv := server.NewServer(ctx, cfg, httpRouter)

	var serveError error

	go func() {
		if err := srv.Serve(); err != nil {
			serveError = err
		}
	}()

	logger.Info("server started",
		zap.String("service", authService.Name))

	if serveError != nil {
		return serveError
	}

	// Wait for shutdown signal
	<-ctx.Done()
	logger.Info("shutdown signal received")

	// Allow some time for graceful shutdown
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Error("server shutdown failed", err)
		return err
	}

	return nil
}

func main() {
	rootCtx := context.Background()
	cfg := config.NewConfig()

	if err := run(
		rootCtx,
		cfg,
	); err != nil {
		log.Fatalf("error: %+v", err)
	}
}
