package sentryslog

import (
	"context"
	"log/slog"

	"github.com/getsentry/sentry-go"
)

// Majority of the code in this package is derived from https://github.com/samber/slog-sentry AND https://github.com/samber/slog-common
// MIT License

// Copyright (c) 2023 Samuel Berthe

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

var (
	_ slog.Handler = (*SentryHandler)(nil)

	// Deprecated: LogLevels is only used by the deprecated event capture functionality.
	// Will be removed in 0.48.0.
	LogLevels = map[slog.Level]sentry.Level{
		slog.LevelDebug: sentry.LevelDebug,
		slog.LevelInfo:  sentry.LevelInfo,
		slog.LevelWarn:  sentry.LevelWarning,
		slog.LevelError: sentry.LevelError,
		LevelFatal:      sentry.LevelFatal,
	}
)

// LevelFatal is a custom [slog.Level] that maps to [sentry.LevelFatal].
const LevelFatal = slog.Level(12)
const SlogOrigin = "auto.log.slog"

type Option struct {
	// Deprecated: Use EventLevel instead. Level is kept for backwards compatibility and defaults to EventLevel.
	Level slog.Leveler

	// Deprecated: EventLevel creates issues/events from log entries. Errors should only be
	// captured using sentry.CaptureException instead of being converted from log entries.
	// Use LogLevel for structured logging. Will be removed in 0.48.0.
	EventLevel []slog.Level

	// LogLevel specifies the exact log levels to capture and send to Sentry as Log entries.
	// Only logs at these specific levels will be processed as log entries.
	// Defaults to []slog.Level{slog.LevelDebug, slog.LevelInfo, slog.LevelWarn, slog.LevelError, LevelFatal}.
	LogLevel []slog.Level

	// Deprecated: Hub is only used by the deprecated event capture functionality. Errors
	// should only be captured using sentry.CaptureException instead of being converted
	// from log entries. Will be removed in 0.48.0.
	Hub *sentry.Hub

	// Deprecated: Converter is only used by the deprecated event capture functionality.
	// Errors should only be captured using sentry.CaptureException instead of being
	// converted from log entries. Will be removed in 0.48.0.
	Converter Converter

	// AttrFromContext is an optional slice of functions that extract attributes
	// from the context. These functions can add additional metadata to the log entry.
	AttrFromContext []func(ctx context.Context) []slog.Attr

	// AddSource is an optional flag that, when set to true, includes the source
	// information (such as file and line number) in the Sentry event.
	// This can be useful for debugging purposes.
	AddSource bool

	// ReplaceAttr is an optional function that allows for the modification or
	// replacement of attributes in the log record. This can be used to filter
	// or transform attributes before they are sent to Sentry.
	ReplaceAttr func(groups []string, a slog.Attr) slog.Attr
}

func (o Option) NewSentryHandler(ctx context.Context) slog.Handler {
	_ = "STUB: not implemented"
	return *

	// backwards compatibility
	new(slog.Handler)
}

type SentryHandler struct {
	eventHandler *eventHandler
	logHandler   *logHandler
}

func (h *SentryHandler) Enabled(ctx context.Context, level slog.Level) bool {
	_ = "STUB: not implemented"
	return false
}

func (h *SentryHandler) Handle(ctx context.Context, record slog.Record) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *SentryHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

func (h *SentryHandler) WithGroup(name string) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

type eventHandler struct {
	ctx    context.Context
	option Option
	attrs  []slog.Attr
	groups []string
}

func (h *eventHandler) Enabled(_ context.Context, level slog.Level) bool {
	_ = "STUB: not implemented"
	return false
}

func (h *eventHandler) Handle(ctx context.Context, record slog.Record) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *eventHandler) WithAttrs(attrs []slog.Attr) *eventHandler {
	_ = "STUB: not implemented"
	// Create a copy of the groups slice to avoid sharing state
	return nil
}

func (h *eventHandler) WithGroup(name string) *eventHandler { _ = "STUB: not implemented"; return nil }

// Create a copy of the groups slice to avoid modifying the original

type logHandler struct {
	option Option
	attrs  []slog.Attr
	groups []string
	logger sentry.Logger
}

func (h *logHandler) Enabled(_ context.Context, level slog.Level) bool {
	_ = "STUB: not implemented"
	return false
}

func (h *logHandler) Handle(ctx context.Context, record slog.Record) error {
	_ = "STUB: not implemented"
	// when logging without context, slog passes `context.Background`. Check for span existence to not overwrite the root context.
	return nil
}

// aggregate all attributes

// Use level ranges instead of exact matches to support custom levels

// Levels below Debug (e.g., Trace)

// Debug level range: -4 to -1

// Info level range: 0 to 3

// Warn level range: 4 to 7

// custom Fatal level, keep +4 increments

// Fatal level range: 12 and above

func (h *logHandler) WithAttrs(attrs []slog.Attr) *logHandler {
	_ = "STUB: not implemented"
	// Create a copy of the groups slice to avoid sharing state
	return nil
}

func (h *logHandler) WithGroup(name string) *logHandler { _ = "STUB: not implemented"; return nil }

// Create a copy of the groups slice to avoid modifying the original

func levelsFromMinimum(minLevel slog.Level) []slog.Level { _ = "STUB: not implemented"; return nil }
