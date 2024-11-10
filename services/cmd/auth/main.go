/*
Package main is the entry point for the application, responsible for initializing
and running the HTTP server that handles incoming requests.

This package performs the following key tasks:

 1. **Configuration and Logging**: It sets up the application configuration and initializes
    a logger for observability.

 2. **Service Initialization**: It creates an instance of the authentication service
    by invoking the NewAuthService function, passing in the necessary context,
    configuration, and logger.

 3. **Server Setup**: The server is initialized with middleware for logging, recovery,
    CORS handling, rate limiting, and health checks. The service is registered with
    the server to handle specific routes.

 4. **Graceful Shutdown**: The application listens for interrupt signals (e.g., SIGINT, SIGTERM)
    to initiate a graceful shutdown of the server, allowing ongoing requests to complete
    before shutting down. This includes handling any errors that may occur during
    the server's operation.

This package serves as the backbone of the application, coordinating the
components required to start and manage the server lifecycle.
*/
package main

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/badrchoubai/services/internal/config"
	"github.com/badrchoubai/services/internal/middleware"
	"github.com/badrchoubai/services/internal/server"
	"github.com/badrchoubai/services/internal/services/auth"
)

func main() {
	ctx := context.Background()

	if err := run(ctx); err != nil {
		log.Fatalf("%+v\n", err)
	}
}

func run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer cancel()

	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}

	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}

	service, err := auth.NewAuthService(
		ctx,
		cfg,
		logger,
	)

	if err != nil {
		return err
	}

	srv := server.NewServer(
		cfg,
		server.WithLogger(logger),
		server.WithMiddleware(
			middleware.RequestLogging(logger),
			middleware.Recover(logger),
			middleware.Cors(cfg.CORSEnabled(), cfg.CORSTrustedOrigins()),
			middleware.RateLimit(cfg.RateLimitEnabled(), cfg.Burst(), cfg.RPS()),
			middleware.Heartbeat("/health"),
		),
		server.WithService(service),
	)

	var serveError error
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		if err := srv.Serve(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serveError = err
			logger.Error("listening and serving", zap.Error(err)) // Log server error
		}
	}()

	// Wait for a cancellation signal
	<-ctx.Done()
	logger.Info("cancellation signal received, shutting down") // Log cancellation

	// Initiate shutdown
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Error("shutting down http server", zap.Error(err)) // Log shutdown error if any
	}

	// Wait for the server goroutine to finish
	wg.Wait()

	if serveError != nil {
		return serveError
	}

	return nil
}
