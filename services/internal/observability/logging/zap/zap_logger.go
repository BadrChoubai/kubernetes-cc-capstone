package zap

import (
	"github.com/badrchoubai/services/internal/observability/logging"
	"go.uber.org/zap"
)

type zapLogger struct {
	logger *zap.Logger
}

// NewZapLogger creates new instance of Logger using go.uber.org/zap as underlying logging library
func NewZapLogger(level int) logging.Logger {
	switch level {
	case logging.DebugLevel:
		l, err := zap.NewDevelopment()
		if err != nil {
			panic(err)
		}

		return &zapLogger{logger: l.WithOptions(zap.AddCallerSkip(1))}
	default:
		l, err := zap.NewProduction()
		if err != nil {
			panic(err)
		}

		return &zapLogger{logger: l.WithOptions(zap.AddCallerSkip(1))}
	}
}

// Info makes call to zap library's Info func
func (l *zapLogger) Info(message string) {
	l.logger.Info(message)
}

// Error makes call to zap library's Error func
func (l *zapLogger) Error(err error, whatWasHappening string) {
	l.logger.Error(err.Error(), zap.String("whatWasHappening", whatWasHappening))
}

// WithName gives a name to a used instance of Logger
func (l *zapLogger) WithName(name string) logging.Logger {
	const (
		loggerNameKey = "_name_"
	)

	l2 := l.logger.With(zap.String(loggerNameKey, name))
	return &zapLogger{logger: l2}
}
