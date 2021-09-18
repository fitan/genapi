package log

import (
	"context"
	"go.opentelemetry.io/otel/codes"
	otelsdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	//"github.com/uber/jaeger-client-go/log/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Xlog struct {
	traceLevel zapcore.Level
	tp *otelsdk.TracerProvider
	*zap.Logger
}

func (x Xlog) TraceLog(ctx context.Context, spanName string) *TraceLog {
	traceID := trace.SpanFromContext(ctx).SpanContext().TraceID().String()
	hook := NewTraceHook(ctx, spanName, x.tp)
	traceCore := zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), zapcore.AddSync(hook), x.traceLevel)
	wrapCore := zap.WrapCore(
		func(core zapcore.Core) zapcore.Core {
			return zapcore.NewTee(core, traceCore)
		})
	return &TraceLog{
		traceHook: hook,
		Logger:    x.Logger.WithOptions(wrapCore, zap.Fields(zap.String("traceID", traceID))),
		//Logger: x.Logger.WithOptions(
		//	wrapCore,
		//	zap.Fields(zap.String("traceID", traceID))),
	}
}

type TraceLog struct {
	traceHook *TraceHook
	*zap.Logger
}

func (t *TraceLog) With(fields ...zap.Field) {
	l := t.Logger.With(fields...)
	t.Logger = l
}

func (t *TraceLog) Context() context.Context {
	return t.traceHook.traceOption.spanCtx
}

func (t *TraceLog) End() {
	t.Sync()
	t.traceHook.traceOption.span.End()
}

func (t *TraceLog) Error(msg string, fields ...zap.Field) {
	t.traceHook.traceOption.span.SetStatus(codes.Error, "")
	t.Logger.Error(msg, fields...)
}