package testutils

import (
	"context"
	"sync"
	"time"

	"github.com/getsentry/sentry-go/internal/protocol"
	"github.com/getsentry/sentry-go/internal/ratelimit"
)

type MockTelemetryTransport struct {
	sentEnvelopes    []*protocol.Envelope
	rateLimited      map[string]bool
	sendError        error
	mu               sync.Mutex
	sendCount        int64
	rateLimitedCalls int64
	capacity         int
}

func (m *MockTelemetryTransport) SendEnvelope(envelope *protocol.Envelope) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockTelemetryTransport) IsRateLimited(category ratelimit.Category) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *MockTelemetryTransport) HasCapacity() bool { _ = "STUB: not implemented"; return false }

func (m *MockTelemetryTransport) Flush(_ time.Duration) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *MockTelemetryTransport) FlushWithContext(_ context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

func (m *MockTelemetryTransport) Configure(_ interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockTelemetryTransport) Close() { _ = "STUB: not implemented"; return }

func (m *MockTelemetryTransport) GetSentEnvelopes() []*protocol.Envelope {
	_ = "STUB: not implemented"
	return nil
}

func (m *MockTelemetryTransport) SetRateLimited(category string, limited bool) {
	_ = "STUB: not implemented"
	return
}

func (m *MockTelemetryTransport) GetSendCount() int64 { _ = "STUB: not implemented"; return 0 }

func (m *MockTelemetryTransport) GetRateLimitedCalls() int64 { _ = "STUB: not implemented"; return 0 }

func (m *MockTelemetryTransport) Reset() { _ = "STUB: not implemented"; return }
