package sentry

import (
	"context"
	"encoding/json"
	"net/http"
	"regexp"
	"sync"
	"time"
)

const (
	SentryTraceHeader   = "sentry-trace"
	SentryBaggageHeader = "baggage"
	TraceparentHeader   = "traceparent"
)

// SpanOrigin indicates what created a trace or a span. See: https://develop.sentry.dev/sdk/performance/trace-origin/
type SpanOrigin string

const (
	SpanOriginManual   = "manual"
	SpanOriginEcho     = "auto.http.echo"
	SpanOriginFastHTTP = "auto.http.fasthttp"
	SpanOriginFiber    = "auto.http.fiber"
	SpanOriginGin      = "auto.http.gin"
	SpanOriginStdLib   = "auto.http.stdlib"
	SpanOriginIris     = "auto.http.iris"
	SpanOriginNegroni  = "auto.http.negroni"
	SpanOriginGrpc     = "auto.rpc.grpc"
)

// A Span is the building block of a Sentry transaction. Spans build up a tree
// structure of timed operations. The span tree makes up a transaction event
// that is sent to Sentry when the root span is finished.
//
// Spans must be started with either StartSpan or Span.StartChild.
type Span struct { //nolint: maligned // prefer readability over optimal memory layout (see note below *)
	TraceID      TraceID                `json:"trace_id"`
	SpanID       SpanID                 `json:"span_id"`
	ParentSpanID SpanID                 `json:"parent_span_id,omitzero"`
	Name         string                 `json:"name,omitempty"`
	Op           string                 `json:"op,omitempty"`
	Description  string                 `json:"description,omitempty"`
	Status       SpanStatus             `json:"status,omitempty"`
	Tags         map[string]string      `json:"tags,omitempty"`
	StartTime    time.Time              `json:"start_timestamp,omitzero"`
	EndTime      time.Time              `json:"timestamp,omitzero"`
	Data         map[string]interface{} `json:"data,omitempty"`
	Sampled      Sampled                `json:"-"`
	Source       TransactionSource      `json:"-"`
	Origin       SpanOrigin             `json:"origin,omitempty"`

	// mu protects concurrent writes to map fields
	mu sync.RWMutex
	// sample rate the span was sampled with.
	sampleRate float64
	// ctx is the context where the span was started. Always non-nil.
	ctx context.Context
	// Dynamic Sampling context
	dynamicSamplingContext DynamicSamplingContext
	// parent refers to the immediate local parent span. A remote parent span is
	// only referenced by setting ParentSpanID.
	parent *Span
	// recorder stores all spans in a transaction. Guaranteed to be non-nil.
	recorder *spanRecorder
	// span context, can only be set on transactions
	contexts map[string]Context
	// a Once instance to make sure that Finish() is only called once.
	finishOnce sync.Once
	// explicitSampled is a flag for configuring sampling by using `WithSpanSampled` option.
	explicitSampled Sampled
	// Pre-serialized copies of mutable fields, set by MakeSerializationSafe.
	serializedTags    json.RawMessage
	serializedData    json.RawMessage
	serializationSafe bool
}

// TraceParentContext describes the context of a (remote) parent span.
//
// The context is normally extracted from a received "sentry-trace" header and
// used to initialize a new transaction.
//
// Note: the name might be not the best one. It was taken mostly to stay aligned
// with other SDKs, and it alludes to W3C "traceparent" header (https://www.w3.org/TR/trace-context/),
// which serves a similar purpose to "sentry-trace". We should eventually consider
// making this type internal-only and give it a better name.
type TraceParentContext struct {
	TraceID      TraceID
	ParentSpanID SpanID
	Sampled      Sampled
}

// (*) Note on maligned:
//
// We prefer readability over optimal memory layout. If we ever decide to
// reorder fields, we can use a tool:
//
// go run honnef.co/go/tools/cmd/structlayout -json . Span | go run honnef.co/go/tools/cmd/structlayout-optimize
//
// Other structs would deserve reordering as well, for example Event.

// TODO: make Span.Tags and Span.Data opaque types (struct{unexported []slice}).
// An opaque type allows us to add methods and make it more convenient to use
// than maps, because maps require careful nil checks to use properly or rely on
// explicit initialization for every span, even when there might be no
// tags/data. For Span.Data, must gracefully handle values that cannot be
// marshaled into JSON (see transport.go:getRequestBodyFromEvent).

// StartSpan starts a new span to describe an operation. The new span will be a
// child of the last span stored in ctx, if any.
//
// One or more options can be used to modify the span properties. Typically one
// option as a function literal is enough. Combining multiple options can be
// useful to define and reuse specific properties with named functions.
//
// Caller should call the Finish method on the span to mark its end. Finishing a
// root span sends the span and all of its children, recursively, as a
// transaction to Sentry.
func StartSpan(ctx context.Context, operation string, options ...SpanOption) *Span {
	_ = "STUB: not implemented"
	return nil
}

// defaults

// Only set the Source if this is a transaction

// Implementation note:
//
// While math/rand is ~2x faster than crypto/rand (exact
// difference depends on hardware / OS), crypto/rand is probably
// fast enough and a safer choice.
//
// For reference, OpenTelemetry [1] uses crypto/rand to seed
// math/rand. AFAICT this approach does not preserve the
// properties from crypto/rand that make it suitable for
// cryptography. While it might be debatable whether those
// properties are important for us here, again, we're taking the
// safer path.
//
// See [2a] & [2b] for a discussion of some of the properties we
// obtain by using crypto/rand and [3a] & [3b] for why we avoid
// math/rand.
//
// Because the math/rand seed has only 64 bits (int64), if the
// first thing we do after seeding an RNG is to read in a random
// TraceID, there are only 2^64 possible values. Compared to
// UUID v4 that have 122 random bits, there is a much greater
// chance of collision [4a] & [4b].
//
// [1]:  https://github.com/open-telemetry/opentelemetry-go/blob/958041ddf619a128/sdk/trace/trace.go#L25-L31
// [2a]: https://security.stackexchange.com/q/120352/246345
// [2b]: https://security.stackexchange.com/a/120365/246345
// [3a]: https://github.com/golang/go/issues/11871#issuecomment-126333686
// [3b]: https://github.com/golang/go/issues/11871#issuecomment-126357889
// [4a]: https://en.wikipedia.org/wiki/Universally_unique_identifier#Collisions
// [4b]: https://www.wolframalpha.com/input/?i=sqrt%282*2%5E64*ln%281%2F%281-0.5%29%29%29

// Apply options to override defaults.

// Finish sets the span's end time, unless already set. If the span is the root
// of a span tree, Finish sends the span tree to Sentry as a transaction.
//
// The logic is executed at most once per span, so that (incorrectly) calling it twice
// never double sends to Sentry.
func (s *Span) Finish() { _ = "STUB: not implemented"; return }

// Context returns the context containing the span.
func (s *Span) Context() context.Context {
	_ = "STUB: not implemented"

	// StartChild starts a new child span.
	//
	// The call span.StartChild(operation, options...) is a shortcut for
	// StartSpan(span.Context(), operation, options...).
	return *new(context.Context)
}

func (s *Span) StartChild(operation string, options ...SpanOption) *Span {
	_ = "STUB: not implemented"
	return nil
}

// makeSerializationSafe pre-serializes user mutable fields.
func (s *Span) makeSerializationSafe() { _ = "STUB: not implemented"; return }

// MarshalJSON emits the span using pre-serialized fields if available.
func (s *Span) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// SetTag sets a tag on the span. It is recommended to use SetTag instead of
// accessing the tags map directly as SetTag takes care of initializing the map
// when necessary.
func (s *Span) SetTag(name, value string) { _ = "STUB: not implemented"; return }

// SetData sets a data on the span. It is recommended to use SetData instead of
// accessing the data map directly as SetData takes care of initializing the map
// when necessary.
func (s *Span) SetData(name string, value interface{}) { _ = "STUB: not implemented"; return }

// SetContext sets a context on the span. It is recommended to use SetContext instead of
// accessing the contexts map directly as SetContext takes care of initializing the map
// when necessary.
func (s *Span) SetContext(key string, value Context) { _ = "STUB: not implemented"; return }

// IsTransaction checks if the given span is a transaction.
func (s *Span) IsTransaction() bool { _ = "STUB: not implemented"; return false }

// GetTransaction returns the transaction that contains this span.
//
// For transaction spans it returns itself. For spans that were created manually
// the method returns "nil".
func (s *Span) GetTransaction() *Span { _ = "STUB: not implemented"; return nil }

// This probably means that the Span was created manually (not via
// StartTransaction/StartSpan or StartChild).
// Return "nil" to indicate that it's not a normal situation.

// Same as above: manually created Span.

// TODO(tracing): maybe add shortcuts to get/set transaction name. Right now the
// transaction name is in the Scope, as it has existed there historically, prior
// to tracing.
//
// See Scope.Transaction() and Scope.SetTransaction().
//
// func (s *Span) TransactionName() string
// func (s *Span) SetTransactionName(name string)

// ToSentryTrace returns the serialized TraceParentContext from a transaction/span.
// Use this function to propagate the TraceParentContext to a downstream SDK,
// either as the value of the "sentry-trace" HTTP header, or as an html "sentry-trace" meta tag.
func (s *Span) ToSentryTrace() string {
	_ = "STUB: not implemented"
	// TODO(tracing): add instrumentation for outgoing HTTP requests using
	// ToSentryTrace.
	return ""
}

// ToTraceparent returns the W3C traceparent header value for the span.
func (s *Span) ToTraceparent() string { _ = "STUB: not implemented"; return "" }

// ToBaggage returns the serialized DynamicSamplingContext from a transaction.
// Use this function to propagate the DynamicSamplingContext to a downstream SDK,
// either as the value of the "baggage" HTTP header, or as an html "baggage" meta tag.
func (s *Span) ToBaggage() string { _ = "STUB: not implemented"; return "" }

// In case there is currently no frozen DynamicSamplingContext attached to the transaction,
// create one from the properties of the transaction.

// This will return a frozen DynamicSamplingContext.

// SetDynamicSamplingContext sets the given dynamic sampling context on the
// current transaction.
func (s *Span) SetDynamicSamplingContext(dsc DynamicSamplingContext) {
	_ = "STUB: not implemented"
	return
}

// shouldIgnoreStatusCode checks if the transaction should be ignored based on HTTP status code.
func (s *Span) shouldIgnoreStatusCode() bool { _ = "STUB: not implemented"; return false }

// Single status code

// Range of status codes [min, max]

// doFinish runs the actual Span.Finish() logic.
func (s *Span) doFinish() { _ = "STUB: not implemented"; return }

// we count the sampled spans from the transaction root. it is guaranteed that the whole transaction
// would be sampled

// TODO(tracing): add breadcrumbs
// (see https://github.com/getsentry/sentry-python/blob/f6f3525f8812f609/sentry_sdk/tracing.py#L372)

// sentryTracePattern matches either
//
//	TRACE_ID - SPAN_ID
//	[[:xdigit:]]{32}-[[:xdigit:]]{16}
//
// or
//
//	TRACE_ID - SPAN_ID - SAMPLED
//	[[:xdigit:]]{32}-[[:xdigit:]]{16}-[01]
var sentryTracePattern = regexp.MustCompile(`^([[:xdigit:]]{32})-([[:xdigit:]]{16})(?:-([01]))?$`)

// updateFromSentryTrace parses a sentry-trace HTTP header (as returned by
// ToSentryTrace) and updates fields of the span. If the header cannot be
// recognized as valid, the span is left unchanged. The returned value indicates
// whether the span was updated.
func (s *Span) updateFromSentryTrace(header []byte) (updated bool) {
	_ = "STUB: not implemented"
	return false
}

// no match

func (s *Span) updateFromBaggage(header []byte) { _ = "STUB: not implemented"; return }

func (s *Span) clientOptions() *ClientOptions { _ = "STUB: not implemented"; return nil }

func (s *Span) sample() Sampled { _ = "STUB: not implemented"; return *new(Sampled) }

// https://develop.sentry.dev/sdk/performance/#sampling
// #1 tracing is not enabled.

// #2 explicit sampling decision via StartSpan/StartTransaction options.

// Variant for non-transaction spans: they inherit the parent decision.
// Note: non-transaction should always have a parent, but we check both
// conditions anyway -- the first for semantic meaning, the second to
// avoid a nil pointer dereference.

// #3 use TracesSampler from ClientOptions.

// tracesSampler can update the sample_rate on frozen DSC

// #4 inherit parent decision.

// #5 use TracesSampleRate from ClientOptions.

// tracesSampleRate can update the sample_rate on frozen DSC

func (s *Span) toEvent() *Event { _ = "STUB: not implemented"; return nil }

// only transactions can be transformed into events

// Create and attach a DynamicSamplingContext to the transaction.
// If the DynamicSamplingContext is not frozen at this point, we can assume being head of trace.

// Make sure that the transaction source is valid

// traceContext returns a TraceContext snapshot for an active span.
//
// It needs to clone span Data to avoid holding any user mutable state.
func (s *Span) traceContext() *TraceContext { _ = "STUB: not implemented"; return nil }

// traceContextLockFree is the lock-free variant of traceContext.
//
// This is supposed to be used only when span Finish() was already called.
func (s *Span) traceContextLockFree() *TraceContext { _ = "STUB: not implemented"; return nil }

// spanRecorder stores the span tree. Guaranteed to be non-nil.
func (s *Span) spanRecorder() *spanRecorder {
	_ = "STUB: not implemented"

	// ParseTraceParentContext parses a sentry-trace header and builds a TraceParentContext from the
	// parsed values. If the header was parsed correctly, the second returned argument
	// ("valid") will be set to true, otherwise (e.g., empty or malformed header) it will
	// be false.
	return nil
}

func ParseTraceParentContext(header []byte) (traceParentContext TraceParentContext, valid bool) {
	_ = "STUB: not implemented"
	return *new(TraceParentContext), false
}

// TraceID identifies a trace.
type TraceID [16]byte

func (id TraceID) Hex() []byte { _ = "STUB: not implemented"; return nil }

func (id TraceID) String() string { _ = "STUB: not implemented"; return "" }

func (id TraceID) MarshalText() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// SpanID identifies a span.
		nil
}

type SpanID [8]byte

func (id SpanID) Hex() []byte { _ = "STUB: not implemented"; return nil }

func (id SpanID) String() string { _ = "STUB: not implemented"; return "" }

func (id SpanID) MarshalText() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// Zero values of TraceID and SpanID used for comparisons.
		nil
}

var (
	zeroTraceID TraceID
	zeroSpanID  SpanID
)

// Contains information about how the name of the transaction was determined.
type TransactionSource string

const (
	SourceCustom    TransactionSource = "custom"
	SourceURL       TransactionSource = "url"
	SourceRoute     TransactionSource = "route"
	SourceView      TransactionSource = "view"
	SourceComponent TransactionSource = "component"
	SourceTask      TransactionSource = "task"
)

// A set of all valid transaction sources.
var allTransactionSources = map[TransactionSource]struct{}{
	SourceCustom:    {},
	SourceURL:       {},
	SourceRoute:     {},
	SourceView:      {},
	SourceComponent: {},
	SourceTask:      {},
}

// isValid returns 'true' if the given transaction source is a valid
// source as recognized by the envelope protocol:
// https://develop.sentry.dev/sdk/event-payloads/transaction/#transaction-annotations
func (ts TransactionSource) isValid() bool { _ = "STUB: not implemented"; return false }

// SpanStatus is the status of a span.
type SpanStatus uint8

// Implementation note:
//
// In Relay (ingestion), the SpanStatus type is an enum used as
// Annotated<SpanStatus> when embedded in structs, making it effectively
// Option<SpanStatus>. It means the status is either null or one of the known
// string values.
//
// In Snuba (search), the SpanStatus is stored as an uint8 and defaulted to 2
// ("unknown") when not set. It means that Discover searches for
// `transaction.status:unknown` return both transactions/spans with status
// `null` or `"unknown"`. Searches for `transaction.status:""` return nothing.
//
// With that in mind, the Go SDK default is SpanStatusUndefined, which is
// null/omitted when serializing to JSON, but integrations may update the status
// automatically based on contextual information.

const (
	SpanStatusUndefined SpanStatus = iota
	SpanStatusOK
	SpanStatusCanceled
	SpanStatusUnknown
	SpanStatusInvalidArgument
	SpanStatusDeadlineExceeded
	SpanStatusNotFound
	SpanStatusAlreadyExists
	SpanStatusPermissionDenied
	SpanStatusResourceExhausted
	SpanStatusFailedPrecondition
	SpanStatusAborted
	SpanStatusOutOfRange
	SpanStatusUnimplemented
	SpanStatusInternalError
	SpanStatusUnavailable
	SpanStatusDataLoss
	SpanStatusUnauthenticated
	maxSpanStatus
)

var spanStatuses = [maxSpanStatus]string{
	"",
	"ok",
	"cancelled", // [sic]
	"unknown",
	"invalid_argument",
	"deadline_exceeded",
	"not_found",
	"already_exists",
	"permission_denied",
	"resource_exhausted",
	"failed_precondition",
	"aborted",
	"out_of_range",
	"unimplemented",
	"internal_error",
	"unavailable",
	"data_loss",
	"unauthenticated",
}

func (ss SpanStatus) String() string { _ = "STUB: not implemented"; return "" }

func (ss SpanStatus) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// A TraceContext carries information about an ongoing trace and is meant to be
// stored in Event.Contexts (as *TraceContext).
type TraceContext struct {
	TraceID      TraceID                `json:"trace_id"`
	SpanID       SpanID                 `json:"span_id"`
	ParentSpanID SpanID                 `json:"parent_span_id,omitzero"`
	Op           string                 `json:"op,omitempty"`
	Description  string                 `json:"description,omitempty"`
	Status       SpanStatus             `json:"status,omitempty"`
	Data         map[string]interface{} `json:"data,omitempty"`
}

func (tc TraceContext) Map() map[string]interface{} { _ = "STUB: not implemented"; return nil }

// Sampled signifies a sampling decision.
type Sampled int8

// The possible trace sampling decisions are: SampledFalse, SampledUndefined
// (default) and SampledTrue.
const (
	SampledFalse     Sampled = -1
	SampledUndefined Sampled = 0
	SampledTrue      Sampled = 1
)

func (s Sampled) String() string { _ = "STUB: not implemented"; return "" }

// Bool returns true if the sample decision is SampledTrue, false otherwise.
func (s Sampled) Bool() bool { _ = "STUB: not implemented"; return false }

// A SpanOption is a function that can modify the properties of a span.
type SpanOption func(s *Span)

// WithTransactionName option sets the name of the current transaction.
//
// A span tree has a single transaction name, therefore using this option when
// starting a span affects the span tree as a whole, potentially overwriting a
// name set previously.
func WithTransactionName(name string) SpanOption {
	_ = "STUB: not implemented"
	return *new(SpanOption)
}

// WithDescription sets the description of a span.
func WithDescription(description string) SpanOption {
	_ = "STUB: not implemented"
	return *new(SpanOption)
}

// WithOpName sets the operation name for a given span.
func WithOpName(name string) SpanOption { _ = "STUB: not implemented"; return *new(SpanOption) }

// WithTransactionSource sets the source of the transaction name.
//
// Note: if the transaction source is not a valid source (as described
// by the spec https://develop.sentry.dev/sdk/event-payloads/transaction/#transaction-annotations),
// it will be corrected to "custom" eventually, before the transaction is sent.
func WithTransactionSource(source TransactionSource) SpanOption {
	_ = "STUB: not implemented"
	return *new(SpanOption)
}

// WithSpanSampled updates the sampling flag for a given span.
func WithSpanSampled(sampled Sampled) SpanOption {
	_ = "STUB: not implemented"
	return *new(SpanOption)
}

// WithSpanOrigin sets the origin of the span.
func WithSpanOrigin(origin SpanOrigin) SpanOption {
	_ = "STUB: not implemented"
	return *new(SpanOption)
}

// ContinueTrace continues a trace based on traceparent and baggage values.
// If the SDK is configured with tracing enabled,
// this function returns populated SpanOption.
// In any other cases, it populates the propagation context on the scope.
func ContinueTrace(hub *Hub, traceparent, baggage string) SpanOption {
	_ = "STUB: not implemented"
	return *new(SpanOption)
}

// ContinueFromRequest returns a span option that updates the span to continue
// an existing trace. If it cannot detect an existing trace in the request, the
// span will be left unchanged.
//
// ContinueFromRequest is an alias for:
//
// ContinueFromHeaders(r.Header.Get(SentryTraceHeader), r.Header.Get(SentryBaggageHeader)).
func ContinueFromRequest(r *http.Request) SpanOption {
	_ = "STUB: not implemented"
	return *new(SpanOption)
}

// ContinueFromHeaders returns a span option that updates the span to continue
// an existing TraceID and propagates the Dynamic Sampling context.
func ContinueFromHeaders(trace, baggage string) SpanOption {
	_ = "STUB: not implemented"
	return *new(SpanOption)
}

// Parse baggage first to get org_id for comparison

// leave span unchanged → behaves as head of trace

// In case a sentry-trace header is present but there are no sentry-related
// values in the baggage, create an empty, frozen DynamicSamplingContext.

// ContinueFromTrace returns a span option that updates the span to continue
// an existing TraceID.
func ContinueFromTrace(trace string) SpanOption { _ = "STUB: not implemented"; return *new(SpanOption) }

// spanContextKey is used to store span values in contexts.
type spanContextKey struct{}

// TransactionFromContext returns the root span of the current transaction. It
// returns nil if no transaction is tracked in the context.
func TransactionFromContext(ctx context.Context) *Span { _ = "STUB: not implemented"; return nil }

// SpanFromContext returns the last span stored in the context, or nil if no span
// is set on the context.
func SpanFromContext(ctx context.Context) *Span { _ = "STUB: not implemented"; return nil }

// StartTransaction will create a transaction (root span) if there's no existing
// transaction in the context otherwise, it will return the existing transaction.
func StartTransaction(ctx context.Context, name string, options ...SpanOption) *Span {
	_ = "STUB: not implemented"
	return nil
}

// HTTPtoSpanStatus converts an HTTP status code to a SpanStatus.
func HTTPtoSpanStatus(code int) SpanStatus { _ = "STUB: not implemented"; return *new(SpanStatus) }

func shouldContinueTrace(client *Client, dsc DynamicSamplingContext) bool {
	_ = "STUB: not implemented"
	return false
}

// we reject non-matching orgs regardless of strict mode

// If strict mode is on, both must be present and match
