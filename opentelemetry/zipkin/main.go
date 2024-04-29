package main

import (
	"context"
	"errors"
	"flag"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"log"
)

const (
	serverName  = "trace-demo-zipkin"
	environment = "production"
	id          = 1
)

func traceProvider(url string) (*tracesdk.TracerProvider, error) {
	exp, err := zipkin.New(url)
	if err != nil {
		log.Fatal(err)
	}
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceName(serverName),
				attribute.String("environment", environment),
				attribute.Int("id", id),
			),
		),
	)
	return tp, nil
}

func main() {
	url := flag.String("zipkin", "http://localhost:9411/api/v2/spans", "")
	tp, err := traceProvider(*url)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	defer func(ctx context.Context) {
		err := tp.Shutdown(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(ctx)
	otel.SetTracerProvider(tp)

	tr := tp.Tracer("component-main")
	ctx, span := tr.Start(ctx, "foo")
	defer span.End()
	err = bar(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}
}

func bar(ctx context.Context) error {
	// 数据采集固定逻辑
	_, span := otel.Tracer("component-bar").Start(ctx, "bar")
	defer span.End()
	// 业务逻辑
	span.SetAttributes(attribute.Key("testSet").String("value"))
	err := errors.New("出现了错误")
	span.AddEvent(err.Error())
	span.SetStatus(codes.Error, err.Error())
	return err

}
