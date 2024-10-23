package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

var _ Config = (*AppConfig)(nil)

type (
	AppConfig struct {
		environment string
		httpHost    string
		httpPort    int
		logLevel    int
		dbConn      *DatabaseSettings
	}

	DatabaseSettings struct {
		DbConnectionString string
		DbMaxOpenConns     int
		DbMaxIdleConns     int
		DbConnMaxIdleTime  time.Duration
		DbConnMaxLifetime  time.Duration
	}

	Config interface {
		Environment() string
		HttpHost() string
		HttpPort() int
		LogLevel() int

		DbConn() *DatabaseSettings
	}
)

func (c *AppConfig) Environment() string {
	return c.environment
}

func (c *AppConfig) HttpHost() string {
	return c.httpHost
}

func (c *AppConfig) HttpPort() int {
	return c.httpPort
}

func (c *AppConfig) LogLevel() int {
	return c.logLevel
}

func (c *AppConfig) DbConn() *DatabaseSettings {
	return c.dbConn
}

func NewConfig() *AppConfig {
	env := getenv("ENVIRONMENT", "development")
	host := getenv("HTTP_HOST", "0.0.0.0")
	port := getenvInt("HTTP_PORT", 8080)
	logLevel := getenvInt("LOG_LEVEL", 1)

	dbConnSettings := &DatabaseSettings{
		DbConnectionString: getenv("DB_CONNECTION_STRING", ""),
		DbMaxOpenConns:     getenvInt("DB_MAX_OPEN_CONNS", 5),
		DbMaxIdleConns:     getenvInt("DB_MAX_IDLE_CONNS", 2),
		DbConnMaxIdleTime:  time.Duration(getenvInt("DB_CONN_MAX_IDLE_TIME", 60)) * time.Second,
		DbConnMaxLifetime:  time.Duration(getenvInt("DB_CONN_MAX_LIFETIME", 300)) * time.Second,
	}

	return &AppConfig{
		environment: env,
		httpHost:    host,
		httpPort:    port,
		logLevel:    logLevel,
		dbConn:      dbConnSettings,
	}
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Printf("ENV %s is empty, using %s", key, fallback)
		return fallback
	}
	return value
}

func getenvInt(key string, fallback int) int {
	value := os.Getenv(key)
	if value == "" {
		log.Printf("ENV %s is empty, using %d", key, fallback)
		return fallback
	}
	valueInt, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return valueInt
}
