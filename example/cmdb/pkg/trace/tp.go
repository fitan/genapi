package trace

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/trace"
)

//func TracerProvider(url string) (*trace.TracerProvider, error) {
//	// Create the Jaeger exporter
//	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
//	if err != nil {
//		return nil, err
//	}
//	tp := trace.NewTracerProvider(
//		// Always be sure to batch in production.
//		trace.WithBatcher(exp),
//		// Record information about this application in an Resource.
//		trace.WithResource(resource.NewWithAttributes(
//			semconv.SchemaURL,
//			semconv.ServiceNameKey.String(service),
//			attribute.String("environment", environment),
//			attribute.Int64("ID", id),
//		)),
//	)
//	return tp, nil
//}
