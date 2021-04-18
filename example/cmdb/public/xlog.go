package public

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path"
)

var XLog *xLog
var buildLog bool

func init() {
	XLog = NewXLog(GetConf().Log.Dir, GetConf().App.Name)
}

func NewXLog(dir string, mark string) *xLog {
	if buildLog {
		MaxSize := GetConf().Log.MaxSize
		MaxBackups := GetConf().Log.MaxBackups
		MaxAge := GetConf().Log.MaxAge
		Compress := GetConf().Log.Compress
		fileName := func(mark string, name string) string {
			if mark == "" {
				return name
			} else {
				return mark + "_" + name
			}
		}
		infoLog := zerolog.New(&lumberjack.Logger{
			Filename:   path.Join(dir, fileName(mark, "info.log")),
			MaxSize:    MaxSize, // megabytes
			MaxBackups: MaxBackups,
			MaxAge:     MaxAge,   //days
			Compress:   Compress, // disabled by default
		}).With().Logger()
		warnLog := zerolog.New(&lumberjack.Logger{
			Filename:   path.Join(dir, fileName(mark, "warn.log")),
			MaxSize:    MaxSize, // megabytes
			MaxBackups: MaxBackups,
			MaxAge:     MaxAge,   //days
			Compress:   Compress, // disabled by default
		}).With().Logger()
		errLog := zerolog.New(&lumberjack.Logger{
			Filename:   path.Join(dir, fileName(mark, "err.log")),
			MaxSize:    MaxSize, // megabytes
			MaxBackups: MaxBackups,
			MaxAge:     MaxAge,   //days
			Compress:   Compress, // disabled by default
		}).With().Logger()
		fatalLog := zerolog.New(&lumberjack.Logger{
			Filename:   path.Join(dir, fileName(mark, "fatal.log")),
			MaxSize:    MaxSize, // megabytes
			MaxBackups: MaxBackups,
			MaxAge:     MaxAge,   //days
			Compress:   Compress, // disabled by default
		}).With().Logger()

		return &xLog{
			infoLog:  infoLog,
			wainLog:  warnLog,
			errLog:   errLog,
			fatalLog: fatalLog,
		}
	}

	l := zerolog.New(os.Stdout).With().Logger()
	return &xLog{
		infoLog:  l,
		wainLog:  l,
		errLog:   l,
		fatalLog: l,
	}
}

type xLog struct {
	infoLog  zerolog.Logger
	wainLog  zerolog.Logger
	errLog   zerolog.Logger
	fatalLog zerolog.Logger
}

func (x *xLog) Info() *zerolog.Event {
	return x.infoLog.Info().Caller(1)
}

func (x *xLog) Warn() *zerolog.Event {
	return x.wainLog.Warn().Caller(1)
}

func (x *xLog) Error() *zerolog.Event {
	return x.errLog.Error().Caller(1)
}

func (x *xLog) Fatal() *zerolog.Event {
	return x.fatalLog.Fatal().Caller(1)
}

type Result struct {
	Code int
	Data interface{}
	Err  string
}

func HttpResultTmpl(data interface{}, err error) Result {
	res := Result{Data: data}
	if err != nil {
		XLog.Error().Err(err).Msg("")
		res.Code = 503
		res.Err = err.Error()
	} else {
		res.Code = 200
	}
	return res
}

func GinHttpResult(c *gin.Context, data interface{}, err error) {
	c.JSON(200, HttpResultTmpl(data, err))
}
