package sentryotel

import (
	"sync/atomic"

	"github.com/getsentry/sentry-go"
	"github.com/getsentry/sentry-go/internal/util"
	otelTrace "go.opentelemetry.io/otel/trace"
)

// TransactionEntry holds a reference to the root transaction span and
// tracks the number of active spans belonging to this trace.
//
// Deprecated: Prefer OTLP export via sentryotlp.NewTraceExporter.
// Will be removed in 0.47.0 alongside [NewSentrySpanProcessor].
type TransactionEntry struct {
	root        *sentry.Span
	activeCount atomic.Int64
	// spans holds active (not yet finished) spans for Get lookups.
	spans util.SyncMap[otelTrace.SpanID, *sentry.Span]
	// knownSpanIDs tracks all span IDs ever added to this transaction.
	knownSpanIDs util.SyncMap[otelTrace.SpanID, struct{}]
}

// HasSpan returns true if the given spanID was ever part of this transaction.
func (te *TransactionEntry) HasSpan(spanID otelTrace.SpanID) bool {
	_ = "STUB: not implemented"
	return false
}

// SentrySpanMap is a mapping between OpenTelemetry spans and Sentry spans.
// It stores spans per transaction for lookup by the propagator and event processor,
// and manages transaction entries for creating child spans via the shared spanRecorder.
//
// Deprecated: Prefer OTLP export via sentryotlp.NewTraceExporter.
// Will be removed in 0.47.0 alongside [NewSentrySpanProcessor].
type SentrySpanMap struct {
	transactions util.SyncMap[otelTrace.TraceID, *TransactionEntry]
}

// Get returns the current sentry.Span associated with the given OTel traceID and spanID.
func (ssm *SentrySpanMap) Get(traceID otelTrace.TraceID, spanID otelTrace.SpanID) (*sentry.Span, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// GetTransaction returns the transaction information for the given OTel traceID.
func (ssm *SentrySpanMap) GetTransaction(traceID otelTrace.TraceID) (*TransactionEntry, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Set stores the span and transaction information on the map. It handles both root and child spans automatically.
//
// If there is a cache miss on the given traceID, a transaction entry is created. Subsequent calls for the same traceID
// just increment the active span count and store the span in the entry.
func (ssm *SentrySpanMap) Set(spanID otelTrace.SpanID, span *sentry.Span, traceID otelTrace.TraceID) {
	_ = "STUB: not implemented"
	return
}

// MarkFinished removes a span from the active set and decrements the transaction's active count.
// When the count reaches zero, the transaction entry is removed.
// The span ID is kept in knownSpanIDs so that HasSpan continues to work for child span creation.
func (ssm *SentrySpanMap) MarkFinished(spanID otelTrace.SpanID, traceID otelTrace.TraceID) {
	_ = "STUB: not implemented"
	return
}

// CompareAndSwap(CAS) is used to prevent a race between Set and MarkFinished.
// The race has two windows:
// 1. MarkFinished decremented activeCount to 0 but hasn't CAS'd yet -> Set Adds(1), and CAS fails keeping the
// transaction, since we just added a new span.
// 2. MarkFinished already CAS'd -> Set will store on the transaction marked for deletion (better than
// creating a new orphaned span).

// Clear removes all spans stored on the map.
func (ssm *SentrySpanMap) Clear() { _ = "STUB: not implemented"; return }

// Len returns the number of spans on the map.
//
// This should only be used in tests, since computing the map length is fairly expensive.
func (ssm *SentrySpanMap) Len() int { _ = "STUB: not implemented"; return 0 }

var sentrySpanMap = SentrySpanMap{}
