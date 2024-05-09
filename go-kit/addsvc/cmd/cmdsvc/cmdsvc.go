package main

import (
	"addsvc/pkg/endpoint"
	service2 "addsvc/pkg/service"
	"addsvc/pkg/transport"
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.25.0"
	"net"
	"net/http"
	"os"
)

func main() {
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	httpAddress := ":8080"
	jaegerUrl := "http://localhost:14268/api/traces"
	logger := getLogger()
	var err error
	var exp sdktrace.SpanExporter
	exp, err = getJaegerExporter(jaegerUrl)
	if err != nil {
		level.Error(logger).Log(err)
		os.Exit(1)
	}
	var tp *sdktrace.TracerProvider
	tp, err = getTracerProvider(exp)
	defer func() {
		tp.Shutdown(context.Background())
	}()
	otel.SetTracerProvider(tp)
	//var tracer trace.Tracer
	//tracer = tp.Tracer("addsvc")
	var service = service2.New(logger)
	var endpoints = endpoint.New(service, logger, tp)
	var httpHandler = transport.NewHttpHandler(endpoints, logger)
	httpListener, err := net.Listen("tcp", httpAddress)
	if err != nil {
		logger.Log(err)
		os.Exit(1)
	}
	http.Serve(httpListener, httpHandler)
}

func getTracerProvider(exp sdktrace.SpanExporter) (*sdktrace.TracerProvider, error) {
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(getSvcResources()),
	)
	return tp, nil
}

func getSvcResources() *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("addsvc"),
			semconv.ServiceVersion("v0.1.0"),
			attribute.String("environment", "dev"),
		))
	return r
}

func getJaegerExporter(url string) (sdktrace.SpanExporter, error) {
	return jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
}

func getLogger() (logger log.Logger) {
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)
	return
}
