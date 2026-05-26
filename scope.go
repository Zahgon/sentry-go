package sentry

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"sync"

	"github.com/getsentry/sentry-go/attribute"
)

// Scope holds contextual data for the current scope.
//
// The scope is an object that can cloned efficiently and stores data that is
// locally relevant to an event. For instance the scope will hold recorded
// breadcrumbs and similar information.
//
// The scope can be interacted with in two ways. First, the scope is routinely
// updated with information by functions such as AddBreadcrumb which will modify
// the current scope. Second, the current scope can be configured through the
// ConfigureScope function or Hub method of the same name.
//
// The scope is meant to be modified but not inspected directly. When preparing
// an event for reporting, the current client adds information from the current
// scope into the event.
type Scope struct {
	mu          sync.RWMutex
	attributes  map[string]attribute.Value
	breadcrumbs []*Breadcrumb
	attachments []*Attachment
	user        User
	tags        map[string]string
	contexts    map[string]Context
	fingerprint []string
	level       Level
	request     *http.Request
	// requestBody holds a reference to the original request.Body.
	requestBody interface {
		// Bytes returns bytes from the original body, lazily buffered as the
		// original body is read.
		Bytes() []byte
		// Overflow returns true if the body is larger than the maximum buffer
		// size.
		Overflow() bool
	}
	eventProcessors []EventProcessor

	propagationContext PropagationContext
	span               *Span
}

// NewScope creates a new Scope.
func NewScope() *Scope { _ = "STUB: not implemented"; return nil }

// AddBreadcrumb adds new breadcrumb to the current scope
// and optionally throws the old one if limit is reached.
func (scope *Scope) AddBreadcrumb(breadcrumb *Breadcrumb, limit int) {
	_ = "STUB: not implemented"
	return
}

// ClearBreadcrumbs clears all breadcrumbs from the current scope.
func (scope *Scope) ClearBreadcrumbs() { _ = "STUB: not implemented"; return }

// AddAttachment adds new attachment to the current scope.
func (scope *Scope) AddAttachment(attachment *Attachment) { _ = "STUB: not implemented"; return }

// ClearAttachments clears all attachments from the current scope.
func (scope *Scope) ClearAttachments() { _ = "STUB: not implemented"; return }

// SetUser sets the user for the current scope.
func (scope *Scope) SetUser(user User) { _ = "STUB: not implemented"; return }

// SetRequest sets the request for the current scope.
func (scope *Scope) SetRequest(r *http.Request) { _ = "STUB: not implemented"; return }

// Don't buffer request body if we know it is oversized.

// Don't buffer if there is no body.

// SetRequestBody sets the request body for the current scope.
//
// This method should only be called when the body bytes are already available
// in memory. Typically, the request body is buffered lazily from the
// Request.Body from SetRequest.
func (scope *Scope) SetRequestBody(b []byte) { _ = "STUB: not implemented"; return }

// maxRequestBodyBytes is the default maximum request body size to send to
// Sentry.
const maxRequestBodyBytes = 10 * 1024

// A limitedBuffer is like a bytes.Buffer, but limited to store at most Capacity
// bytes. Any writes past the capacity are silently discarded, similar to
// io.Discard.
type limitedBuffer struct {
	Capacity int

	bytes.Buffer
	overflow bool
}

// Write implements io.Writer.
func (b *limitedBuffer) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	// Silently ignore writes after overflow.
	return 0, nil
}

// Overflow returns true if the limitedBuffer discarded bytes written to it.
func (b *limitedBuffer) Overflow() bool {
	_ = "STUB: not implemented"

	// readCloser combines an io.Reader and an io.Closer to implement io.ReadCloser.
	return false
}

type readCloser struct {
	io.Reader
	io.Closer
}

// SetAttributes adds attributes to the current scope.
func (scope *Scope) SetAttributes(attrs ...attribute.Builder) { _ = "STUB: not implemented"; return }

// RemoveAttribute removes an attribute from the current scope.
func (scope *Scope) RemoveAttribute(key string) { _ = "STUB: not implemented"; return }

// SetTag adds a tag to the current scope.
func (scope *Scope) SetTag(key, value string) { _ = "STUB: not implemented"; return }

// SetTags assigns multiple tags to the current scope.
func (scope *Scope) SetTags(tags map[string]string) { _ = "STUB: not implemented"; return }

// RemoveTag removes a tag from the current scope.
func (scope *Scope) RemoveTag(key string) { _ = "STUB: not implemented"; return }

// SetContext adds a context to the current scope.
func (scope *Scope) SetContext(key string, value Context) { _ = "STUB: not implemented"; return }

// SetContexts assigns multiple contexts to the current scope.
func (scope *Scope) SetContexts(contexts map[string]Context) { _ = "STUB: not implemented"; return }

// RemoveContext removes a context from the current scope.
func (scope *Scope) RemoveContext(key string) { _ = "STUB: not implemented"; return }

// SetFingerprint sets new fingerprint for the current scope.
func (scope *Scope) SetFingerprint(fingerprint []string) { _ = "STUB: not implemented"; return }

// SetLevel sets new level for the current scope.
func (scope *Scope) SetLevel(level Level) { _ = "STUB: not implemented"; return }

// SetPropagationContext sets the propagation context for the current scope.
func (scope *Scope) SetPropagationContext(propagationContext PropagationContext) {
	_ = "STUB: not implemented"
	return
}

func (scope *Scope) propagationContextSnapshot() PropagationContext {
	_ = "STUB: not implemented"
	return *new(PropagationContext)
}

// GetSpan returns the span from the current scope.
func (scope *Scope) GetSpan() *Span { _ = "STUB: not implemented"; return nil }

// SetSpan sets a span for the current scope.
func (scope *Scope) SetSpan(span *Span) { _ = "STUB: not implemented"; return }

// Clone returns a copy of the current scope with all data copied over.
func (scope *Scope) Clone() *Scope { _ = "STUB: not implemented"; return nil }

// Clear removes the data from the current scope. Not safe for concurrent use.
func (scope *Scope) Clear() { _ = "STUB: not implemented"; return }

// AddEventProcessor adds an event processor to the current scope.
func (scope *Scope) AddEventProcessor(processor EventProcessor) { _ = "STUB: not implemented"; return }

// ApplyToEvent takes the data from the current scope and attaches it to the event.
func (scope *Scope) ApplyToEvent(event *Event, hint *EventHint, client *Client) *Event {
	_ = "STUB: not implemented" //nolint:gocyclo
	return nil
}

// Do not override trace context of
// transactions, otherwise it breaks the
// transaction event representation.
// For error events, the trace context is used
// to link errors and traces/spans in Sentry.

// Ensure we are not overwriting event fields

// If an external trace resolver is registered (e.g. OTel), override
// trace/span IDs from the hint context or the scope's request context.

// NOTE: The SDK does not attempt to send partial request body data.
//
// The reason being that Sentry's ingest pipeline and UI are optimized
// to show structured data. Additionally, tooling around PII scrubbing
// relies on structured data; truncated request bodies would create
// invalid payloads that are more prone to leaking PII data.
//
// Users can still send more data along their events if they want to,
// for example using Event.Contexts.

// cloneContext returns a new context with keys and values copied from the passed one.
//
// Note: a new Context (map) is returned, but the function does NOT do
// a proper deep copy: if some context values are pointer types (e.g. maps),
// they won't be properly copied.
func cloneContext(c Context) Context { _ = "STUB: not implemented"; return *new(Context) }

func (scope *Scope) populateAttrs(attrs map[string]attribute.Value) {
	_ = "STUB: not implemented"
	return
}

// Add user-related attributes

// hubFromContexts is a helper to return the first hub found in the given contexts.
func hubFromContexts(ctxs ...context.Context) *Hub { _ = "STUB: not implemented"; return nil }

// resolveTrace resolves trace ID and span ID from the given scope and contexts.
//
// The resolution order follows a most-specific-to-least-specific pattern:
//  1. If an external trace resolver was registered (eg. OTel), we prioritise trace context
//     information from that
//  2. Check for span directly in contexts (SpanFromContext) - this is the most specific
//     source as it represents a span explicitly attached to the current operation's context
//  3. Check scope's span - provides access to span set on the hub's scope
//  4. Fall back to scope's propagation context trace ID
//
// This ordering ensures we always use the most contextually relevant tracing information.
// For example, if a specific span is active for an operation, we use that span's trace/span IDs
// rather than accidentally using a different span that might be set on the hub's scope.
func resolveTrace(scope *Scope, client *Client, ctxs ...context.Context) (traceID TraceID, spanID SpanID) {
	_ = "STUB: not implemented"
	return *new(TraceID), *new(SpanID)
}
