package utils

import (
	"context"

	"go.opentelemetry.io/otel/sdk/trace"
)

func IsSentryRequestSpan(ctx context.Context, s trace.ReadOnlySpan) bool {
	_ = "STUB: not implemented"
	return false

	// TODO(michi): can we access the attribute directly?
}

func isSentryRequestURL(ctx context.Context, url string) bool {
	_ = "STUB: not implemented"
	return false
}
