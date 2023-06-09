// Package tracer is implements an adapter to talks low-level trace observability.
// # This manifest was generated by ymir. DO NOT EDIT.
package tracer

import (
	"net/http"

	"go.opentelemetry.io/otel/propagation"
	otelTrace "go.opentelemetry.io/otel/trace"
)

// config is used to configure the mux middleware.
type config struct {
	TracerProvider    otelTrace.TracerProvider
	Propagators       propagation.TextMapPropagator
	spanNameFormatter func(string, *http.Request) string
}

// Option specifies instrumentation configuration options.
type Option interface {
	apply(*config)
}

type optionFunc func(*config)

func (o optionFunc) apply(c *config) {
	o(c)
}

// WithPropagators specifies propagators to use for extracting
// information from the HTTP requests. If none are specified, global
// ones will be used.
func WithPropagators(propagators propagation.TextMapPropagator) Option {
	return optionFunc(func(cfg *config) {
		if propagators != nil {
			cfg.Propagators = propagators
		}
	})
}

// WithTracerProvider specifies a tracer provider to use for creating a tracer.
// If none is specified, the global provider is used.
func WithTracerProvider(provider otelTrace.TracerProvider) Option {
	return optionFunc(func(cfg *config) {
		if provider != nil {
			cfg.TracerProvider = provider
		}
	})
}

// WithSpanNameFormatter specifies a function to use for generating a custom span
// name. By default, the route name (path template or regexp) is used. The route
// name is provided, so you can use it in the span name without needing to
// duplicate the logic for extracting it from the request.
func WithSpanNameFormatter(fn func(routeName string, r *http.Request) string) Option {
	return optionFunc(func(cfg *config) {
		cfg.spanNameFormatter = fn
	})
}
