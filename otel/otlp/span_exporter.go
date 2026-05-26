package sentryotlp

import (
	"context"
	"crypto/tls"
	"net/url"
	"time"

	"github.com/getsentry/sentry-go"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

const apiVersion = "7"

// Option configures safe OTLP HTTP client behavior for the Sentry exporter.
// Target selection and auth remain derived from the DSN.
type Option func(*config)

type config struct {
	otlpOptions []otlptracehttp.Option
}

// WithCompression configures OTLP payload compression.
func WithCompression(compression otlptracehttp.Compression) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithTLSClientConfig configures the HTTP client's TLS settings.
func WithTLSClientConfig(tlsCfg *tls.Config) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTimeout configures the OTLP export timeout.
func WithTimeout(duration time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithRetry configures OTLP retry behavior.
func WithRetry(rc otlptracehttp.RetryConfig) Option { _ = "STUB: not implemented"; return *new(Option) }

type sentryOTLPExporter struct {
	inner sdktrace.SpanExporter
}

// NewTraceExporter creates a new SpanExporter that sends spans to Sentry via the OTLP HTTP protocol.
//
// The endpoint, URL path, headers, and HTTP/HTTPS mode are derived from the DSN.
func NewTraceExporter(ctx context.Context, dsn string, opts ...Option) (sdktrace.SpanExporter, error) {
	_ = "STUB: not implemented"
	return *new(sdktrace.SpanExporter), nil
}

func buildOTLPOptions(dsn string, opts ...otlptracehttp.Option) ([]otlptracehttp.Option, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func otlpTracesURL(dsn *sentry.Dsn) *url.URL { _ = "STUB: not implemented"; return nil }

// ExportSpans exports a batch of spans to Sentry via OTLP HTTP.
func (e *sentryOTLPExporter) ExportSpans(ctx context.Context, spans []sdktrace.ReadOnlySpan) error {
	_ = "STUB: not implemented"
	return nil
}

// Shutdown shuts down the exporter, flushing any remaining spans.
func (e *sentryOTLPExporter) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// sentryAuthHeaders builds the X-Sentry-Auth header map for OTLP requests.
func sentryAuthHeaders(dsn *sentry.Dsn) map[string]string { _ = "STUB: not implemented"; return nil }
