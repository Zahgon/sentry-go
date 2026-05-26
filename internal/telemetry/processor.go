package telemetry

import (
	"context"
	"time"

	"github.com/getsentry/sentry-go/internal/protocol"
	"github.com/getsentry/sentry-go/internal/ratelimit"
	"github.com/getsentry/sentry-go/report"
)

// Processor is the top-level object that wraps the scheduler and buffers.
type Processor struct {
	scheduler *Scheduler
}

// NewProcessor creates a new Processor with the given configuration.
func NewProcessor(
	buffers map[ratelimit.Category]Buffer[protocol.TelemetryItem],
	transport protocol.TelemetryTransport,
	dsn *protocol.Dsn,
	sdkInfo func() *protocol.SdkInfo,
	recorder report.ClientReportRecorder,
) *Processor {
	_ = "STUB: not implemented"
	return nil
}

// Add adds a TelemetryItem to the appropriate buffer based on its category.
//
// The processor should call MakeSerializationSafe to eliminate any race on user mutable fields,
// since the serialization happens on a background goroutine.
func (b *Processor) Add(item protocol.TelemetryItem) bool { _ = "STUB: not implemented"; return false }

// Flush forces all buffers to flush within the given timeout.
func (b *Processor) Flush(timeout time.Duration) bool { _ = "STUB: not implemented"; return false }

// FlushWithContext flushes with a custom context for cancellation.
func (b *Processor) FlushWithContext(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

// Close stops the buffer, flushes remaining data, and releases resources.
func (b *Processor) Close(timeout time.Duration) { _ = "STUB: not implemented"; return }
