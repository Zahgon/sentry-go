package sentry

import (
	"runtime/debug"
	"time"
)

func uuid() string { _ = "STUB: not implemented"; return "" }

func fileExists(fileName string) bool { _ = "STUB: not implemented"; return false }

// monotonicTimeSince replaces uses of time.Now() to take into account the
// monotonic clock reading stored in start, such that duration = end - start is
// unaffected by changes in the system wall clock.
func monotonicTimeSince(start time.Time) (end time.Time) {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// nolint: unused
func prettyPrint(data interface{}) { _ = "STUB: not implemented"; return }

// defaultRelease attempts to guess a default release for the currently running
// program.
func defaultRelease() (release string) {
	_ = "STUB: not implemented"
	// Return first non-empty environment variable known to hold release info, if any.
	return ""
}

// Deprecated, kept for backwards compatibility

// GitHub Actions - https://help.github.com/en/actions
// Netlify - https://docs.netlify.com/
// Vercel - https://vercel.com/
// Zeit (now known as Vercel)

// Derive a version string from Git. Example outputs:
// 	v1.0.1-0-g9de4
// 	v2.0-8-g77df-dirty
// 	4f72d7

// Either Git is not available or the current directory is not a
// Git repository.

func revisionFromBuildInfo(info *debug.BuildInfo) string { _ = "STUB: not implemented"; return "" }

func Pointer[T any](v T) *T { _ = "STUB: not implemented"; return nil }

var sensitiveHeaders = map[string]struct{}{
	"_csrf":               {},
	"_csrf_token":         {},
	"_session":            {},
	"_xsrf":               {},
	"api-key":             {},
	"apikey":              {},
	"auth":                {},
	"authorization":       {},
	"cookie":              {},
	"credentials":         {},
	"csrf":                {},
	"csrf-token":          {},
	"csrftoken":           {},
	"ip-address":          {},
	"passwd":              {},
	"password":            {},
	"private-key":         {},
	"privatekey":          {},
	"proxy-authorization": {},
	"remote-addr":         {},
	"secret":              {},
	"session":             {},
	"sessionid":           {},
	"token":               {},
	"user-session":        {},
	"x-api-key":           {},
	"x-csrftoken":         {},
	"x-forwarded-for":     {},
	"x-real-ip":           {},
	"xsrf-token":          {},
}

// IsSensitiveHeader reports whether a header or metadata key should be treated as sensitive.
func IsSensitiveHeader(key string) bool { _ = "STUB: not implemented"; return false }

// eventIdentifier returns a human-readable identifier for the event to be used in log messages.
// Format: "<description> [<event-id>]".
func eventIdentifier(event *Event) string { _ = "STUB: not implemented"; return "" }
