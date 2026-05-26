package util

import (
	"net/http"

	"github.com/getsentry/sentry-go/internal/protocol"
	"github.com/getsentry/sentry-go/internal/ratelimit"
)

// MaxDrainResponseBytes is the maximum number of bytes that transport
// implementations will read from response bodies when draining them.
const MaxDrainResponseBytes = 16 << 10

// HandleHTTPResponse is a helper method that reads the HTTP response and handles debug output.
func HandleHTTPResponse(response *http.Response, identifier string) bool {
	_ = "STUB: not implemented"
	return false
}

// EnvelopeIdentifier returns a human-readable identifier for the event to be used in log messages.
// Format: "<description> [<event-id>]".
func EnvelopeIdentifier(envelope *protocol.Envelope) string { _ = "STUB: not implemented"; return "" }

// we don't currently support mixed envelope types, so all event types would have the same type.

// SendResult holds the outcome of an HTTP request sent to Sentry.
type SendResult struct {
	Success    bool
	StatusCode int
	Limits     ratelimit.Map
}

// IsSendError returns true if the response indicates a server/client error that should be recorded
// as send_error for client report outcomes (non-429 failures).
func (r *SendResult) IsSendError() bool { _ = "STUB: not implemented"; return false }

// DoSendRequest executes an HTTP request, handles response logging, extracts rate limits, and
// drains+closes the response body.
func DoSendRequest(client *http.Client, request *http.Request, identifier string) (*SendResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // G704: request is constructed internally, not from user input

// Drain body up to a limit and close it, allowing the transport to reuse TCP connections.
