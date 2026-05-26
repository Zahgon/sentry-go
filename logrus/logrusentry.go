// Package sentrylogrus provides a simple Logrus hook for Sentry.
package sentrylogrus

import (
	"context"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
)

const (
	// sdkIdentifier is the identifier of the Logrus SDK.
	sdkIdentifier = "sentry.go.logrus"
	// the name of the logger.
	name = "logrus"

	maxErrorDepth = 100
)

// These default log field keys are used to pass specific metadata in a way that
// Sentry understands. If they are found in the log fields, and the value is of
// the expected datatype, it will be converted from a generic field, into Sentry
// metadata.
//
// These keys may be overridden by calling SetKey on the hook object.
const (
	// FieldRequest holds an *http.Request.
	FieldRequest = "request"
	// FieldUser holds a User or *User value.
	FieldUser = "user"
	// FieldTransaction holds a transaction ID as a string.
	FieldTransaction = "transaction"
	// FieldFingerprint holds a string slice ([]string), used to dictate the
	// grouping of this event.
	FieldFingerprint = "fingerprint"

	// These fields are simply omitted, as they are duplicated by the Sentry SDK.
	FieldGoVersion = "go_version"
	FieldMaxProcs  = "go_maxprocs"

	LogrusOrigin = "auto.log.logrus"
)

var levelMap = map[logrus.Level]sentry.Level{
	logrus.TraceLevel: sentry.LevelDebug,
	logrus.DebugLevel: sentry.LevelDebug,
	logrus.InfoLevel:  sentry.LevelInfo,
	logrus.WarnLevel:  sentry.LevelWarning,
	logrus.ErrorLevel: sentry.LevelError,
	logrus.FatalLevel: sentry.LevelFatal,
	logrus.PanicLevel: sentry.LevelFatal,
}

// Hook is the logrus hook for Sentry.
//
// It is not safe to configure the hook while logging is happening. Please
// perform all configuration before using it.
type Hook interface {
	// SetHubProvider sets a function to provide a hub for each log entry.
	SetHubProvider(provider func() *sentry.Hub)
	// AddTags adds tags to the hook's scope.
	AddTags(tags map[string]string)
	// SetFallback sets a fallback function for the eventHook.
	SetFallback(fb FallbackFunc)
	// SetKey sets an alternate field key for the eventHook.
	SetKey(oldKey, newKey string)
	// Levels returns the list of logging levels that will be sent to Sentry as events.
	Levels() []logrus.Level
	// Fire sends entry to Sentry as an event.
	Fire(entry *logrus.Entry) error
	// Flush waits until the underlying Sentry transport sends any buffered events.
	Flush(timeout time.Duration) bool
	// FlushWithContext waits for the underlying Sentry transport to send any buffered
	// events, blocking until the context's deadline is reached or the context is canceled.
	// It returns false if the context is canceled or its deadline expires before the events
	// are sent, meaning some events may not have been sent.
	FlushWithContext(ctx context.Context) bool
}

// Deprecated: New creates issues/events from log entries. Errors should only be captured
// using sentry.CaptureException instead of being converted from log entries.
// Use [NewLogHook] for structured logging. Will be removed in 0.48.0.
func New(levels []logrus.Level, opts sentry.ClientOptions) (Hook, error) {
	_ = "STUB: not implemented"
	return *new(Hook), nil
}

// Deprecated: NewFromClient creates issues/events from log entries. Errors should only be
// captured using sentry.CaptureException instead of being converted from log entries.
// Use [NewLogHookFromClient] for structured logging. Will be removed in 0.48.0.
func NewFromClient(levels []logrus.Level, client *sentry.Client) Hook {
	_ = "STUB: not implemented"
	return *new(Hook)
}

// A FallbackFunc can be used to attempt to handle any errors in logging, before
// resorting to Logrus's standard error reporting.
type FallbackFunc func(*logrus.Entry) error

type eventHook struct {
	hubProvider func() *sentry.Hub
	fallback    FallbackFunc
	keys        map[string]string
	levels      []logrus.Level
}

var _ Hook = &eventHook{}
var _ logrus.Hook = &eventHook{} // eventHook still needs to be a logrus.Hook

func (h *eventHook) SetHubProvider(provider func() *sentry.Hub) { _ = "STUB: not implemented"; return }

func (h *eventHook) AddTags(tags map[string]string) { _ = "STUB: not implemented"; return }

func (h *eventHook) SetFallback(fb FallbackFunc) { _ = "STUB: not implemented"; return }

func (h *eventHook) SetKey(oldKey, newKey string) { _ = "STUB: not implemented"; return }

func (h *eventHook) key(key string) string { _ = "STUB: not implemented"; return "" }

func (h *eventHook) Levels() []logrus.Level { _ = "STUB: not implemented"; return nil }

func (h *eventHook) Fire(entry *logrus.Entry) error { _ = "STUB: not implemented"; return nil }

func (h *eventHook) entryToEvent(l *logrus.Entry) *sentry.Event {
	_ = "STUB: not implemented"
	return nil
}

func (h *eventHook) Flush(timeout time.Duration) bool { _ = "STUB: not implemented"; return false }

func (h *eventHook) FlushWithContext(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

// Deprecated: NewEventHook creates issues/events from log entries. Errors should only be
// captured using sentry.CaptureException instead of being converted from log entries.
// Use [NewLogHook] for structured logging. Will be removed in 0.48.0.
func NewEventHook(levels []logrus.Level, opts sentry.ClientOptions) (Hook, error) {
	_ = "STUB: not implemented"
	return *new(Hook), nil
}

// Deprecated: NewEventHookFromClient creates issues/events from log entries. Errors should
// only be captured using sentry.CaptureException instead of being converted from log
// entries. Use [NewLogHookFromClient] for structured logging. Will be removed in 0.48.0.
func NewEventHookFromClient(levels []logrus.Level, client *sentry.Client) Hook {
	_ = "STUB: not implemented"
	return *new(Hook)
}

// Default to using the same hub if no specific provider is set

type logHook struct {
	hubProvider func() *sentry.Hub
	fallback    FallbackFunc
	keys        map[string]string
	levels      []logrus.Level
	logger      sentry.Logger
}

var _ Hook = &logHook{}
var _ logrus.Hook = &logHook{} // logHook also needs to be a logrus.Hook

func (h *logHook) SetHubProvider(provider func() *sentry.Hub) { _ = "STUB: not implemented"; return }

func (h *logHook) AddTags(tags map[string]string) {
	_ = "STUB: not implemented"
	// for logs convert tags to attributes
	return
}

func (h *logHook) SetFallback(fb FallbackFunc) { _ = "STUB: not implemented"; return }

func (h *logHook) SetKey(oldKey, newKey string) { _ = "STUB: not implemented"; return }

func (h *logHook) key(key string) string { _ = "STUB: not implemented"; return "" }

// uint64LogEntry is used to pass uint64 values without conversion.
// The concrete sentry.logEntry type satisfies this interface,
// but it is intentionally not part of the public sentry.LogEntry API.
type uint64LogEntry interface {
	Uint64(key string, value uint64) sentry.LogEntry
}

func logrusFieldToLogEntry(logEntry sentry.LogEntry, key string, value interface{}) sentry.LogEntry {
	_ = "STUB: not implemented"
	return *new(sentry.LogEntry)
}

// Fallback to string conversion for unknown types

func (h *logHook) Fire(entry *logrus.Entry) error { _ = "STUB: not implemented"; return nil }

// Create the base log entry for the appropriate level

// Add all the fields as attributes to this specific log entry

// Skip specific fields that might be handled separately

// Emit the log entry with the message

func (h *logHook) Levels() []logrus.Level { _ = "STUB: not implemented"; return nil }

func (h *logHook) Flush(timeout time.Duration) bool { _ = "STUB: not implemented"; return false }

func (h *logHook) FlushWithContext(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

// NewLogHook initializes a new Logrus hook which sends logs to a new Sentry client
// configured according to opts.
func NewLogHook(levels []logrus.Level, opts sentry.ClientOptions) (Hook, error) {
	_ = "STUB: not implemented"
	return *new(Hook), nil
}

// NewLogHookFromClient initializes a new Logrus hook which sends logs to the provided
// sentry client.
func NewLogHookFromClient(levels []logrus.Level, client *sentry.Client) Hook {
	_ = "STUB: not implemented"
	return *new(Hook)
}

// Default to using the same hub if no specific provider is set
