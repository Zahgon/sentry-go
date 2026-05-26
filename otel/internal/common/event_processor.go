package common

import (
	"context"

	"github.com/getsentry/sentry-go"
)

// NewEventProcessor creates a Sentry event processor that attaches OTel trace
// information from the active SpanContext to an error event.
func NewEventProcessor() sentry.EventProcessor {
	_ = "STUB: not implemented"
	return *new(sentry.EventProcessor)
}

func linkTraceContextToErrorEvent(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
	_ = "STUB: not implemented"
	return nil
}

// ResolveTraceContext returns Sentry trace/span IDs from the active OTel span in ctx.
func ResolveTraceContext(ctx context.Context) (sentry.TraceID, sentry.SpanID, bool) {
	_ = "STUB: not implemented"
	return *new(sentry.TraceID), *new(sentry.SpanID), false
}
