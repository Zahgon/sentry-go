// Package sentrytest provides test fixtures for the Sentry Go SDK.
//
// Fixture bundles a Hub, Client, and MockTransport with assertion helpers,
// eliminating boilerplate across integration and unit tests.
//
// Isolated mode (default) creates a cloned hub safe for parallel tests.
// Global mode (WithGlobal) calls sentry.Init for middleware tests that read
// from sentry.CurrentHub.
package sentrytest

import (
	"context"
	"testing"

	"github.com/getsentry/sentry-go"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

const testDsn = "https://whatever@sentry.io/1337"

// DefaultEventCmpOpts are [cmp.Options] for comparing [sentry.Event] values.
// They ignore fields that vary between runs (IDs, timestamps, server metadata)
// and all unexported fields.
var DefaultEventCmpOpts = cmp.Options{
	cmpopts.IgnoreFields(sentry.Event{},
		"Contexts",
		"EventID",
		"Modules",
		"Platform",
		"Release",
		"Sdk",
		"ServerName",
		"Timestamp",
	),
	cmpopts.IgnoreFields(sentry.Request{}, "Env"),
	cmpopts.IgnoreUnexported(sentry.Event{}),
	cmpopts.EquateEmpty(),
}

// Option configures a [Fixture].
type Option func(*config)

type config struct {
	opts   sentry.ClientOptions
	global bool
}

// WithGlobal makes the fixture call [sentry.Init] to set the global hub
// instead of creating an isolated hub. Use this for middleware tests where
// the middleware reads from [sentry.CurrentHub].
//
// Tests using global mode must not run in parallel.
func WithGlobal() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithClientOptions sets the [sentry.ClientOptions] for the fixture.
// Transport is always overridden to use the fixture's mock transport. If Dsn is
// empty, the fixture sets a placeholder test DSN.
func WithClientOptions(opts sentry.ClientOptions) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// Fixture provides an isolated Sentry environment for testing.
// It bundles a Hub, Client, and MockTransport with assertion helpers.
type Fixture struct {
	// T is the test context.
	T testing.TB

	// Hub is the fixture's hub. In global mode this is [sentry.CurrentHub].
	// In isolated mode this is a clone with its own client.
	Hub *sentry.Hub

	// Client is the fixture's client, configured with the MockTransport.
	Client *sentry.Client

	// Transport captures all events sent through the client.
	Transport *sentry.MockTransport

	// useSynctest indicates the fixture is running inside a synctest bubble.
	useSynctest bool
}

// Run creates a [Fixture] inside a [synctest.Test] bubble and calls fn
// with it. All background goroutines (batch processors) use fake time, so
// [Fixture.Flush] completes instantly. This is the preferred way to
// create fixtures.
//
// For the rare case where synctest is incompatible (e.g. third-party code that
// leaks goroutines), use [NewFixture] directly.
func Run(t *testing.T, fn func(t *testing.T, f *Fixture), opts ...Option) {
	_ = "STUB: not implemented"
	return
}

// NewFixture creates a new test fixture without a synctest bubble.
// Prefer [Run] for most tests; use this only when synctest is incompatible.
func NewFixture(t testing.TB, opts ...Option) *Fixture { _ = "STUB: not implemented"; return nil }

// NewContext creates a context backed by a new [Fixture] without a synctest
// bubble. If parent is nil, [context.Background] is used.
func NewContext(parent context.Context, t testing.TB, opts ...Option) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func newFixture(t testing.TB, useSynctest bool, opts ...Option) *Fixture {
	_ = "STUB: not implemented"
	return nil
}

// Ensure background goroutines (batch processors) are stopped when the test finishes.
// This is required for synctest bubbles which will panic if blocked goroutines remain
// after the bubble's root goroutine exits.

// Flush flushes the fixture's client.
//
// Inside a [Run] bubble it calls [synctest.Wait] first to let all background
// goroutines settle, then flushes under fake time (completing instantly).
// Outside a bubble it awaits for a real Flush.
func (f *Fixture) Flush() { _ = "STUB: not implemented"; return }

// NewContext returns parent with the fixture's hub attached. If parent is nil,
// [context.Background] is used.
func (f *Fixture) NewContext(parent context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// Events returns all captured events, including transactions.
//
// TODO: Add typed helper views (errors, transactions, logs, metrics,
// check-ins) when the telemetry processor path is enabled in tests.
func (f *Fixture) Events() []*sentry.Event { _ = "STUB: not implemented"; return nil }

// AssertEventCount flushes and asserts the number of captured events.
func (f *Fixture) AssertEventCount(want int) { _ = "STUB: not implemented"; return }

func (f *Fixture) assertCount(kind string, want int) { _ = "STUB: not implemented"; return }

// DiffEvents flushes and returns a [cmp.Diff] of captured events
// against want. Uses [DefaultEventCmpOpts] merged with any additional opts.
// Returns "" when events match.
func (f *Fixture) DiffEvents(want []*sentry.Event, opts ...cmp.Option) string {
	_ = "STUB: not implemented"
	return ""
}

// AssertHubIsolation verifies that requestHub is a distinct clone from the
// fixture's hub, confirming the middleware properly cloned the hub per request.
func (f *Fixture) AssertHubIsolation(requestHub *sentry.Hub) { _ = "STUB: not implemented"; return }

// Apply the request scope to a probe event to read its tags.
