package main

import (
	"context"
	"github.com/badrchoubai/services/internal/encoding"
	"github.com/badrchoubai/services/internal/services"
	"github.com/badrchoubai/services/internal/services/users"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/badrchoubai/services/internal/config"
	"github.com/badrchoubai/services/internal/observability/logging/zap"
	"github.com/badrchoubai/services/internal/server"
)

func run(ctx context.Context, cfg *config.AppConfig) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer cancel()

	logger := zap.NewZapLogger(cfg.LogLevel())

	authService := users.NewUsersService(
		services.WithName("UserService"),
		services.WithLogger(&logger),
		services.WithEncoderDecoder(encoding.NewEncoderDecoder(logger)),
	)

	router := server.NewRouter(logger, authService)
	srv := server.NewServer(ctx, cfg, logger, router)

	var serveError error

	go func() {
		if err := srv.Serve(); err != nil {
			serveError = err
		}
	}()

	logger.Info("server started")

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
		logger.Error(err, "server shutdown failed")
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
