package trace

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"log"
)

var (
	tp    *trace.TracerProvider
	stdTp *trace.TracerProvider
)

func StdTracerProvider() {
	var err error
	exp, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		log.Panicf("failed to initialize stdouttrace exporter %v\n", err)
		return
	}
	bsp := trace.NewBatchSpanProcessor(exp)
	stdTp = trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithSpanProcessor(bsp),
	)
	otel.SetTracerProvider(stdTp)
}

func TracerProvider(serviceName string, url string) (*trace.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := trace.NewTracerProvider(
		// Always be sure to batch in production.
		trace.WithBatcher(exp),
		// Record information about this application in an Resource.
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
			attribute.String("environment", "fdsf"),
			attribute.Int64("ID", 1),
		)),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp, nil
}

func Init() context.Context {
	var err error
	tp, err = TracerProvider("demo", "http://localhost:14268/api/traces")
	if err != nil {
		log.Fatal(err)
	}
	tr := tp.Tracer("cmdb")
	ctx, _ := context.WithCancel(context.Background())
	ctx, span := tr.Start(ctx, "foo")
	defer span.End()
	tr = otel.Tracer("new-")
	ctx, span = tr.Start(ctx, "bar")
	fmt.Println("执行到 这里了")
	defer span.End()
	return ctx
}

func GetTr() context.Context {
	StdTracerProvider()
	tr := stdTp.Tracer("cmdb")
	ctx := context.Background()
	ctx, span := tr.Start(ctx, "ent")
	defer span.End()
	tr = otel.Tracer("new_cmdb")
	tr.Start(ctx, "fsdf")
	return ctx
}
