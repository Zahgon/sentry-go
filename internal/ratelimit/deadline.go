package ratelimit

import "time"

// A Deadline is a time instant when a rate limit expires.
type Deadline time.Time

// After reports whether the deadline d is after other.
func (d Deadline) After(other Deadline) bool { _ = "STUB: not implemented"; return false }

// Equal reports whether d and e represent the same deadline.
func (d Deadline) Equal(e Deadline) bool { _ = "STUB: not implemented"; return false }

// String returns the deadline formatted for debugging.
func (d Deadline) String() string {
	_ = "STUB: not implemented"
	// Like time.Time.String, but without the monotonic clock reading.
	return ""
}
