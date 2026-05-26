package sentryotel

import (
	"github.com/getsentry/sentry-go"
)

type integration struct{}

// NewOtelIntegration registers OpenTelemetry linking with Sentry.
//
// It links captured Sentry errors, logs, and metrics to the active
// OpenTelemetry trace when a context carrying an active OTel span is used.
func NewOtelIntegration() sentry.Integration {
	_ = "STUB: not implemented"
	return *new(sentry.Integration)
}

func (integration) Name() string { _ = "STUB: not implemented"; return "" }

func (integration) SetupOnce(client *sentry.Client) { _ = "STUB: not implemented"; return }
