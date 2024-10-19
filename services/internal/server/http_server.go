package server

import (
	"context"
	"errors"
	"github.com/badrchoubai/services/internal/config"
	"net"
	"net/http"
	"strconv"
	"time"
)

type (
	Server interface {
		Addr() string
		Serve() error
		Shutdown(ctx context.Context) error
	}

	server struct {
		ctx        context.Context
		httpServer *http.Server
	}
)

func (s *server) Addr() string {
	return s.httpServer.Addr
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *server) Serve() error {
	if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func NewServer(ctx context.Context, cfg *config.AppConfig, handler http.Handler) Server {
	const (
		maxTimeout   = 60 * time.Second
		readTimeout  = 5 * time.Second
		writeTimeout = 2 * readTimeout
		idleTimeout  = maxTimeout
	)

	httpserver := &http.Server{
		Addr:         net.JoinHostPort(cfg.HttpHost(), strconv.Itoa(cfg.HttpPort())),
		Handler:      handler,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
	}

	srv := &server{
		ctx:        ctx,
		httpServer: httpserver,
	}

	return srv
}
