package sentry

import (
	"context"
	"sync"

	"github.com/getsentry/sentry-go/attribute"
)

type LogLevel string

const (
	LogLevelTrace LogLevel = "trace"
	LogLevelDebug LogLevel = "debug"
	LogLevelInfo  LogLevel = "info"
	LogLevelWarn  LogLevel = "warn"
	LogLevelError LogLevel = "error"
	LogLevelFatal LogLevel = "fatal"
)

const (
	LogSeverityTrace   int = 1
	LogSeverityDebug   int = 5
	LogSeverityInfo    int = 9
	LogSeverityWarning int = 13
	LogSeverityError   int = 17
	LogSeverityFatal   int = 21
)

type sentryLogger struct {
	ctx               context.Context
	hub               *Hub
	attributes        map[string]attribute.Value
	defaultAttributes map[string]attribute.Value
	mu                sync.RWMutex
}

type logEntry struct {
	logger      *sentryLogger
	ctx         context.Context
	level       LogLevel
	severity    int
	attributes  map[string]attribute.Value
	shouldPanic bool
	shouldFatal bool
}

// NewLogger returns a Logger that emits logs to Sentry. If logging is turned off, all logs get discarded.
func NewLogger(ctx context.Context) Logger {
	_ = "STUB: not implemented" // nolint: dupl
	return *new(Logger)
}

// Build default attrs

func (l *sentryLogger) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (l *sentryLogger) log(ctx context.Context, level LogLevel, severity int, message string, entryAttrs map[string]attribute.Value, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// Pre-allocate with capacity hint to avoid map growth reallocations
// scope ~3 + instance ~5

// attribute precedence: default -> scope -> instance (from SetAttrs) -> entry-specific

func (l *sentryLogger) SetAttributes(attrs ...attribute.Builder) { _ = "STUB: not implemented"; return }

func (l *sentryLogger) Trace() LogEntry { _ = "STUB: not implemented"; return *new(LogEntry) }

func (l *sentryLogger) Debug() LogEntry { _ = "STUB: not implemented"; return *new(LogEntry) }

func (l *sentryLogger) Info() LogEntry { _ = "STUB: not implemented"; return *new(LogEntry) }

func (l *sentryLogger) Warn() LogEntry { _ = "STUB: not implemented"; return *new(LogEntry) }

func (l *sentryLogger) Error() LogEntry { _ = "STUB: not implemented"; return *new(LogEntry) }

func (l *sentryLogger) Fatal() LogEntry { _ = "STUB: not implemented"; return *new(LogEntry) }

func (l *sentryLogger) Panic() LogEntry { _ = "STUB: not implemented"; return *new(LogEntry) }

func (l *sentryLogger) LFatal() LogEntry { _ = "STUB: not implemented"; return *new(LogEntry) }

func (l *sentryLogger) GetCtx() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (e *logEntry) WithCtx(ctx context.Context) LogEntry {
	_ = "STUB: not implemented"
	return *new(LogEntry)
}

func (e *logEntry) String(key, value string) LogEntry {
	_ = "STUB: not implemented"
	return *new(LogEntry)
}

func (e *logEntry) StringSlice(key string, value []string) LogEntry {
	_ = "STUB: not implemented"
	return *new(LogEntry)
}

func (e *logEntry) Int(key string, value int) LogEntry {
	_ = "STUB: not implemented"
	return *new(LogEntry)
}

func (e *logEntry) Int64Slice(key string, value []int64) LogEntry {
	_ = "STUB: not implemented"
	return *new(LogEntry)
}

func (e *logEntry) Int64(key string, value int64) LogEntry {
	_ = "STUB: not implemented"
	return *new(LogEntry)
}

func (e *logEntry) Float64Slice(key string, value []float64) LogEntry {
	_ = "STUB: not implemented"
	return *new(LogEntry)
}

func (e *logEntry) Float64(key string, value float64) LogEntry {
	_ = "STUB: not implemented"
	return *new(LogEntry)
}

func (e *logEntry) BoolSlice(key string, value []bool) LogEntry {
	_ = "STUB: not implemented"
	return *new(LogEntry)
}

func (e *logEntry) Bool(key string, value bool) LogEntry {
	_ = "STUB: not implemented"
	return *new(LogEntry)
}

// Uint64 adds uint64 attributes to the log entry.
//
// This method is intentionally not part of the LogEntry interface to avoid exposing uint64 in the public API.
func (e *logEntry) Uint64(key string, value uint64) LogEntry {
	_ = "STUB: not implemented"
	return *new(LogEntry)
}

func (e *logEntry) Emit(args ...interface{}) { _ = "STUB: not implemented"; return }

func (e *logEntry) Emitf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }
