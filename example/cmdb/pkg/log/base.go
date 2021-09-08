package log

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type SeverityHook struct{}

func (h SeverityHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	e.
	if level != zerolog.NoLevel {
		e.Str("severity", level.String())
	}
}

type Logger struct {
	underlying zerolog.Logger
}

func (l *Logger) TraceContext(ctx context.Context) *zerolog.Logger {
	// TODO The switching of pointer and value semantic might go against what zerolog was trying to convey,
	// but we need figure out if hooks are meant to be only initialized once vs being added dynamically
	return &logger
}

func NewLog()  {
	l := Logger{}
	l.TraceContext(context.TODO())
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Trace()
	log.
	hooked := log.Hook(SeverityHook{})
	hooked.Info()
	hooked.con
	log.Fatal()
}