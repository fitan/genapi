package log

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func NewTraceHook(ctx context.Context, spanName string) *TraceHook {
	return &TraceHook{
		spanName: spanName,
		ctx:      ctx,
	}
}

type TraceHook struct {
	spanName string
	ctx      context.Context
	spanCtx  context.Context
}

func (t TraceHook) Write(p []byte) (n int, err error) {
	tr := otel.Tracer(t.spanName)
	spanCtx, span := tr.Start(t.ctx, t.spanName, trace.WithAttributes(attribute.String("zap_log", string(p))))
	defer span.End()
	fmt.Println(span.IsRecording())
	t.spanCtx = spanCtx
	return 0, nil
}

func (t TraceHook) Sync() error {
	return nil
}

func (t TraceHook) GetSpanCtx() context.Context {
	return t.spanCtx
}
