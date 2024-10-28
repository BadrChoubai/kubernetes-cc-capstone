package main

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/badrchoubai/services/internal/config"
	"github.com/badrchoubai/services/internal/database"
	"github.com/badrchoubai/services/internal/middleware"
	"github.com/badrchoubai/services/internal/observability"
	"github.com/badrchoubai/services/internal/observability/logging"
	"github.com/badrchoubai/services/internal/server"
	"github.com/badrchoubai/services/internal/service"
)

func main() {
	ctx := context.Background()
	cfg := config.NewConfig()

	if err := run(ctx, cfg); err != nil {
		log.Fatalf("%+v\n", err)
	}
}

func run(ctx context.Context, cfg *config.AppConfig) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer cancel()

	logger, err := logging.NewLogger()
	if err != nil {
		return err
	}

	db, err := database.NewDatabase(cfg)
	if err != nil {
		logger.Error("establishing database connection", err)
	}

	svc := service.NewService(
		"auth-service",
		service.WithURL("/api/v1/auth"),
		service.WithLogger(logger),
		service.WithDatabase(db),
	)
	svc.RegisterRoute("", svc.Index())

	srv := server.NewServer(
		cfg,
		server.WithLogger(logger),
		server.WithMiddleware(middleware.Recover(logger)),
		server.WithMiddleware(observability.RequestLoggingMiddleware(logger)),
		server.WithMiddleware(middleware.Heartbeat("/health")),
		server.WithService(svc),
	)

	var serveError error
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		logger.Info("starting HTTP server", zap.String("serverUrl", fmt.Sprintf("http://%s", srv.HttpServer().Addr))) // Log server start
		if err := srv.HttpServer().ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serveError = err
			logger.Error("listening and serving", err) // Log server error
		}
	}()

	// Wait for a cancellation signal
	<-ctx.Done()
	logger.Info("cancellation signal received, shutting down") // Log cancellation

	// Initiate shutdown
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := srv.HttpServer().Shutdown(shutdownCtx); err != nil {
		logger.Error("shutting down http server", err) // Log shutdown error if any
	}

	// Wait for the server goroutine to finish
	wg.Wait()

	if serveError != nil {
		return serveError
	}

	return nil
}
