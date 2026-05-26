package testutils

import (
	"time"
)

func IsCI() bool { _ = "STUB: not implemented"; return false }

func FlushTimeout() time.Duration {
	_ = "STUB: not implemented"

	// CI is very overloaded so we need to allow for a long wait time.
	return *new(time.Duration)
}
