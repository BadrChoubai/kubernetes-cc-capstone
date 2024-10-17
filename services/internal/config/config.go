package config

import (
	"log"
	"os"
	"strconv"
)

var _ Config = (*AppConfig)(nil)

type (
	AppConfig struct {
		environment string
		httpHost    string
		httpPort    int
		logLevel    int
	}

	Config interface {
		Environment() string
		HttpHost() string
		HttpPort() int
		LogLevel() int
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

func NewConfig() *AppConfig {
	env := getenv("ENVIRONMENT", "development")
	host := getenv("HTTP_HOST", "0.0.0.0")
	port := getenvInt("HTTP_PORT", 8080)
	logLevel := getenvInt("LOG_LEVEL", 1)

	return &AppConfig{
		environment: env,
		httpHost:    host,
		httpPort:    port,
		logLevel:    logLevel,
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
