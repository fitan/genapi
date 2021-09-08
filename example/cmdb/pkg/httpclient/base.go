package httpclient

import (
	"github.com/go-resty/resty/v2"
	"time"

	"go.opentelemetry.io/otel"
)

type option struct {
	Host string

	TraceName string
	// 记录所有的详细的http info, 否则只记录发生错误时的http info
	TraceDebug bool

	Debug            bool
	TimeOut          time.Duration
	RetryCount       int
	RetryWaitTime    time.Duration
	RetryMaxWaitTime time.Duration
}

type Option func(*option)

func NewClient(fs ...Option) *resty.Client {
	o := option{
		Debug:            false,
		TimeOut:          30 * time.Second,
		RetryCount:       0,
		RetryWaitTime:    10 * time.Second,
		RetryMaxWaitTime: 30 * time.Second,
	}
	for _, f := range fs {
		f(&o)
	}

	client := resty.New().SetDebug(o.Debug).SetTimeout(o.TimeOut).SetRetryCount(o.RetryCount).SetRetryWaitTime(o.RetryWaitTime).SetRetryMaxWaitTime(o.RetryMaxWaitTime)
	if o.TraceName != "" {
		tr := otel.Tracer(o.TraceName)
		client = client.EnableTrace()
		client = client.OnBeforeRequest(BeforeTrace())
		if o.Debug {
			client = client.OnAfterResponse(AfterTraceDebug(tr))
		} else {
			client = client.OnError(ErrorTrace(tr))
		}
	}

	if o.Host != "" {
		client = client.SetHostURL(o.Host)
	}
	return client
}

func WithTrace(traceName string, traceDebug bool) func(o *option) {
	return func(o *option) {
		o.TraceName = traceName
		o.TraceDebug = traceDebug
	}
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

func WithRetry(retryCount int, retryWaitTime, retryMaxWaitTime time.Duration) Option {
	return func(o *option) {
		o.RetryCount = retryCount
		o.RetryWaitTime = retryWaitTime
		o.RetryMaxWaitTime = retryMaxWaitTime
	}
}
