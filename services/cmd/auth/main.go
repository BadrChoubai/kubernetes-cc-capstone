package main

import (
	"context"
	"github.com/badrchoubai/services/internal/config"
	"github.com/badrchoubai/services/internal/services/auth"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/badrchoubai/services/internal/server"
)

func run(ctx context.Context, cfg *config.AppConfig) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer cancel()

	infoLog := log.New(os.Stdout, "", log.LstdFlags)
	errLog := log.New(os.Stderr, "", log.LstdFlags)

	infoLog.Printf("%+v\n", cfg)

	service := auth.NewAuthService()
	router := server.NewRouter(service)
	srv := server.NewServer(ctx, cfg.HttpHost(), cfg.HttpPort(), router)

	var serveError error

	go func() {
		if err := srv.Serve(); err != nil {
			serveError = err
		}
	}()

	infoLog.Printf("http://%s/health", srv.Addr())

	if serveError != nil {
		return serveError
	}

	// Wait for shutdown signal
	<-ctx.Done()
	infoLog.Print("shutdown signal received")

	// Allow some time for graceful shutdown
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		errLog.Print(err, "server shutdown failed")
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
