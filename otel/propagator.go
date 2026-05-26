package sentryotel

import (
	"context"

	"go.opentelemetry.io/otel/propagation"
)

type sentryPropagator struct{}

// NewSentryPropagator creates a propagator for Sentry trace headers.
//
// Deprecated: This propagator depends on the span-processor integration and does
// not work for OTLP-based setups. Prefer the standard OpenTelemetry propagators.
// Will be removed in 0.47.0.
//
// For more details: https://github.com/getsentry/sentry-go/issues/1221
func NewSentryPropagator() propagation.TextMapPropagator {
	_ = "STUB: not implemented"
	return *new(propagation.TextMapPropagator)
}

// Inject sets Sentry-related values from the Context into the carrier.
//
// https://opentelemetry.io/docs/reference/specification/context/api-propagators/#inject
func (p sentryPropagator) Inject(ctx context.Context, carrier propagation.TextMapCarrier) {
	_ = "STUB: not implemented"
	return
}

// Propagate sentry-trace header

// No span => propagate the incoming sentry-trace header, if exists

// Sentry span exists => generate "sentry-trace" from it

// Propagate baggage header

// FIXME(anton): We're basically reparsing the header again, because in sentry-go
// we currently don't expose a method to get only DSC or its baggage (only a string).
// This is not optimal and we should consider other approaches.

// Merge the baggage values

// Extract reads cross-cutting concerns from the carrier into a Context.
//
// https://opentelemetry.io/docs/reference/specification/context/api-propagators/#extract
func (p sentryPropagator) Extract(ctx context.Context, carrier propagation.TextMapCarrier) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// Save traceParentContext because we'll at least need to know the original "sampled"
// value in the span processor.

// Preserve the original baggage

// The following cases should be already covered below:
// * We can extract a valid dynamic sampling context (DSC) from the baggage
// * No baggage header is present
// * No Sentry-related values are present
// * We cannot parse the baggage header for whatever reason

// If there are any errors, create a new non-frozen one.

// Fields returns a list of fields that will be used by the propagator.
//
// https://opentelemetry.io/docs/reference/specification/context/api-propagators/#fields
func (p sentryPropagator) Fields() []string { _ = "STUB: not implemented"; return nil }
