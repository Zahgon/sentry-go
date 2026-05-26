package report

import (
	"sync"
	"sync/atomic"

	"github.com/getsentry/sentry-go/internal/protocol"
	"github.com/getsentry/sentry-go/internal/ratelimit"
)

// Aggregator collects discarded event outcomes for client reports.
// Uses atomic operations to be safe for concurrent use.
type Aggregator struct {
	mu       sync.Mutex
	outcomes map[OutcomeKey]*atomic.Int64
}

// NewAggregator creates a new client report Aggregator.
func NewAggregator() *Aggregator { _ = "STUB: not implemented"; return nil }

// Record records a discarded event outcome.
func (a *Aggregator) Record(reason DiscardReason, category ratelimit.Category, quantity int64) {
	_ = "STUB: not implemented"
	return
}

// RecordOne is a helper method to record one discarded event outcome.
func (a *Aggregator) RecordOne(reason DiscardReason, category ratelimit.Category) {
	_ = "STUB: not implemented"
	return
}

// TakeReport atomically takes all accumulated outcomes and returns a ClientReport.
func (a *Aggregator) TakeReport() *ClientReport { _ = "STUB: not implemented"; return nil }

// Clear empty counters to prevent unbounded growth

// RecordForEnvelope records client report outcomes for all items in the envelope.
// It inspects envelope item headers to derive categories, span counts, and log byte sizes.
func (a *Aggregator) RecordForEnvelope(reason DiscardReason, envelope *protocol.Envelope) {
	_ = "STUB: not implemented"
	return
}

// Skip — not reportable categories

// RecordItem records outcomes for a telemetry item, including supplementary
// categories (span outcomes for transactions, byte size for logs).
func (a *Aggregator) RecordItem(reason DiscardReason, item ReportableItem) {
	_ = "STUB: not implemented"
	return
}

// Span outcomes for transactions

// Byte size outcomes for logs

// AttachToEnvelope adds a client report to the envelope if the Aggregator has outcomes available.
func (a *Aggregator) AttachToEnvelope(envelope *protocol.Envelope) {
	_ = "STUB: not implemented"
	return
}
