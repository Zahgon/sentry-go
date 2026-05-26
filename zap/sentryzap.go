// Package sentryzap provides a zap Core implementation for sending logs to Sentry.
package sentryzap

import (
	"context"
	"time"

	"github.com/getsentry/sentry-go"
	"go.uber.org/zap/zapcore"
)

const (
	// ZapOrigin is the Sentry origin attribute value for zap logs.
	ZapOrigin = "auto.log.zap"
)

// Ensure SentryCore implements zapcore.Core.
var _ zapcore.Core = (*SentryCore)(nil)

// Option configures the Sentry zap Core.
type Option struct {
	// Level specifies the zap levels to capture and send to Sentry as log entries.
	// Only logs at these specific levels will be processed.
	// Defaults to all levels: Debug, Info, Warn, Error, DPanic, Panic, Fatal.
	Level []zapcore.Level

	// AddCaller includes caller information (file, line, function) in logs.
	// Defaults to false.
	AddCaller bool

	// FlushTimeout specifies how long to wait when syncing/flushing logs.
	// Defaults to 5 seconds.
	FlushTimeout time.Duration
}

// SentryCore is a zapcore.Core implementation that sends logs to Sentry.
type SentryCore struct {
	option Option
	logger sentry.Logger
	fields []zapcore.Field
	ctx    context.Context
}

// NewSentryCore creates a new zapcore.Core that sends logs to Sentry.
func NewSentryCore(ctx context.Context, opts Option) *SentryCore {
	_ = "STUB: not implemented"
	return nil
}

// Context returns a zapcore.Field that can be used with logger.With() to link
// traces with the provided context. This allows propagating Sentry trace information
// from the context to logs without needing to pass a Hub.
//
// Example:
//
//	logger := zap.New(sentryzap.NewSentryCore(ctx, sentryzap.Option{}))
//	logger = logger.With(sentryzap.Context(requestCtx))
//	logger.Info("handling request") // This log will be linked to the trace in requestCtx
func Context(ctx context.Context) zapcore.Field {
	_ = "STUB: not implemented"
	return *new(zapcore.Field)
}

// Enabled returns true if the given level is in the configured Level list.
func (c *SentryCore) Enabled(level zapcore.Level) bool { _ = "STUB: not implemented"; return false }

// With returns a new Core with the given fields added to the context.
func (c *SentryCore) With(fields []zapcore.Field) zapcore.Core {
	_ = "STUB: not implemented"
	return *new(zapcore.Core)
}

// Check determines whether the supplied Entry should be logged.
func (c *SentryCore) Check(entry zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	_ = "STUB: not implemented"
	return nil
}

// Write serializes the Entry and any Fields and sends them to Sentry.
func (c *SentryCore) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	_ = "STUB: not implemented"
	return nil
}

// Convert and add accumulated fields from With()

// Convert and add fields from this specific log call

// Sync flushes any buffered log entries to Sentry.
func (c *SentryCore) Sync() error { _ = "STUB: not implemented"; return nil }

// getLogEntry returns the appropriate sentry.LogEntry for the given zap level.
func (c *SentryCore) getLogEntry(ctx context.Context, level zapcore.Level) sentry.LogEntry {
	_ = "STUB: not implemented"
	return *new(sentry.LogEntry)
}

// DPanic is treated as Error in production

// For any other level, use Info
