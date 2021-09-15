package zlog

import (
	"context"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/codes"
	otelsdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

func NewTraceHook(ctx context.Context, spanName string, provider *otelsdk.TracerProvider) *TraceHook {
	spanCtx, span := provider.Tracer(spanName).Start(ctx, spanName)
	return &TraceHook{
		traceOption: &TraceOption{
			span: span,
			spanCtx: spanCtx,
		},
	}
}

type TraceOption struct {
	span trace.Span
	spanCtx context.Context
}

type TraceHook struct {
	traceOption *TraceOption
}

func (t TraceHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	t.traceOption.span.AddEvent(semconv.ExceptionEventName,trace.WithAttributes(semconv.ExceptionTypeKey.String("log"),semconv.ExceptionMessageKey.String(msg)))
	if level == zerolog.ErrorLevel {
		t.traceOption.span.SetStatus(codes.Error, "")
	}
}
