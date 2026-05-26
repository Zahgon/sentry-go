package sentry

import (
	"context"

	"github.com/getsentry/sentry-go/attribute"
)

// Fallback, no-op logger if logging is disabled.
type noopLogger struct{}

// noopLogEntry implements LogEntry for the no-op logger.
type noopLogEntry struct {
	level       LogLevel
	shouldPanic bool
	shouldFatal bool
}

func (n *noopLogEntry) WithCtx(_ context.Context) LogEntry {
	_ = "STUB: not implemented"
	return *new(LogEntry)
}

func (n *noopLogEntry) String(_, _ string) LogEntry {
	_ = "STUB: not implemented"
	return *new(LogEntry)
}

func (n *noopLogEntry) Int(_ string, _ int) LogEntry {
	_ = "STUB: not implemented"
	return *new(LogEntry)
}

func (n *noopLogEntry) Int64(_ string, _ int64) LogEntry {
	_ = "STUB: not implemented"
	return *new(LogEntry)
}

func (n *noopLogEntry) Float64(_ string, _ float64) LogEntry {
	_ = "STUB: not implemented"
	return *new(LogEntry)
}

func (n *noopLogEntry) Bool(_ string, _ bool) LogEntry {
	_ = "STUB: not implemented"
	return *new(LogEntry)
}

func (n *noopLogEntry) StringSlice(_ string, _ []string) LogEntry {
	_ = "STUB: not implemented"
	return *new(LogEntry)
}

func (n *noopLogEntry) Int64Slice(_ string, _ []int64) LogEntry {
	_ = "STUB: not implemented"
	return *new(LogEntry)
}

func (n *noopLogEntry) Float64Slice(_ string, _ []float64) LogEntry {
	_ = "STUB: not implemented"
	return *new(LogEntry)
}

func (n *noopLogEntry) BoolSlice(_ string, _ []bool) LogEntry {
	_ = "STUB: not implemented"
	return *new(LogEntry)
}

func (n *noopLogEntry) Attributes(_ ...attribute.Builder) LogEntry {
	_ = "STUB: not implemented"
	return *new(LogEntry)
}

func (n *noopLogEntry) Emit(args ...interface{}) { _ = "STUB: not implemented"; return }

func (n *noopLogEntry) Emitf(message string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (n *noopLogger) GetCtx() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (*noopLogger) Trace() LogEntry { _ = "STUB: not implemented"; return *new(LogEntry) }

func (*noopLogger) Debug() LogEntry { _ = "STUB: not implemented"; return *new(LogEntry) }

func (*noopLogger) Info() LogEntry { _ = "STUB: not implemented"; return *new(LogEntry) }

func (*noopLogger) Warn() LogEntry { _ = "STUB: not implemented"; return *new(LogEntry) }

func (*noopLogger) Error() LogEntry { _ = "STUB: not implemented"; return *new(LogEntry) }

func (*noopLogger) Fatal() LogEntry { _ = "STUB: not implemented"; return *new(LogEntry) }

func (*noopLogger) Panic() LogEntry { _ = "STUB: not implemented"; return *new(LogEntry) }

func (*noopLogger) LFatal() LogEntry { _ = "STUB: not implemented"; return *new(LogEntry) }

func (*noopLogger) SetAttributes(...attribute.Builder) { _ = "STUB: not implemented"; return }

func (*noopLogger) Write(_ []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }
