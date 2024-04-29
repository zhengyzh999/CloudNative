package main

import (
	"context"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/prometheus"
	metric2 "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/metric"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	ctx := context.Background()
	// 既是数据导出器，也是数据采集器
	exp, err := prometheus.New()
	if err != nil {
		log.Fatal(err)
	}
	provider := metric.NewMeterProvider(metric.WithReader(exp))
	meter := provider.Meter("prometheus")

	// 启动一个prometheus服务协程
	go serveMetrics()
	attrs := []attribute.KeyValue{
		attribute.Key("A").String("B"),
		attribute.Key("C").String("D"),
	}
	// counter
	counter, err := meter.Float64Counter("foo", metric2.WithDescription("foo counter 描述"))
	if err != nil {
		log.Fatal(err)
	}
	counter.Add(ctx, 5, metric2.WithAttributes(attrs...))
	// observer
	gauge, err := meter.Float64ObservableGauge("bar", metric2.WithDescription("bar observer 描述"))
	if err != nil {
		log.Fatal(err)
	}
	_, err = meter.RegisterCallback(func(ctx context.Context, observer metric2.Observer) error {
		n := -10. + rand.Float64()*(90.)
		observer.ObserveFloat64(gauge, n, metric2.WithAttributes(attrs...))
		return nil
	}, gauge)
	// measure
	histogram, err := meter.Float64Histogram("baz", metric2.WithDescription("baz measure 描述"))
	if err != nil {
		log.Fatal(err)
	}
	histogram.Record(ctx, 27, metric2.WithAttributes(attrs...))
	histogram.Record(ctx, 50, metric2.WithAttributes(attrs...))
	histogram.Record(ctx, 7, metric2.WithAttributes(attrs...))
	histogram.Record(ctx, 100, metric2.WithAttributes(attrs...))
	histogram.Record(ctx, 78, metric2.WithAttributes(attrs...))
	histogram.Record(ctx, 101, metric2.WithAttributes(attrs...))
	histogram.Record(ctx, 46, metric2.WithAttributes(attrs...))
	histogram.Record(ctx, 105, metric2.WithAttributes(attrs...))

	ctx, _ = signal.NotifyContext(ctx, os.Interrupt)
	<-ctx.Done()
}

func serveMetrics() {
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":2223", nil)
	if err != nil {
		log.Fatal(err)
	}
}
