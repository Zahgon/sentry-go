package report

import (
	"github.com/getsentry/sentry-go/internal/protocol"
	"github.com/getsentry/sentry-go/internal/ratelimit"
)

// ReportableItem is the minimal surface needed for client report accounting.
type ReportableItem interface {
	GetCategory() ratelimit.Category
}

// ClientReportRecorder is used by components that need to record lost/discarded events.
type ClientReportRecorder interface {
	Record(reason DiscardReason, category ratelimit.Category, quantity int64)
	RecordOne(reason DiscardReason, category ratelimit.Category)
	RecordForEnvelope(reason DiscardReason, envelope *protocol.Envelope)
	RecordItem(reason DiscardReason, item ReportableItem)
}

// ClientReportProvider is used by the single component responsible for sending client reports.
type ClientReportProvider interface {
	TakeReport() *ClientReport
	AttachToEnvelope(envelope *protocol.Envelope)
}

// noopRecorder is a no-op implementation of ClientReportRecorder.
type noopRecorder struct{}

func (noopRecorder) Record(DiscardReason, ratelimit.Category, int64) {
	_ = "STUB: not implemented"
	return
}
func (noopRecorder) RecordOne(DiscardReason, ratelimit.Category) { _ = "STUB: not implemented"; return }
func (noopRecorder) RecordForEnvelope(DiscardReason, *protocol.Envelope) {
	_ = "STUB: not implemented"
	return
}
func (noopRecorder) RecordItem(DiscardReason, ReportableItem) {
	_ = "STUB: not implemented"

	// noopProvider is a no-op implementation of ClientReportProvider.
	return
}

type noopProvider struct{}

func (noopProvider) TakeReport() *ClientReport { _ = "STUB: not implemented"; return nil }
func (noopProvider) AttachToEnvelope(_ *protocol.Envelope) {
	_ = "STUB: not implemented"

	// NoopRecorder returns a no-op ClientReportRecorder that silently discards all records.
	return
}

func NoopRecorder() ClientReportRecorder {
	_ = "STUB: not implemented"
	return *

	// NoopProvider returns a no-op ClientReportProvider that always returns nil reports.
	new(ClientReportRecorder)
}

func NoopProvider() ClientReportProvider {
	_ = "STUB: not implemented"
	return *new(ClientReportProvider)
}
