package config

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var _ Config = (*AppConfig)(nil)

func (c *AppConfig) Burst() int                     { return c.rateLimiterSettings.burst }
func (c *AppConfig) CORSEnabled() bool              { return c.corsSettings.corsEnabled }
func (c *AppConfig) CORSTrustedOrigins() []string   { return c.corsSettings.trustedOrigins }
func (c *AppConfig) ConnMaxIdleTime() time.Duration { return c.databaseSettings.connMaxIdleTime }
func (c *AppConfig) ConnMaxLifetime() time.Duration { return c.databaseSettings.connMaxLifetime }
func (c *AppConfig) DbConnectionString() string     { return c.databaseSettings.dbConnectionString }
func (c *AppConfig) Environment() string            { return c.environment }
func (c *AppConfig) HTTPHost() string               { return c.httpHost }
func (c *AppConfig) HTTPPort() int                  { return c.httpPort }
func (c *AppConfig) LogLevel() int                  { return c.logLevel }
func (c *AppConfig) MaxIdleConns() int              { return c.databaseSettings.maxIdleConns }
func (c *AppConfig) MaxOpenConns() int              { return c.databaseSettings.maxOpenConns }
func (c *AppConfig) RPS() int                       { return c.rateLimiterSettings.rps }
func (c *AppConfig) RateLimitEnabled() bool         { return c.rateLimiterSettings.enabled }

// Builder builds the AppConfig instance with environment variables and defaults
type Builder struct{}

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
	}

	return cfg
}

// Environment variable helpers
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

func NewConfig() *AppConfig {
	return (&Builder{}).Build()
}
