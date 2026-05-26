// Package sentryhttpclient provides Sentry integration for Requests modules to enable distributed tracing between services.
// It is compatible with `net/http.RoundTripper`.
//
//	import sentryhttpclient "github.com/getsentry/sentry-go/httpclient"
//
//	roundTrippper := sentryhttpclient.NewSentryRoundTripper(nil, nil)
//	client := &http.Client{
//		Transport: roundTripper,
//	}
//
//	request, err := client.Do(request)
package sentryhttpclient

import (
	"net/http"
)

// SentryRoundTripTracerOption provides a specific type in which defines the option for SentryRoundTripper.
type SentryRoundTripTracerOption func(*SentryRoundTripper)

// WithTracePropagationTargets configures additional trace propagation targets URL for the RoundTripper.
// Does not support regex patterns.
func WithTracePropagationTargets(targets []string) SentryRoundTripTracerOption {
	_ = "STUB: not implemented"
	return *new(SentryRoundTripTracerOption)
}

// NewSentryRoundTripper provides a wrapper to existing http.RoundTripper to have required span data and trace headers for outgoing HTTP requests.
//
//   - If `nil` is passed to `originalRoundTripper`, it will use http.DefaultTransport instead.
func NewSentryRoundTripper(originalRoundTripper http.RoundTripper, opts ...SentryRoundTripTracerOption) http.RoundTripper {
	_ = "STUB: not implemented"
	return *new(http.RoundTripper)
}

// Configure trace propagation targets

// SentryRoundTripper provides a http.RoundTripper implementation for Sentry Requests module.
type SentryRoundTripper struct {
	originalRoundTripper http.RoundTripper

	propagateTraceparent    bool
	tracePropagationTargets []string
}

func (s *SentryRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	// Respect trace propagation targets
	return nil, nil
}

// Only create the `http.client` span only if there is a parent span.

// Always add `Baggage` and `Sentry-Trace` headers.
