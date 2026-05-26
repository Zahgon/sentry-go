package sentryotel

import (
	"context"

	"github.com/getsentry/sentry-go"
	otelSdkTrace "go.opentelemetry.io/otel/sdk/trace"
)

type sentrySpanProcessor struct{}

// Singleton instance of the Sentry span processor.
// At the moment we do not support multiple instances.
var sentrySpanProcessorInstance *sentrySpanProcessor

// NewSentrySpanProcessor creates an OpenTelemetry span processor that mirrors
// OTel spans into Sentry's native span model.
//
// Deprecated: Prefer OTLP export via sentryotlp.NewTraceExporter.
// For collector-based setups, use the standard OTel exporter and register
// sentryotel.NewOtelIntegration for linking.
// Will be removed in 0.47.0.
func NewSentrySpanProcessor() otelSdkTrace.SpanProcessor {
	_ = "STUB: not implemented"
	return *new(otelSdkTrace.SpanProcessor)
}

// https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/trace/sdk.md#onstart
func (ssp *sentrySpanProcessor) OnStart(parent context.Context, s otelSdkTrace.ReadWriteSpan) {
	_ = "STUB: not implemented"
	return
}

// https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/trace/sdk.md#onendspan
func (ssp *sentrySpanProcessor) OnEnd(s otelSdkTrace.ReadOnlySpan) {
	_ = "STUB: not implemented"
	return
}

// https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/trace/sdk.md#shutdown-1
func (ssp *sentrySpanProcessor) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil

	// Note: according to the spec, "Shutdown MUST include the effects of ForceFlush".
}

// https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/trace/sdk.md#forceflush-1
func (ssp *sentrySpanProcessor) ForceFlush(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func flushSpanProcessor(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// TODO(michi) should we make this configurable?

func getTraceParentContext(ctx context.Context) sentry.TraceParentContext {
	_ = "STUB: not implemented"
	return *new(sentry.TraceParentContext)
}

func updateTransactionWithOtelData(transaction *sentry.Span, s otelSdkTrace.ReadOnlySpan) {
	_ = "STUB: not implemented"
	// TODO(michi) This is crazy inefficient
	return
}

func updateSpanWithOtelData(span *sentry.Span, s otelSdkTrace.ReadOnlySpan) {
	_ = "STUB: not implemented"
	return
}
