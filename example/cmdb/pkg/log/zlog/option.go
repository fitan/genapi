package zlog

import (
	"github.com/rs/zerolog"
	otelsdk "go.opentelemetry.io/otel/sdk/trace"
	"gopkg.in/natefinch/lumberjack.v2"
)

type option struct {
	traceLevel zerolog.Level
	tp *otelsdk.TracerProvider
	logLevel zerolog.Level
	fileName string
	*zerolog.Logger
}

type Option func(o *option)


func getLogWriter(fileName string) *lumberjack.Logger {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    100,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return lumberJackLogger
}

func WithLogger(logger *zerolog.Logger) Option {
	return func(o *option) {
		o.Logger = logger
	}
}

func WithLogFileName(fileName string) Option {
	return func(o *option) {
		o.fileName = fileName
	}
}

func WithLogLevel(level zerolog.Level) Option {
	return func(o *option) {
		o.logLevel = level
	}
}

func WithTrace(tp *otelsdk.TracerProvider,level zerolog.Level) Option {
	return func(o *option) {
		o.traceLevel = level
		o.tp = tp
	}
}



func NewZlog(fs ...Option) *zlog {
	zlog := new(zlog)
	o := new(option)
	for _, f := range fs {
		f(o)
	}

	if o.Logger == nil {
		fileName := o.fileName
		logLever := o.logLevel
		if fileName == "" {
			fileName = "./logs/zlog.log"
		}
		zerolog.MultiLevelWriter()
		l := zerolog.New(getLogWriter(fileName)).With().Timestamp().Caller().Logger().Level(logLever)
		zlog.Logger = &l
	} else {
		zlog.Logger = o.Logger
	}
	zlog.traceLevel = o.traceLevel
	zlog.tp = o.tp
	return zlog
}
