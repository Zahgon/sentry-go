package ratelimit

import (
	"net/http"
	"time"
)

// Map maps categories to rate limit deadlines.
//
// A rate limit is in effect for a given category if either the category's
// deadline or the deadline for the special CategoryAll has not yet expired.
//
// Use IsRateLimited to check whether a category is rate-limited.
type Map map[Category]Deadline

// IsRateLimited returns true if the category is currently rate limited.
func (m Map) IsRateLimited(c Category) bool { _ = "STUB: not implemented"; return false }

func (m Map) isRateLimited(c Category, now time.Time) bool { _ = "STUB: not implemented"; return false }

// Deadline returns the deadline when the rate limit for the given category or
// the special CategoryAll expire, whichever is furthest into the future.
func (m Map) Deadline(c Category) Deadline { _ = "STUB: not implemented"; return *new(Deadline) }

// Merge merges the other map into m.
//
// If a category appears in both maps, the deadline that is furthest into the
// future is preserved.
func (m Map) Merge(other Map) { _ = "STUB: not implemented"; return }

// FromResponse returns a rate limit map from an HTTP response.
func FromResponse(r *http.Response) Map { _ = "STUB: not implemented"; return *new(Map) }

func fromResponse(r *http.Response, now time.Time) Map { _ = "STUB: not implemented"; return *new(Map) }
