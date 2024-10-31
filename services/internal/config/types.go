package config

import "time"

type (
	AppConfig struct {
		environment string
		httpHost    string
		httpPort    int
		logLevel    int

		corsSettings        CORSSettings
		databaseSettings    DatabaseSettings
		rateLimiterSettings RateLimiterSettings
	}

	CORSSettings struct {
		corsEnabled    bool
		trustedOrigins []string
	}

	DatabaseSettings struct {
		connMaxIdleTime    time.Duration
		connMaxLifetime    time.Duration
		dbConnectionString string
		maxOpenConns       int
		maxIdleConns       int
	}

	RateLimiterSettings struct {
		burst   int
		enabled bool
		rps     int
	}

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
	}
)
