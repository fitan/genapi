package log

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func NewTraceHook(tr trace.Tracer, ctx context.Context, spanName string) *TraceHook {
	return &TraceHook{
		spanName: spanName,
		tr: tr,
		ctx: ctx,
	}
}

type TraceHook struct {
	spanName string
	tr trace.Tracer
	ctx context.Context
	spanCtx context.Context
	write string
}

func (t TraceHook) Write(p []byte) (n int, err error) {
	t.write = string(p)
	fmt.Println("this write: ", t.write)
	return 0, nil
}

func (t TraceHook) Sync() error {
	spanCtx, span := t.tr.Start(t.ctx, t.spanName, trace.WithAttributes(attribute.String("zap_log", t.write)))
	fmt.Println(span.IsRecording(), "可以发送")
	defer span.End()
	t.spanCtx = spanCtx
	return nil
}

func (t TraceHook) GetSpanCtx() context.Context {
	return t.spanCtx
}
