package sentry

import (
	"context"
	"sync"
	"time"
)

// MockScope implements [Scope] for use in tests.
type MockScope struct {
	breadcrumb      *Breadcrumb
	shouldDropEvent bool
}

func (scope *MockScope) AddBreadcrumb(breadcrumb *Breadcrumb, _ int) {
	_ = "STUB: not implemented"
	return
}

func (scope *MockScope) ApplyToEvent(event *Event, _ *EventHint, _ *Client) *Event {
	_ = "STUB: not implemented"
	return nil
}

// MockTransport implements [Transport] for use in tests.
type MockTransport struct {
	mu        sync.Mutex
	events    []*Event
	lastEvent *Event
}

func (t *MockTransport) Configure(_ ClientOptions) { _ = "STUB: not implemented"; return }
func (t *MockTransport) SendEvent(event *Event)    { _ = "STUB: not implemented"; return }

func (t *MockTransport) Flush(_ time.Duration) bool { _ = "STUB: not implemented"; return false }

func (t *MockTransport) FlushWithContext(_ context.Context) bool {
	_ = "STUB: not implemented"
	return false
}
func (t *MockTransport) Events() []*Event { _ = "STUB: not implemented"; return nil }

func (t *MockTransport) Close() {
	_ = "STUB: not implemented"

	// MockLogEntry implements [sentry.LogEntry] for use in tests.
	return
}

type MockLogEntry struct {
	Attributes map[string]any
}

func (m *MockLogEntry) StringSlice(key string, value []string) LogEntry {
	_ = "STUB: not implemented"
	return *new(LogEntry)
}

func (m *MockLogEntry) Int64Slice(key string, value []int64) LogEntry {
	_ = "STUB: not implemented"
	return *new(LogEntry)
}

func (m *MockLogEntry) Float64Slice(key string, value []float64) LogEntry {
	_ = "STUB: not implemented"
	return *new(LogEntry)
}

func (m *MockLogEntry) BoolSlice(key string, value []bool) LogEntry {
	_ = "STUB: not implemented"
	return *new(LogEntry)
}

func NewMockLogEntry() *MockLogEntry { _ = "STUB: not implemented"; return nil }

func (m *MockLogEntry) WithCtx(_ context.Context) LogEntry {
	_ = "STUB: not implemented"
	return *new(LogEntry)
}
func (m *MockLogEntry) String(key, value string) LogEntry {
	_ = "STUB: not implemented"
	return *new(LogEntry)
}
func (m *MockLogEntry) Int(key string, value int) LogEntry {
	_ = "STUB: not implemented"
	return *new(LogEntry)
}

func (m *MockLogEntry) Int64(key string, value int64) LogEntry {
	_ = "STUB: not implemented"
	return *new(LogEntry)
}

func (m *MockLogEntry) Float64(key string, value float64) LogEntry {
	_ = "STUB: not implemented"
	return *new(LogEntry)
}

func (m *MockLogEntry) Bool(key string, value bool) LogEntry {
	_ = "STUB: not implemented"
	return *new(LogEntry)
}

func (m *MockLogEntry) Emit(...any)          { _ = "STUB: not implemented"; return }
func (m *MockLogEntry) Emitf(string, ...any) { _ = "STUB: not implemented"; return }
