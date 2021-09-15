package zlog

import (
	"context"
	"github.com/rs/zerolog"
	otelsdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

type zlog struct {
	traceLevel zerolog.Level
	tp *otelsdk.TracerProvider
	*zerolog.Logger
}

func (z *zlog) TraceLog(ctx context.Context, spanName string) *zTraceLog {
	traceID := trace.SpanFromContext(ctx).SpanContext().TraceID().String()
	ztl := new(zTraceLog)
	ztl.traceHook = NewTraceHook(ctx, spanName, z.tp)
	ztl.Logger = z.Logger.With().Str("traceID", traceID).Logger().Hook(ztl.traceHook)
	return ztl
}

type zTraceLog struct {
	traceHook *TraceHook
	zerolog.Logger
}

func (z *zTraceLog) Context() context.Context {
	return z.traceHook.traceOption.spanCtx
}

func (z *zTraceLog) End() {
	z.traceHook.traceOption.span.End()
}
