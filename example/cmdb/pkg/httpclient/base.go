package httpclient

import (
	"context"
	"github.com/go-resty/resty/v2"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"time"

	"go.opentelemetry.io/otel"
)

type option struct {
	Host string
	TraceContext context.Context
	TraceName string
	Debug bool
	TimeOut time.Duration
	RetryCount int
	RetryWaitTime time.Duration
	RetryMaxWaitTime time.Duration

}


type Option func(*option)


func NewClient(fs ...Option) *resty.Client {
	o := option{
		Host:             "",
		TraceContext:     nil,
		Debug:            false,
		TimeOut:          10 * time.Second,
		RetryCount:       3,
		RetryWaitTime:    10 * time.Second,
		RetryMaxWaitTime: 30 * time.Second,
	}
	for _, f := range fs {
		f(&o)
	}

	if o.TraceName != "" {
		otel.Tracer(o.TraceName)
	}

	client := resty.New().SetDebug(o.Debug).SetHostURL(o.Host).SetTimeout(o.TimeOut).SetRetryCount(o.RetryCount).SetRetryWaitTime(o.RetryWaitTime).SetRetryMaxWaitTime(o.RetryMaxWaitTime)
	if o.TraceName != "" {
		tr := otel.Tracer(o.TraceName)
		client.OnAfterResponse(TraceAfter(tr))
		client.OnBeforeRequest(TraceBefor())
	}
	return client
}

func WithHost(host string) Option {
	return func(o *option) {
		o.Host = host
	}
}

func WithDebug(debug bool) Option {
	return func(o *option) {
		o.Debug = debug
	}
}

func WithTimeOut(timeOut time.Duration) Option {
	return func(o *option) {
		o.TimeOut = timeOut
	}
}

func WithRetry(retryCount int, retryWaitTime,retryMaxWaitTime time.Duration) Option {
	return func(o *option) {
		o.RetryCount = retryCount
		o.RetryWaitTime = retryWaitTime
		o.RetryMaxWaitTime = retryMaxWaitTime
	}
}

func WithTraceContext(traceName string, c context.Context) Option {
	return func(o *option) {
		o.TraceName = traceName
		o.TraceContext = c
	}
}


func TraceBefor() resty.RequestMiddleware {
	return func(client *resty.Client, request *resty.Request) error {
		client.SetTransport(otelhttp.NewTransport(http.DefaultTransport))
		return nil
	}
}

func TraceAfter(tr trace.Tracer) resty.ResponseMiddleware {
	return func(client *resty.Client, response *resty.Response) error {
		subContext, span := tr.Start(response.Request.Context(), response.Request.URL, trace.WithAttributes(attribute.String("http_info", SetTraceInfo(response))))
		defer span.End()
		context.WithValue(subContext, "sub_trace", subContext)
		return nil
	}
}
