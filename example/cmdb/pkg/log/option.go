package log

import (
	"go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Option func(xlog *Xlog)

func WithLogger(logger *zap.Logger) Option {
	return func(xlog *Xlog) {
		xlog.Logger = logger
	}
}

func WithTrace(tp *trace.TracerProvider,level zapcore.Level) Option {
	return func(xlog *Xlog) {
		xlog.traceLevel = level
		xlog.tp = tp
	}
}

func NewXlog(fs ...Option) (*Xlog, error) {
	xlog := new(Xlog)
	for _, f := range fs {
		f(xlog)
	}
	if xlog.Logger == nil {
		log, err := zap.NewProduction()
		xlog.Logger = log
		return xlog, err
	}
	return xlog, nil
}
