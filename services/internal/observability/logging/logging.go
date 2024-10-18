package logging

import "go.uber.org/zap"

// Logger is a simple interface we can use to integrate calls between a preferred library [zap] and our code
type (
	Level uint8

	// Logger defines the contract between a logging library and caller
	Logger interface {
		Info(message string, fields ...zap.Field)
		Error(err error, whatWasHappening string)

		WithName(name string) Logger
	}
)

const (
	DebugLevel int = iota + 1
	InfoLevel
)
