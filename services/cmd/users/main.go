package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/badrchoubai/services/internal/server"
)

func run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer cancel()

	infoLog := log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds)
	errLog := log.New(os.Stderr, "", log.LstdFlags|log.Lmicroseconds)
	router := server.NewRouter()
	srv := server.NewServer(ctx, router)

	var serveError error

	go func() {
		if err := srv.Serve(); err != nil {
			serveError = err
		}
	}()

	infoLog.Printf("https://%s", srv.Addr())

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

	if err := run(
		rootCtx,
	); err != nil {
		log.Fatalf("error: %+v", err)
	}
}
