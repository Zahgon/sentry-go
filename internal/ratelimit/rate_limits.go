package ratelimit

import (
	"errors"
	"time"
)

var errInvalidXSRLRetryAfter = errors.New("invalid retry-after value")

// parseXSentryRateLimits returns a RateLimits map by parsing an input string in
// the format of the X-Sentry-Rate-Limits header.
//
// Example
//
//	X-Sentry-Rate-Limits: 60:transaction, 2700:default;error;security
//
// This will rate limit transactions for the next 60 seconds and errors for the
// next 2700 seconds.
//
// Limits for unknown categories are ignored.
func parseXSentryRateLimits(s string, now time.Time) Map {
	_ = "STUB: not implemented"
	// https://github.com/getsentry/relay/blob/0424a2e017d193a93918053c90cdae9472d164bf/relay-server/src/utils/rate_limits.rs#L44-L82
	return *new(Map)
}

// skip unknown categories, keep m small

// always keep the deadline furthest into the future

// parseXSRLRetryAfter parses a string into a retry-after rate limit deadline.
//
// Valid input is a number, possibly signed and possibly floating-point,
// indicating the number of seconds to wait before sending another request.
// Negative values are treated as zero. Fractional values are rounded to the
// next integer.
func parseXSRLRetryAfter(s string, now time.Time) (Deadline, error) {
	_ = "STUB: not implemented"
	// https://github.com/getsentry/relay/blob/0424a2e017d193a93918053c90cdae9472d164bf/relay-quotas/src/rate_limit.rs#L88-L96
	return *new(Deadline), nil
}
