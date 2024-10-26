package main

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/badrchoubai/services/internal/config"
	"github.com/badrchoubai/services/internal/middleware"
	"github.com/badrchoubai/services/internal/observability"
	"github.com/badrchoubai/services/internal/observability/logging"
	"github.com/badrchoubai/services/internal/server"
	"github.com/badrchoubai/services/internal/service"
	services "github.com/badrchoubai/services/internal/services/auth"
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

	authService := services.NewAuthService(
		"auth-service",
		service.WithLogger(logger),
	)

	srv := server.NewServer(
		cfg,
		authService.Service(),
		server.WithLogger(logger),
		server.WithMiddleware(observability.RequestLoggingMiddleware(logger)),
		server.WithMiddleware(middleware.Heartbeat("/health")),
	)

	httpServer := &http.Server{
		Addr:    net.JoinHostPort(cfg.HTTPHost(), strconv.Itoa(cfg.HTTPPort())),
		Handler: srv.ApplyMiddleware(srv.Handler()),
	}

	var serveError error
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		logger.Info("starting HTTP server", zap.String("serverUrl", fmt.Sprintf("http://%s", httpServer.Addr))) // Log server start
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
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

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		logger.Error("shutting down http server", err) // Log shutdown error if any
	}

	// Wait for the server goroutine to finish
	wg.Wait()

	if serveError != nil {
		return serveError
	}

	return nil
}
