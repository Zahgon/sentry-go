package sentry

import (
	"sync"
)

// A spanRecorder stores a span tree that makes up a transaction. Safe for
// concurrent use. It is okay to add child spans from multiple goroutines.
type spanRecorder struct {
	mu           sync.Mutex
	spans        []*Span
	overflowOnce sync.Once
}

// record stores a span. The first stored span is assumed to be the root of a
// span tree.
func (r *spanRecorder) record(s *Span) { _ = "STUB: not implemented"; return }

// TODO(tracing): mark the transaction event in some way to
// communicate that spans were dropped.

// root returns the first recorded span. Returns nil if none have been recorded.
func (r *spanRecorder) root() *Span { _ = "STUB: not implemented"; return nil }

// children returns a list of all recorded spans, except the root. Returns nil
// if there are no children.
func (r *spanRecorder) children() []*Span { _ = "STUB: not implemented"; return nil }
