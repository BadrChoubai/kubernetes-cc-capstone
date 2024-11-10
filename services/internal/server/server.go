// Package server provides an HTTP server implementation with configurable  options for middleware, routing, and service
// registration. It supports starting and shutting down the server gracefully, applying middleware, and registering
// service-specific routes.
package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"go.uber.org/zap"
	"net"
	"net/http"
	"strconv"

	"github.com/badrchoubai/services/internal/config"
	"github.com/badrchoubai/services/internal/service"
)

var _ HTTPServer = (*Server)(nil)

// Server represents an HTTP server with configuration, middleware, and service management.
type Server struct {
	config      config.AppConfig
	httpServer  *http.Server
	logger      *zap.Logger
	mux         *http.ServeMux
	middlewares []func(http.Handler) http.Handler
	services    []*service.Service
}

// HTTPServer defines the interface for managing HTTP servers, allowing for middleware
// application and option configuration.
type HTTPServer interface {
	ApplyMiddleware(http.Handler) http.Handler
	WithOptions(opts ...Option) *Server

	clone() *Server
}

// NewServer initializes a new Server instance with the provided configuration and options.
// It sets up the HTTP server, registers services with the router, and applies middleware to handle requests.kkk
func NewServer(cfg config.AppConfig, opts ...Option) *Server {
	server := &Server{
		config:     cfg,
		mux:        http.NewServeMux(),
		httpServer: createStdLibHTTPServer(cfg),
	}
	server = server.WithOptions(opts...)

	for _, svc := range server.services {
		server.mux.Handle(svc.Path()+"/", http.StripPrefix(svc.Path(), svc.Mux())) // Register with service Path prefix
	}

	server.httpServer.Handler = server.ApplyMiddleware(server.mux)

	return server
}

func createStdLibHTTPServer(cfg config.AppConfig) *http.Server {
	return &http.Server{
		Addr:         net.JoinHostPort(cfg.HTTPHost(), strconv.Itoa(cfg.HTTPPort())),
		IdleTimeout:  cfg.IdleTimeout(),
		ReadTimeout:  cfg.WriteTimeout(),
		WriteTimeout: cfg.WriteTimeout(),
		TLSConfig: &tls.Config{
			CipherSuites: []uint16{
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			},
			CurvePreferences: []tls.CurveID{
				tls.CurveP256,
				tls.X25519,
			},
			MinVersion: tls.VersionTLS12,
		},
	}
}

// ApplyMiddleware takes a http.Handler and applies all registered middleware functions
// to it in reverse order. This ensures that the first middleware added becomes the outermost
// handler in the chain, allowing it to wrap around the subsequent middleware and the original
// handler.
func (s *Server) ApplyMiddleware(handler http.Handler) http.Handler {
	for i := len(s.middlewares) - 1; i >= 0; i-- {
		handler = s.middlewares[i](handler)
	}

	return handler
}

// WithOptions clones the current Router, applies the supplied Options, and
// returns the resulting Router. It's safe to use concurrently.
func (s *Server) WithOptions(opts ...Option) *Server {
	c := s.clone()
	for _, opt := range opts {
		opt.apply(c)
	}

	return c
}

func (s *Server) clone() *Server {
	clone := *s
	return &clone
}

// Serve starts the HTTP server and listens for incoming requests.
// It logs the server's URL and returns any error encountered while starting the server.
func (s *Server) Serve() error {
	s.logger.Info("starting server")

	if s.config.HTTPSCertificateFilePath() != "" && s.config.HTTPSCertificateKeyFilePath() != "" {
		s.logger.Info(
			"serving HTTPS",
			zap.String("serverUrl", fmt.Sprintf("https://%s", s.httpServer.Addr)),
		)

		if err := s.httpServer.ListenAndServeTLS(
			s.config.HTTPSCertificateFilePath(),
			s.config.HTTPSCertificateKeyFilePath(),
		); err != nil {
			return err
		}
	} else {
		s.logger.Info(
			"serving HTTP",
			zap.String("serverUrl", fmt.Sprintf("http://%s", s.httpServer.Addr)),
		)

		if err := s.httpServer.ListenAndServe(); err != nil {
			return err
		}
	}

	return nil
}

// Shutdown gracefully shuts down the HTTP server, allowing existing connections to finish.
// It logs the shutdown event and returns any error encountered during the shutdown process.
func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Info("HTTP server shut down")

	err := s.httpServer.Shutdown(ctx)
	if err != nil {
		return err
	}

	return nil
}
