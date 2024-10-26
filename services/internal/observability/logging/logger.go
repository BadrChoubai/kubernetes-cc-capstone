package logging

import "go.uber.org/zap"

var _ Log = (*Logger)(nil)

type Log interface {
	Info(message string, fields ...zap.Field)
	Error(whatWasHappening string, err error, fields ...zap.Field)
	Debug(message string, fields ...zap.Field)
}

type Logger struct {
	log *zap.Logger
}

func (l *Logger) Info(message string, fields ...zap.Field) {
	l.log.Info(message, fields...)
}

func (l *Logger) Error(whatWasHappening string, err error, fields ...zap.Field) {
	l.log.Error(whatWasHappening, append(fields, zap.Error(err))...)
}

func (l *Logger) Debug(message string, fields ...zap.Field) {
	l.log.Debug(message, fields...)
}

func NewLogger() (*Logger, error) {
	zapLogger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	return &Logger{
		log: zapLogger,
	}, nil
}
