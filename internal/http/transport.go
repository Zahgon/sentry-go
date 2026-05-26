package http

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/getsentry/sentry-go/internal/protocol"
	"github.com/getsentry/sentry-go/internal/ratelimit"
	"github.com/getsentry/sentry-go/report"
)

const (
	apiVersion = 7

	defaultTimeout           = time.Second * 30
	defaultQueueSize         = 1000
	defaultClientReportsTick = time.Second * 30
)

var (
	ErrTransportQueueFull = errors.New("transport queue full")
	ErrTransportClosed    = errors.New("transport is closed")
	ErrEmptyEnvelope      = errors.New("empty envelope provided")
)

type TransportOptions struct {
	Dsn           string
	HTTPClient    *http.Client
	HTTPTransport http.RoundTripper
	HTTPProxy     string
	HTTPSProxy    string
	CaCerts       *x509.CertPool
	Recorder      report.ClientReportRecorder
	Provider      report.ClientReportProvider
	SdkInfo       func() *protocol.SdkInfo
}

func getProxyConfig(options TransportOptions) func(*http.Request) (*url.URL, error) {
	_ = "STUB: not implemented"
	return nil
}

func getTLSConfig(options TransportOptions) *tls.Config { _ = "STUB: not implemented"; return nil }

func getSentryRequestFromEnvelope(ctx context.Context, dsn *protocol.Dsn, envelope *protocol.Envelope) (r *http.Request, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func categoryFromEnvelope(envelope *protocol.Envelope) ratelimit.Category {
	_ = "STUB: not implemented"
	return *new(ratelimit.Category)
}

// SyncTransport is a blocking implementation of Transport.
//
// Clients using this transport will send requests to Sentry sequentially and
// block until a response is returned.
//
// The blocking behavior is useful in a limited set of use cases. For example,
// use it when deploying code to a Function as a Service ("Serverless")
// platform, where any work happening in a background goroutine is not
// guaranteed to execute.
//
// For most cases, prefer AsyncTransport.
type SyncTransport struct {
	dsn       *protocol.Dsn
	client    *http.Client
	transport http.RoundTripper
	recorder  report.ClientReportRecorder
	provider  report.ClientReportProvider
	sdkInfo   func() *protocol.SdkInfo

	mu     sync.Mutex
	limits ratelimit.Map

	Timeout time.Duration
}

func NewSyncTransport(options TransportOptions) protocol.TelemetryTransport {
	_ = "STUB: not implemented"
	return *new(protocol.TelemetryTransport)
}

func (t *SyncTransport) SendEnvelope(envelope *protocol.Envelope) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *SyncTransport) Close() { _ = "STUB: not implemented"; return }

func (t *SyncTransport) IsRateLimited(category ratelimit.Category) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *SyncTransport) HasCapacity() bool { _ = "STUB: not implemented"; return false }

func (t *SyncTransport) SendEnvelopeWithContext(ctx context.Context, envelope *protocol.Envelope) error {
	_ = "STUB: not implemented"
	return nil
}

// the sync transport needs to attach client reports when available

func (t *SyncTransport) Flush(_ time.Duration) bool { _ = "STUB: not implemented"; return false }

func (t *SyncTransport) FlushWithContext(_ context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *SyncTransport) disabled(c ratelimit.Category) bool {
	_ = "STUB: not implemented"
	return false
}

// AsyncTransport is the default, non-blocking, implementation of Transport.
//
// Clients using this transport will enqueue requests in a queue and return to
// the caller before any network communication has happened. Requests are sent
// to Sentry sequentially from a background goroutine.
type AsyncTransport struct {
	dsn       *protocol.Dsn
	client    *http.Client
	transport http.RoundTripper
	recorder  report.ClientReportRecorder
	provider  report.ClientReportProvider
	sdkInfo   func() *protocol.SdkInfo

	queue chan *protocol.Envelope

	mu     sync.RWMutex
	limits ratelimit.Map

	done chan struct{}
	wg   sync.WaitGroup

	flushRequest chan chan struct{}

	closeMu sync.RWMutex

	QueueSize int
	Timeout   time.Duration

	startOnce sync.Once
	closeOnce sync.Once
}

func NewAsyncTransport(options TransportOptions) protocol.TelemetryTransport {
	_ = "STUB: not implemented"
	return *new(protocol.TelemetryTransport)
}

func (t *AsyncTransport) start() { _ = "STUB: not implemented"; return }

// HasCapacity reports whether the async transport queue appears to have space
// for at least one more envelope. This is a best-effort, non-blocking check.
func (t *AsyncTransport) HasCapacity() bool { _ = "STUB: not implemented"; return false }

func (t *AsyncTransport) SendEnvelope(envelope *protocol.Envelope) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *AsyncTransport) Flush(timeout time.Duration) bool { _ = "STUB: not implemented"; return false }

func (t *AsyncTransport) FlushWithContext(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *AsyncTransport) Close() { _ = "STUB: not implemented"; return }

func (t *AsyncTransport) IsRateLimited(category ratelimit.Category) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *AsyncTransport) resolveSdkInfo() *protocol.SdkInfo { _ = "STUB: not implemented"; return nil }

func (t *AsyncTransport) worker() { _ = "STUB: not implemented"; return }

// sendClientReport sends a standalone envelope containing only a client report.
func (t *AsyncTransport) sendClientReport() { _ = "STUB: not implemented"; return }

func (t *AsyncTransport) drainQueue() { _ = "STUB: not implemented"; return }

func (t *AsyncTransport) sendEnvelopeHTTP(envelope *protocol.Envelope) bool {
	_ = "STUB: not implemented" //nolint: unparam
	return false
}

// attach to envelope after rate-limit check

func (t *AsyncTransport) isRateLimited(category ratelimit.Category) bool {
	_ = "STUB: not implemented"
	return false
}

// NoopTransport is a transport implementation that drops all events.
// Used internally when an empty or invalid DSN is provided.
type NoopTransport struct{}

func NewNoopTransport() *NoopTransport { _ = "STUB: not implemented"; return nil }

func (t *NoopTransport) SendEnvelope(_ *protocol.Envelope) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *NoopTransport) IsRateLimited(_ ratelimit.Category) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *NoopTransport) Flush(_ time.Duration) bool { _ = "STUB: not implemented"; return false }

func (t *NoopTransport) FlushWithContext(_ context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

func (t *NoopTransport) Close() {
	_ = "STUB: not implemented"
	// Nothing to close
	return
}

func (t *NoopTransport) HasCapacity() bool { _ = "STUB: not implemented"; return false }
