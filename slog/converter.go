package sentryslog

import (
	"log/slog"

	"github.com/getsentry/sentry-go"
)

const maxErrorDepth = 100

var (
	sourceKey = "source"
	errorKeys = map[string]struct{}{
		"error": {},
		"err":   {},
	}
	name = "slog"
)

// Deprecated: Converter is only used by the deprecated event capture functionality.
// Errors should only be captured using sentry.CaptureException instead of being converted
// from log entries. Will be removed in 0.48.0.
type Converter func(addSource bool, replaceAttr func(groups []string, a slog.Attr) slog.Attr, loggerAttr []slog.Attr, groups []string, record *slog.Record, hub *sentry.Hub) *sentry.Event

// Deprecated: DefaultConverter is only used by the deprecated event capture functionality.
// Errors should only be captured using sentry.CaptureException instead of being converted
// from log entries. Will be removed in 0.48.0.
func DefaultConverter(addSource bool, replaceAttr func(groups []string, a slog.Attr) slog.Attr, loggerAttr []slog.Attr, groups []string, record *slog.Record, hub *sentry.Hub) *sentry.Event {
	_ = "STUB: not implemented"
	// aggregate all attributes
	return nil
}

// developer formatters

// handler formatter

func attrToSentryEvent(attr slog.Attr, event *sentry.Event) { _ = "STUB: not implemented"; return }

func handleUserAttributes(v slog.Value, event *sentry.Event) { _ = "STUB: not implemented"; return }

func handleRequestAttributes(v slog.Value, event *sentry.Event) { _ = "STUB: not implemented"; return }

func handleFingerprint(v slog.Value, event *sentry.Event) { _ = "STUB: not implemented"; return }

// uint64LogEntry is used to pass uint64 values without conversion.
// The concrete sentry.logEntry type satisfies this interface,
// but it is intentionally not part of the public sentry.LogEntry API.
type uint64LogEntry interface {
	Uint64(key string, value uint64) sentry.LogEntry
}

func slogAttrToLogEntry(logEntry sentry.LogEntry, group string, a slog.Attr) sentry.LogEntry {
	_ = "STUB: not implemented"
	return *new(sentry.LogEntry)
}

// Handle nested group attributes
