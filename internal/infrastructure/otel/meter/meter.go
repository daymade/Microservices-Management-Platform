package meter

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.25.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Config struct {
	AppName           string
	Env               string
	CollectorEndpoint string
}

const (
	envKey      = "environment"
	dialTimeout = 5 * time.Second
)

// InitMeter 初始化带有 OTLP 和 Prometheus 导出器的度量提供器
func InitMeter(ctx context.Context, config Config) (*metric.MeterProvider, error) {
	ctx, cancel := context.WithTimeout(ctx, dialTimeout)
	defer cancel()

	otlpExporter, err := initOTLPExporter(ctx, config.CollectorEndpoint)
	if err != nil {
		return initPrometheusOnly(config)
	}

	promExporter, err := prometheus.New()
	if err != nil {
		return nil, fmt.Errorf("创建 Prometheus 导出器失败: %w", err)
	}

	mp := createMeterProvider(config, metric.NewPeriodicReader(otlpExporter), promExporter)
	otel.SetMeterProvider(mp)
	return mp, nil
}

// initOTLPExporter 初始化OTLP导出器
func initOTLPExporter(ctx context.Context, endpoint string) (*otlpmetricgrpc.Exporter, error) {
	conn, err := grpc.DialContext(ctx, endpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("连接 OpenTelemetry Collector 失败: %w", err)
	}
	defer conn.Close()

	return otlpmetricgrpc.New(ctx,
		otlpmetricgrpc.WithEndpoint(endpoint),
		otlpmetricgrpc.WithInsecure(),
	)
}

// initPrometheusOnly initializes the metric provider with only the Prometheus exporter.
func initPrometheusOnly(config Config) (*metric.MeterProvider, error) {
	promExporter, err := prometheus.New()
	if err != nil {
		return nil, fmt.Errorf("failed to create Prometheus exporter: %w", err)
	}

	mp := createMeterProvider(config, promExporter)
	otel.SetMeterProvider(mp)
	return mp, nil
}

// createMeterProvider creates a new MeterProvider with the given readers.
func createMeterProvider(config Config, readers ...metric.Reader) *metric.MeterProvider {
	opts := []metric.Option{
		metric.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(config.AppName),
			attribute.String(envKey, config.Env),
		)),
	}

	for _, reader := range readers {
		opts = append(opts, metric.WithReader(reader))
	}

	return metric.NewMeterProvider(opts...)
}
