package services

import (
	databaes "github.com/badrchoubai/services/internal/database"
	"sync"

	"github.com/badrchoubai/services/internal/encoding"
	logging "github.com/badrchoubai/services/internal/observability/logging/zap"
)

type Options struct {
	Name           string
	Logger         *logging.Logger
	ServiceMutex   *sync.Mutex
	EncoderDecoder *encoding.ServerEncoderDecoder
	DbConnection   *databaes.Database
}

type Option interface {
	Apply(*Options)
}

type nameOption string
type loggerOption struct {
	Log *logging.Logger
}
type serviceMutexOption struct {
	ServiceMutex *sync.Mutex
}
type encoderDecoderOption struct {
	EncoderDecoder *encoding.ServerEncoderDecoder
}
type dbConnectionOption struct {
	DbConnection *databaes.Database
}

func (n nameOption) Apply(opts *Options) {
	opts.Name = string(n)
}

func (l *loggerOption) Apply(opts *Options) {
	opts.Logger = l.Log
}

func (sm *serviceMutexOption) Apply(opts *Options) {
	opts.ServiceMutex = sm.ServiceMutex
}

func (edc *encoderDecoderOption) Apply(opts *Options) {
	opts.EncoderDecoder = edc.EncoderDecoder
}

func (dbc *dbConnectionOption) Apply(opts *Options) {
	opts.DbConnection = dbc.DbConnection
}

func WithName(name string) Option {
	return nameOption(name)
}

func WithLogger(logger *logging.Logger) Option {
	return &loggerOption{
		Log: logger,
	}
}

func WithServiceMutex(mutex *sync.Mutex) Option {
	return &serviceMutexOption{
		ServiceMutex: mutex,
	}
}

func WithEncoderDecoder(edc *encoding.ServerEncoderDecoder) Option {
	return &encoderDecoderOption{
		EncoderDecoder: edc,
	}
}

func WithDbConnection(conn *databaes.Database) Option {
	return &dbConnectionOption{
		DbConnection: conn,
	}
}
