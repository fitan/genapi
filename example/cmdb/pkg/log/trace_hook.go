package log

import (
	"context"
	"fmt"
	otelsdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

func NewTraceHook(ctx context.Context, spanName string, provider *otelsdk.TracerProvider) *TraceHook {
	return &TraceHook{
		traceOption: &TraceOption{
			spanName: spanName,
			ctx:      ctx,
			tp:       provider,
		},
	}
}

type TraceOption struct {
	spanName string
	ctx      context.Context
	tp *otelsdk.TracerProvider
}

type TraceHook struct {
	traceOption *TraceOption
}

func (t TraceHook) Write(p []byte) (n int, err error) {
	tr := t.traceOption.tp.Tracer("send log")
	spanCtx, span := tr.Start(t.traceOption.ctx, t.traceOption.spanName)
	fmt.Println(span.SpanContext().SpanID())
	span.AddEvent(semconv.ExceptionEventName,trace.WithAttributes(semconv.ExceptionTypeKey.String("info"),semconv.ExceptionMessageKey.String(string(p))))
	span.End()
	t.traceOption.ctx = spanCtx
	return 0, nil
}

func (t TraceHook) Sync() error {
	return nil
}


