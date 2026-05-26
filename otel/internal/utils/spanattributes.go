package utils

import (
	"github.com/getsentry/sentry-go"
	otelSdkTrace "go.opentelemetry.io/otel/sdk/trace"
)

type SpanAttributes struct {
	Op          string
	Description string
	Source      sentry.TransactionSource
}

func ParseSpanAttributes(s otelSdkTrace.ReadOnlySpan) SpanAttributes {
	_ = "STUB: not implemented"
	return *new(SpanAttributes)
}

// TODO(michi) Check if this works for AWS Lambda and such.

// becomes "default" in Relay

func descriptionForDbSystem(s otelSdkTrace.ReadOnlySpan) SpanAttributes {
	_ = "STUB: not implemented"
	return *new(SpanAttributes)
}

// TODO(michi)
// Note: The value may be sanitized to exclude sensitive information.
// See: https://pkg.go.dev/go.opentelemetry.io/otel/semconv/v1.12.0

func descriptionForHTTPMethod(s otelSdkTrace.ReadOnlySpan) SpanAttributes {
	_ = "STUB: not implemented"
	return *new(SpanAttributes)
}

// Prefer httpRoute if available

// Do not include the query and fragment parts

// This is normally the HTTP-client case

// Do not include the query and fragment parts

// Ex. description="GET /api/users".

// If `httpPath` is a root path, then we can categorize the transaction source as route.
