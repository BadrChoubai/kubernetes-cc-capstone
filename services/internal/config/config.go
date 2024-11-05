// Package config handles application configuration.  It provides structures and methods for loading and managing
// application settings from environment variables and defaults. This package encapsulates the configuration for
// various application components, such as server settings, database settings, and rate limiting.
package config

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var _ Config = (*AppConfig)(nil)

type (
	// AppConfig holds the overall application configuration, including
	// environment, HTTP settings, logging level, and other nested settings.
	AppConfig struct {
		environment string
		httpHost    string
		httpPort    int
		logLevel    int

		corsSettings        CORSSettings
		databaseSettings    DatabaseSettings
		rateLimiterSettings RateLimiterSettings
		serverSettings      ServerSettings
	}

	// CORSSettings defines the settings for Cross-Origin Resource Sharing.
	CORSSettings struct {
		corsEnabled    bool
		trustedOrigins []string
	}

	// DatabaseSettings holds configuration for database connections.
	DatabaseSettings struct {
		connMaxIdleTime    time.Duration
		connMaxLifetime    time.Duration
		dbConnectionString string
		maxOpenConns       int
		maxIdleConns       int
	}

	// RateLimiterSettings configures the rate limiting behavior.
	RateLimiterSettings struct {
		burst   int
		enabled bool
		rps     int
	}

	// ServerSettings defines timeout settings for the server.
	ServerSettings struct {
		idleTimeout  time.Duration
		maxTimeout   time.Duration
		readTimeout  time.Duration
		writeTimeout time.Duration
	}

	// Config interface outlines the methods required for retrieving configuration values.
	Config interface {
		Environment() string
		HTTPHost() string
		HTTPPort() int
		LogLevel() int

		CORSEnabled() bool
		CORSTrustedOrigins() []string

		DbConnectionString() string
		MaxOpenConns() int
		MaxIdleConns() int
		ConnMaxIdleTime() time.Duration
		ConnMaxLifetime() time.Duration

		RateLimitEnabled() bool
		RPS() int
		Burst() int

		IdleTimeout() time.Duration
		MaxTimeout() time.Duration
		ReadTimeout() time.Duration
		WriteTimeout() time.Duration
	}
)

// NewConfig initializes a new AppConfig instance using the Builder pattern. It creates a Builder, builds the
// configuration, and returns a pointer to the constructed AppConfig.
func NewConfig() *AppConfig {
	return (&Builder{}).Build()
}

// Builder builds the AppConfig instance with environment variables and defaults
type Builder struct{}

// Build creates an AppConfig instance, populating it with values from environment variables, falling back to defaults
// when necessary.
func (cb *Builder) Build() *AppConfig {
	cfg := &AppConfig{
		environment: cb.getenv("ENVIRONMENT", "development"),
		httpHost:    cb.getenv("HTTP_HOST", "0.0.0.0"),
		httpPort:    cb.getenvInt("HTTP_PORT", 8080),
		logLevel:    cb.getenvInt("LOG_LEVEL", 1),

		corsSettings: CORSSettings{
			corsEnabled:    cb.getenvBool("CORS_ENABLED", false),
			trustedOrigins: cb.getenvList("CORS_ALLOWED_ORIGINS", []string{"*"}),
		},
		databaseSettings: DatabaseSettings{
			connMaxIdleTime:    time.Duration(cb.getenvInt("DB_CONN_MAX_IDLE_TIME", 60)) * time.Second,
			connMaxLifetime:    time.Duration(cb.getenvInt("DB_CONN_MAX_LIFETIME", 300)) * time.Second,
			dbConnectionString: cb.getenv("DB_CONNECTION_STRING", ""),
			maxIdleConns:       cb.getenvInt("DB_MAX_IDLE_CONNS", 2),
			maxOpenConns:       cb.getenvInt("DB_MAX_OPEN_CONNS", 5),
		},
		rateLimiterSettings: RateLimiterSettings{
			burst:   cb.getenvInt("RATE_LIMIT_BURST", 3),
			enabled: cb.getenvBool("RATE_LIMIT_ENABLED", false),
			rps:     cb.getenvInt("RATE_LIMIT_RPS", 3),
		},
		serverSettings: ServerSettings{
			idleTimeout:  time.Duration(cb.getenvInt("SERVER_IDLE_TIMEOUT", 120)) * time.Second,
			readTimeout:  time.Duration(cb.getenvInt("SERVER_READ_TIMEOUT", 5)) * time.Second,
			writeTimeout: time.Duration(cb.getenvInt("SERVER_WRITE_TIMEOUT", 2)) * time.Second,
		},
	}

	return cfg
}

// Burst returns the burst limit for the rate limiter.
func (c *AppConfig) Burst() int { return c.rateLimiterSettings.burst }

// CORSEnabled returns a boolean indicating if CORS is enabled.
func (c *AppConfig) CORSEnabled() bool { return c.corsSettings.corsEnabled }

// CORSTrustedOrigins returns a slice of trusted origins for CORS.
func (c *AppConfig) CORSTrustedOrigins() []string { return c.corsSettings.trustedOrigins }

// ConnMaxIdleTime returns the maximum idle time for database connections.
func (c *AppConfig) ConnMaxIdleTime() time.Duration { return c.databaseSettings.connMaxIdleTime }

// ConnMaxLifetime returns the maximum lifetime for database connections.
func (c *AppConfig) ConnMaxLifetime() time.Duration { return c.databaseSettings.connMaxLifetime }

// DbConnectionString returns the connection string for the database.
func (c *AppConfig) DbConnectionString() string { return c.databaseSettings.dbConnectionString }

// Environment returns the current application environment (e.g., development, production).
func (c *AppConfig) Environment() string { return c.environment }

// HTTPHost returns the host for the HTTP server.
func (c *AppConfig) HTTPHost() string { return c.httpHost }

// HTTPPort returns the port for the HTTP server.
func (c *AppConfig) HTTPPort() int { return c.httpPort }

// LogLevel returns the log level for the application.
func (c *AppConfig) LogLevel() int { return c.logLevel }

// MaxIdleConns returns the maximum number of idle connections to the database.
func (c *AppConfig) MaxIdleConns() int { return c.databaseSettings.maxIdleConns }

// MaxOpenConns returns the maximum number of open connections to the database.
func (c *AppConfig) MaxOpenConns() int { return c.databaseSettings.maxOpenConns }

// RPS returns the rate limit for requests per second.
func (c *AppConfig) RPS() int { return c.rateLimiterSettings.rps }

// RateLimitEnabled returns a boolean indicating if rate limiting is enabled.
func (c *AppConfig) RateLimitEnabled() bool { return c.rateLimiterSettings.enabled }

// IdleTimeout returns the idle timeout duration for the server.
func (c *AppConfig) IdleTimeout() time.Duration { return c.serverSettings.idleTimeout }

// MaxTimeout returns the maximum timeout duration for the server.
func (c *AppConfig) MaxTimeout() time.Duration { return c.serverSettings.maxTimeout }

// ReadTimeout returns the read timeout duration for the server.
func (c *AppConfig) ReadTimeout() time.Duration { return c.serverSettings.readTimeout }

// WriteTimeout returns the write timeout duration for the server.
func (c *AppConfig) WriteTimeout() time.Duration { return c.serverSettings.writeTimeout }

func (cb *Builder) getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	log.Printf("ENV %s is empty, using fallback", key)
	return fallback
}

func (cb *Builder) getenvInt(key string, fallback int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	log.Printf("ENV %s is empty or invalid, using fallback", key)
	return fallback
}

func (cb *Builder) getenvBool(key string, fallback bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	log.Printf("ENV %s is empty or invalid, using fallback", key)
	return fallback
}

func (cb *Builder) getenvList(key string, fallback []string) []string {
	if value := os.Getenv(key); value != "" {
		return strings.Split(value, ";")
	}
	log.Printf("ENV %s is empty, using fallback", key)
	return fallback
}
