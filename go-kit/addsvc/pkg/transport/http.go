package transport

import (
	"addsvc/pkg/endpoint"
	"addsvc/pkg/service"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/circuitbreaker"
	endpoint2 "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/kit/transport"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/sony/gobreaker"
	"go.opentelemetry.io/contrib/instrumentation/github.com/go-kit/kit/otelkit"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/time/rate"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func NewHttpClient(instance string, tracerProvider trace.TracerProvider) (service.Service, error) {
	// Quickly sanitize the instance string.
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		return nil, err
	}
	limiter := ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 100))
	var options []httpTransport.ClientOption
	var sumEndpoint endpoint2.Endpoint
	{
		sumEndpoint = httpTransport.NewClient(
			"POST",
			copyUrl(u, "/sum"),
			encodeHttpGenericRequest,
			decodeHttpSumResponse,
			append(options, httpTransport.ClientBefore(otelHttpClientToContext()))...,
		).Endpoint()
		sumEndpoint = otelkit.EndpointMiddleware(otelkit.WithOperation("sum"), otelkit.WithTracerProvider(tracerProvider))(sumEndpoint)
		sumEndpoint = limiter(sumEndpoint)
		sumEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "Sum",
			Timeout: 30 * time.Second,
		}))(sumEndpoint)
	}

	var concatEndpoint endpoint2.Endpoint
	{
		concatEndpoint = httpTransport.NewClient(
			"POST",
			copyUrl(u, "/concat"),
			encodeHttpGenericRequest,
			decodeHttpConcatResponse,
			append(options, httpTransport.ClientBefore(otelHttpClientToContext()))...,
		).Endpoint()
		concatEndpoint = otelkit.EndpointMiddleware(otelkit.WithOperation("sum"), otelkit.WithTracerProvider(tracerProvider))(concatEndpoint)
		concatEndpoint = limiter(concatEndpoint)
		concatEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "Concat",
			Timeout: 10 * time.Second,
		}))(concatEndpoint)
	}
	return endpoint.Set{
		SumEndpoint:    sumEndpoint,
		ConcatEndpoint: concatEndpoint,
	}, nil
}

func copyUrl(base *url.URL, path string) *url.URL {
	var target = *base
	target.Path = path
	return &target
}

func encodeHttpGenericRequest(_ context.Context, r *http.Request, request interface{}) error {
	// 将请求对象编码为JSON
	bodyBytes, err := json.Marshal(request)
	if err != nil {
		return err
	}
	// 设置请求的Content-Type为application/json
	r.Header.Set("Content-Type", "application/json")
	// 将编码后的JSON写入请求的Body中
	r.Body = io.NopCloser(bytes.NewReader(bodyBytes))
	// 如果原来的Body有Close方法，需要调用它以释放资源（但在这种情况下，我们通常是创建新的Body，所以这一步是可选的）
	// 注意：如果r.Body原本就是nil，则不需要这一步
	return nil
}
func decodeHttpSumResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var resp endpoint.SumResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}
func decodeHttpConcatResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var resp endpoint.ConcatResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}
func otelHttpClientToContext() httpTransport.RequestFunc {
	return func(ctx context.Context, request *http.Request) context.Context {
		r := request
		propagator := otel.GetTextMapPropagator()
		// 将横切关注点从ctx设置到header
		propagator.Inject(ctx, propagation.HeaderCarrier(r.Header))
		return ctx
	}
}

// ----------------------------------------------------------

func NewHttpHandler(endpoints endpoint.Set, logger log.Logger) http.Handler {
	options := []httpTransport.ServerOption{
		httpTransport.ServerErrorEncoder(errorEncoder),
		httpTransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}
	m := http.NewServeMux()
	m.Handle("/sum", httpTransport.NewServer(
		endpoints.SumEndpoint,
		decodeSumRequest,
		encodeResponse,
		append(options, httpTransport.ServerBefore(otelHttpToContext()))...,
	))
	m.Handle("/concat", httpTransport.NewServer(
		endpoints.ConcatEndpoint,
		decodeConcatRequest,
		encodeResponse,
		append(options, httpTransport.ServerBefore(otelHttpToContext()))...,
	))
	return m
}

func otelHttpToContext() httpTransport.RequestFunc {
	return func(ctx context.Context, r *http.Request) context.Context {
		propagator := otel.GetTextMapPropagator()
		wireContext := propagator.Extract(r.Context(), propagation.HeaderCarrier(r.Header))
		return wireContext
	}
}

func decodeSumRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoint.SumRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
func decodeConcatRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoint.ConcatRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(errWrapper{Error: err.Error()})
}

func error2Code(err error) int {
	switch err {
	case service.ErrTwoZeros, service.ErrIntOverflow, service.ErrMaxSizeExceeded:
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

type errWrapper struct {
	Error string `json:"error"`
}
