package log

import (
	"context"
	"go.opentelemetry.io/otel/trace"

	//"github.com/uber/jaeger-client-go/log/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)



type Xlog struct {
	tr trace.Tracer
	traceLevel zapcore.Level
	*zap.Logger
}

func (x Xlog) TraceLog(ctx context.Context, spanName string) TraceLog {
	//traceID := trace.SpanFromContext(ctx).SpanContext().TraceID().String()
	hook := NewTraceHook(x.tr, ctx, spanName)
	traceCore := zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), hook, x.traceLevel)
	wrapCore := zap.WrapCore(
		func(core zapcore.Core) zapcore.Core {
			return zapcore.NewTee(core, traceCore)
		})
	return TraceLog{
		traceHook: hook,
		Logger: x.Logger.WithOptions(wrapCore),
		//Logger: x.Logger.WithOptions(
		//	wrapCore,
		//	zap.Fields(zap.String("traceID", traceID))),
	}
}

type TraceLog struct {
	traceHook *TraceHook
	*zap.Logger
}

func (t *TraceLog) GetSpanCtx() context.Context {
	return t.traceHook.GetSpanCtx()
}





//var xlog *Xlog

//func NewLog()  {
//
//
//	level := zap.DebugLevel
//	core := zapcore.NewCore(
//		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), // json格式日志（ELK渲染收集）
//		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),  // 打印到控制台和文件
//		level,                                                    // 日志级别
//	)
//
//
//	// 开启文件及行号
//	development := zap.Development()
//	Logger = zap.New(core,
//		zap.AddCaller(),
//		zap.AddStacktrace(zap.ErrorLevel),	// error级别日志，打印堆栈
//		development)
//	xlog = &Xlog{Logger}
//}
//
//func TestLog()  {
//	xlog.log.Debug("fdsdf")
//	xlog.log.
//}