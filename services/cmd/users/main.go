package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/badrchoubai/services/internal/config"
	"github.com/badrchoubai/services/internal/database"
	"github.com/badrchoubai/services/internal/encoding"
	logging "github.com/badrchoubai/services/internal/observability/logging/zap"
	"github.com/badrchoubai/services/internal/server"
	"github.com/badrchoubai/services/internal/service"
	"github.com/badrchoubai/services/internal/services/users"
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

	usersService := users.NewUsersService(
		"users-service",
		service.WithLogger(logger),
		service.WithEncoderDecoder(encoding.NewEncoderDecoder(logger)),
		service.WithDbConnection(db),
	)
	fmt.Println(usersService.Name)

	router := server.NewRouter(logger)
	srv := server.NewServer(ctx, cfg, router)

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
		logger.Error("shutting down server", err)
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
