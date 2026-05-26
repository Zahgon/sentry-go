package sentry

import (
	"context"
	"crypto/x509"
	"io"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/getsentry/sentry-go/internal/debuglog"
	"github.com/getsentry/sentry-go/internal/protocol"
	"github.com/getsentry/sentry-go/internal/telemetry"
	"github.com/getsentry/sentry-go/report"
)

// The identifier of the SDK.
const sdkIdentifier = "sentry.go"

const (
	// maxErrorDepth is the maximum number of errors reported in a chain of errors.
	// This protects the SDK from an arbitrarily long chain of wrapped errors.
	//
	// An additional consideration is that arguably reporting a long chain of errors
	// is of little use when debugging production errors with Sentry. The Sentry UI
	// is not optimized for long chains either. The top-level error together with a
	// stack trace is often the most useful information.
	maxErrorDepth = 100

	// defaultMaxSpans limits the default number of recorded spans per transaction. The limit is
	// meant to bound memory usage and prevent too large transaction events that
	// would be rejected by Sentry.
	defaultMaxSpans = 1000

	// defaultMaxBreadcrumbs is the default maximum number of breadcrumbs added to
	// an event. Can be overwritten with the MaxBreadcrumbs option.
	defaultMaxBreadcrumbs = 100
)

// hostname is the host name reported by the kernel. It is precomputed once to
// avoid syscalls when capturing events.
//
// The error is ignored because retrieving the host name is best-effort. If the
// error is non-nil, there is nothing to do other than retrying. We choose not
// to retry for now.
var hostname, _ = os.Hostname()

// lockedRand is a random number generator safe for concurrent use. Its API is
// intentionally limited and it is not meant as a full replacement for a
// rand.Rand.
type lockedRand struct {
	mu sync.Mutex
	r  *rand.Rand
}

// Float64 returns a pseudo-random number in [0.0,1.0).
func (r *lockedRand) Float64() float64 { _ = "STUB: not implemented"; return 0 }

// rng is the internal random number generator.
//
// We do not use the global functions from math/rand because, while they are
// safe for concurrent use, any package in a build could change the seed and
// affect the generated numbers, for instance making them deterministic. On the
// other hand, the source returned from rand.NewSource is not safe for
// concurrent use, so we need to couple its use with a sync.Mutex.
var rng = &lockedRand{
	// #nosec G404 -- We are fine using transparent, non-secure value here.
	r: rand.New(rand.NewSource(time.Now().UnixNano())),
}

// usageError is used to report to Sentry an SDK usage error.
//
// It is not exported because it is never returned by any function or method in
// the exported API.
type usageError struct {
	error
}

// DebugLogger is an instance of log.Logger that is used to provide debug information about running Sentry Client
// can be enabled by either using debuglog.SetOutput directly or with Debug client option.
var DebugLogger = debuglog.GetLogger()

// EventProcessor is a function that processes an event.
// Event processors are used to change an event before it is sent to Sentry.
type EventProcessor func(event *Event, hint *EventHint) *Event

// externalContextTraceResolver extracts trace and span IDs from an external context source.
//
// This is currently a workaround for extractring trace information from OTel SpanContext without
// needing the otel dependency on the root package.
type externalContextTraceResolver func(ctx context.Context) (traceID TraceID, spanID SpanID, ok bool)

// EventModifier is the interface that wraps the ApplyToEvent method.
//
// ApplyToEvent changes an event based on external data and/or
// an event hint.
type EventModifier interface {
	ApplyToEvent(event *Event, hint *EventHint, client *Client) *Event
}

var globalEventProcessors []EventProcessor

// AddGlobalEventProcessor adds processor to the global list of event
// processors. Global event processors apply to all events.
//
// AddGlobalEventProcessor is deprecated. Most users will prefer to initialize
// the SDK with Init and provide a ClientOptions.BeforeSend function or use
// Scope.AddEventProcessor instead.
func AddGlobalEventProcessor(processor EventProcessor) { _ = "STUB: not implemented"; return }

// Integration allows for registering a functions that modify or discard captured events.
type Integration interface {
	Name() string
	SetupOnce(client *Client)
}

// ClientOptions that configures a SDK Client.
type ClientOptions struct {
	// The DSN to use. If the DSN is not set, the client is effectively
	// disabled.
	Dsn string
	// In debug mode, the debug information is printed to stdout to help you
	// understand what sentry is doing.
	Debug bool
	// Configures whether SDK should generate and attach stacktraces to pure
	// capture message calls.
	AttachStacktrace bool
	// The sample rate for event submission in the range [0.0, 1.0]. By default,
	// all events are sent. Thus, as a historical special case, the sample rate
	// 0.0 is treated as if it was 1.0. To drop all events, set the DSN to the
	// empty string.
	SampleRate float64
	// Enable performance tracing.
	EnableTracing bool
	// The sample rate for sampling traces in the range [0.0, 1.0].
	TracesSampleRate float64
	// Used to customize the sampling of traces, overrides TracesSampleRate.
	TracesSampler TracesSampler
	// Control with URLs trace propagation should be enabled. Does not support regex patterns.
	TracePropagationTargets []string
	// PropagateTraceparent is used to control whether the W3C Trace Context HTTP traceparent header
	// is propagated on outgoing http requests.
	PropagateTraceparent bool
	// StrictTraceContinuation is used to control trace continuation from 3rd party services that happen to be
	// instrumented by Sentry.
	//
	// Enabling the option means that the SDK will require the org ids from baggage to match for continuing the trace.
	StrictTraceContinuation bool
	// OrgID configures the orgID used for trace propagation and features like StrictTraceContinuation.
	//
	// In most cases the orgID is already parsed from the DSN. This option should be used when non-standard Sentry DSNs
	// are used, such as self-hosted or when using a local Relay.
	OrgID uint64
	// List of regexp strings that will be used to match against event's message
	// and if applicable, caught errors type and value.
	// If the match is found, then a whole event will be dropped.
	IgnoreErrors []string
	// List of regexp strings that will be used to match against a transaction's
	// name.  If a match is found, then the transaction  will be dropped.
	IgnoreTransactions []string
	// If this flag is enabled, certain personally identifiable information (PII) is added by active integrations.
	// By default, no such data is sent.
	SendDefaultPII bool
	// BeforeSend is called before error events are sent to Sentry.
	// You can use it to mutate the event or return nil to discard it.
	BeforeSend func(event *Event, hint *EventHint) *Event
	// BeforeSendLong is called before log events are sent to Sentry.
	// You can use it to mutate the log event or return nil to discard it.
	BeforeSendLog func(event *Log) *Log
	// BeforeSendTransaction is called before transaction events are sent to Sentry.
	// Use it to mutate the transaction or return nil to discard the transaction.
	BeforeSendTransaction func(event *Event, hint *EventHint) *Event
	// Before breadcrumb add callback.
	BeforeBreadcrumb func(breadcrumb *Breadcrumb, hint *BreadcrumbHint) *Breadcrumb
	// BeforeSendMetric is called before metric events are sent to Sentry.
	// You can use it to mutate the metric or return nil to discard it.
	BeforeSendMetric func(metric *Metric) *Metric
	// Integrations to be installed on the current Client, receives default
	// integrations.
	Integrations func([]Integration) []Integration
	// io.Writer implementation that should be used with the Debug mode.
	DebugWriter io.Writer
	// The transport to use. Defaults to HTTPTransport.
	Transport Transport
	// The server name to be reported.
	ServerName string
	// The release to be sent with events.
	//
	// Some Sentry features are built around releases, and, thus, reporting
	// events with a non-empty release improves the product experience. See
	// https://docs.sentry.io/product/releases/.
	//
	// If Release is not set, the SDK will try to derive a default value
	// from environment variables or the Git repository in the working
	// directory.
	//
	// If you distribute a compiled binary, it is recommended to set the
	// Release value explicitly at build time. As an example, you can use:
	//
	// 	go build -ldflags='-X main.release=VALUE'
	//
	// That will set the value of a predeclared variable 'release' in the
	// 'main' package to 'VALUE'. Then, use that variable when initializing
	// the SDK:
	//
	// 	sentry.Init(ClientOptions{Release: release})
	//
	// See https://golang.org/cmd/go/ and https://golang.org/cmd/link/ for
	// the official documentation of -ldflags and -X, respectively.
	Release string
	// The dist to be sent with events.
	Dist string
	// The environment to be sent with events.
	Environment string
	// Maximum number of breadcrumbs
	// when MaxBreadcrumbs is negative then ignore breadcrumbs.
	MaxBreadcrumbs int
	// Maximum number of spans.
	//
	// See https://develop.sentry.dev/sdk/envelopes/#size-limits for size limits
	// applied during event ingestion. Events that exceed these limits might get dropped.
	MaxSpans int
	// An optional pointer to http.Client that will be used with a default
	// HTTPTransport. Using your own client will make HTTPTransport, HTTPProxy,
	// HTTPSProxy and CaCerts options ignored.
	HTTPClient *http.Client
	// An optional pointer to http.Transport that will be used with a default
	// HTTPTransport. Using your own transport will make HTTPProxy, HTTPSProxy
	// and CaCerts options ignored.
	HTTPTransport http.RoundTripper
	// An optional HTTP proxy to use.
	// This will default to the HTTP_PROXY environment variable.
	HTTPProxy string
	// An optional HTTPS proxy to use.
	// This will default to the HTTPS_PROXY environment variable.
	// HTTPS_PROXY takes precedence over HTTP_PROXY for https requests.
	HTTPSProxy string
	// An optional set of SSL certificates to use.
	CaCerts *x509.CertPool
	// MaxErrorDepth is the maximum number of errors reported in a chain of errors.
	// This protects the SDK from an arbitrarily long chain of wrapped errors.
	//
	// An additional consideration is that arguably reporting a long chain of errors
	// is of little use when debugging production errors with Sentry. The Sentry UI
	// is not optimized for long chains either. The top-level error together with a
	// stack trace is often the most useful information.
	MaxErrorDepth int
	// Default event tags. These are overridden by tags set on a scope.
	Tags map[string]string
	// DisableLogs controls whether logs should be emitted.
	// By default, logs are enabled. Set to true to disable log emission.
	DisableLogs bool
	// DisableMetrics controls when metrics should be emitted.
	DisableMetrics bool
	// DisableClientReports controls when client reports should be emitted.
	DisableClientReports bool
	// TraceIgnoreStatusCodes is a list of HTTP status codes that should not be traced.
	// Each element can be either:
	// - A single-element slice [code] for a specific status code
	// - A two-element slice [min, max] for a range of status codes (inclusive)
	// When an HTTP request results in a status code that matches any of these codes or ranges,
	// the transaction will not be sent to Sentry.
	//
	// Examples:
	//   [][]int{{404}}                           // ignore only status code 404
	//   [][]int{{400, 405}}                     // ignore status codes 400-405
	//   [][]int{{404}, {500}}                   // ignore status codes 404 and 500
	//   [][]int{{404}, {400, 405}, {500, 599}}  // ignore 404, range 400-405, and range 500-599
	//
	// By default, this ignores 404 status codes.
	//
	// IMPORTANT: to not ignore any status codes, the option should be an empty slice and not nil. The nil option is
	// used for defaulting to 404 ignores.
	TraceIgnoreStatusCodes [][]int
	// DisableTelemetryBuffer disables the telemetry buffer layer for prioritizing events and uses the old transport layer.
	DisableTelemetryBuffer bool
}

// Client is the underlying processor that is used by the main API and Hub
// instances. It must be created with NewClient.
type Client struct {
	mu                    sync.RWMutex
	options               ClientOptions
	dsn                   *protocol.Dsn
	eventProcessors       []EventProcessor
	integrations          []Integration
	externalTraceResolver externalContextTraceResolver
	sdkIdentifier         string
	sdkVersion            string
	// Transport is read-only. Replacing the transport of an existing client is
	// not supported, create a new client instead.
	Transport          Transport
	batchLogger        *logBatchProcessor
	batchMeter         *metricBatchProcessor
	telemetryProcessor *telemetry.Processor
	reportRecorder     report.ClientReportRecorder
	reportProvider     report.ClientReportProvider
}

// NewClient creates and returns an instance of Client configured using
// ClientOptions.
//
// Most users will not create clients directly. Instead, initialize the SDK with
// Init and use the package-level functions (for simple programs that run on a
// single goroutine) or hub methods (for concurrent programs, for example web
// servers).
func NewClient(options ClientOptions) (*Client, error) {
	_ = "STUB: not implemented"
	// The default error event sample rate for all SDKs is 1.0 (send all).
	//
	// In Go, the zero value (default) for float64 is 0.0, which means that
	// constructing a client with NewClient(ClientOptions{}), or, equivalently,
	// initializing the SDK with Init(ClientOptions{}) without an explicit
	// SampleRate would drop all events.
	//
	// To retain the desired default behavior, we exceptionally flip SampleRate
	// from 0.0 to 1.0 here. Setting the sample rate to 0.0 is not very useful
	// anyway, and the same end result can be achieved in many other ways like
	// not initializing the SDK, setting the DSN to the empty string or using an
	// event processor that always returns nil.
	//
	// An alternative API could be such that default options don't need to be
	// the same as Go's zero values, for example using the Functional Options
	// pattern. That would either require a breaking change if we want to reuse
	// the obvious NewClient name, or a new function as an alternative
	// constructor.
	return nil, nil
}

// SENTRYGODEBUG is a comma-separated list of key=value pairs (similar
// to GODEBUG). It is not a supported feature: recognized debug options
// may change any time.
//
// The intended public is SDK developers. It is orthogonal to
// options.Debug, which is also available for SDK users.

// dbgOpt returns true when the given debug option is enabled, for
// example SENTRYGODEBUG=someopt=1.

// We currently disallow using custom Transport with the new Telemetry Processor, due to the difference in transport signatures.
// The option should be enabled when the new Transport interface signature changes.

func (client *Client) setupTransport() { _ = "STUB: not implemented"; return }

// For known transport types, inject the client report interfaces.

func (client *Client) sdkInfo() *protocol.SdkInfo { _ = "STUB: not implemented"; return nil }

func (client *Client) setupTelemetryProcessor() { _ = "STUB: not implemented"; return }

func (client *Client) setupIntegrations() { _ = "STUB: not implemented"; return }

// AddEventProcessor adds an event processor to the client. It must not be
// called from concurrent goroutines. Most users will prefer to use
// ClientOptions.BeforeSend or Scope.AddEventProcessor instead.
//
// Note that typical programs have only a single client created by Init and the
// client is shared among multiple hubs, one per goroutine, such that adding an
// event processor to the client affects all hubs that share the client.
func (client *Client) AddEventProcessor(processor EventProcessor) {
	_ = "STUB: not implemented"
	return
}

// SetExternalContextTraceResolver installs a resolver used to extract trace/span IDs
// from external context implementations.
//
// This is intended for integrations such as OpenTelemetry.
func (client *Client) SetExternalContextTraceResolver(resolver func(ctx context.Context) (TraceID, SpanID, bool)) {
	_ = "STUB: not implemented"
	return
}

func (client *Client) externalTraceContextFromContext(ctx context.Context) (TraceID, SpanID, bool) {
	_ = "STUB: not implemented"
	return *new(TraceID), *new(SpanID), false
}

// Options return ClientOptions for the current Client.
func (client *Client) Options() ClientOptions {
	_ = "STUB: not implemented"
	// Note: internally, consider using `client.options` instead of `client.Options()` to avoid copying the object each time.
	return *new(ClientOptions)
}

// CaptureMessage captures an arbitrary message.
func (client *Client) CaptureMessage(message string, hint *EventHint, scope EventModifier) *EventID {
	_ = "STUB: not implemented"
	return nil
}

// CaptureException captures an error.
func (client *Client) CaptureException(exception error, hint *EventHint, scope EventModifier) *EventID {
	_ = "STUB: not implemented"
	return nil
}

// CaptureCheckIn captures a check in.
func (client *Client) CaptureCheckIn(checkIn *CheckIn, monitorConfig *MonitorConfig, scope EventModifier) *EventID {
	_ = "STUB: not implemented"
	return nil
}

// CaptureEvent captures an event on the currently active client if any.
//
// The event must already be assembled. Typically, code would instead use
// the utility methods like CaptureException. The return value is the
// event ID. In case Sentry is disabled or event was dropped, the return value will be nil.
func (client *Client) CaptureEvent(event *Event, hint *EventHint, scope EventModifier) *EventID {
	_ = "STUB: not implemented"
	return nil
}

func (client *Client) captureLog(log *Log, _ *Scope) bool { _ = "STUB: not implemented"; return false }

// Note: processor tracks client report

func (client *Client) captureMetric(metric *Metric, _ *Scope) bool {
	_ = "STUB: not implemented"
	return false
}

// Note: processor tracks client report

// Recover captures a panic.
// Returns EventID if successfully, or nil if there's no error to recover from.
func (client *Client) Recover(err interface{}, hint *EventHint, scope EventModifier) *EventID {
	_ = "STUB: not implemented"
	return nil
}

// Normally we would not pass a nil Context, but RecoverWithContext doesn't
// use the Context for communicating deadline nor cancelation. All it does
// is store the Context in the EventHint and there nil means the Context is
// not available.
// nolint: staticcheck

// RecoverWithContext captures a panic and passes relevant context object.
// Returns EventID if successfully, or nil if there's no error to recover from.
func (client *Client) RecoverWithContext(
	ctx context.Context,
	err interface{},
	hint *EventHint,
	scope EventModifier,
) *EventID {
	_ = "STUB: not implemented"
	return nil
}

// Flush waits until the underlying Transport sends any buffered events to the
// Sentry server, blocking for at most the given timeout. It returns false if
// the timeout was reached. In that case, some events may not have been sent.
//
// Flush should be called before terminating the program to avoid
// unintentionally dropping events.
//
// Do not call Flush indiscriminately after every call to CaptureEvent,
// CaptureException or CaptureMessage. Instead, to have the SDK send events over
// the network synchronously, configure it to use the HTTPSyncTransport in the
// call to Init.
func (client *Client) Flush(timeout time.Duration) bool { _ = "STUB: not implemented"; return false }

// FlushWithContext waits until the underlying Transport sends any buffered events
// to the Sentry server, blocking for at most the duration specified by the context.
// It returns false if the context is canceled before the events are sent. In such a case,
// some events may not be delivered.
//
// FlushWithContext should be called before terminating the program to ensure no
// events are unintentionally dropped.
//
// Avoid calling FlushWithContext indiscriminately after each call to CaptureEvent,
// CaptureException, or CaptureMessage. To send events synchronously over the network,
// configure the SDK to use HTTPSyncTransport during initialization with Init.

func (client *Client) FlushWithContext(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

// Close clean up underlying Transport resources.
//
// Close should be called after Flush and before terminating the program
// otherwise some events may be lost.
func (client *Client) Close() { _ = "STUB: not implemented"; return }

// EventFromMessage creates an event from the given message string.
func (client *Client) EventFromMessage(message string, level Level) *Event {
	_ = "STUB: not implemented"
	return nil
}

// EventFromException creates a new Sentry event from the given `error` instance.
func (client *Client) EventFromException(exception error, level Level) *Event {
	_ = "STUB: not implemented"
	return nil
}

// EventFromCheckIn creates a new Sentry event from the given `check_in` instance.
func (client *Client) EventFromCheckIn(checkIn *CheckIn, monitorConfig *MonitorConfig) *Event {
	_ = "STUB: not implemented"
	return nil
}

func (client *Client) SetSDKIdentifier(identifier string) { _ = "STUB: not implemented"; return }

func (client *Client) GetSDKIdentifier() string { _ = "STUB: not implemented"; return "" }

func (client *Client) processEvent(event *Event, hint *EventHint, scope EventModifier) *EventID {
	_ = "STUB: not implemented"
	return nil
}

// Transactions are sampled by options.TracesSampleRate or
// options.TracesSampler when they are started. Other events
// (errors, messages) are sampled here. Does not apply to check-ins.

// Apply beforeSend* processors

// Track spans removed by the callback

// not a default case, since we shouldn't apply BeforeSend on check-in events

func (client *Client) prepareEvent(event *Event, hint *EventHint, scope EventModifier) *Event {
	_ = "STUB: not implemented"
	return nil

	// TODO set EventID when the event is created, same as in other SDKs. It's necessary for profileTransaction.ID.
}

// Track spans removed by the processor

// Track spans removed by the processor

func (client *Client) listIntegrations() []string { _ = "STUB: not implemented"; return nil }

func (client *Client) integrationAlreadyInstalled(name string) bool {
	_ = "STUB: not implemented"
	return false
}

// sample returns true with the given probability, which must be in the range
// [0.0, 1.0].
func sample(probability float64) bool { _ = "STUB: not implemented"; return false }
