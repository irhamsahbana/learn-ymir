// Package infrastructure is implements an adapter to talks low-level modules.
// # This manifest was generated by ymir. DO NOT EDIT.
package infrastructure

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	stdoutTrace "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
	semConv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// TracerReturnFunc return type func with error.
type TracerReturnFunc func(ctx context.Context) error

// InitTracer create tracer provider from collector exporter.
func InitTracer(exporter sdkTrace.SpanExporter) TracerReturnFunc {
	tp := sdkTrace.NewTracerProvider(
		// Always be sure to batch in production.
		sdkTrace.WithBatcher(exporter),
		// Record information about this application in a Resource.
		sdkTrace.WithResource(resource.NewWithAttributes(
			semConv.SchemaURL,
			semConv.ServiceNameKey.String(Envs.App.ServiceName),
			attribute.String("version", "ver.1"),
		)),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{}, propagation.Baggage{},
		),
	)
	return func(ctx context.Context) error {
		return tp.Shutdown(ctx)
	}
}

// TraceExporter will push to the collector.
func TraceExporter(ctx context.Context,
	debug bool, addr string,
	timeout time.Duration) (sdkTrace.SpanExporter, error) {
	if debug {
		return stdoutTrace.New(stdoutTrace.WithPrettyPrint())
	}
	return GrpcTraceExporter(ctx, addr, timeout)
}

// GrpcTraceExporter is open connection to open telemetry collector.
func GrpcTraceExporter(ctx context.Context, addr string, timeout time.Duration) (*otlptrace.Exporter, error) {
	c, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	var (
		err  error
		conn *grpc.ClientConn
	)
	if Envs.Telemetry.CollectorEnable {
		conn, err = grpc.DialContext(c, addr,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
			grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()))
	} else {
		conn, err = grpc.DialContext(c, addr,
			grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	if err != nil {
		log.Error().Err(err).Msgf("exporter grpc: failed to connect %s", addr)
		return nil, err
	}
	return otlptracegrpc.New(c, otlptracegrpc.WithGRPCConn(conn))
}
