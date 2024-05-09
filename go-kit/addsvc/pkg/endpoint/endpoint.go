package endpoint

import (
	"addsvc/pkg/service"
	"context"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/ratelimit"
	"github.com/sony/gobreaker"
	"go.opentelemetry.io/contrib/instrumentation/github.com/go-kit/kit/otelkit"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/time/rate"
	"time"
)

type SumRequest struct {
	A, B int
}
type ConcatRequest struct {
	A, B string
}

type SumResponse struct {
	V   int   `json:"v"`
	Err error `json:"-"`
}
type ConcatResponse struct {
	V   string `json:"v"`
	Err error  `json:"-"`
}

func (r SumResponse) Failed() error {
	return r.Err
}

func (r ConcatResponse) Failed() error {
	return r.Err
}

var (
	_ endpoint.Failer = &SumResponse{}
	_ endpoint.Failer = &ConcatResponse{}
)

type Set struct {
	SumEndpoint    endpoint.Endpoint
	ConcatEndpoint endpoint.Endpoint
}

func (s Set) Sum(ctx context.Context, a, b int) (int, error) {
	res, err := s.SumEndpoint(ctx, SumRequest{A: a, B: b})
	if err != nil {
		return 0, err
	}
	response := res.(SumResponse)
	return response.V, response.Err
}
func (s Set) Concat(ctx context.Context, a, b string) (string, error) {
	res, err := s.ConcatEndpoint(ctx, ConcatRequest{A: a, B: b})
	if err != nil {
		return "", err
	}
	response := res.(ConcatResponse)
	return response.V, response.Err
}

func MakeSumEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SumRequest)
		v, err := s.Sum(ctx, req.A, req.B)
		return SumResponse{
			V:   v,
			Err: err,
		}, nil
	}
}

func MakeConcatEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ConcatRequest)
		v, err := s.Concat(ctx, req.A, req.B)
		return ConcatResponse{
			V:   v,
			Err: err,
		}, nil
	}
}

func New(svc service.Service, logger log.Logger, traceProvider trace.TracerProvider) Set {
	var sumEndpoint endpoint.Endpoint
	{
		sumEndpoint = MakeSumEndpoint(svc)
		// 限流
		sumEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 1))(sumEndpoint)
		// 断路器
		sumEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(sumEndpoint)
		// 链路追踪
		if traceProvider != nil {
			sumEndpoint = otelkit.EndpointMiddleware(otelkit.WithOperation("sum"), otelkit.WithTracerProvider(traceProvider))(sumEndpoint)
		}
		// 日志中间件
		sumEndpoint = LoggingMiddleware(log.With(logger, "method", "sum"))(sumEndpoint)
	}
	var concatEndpoint endpoint.Endpoint
	{
		concatEndpoint = MakeConcatEndpoint(svc)
		// 限流
		concatEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 1))(concatEndpoint)
		// 断路器
		concatEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(concatEndpoint)
		// 链路追踪
		if traceProvider != nil {
			concatEndpoint = otelkit.EndpointMiddleware(otelkit.WithOperation("concat"), otelkit.WithTracerProvider(traceProvider))(concatEndpoint)
		}
		// 日志中间件
		concatEndpoint = LoggingMiddleware(log.With(logger, "method", "concat"))(concatEndpoint)
	}
	return Set{
		SumEndpoint:    sumEndpoint,
		ConcatEndpoint: concatEndpoint,
	}
}
