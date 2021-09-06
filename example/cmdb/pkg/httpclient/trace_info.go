package httpclient

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"time"
)

type TraceInfo struct {
	RequestMethod string `json:"request_method"`
	RequestBody interface{} `json:"request_body"`
	RequestUrl string `json:"request_url"`
	RequestQueryParam string `json:"request_query_param"`
	RequestFormData string `json:"request_form_data"`
	ResponseData string `json:"response_data"`
	ResponseCode int `json:"response_code"`

	DNSLookup time.Duration `json:"dns_lookup"`
	ConnTime   time.Duration `json:"conn_time"`
	TCPConnTime   time.Duration `json:"tcp_conn_time"`
	TLSHandshake  time.Duration `json:"tls_handshake"`
	ServerTime    time.Duration `json:"server_time"`
	ResponseTime  time.Duration `json:"response_time"`
	TotalTime     time.Duration `json:"total_time"`
	IsConnReused  bool `json:"is_conn_reused"`
	IsConnWasIdle bool `json:"is_conn_was_idle"`
	ConnIdleTime  time.Duration `json:"conn_idle_time"`
	RequestAttempt int `json:"request_attempt"`
	RemoteAddr    string `json:"remote_addr"`
}

func SetTraceInfo(resp *resty.Response) string {
	traceInfo := resp.Request.TraceInfo()
	Method := resp.Request.Method
	url := resp.Request.URL
	body := resp.Request.Body
	queryParam := resp.Request.QueryParam
	formData := resp.Request.FormData
	responseData := resp.String()
	responseCode := resp.StatusCode()
	info :=  TraceInfo{
		RequestMethod:     Method,
		RequestBody:       body,
		RequestUrl: url,
		RequestQueryParam: queryParam.Encode(),
		RequestFormData:   formData.Encode(),
		ResponseData:      responseData,
		ResponseCode: responseCode,

		DNSLookup:         traceInfo.DNSLookup,
		ConnTime:          traceInfo.ConnTime,
		TCPConnTime:       traceInfo.TCPConnTime,
		TLSHandshake:      traceInfo.TLSHandshake,
		ServerTime:        traceInfo.ServerTime,
		ResponseTime:      traceInfo.ResponseTime,
		TotalTime:         traceInfo.TotalTime,
		IsConnReused:      traceInfo.IsConnReused,
		IsConnWasIdle:     traceInfo.IsConnWasIdle,
		ConnIdleTime:      traceInfo.ConnIdleTime,
		RequestAttempt:    traceInfo.RequestAttempt,
		RemoteAddr:        traceInfo.RemoteAddr.String(),
	}

	b, _ := json.Marshal(info)
	return string(b)
}
