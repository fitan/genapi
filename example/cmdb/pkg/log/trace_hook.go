package log

import (
	"context"
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

func (t TraceHook) Write(p []byte) (n int, err error) {
	t.traceOption.span.AddEvent(semconv.ExceptionEventName,trace.WithAttributes(semconv.ExceptionTypeKey.String("log"),semconv.ExceptionMessageKey.String(string(p))))
	return 0, nil
}

func (t TraceHook) Sync() error {
	return nil
}


