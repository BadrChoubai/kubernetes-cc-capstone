package server

import (
	"context"
	"errors"
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

func NewServer(ctx context.Context, handler http.Handler) Server {
	const (
		maxTimeout   = 120 * time.Second
		readTimeout  = 5 * time.Second
		writeTimeout = 2 * readTimeout
		idleTimeout  = maxTimeout
	)

	httpserver := &http.Server{
		Addr:         net.JoinHostPort("0.0.0.0", strconv.Itoa(8080)),
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
