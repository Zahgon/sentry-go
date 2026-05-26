package traceutils

import (
	"net/http"
)

// GetHTTPSpanName grab needed fields from *http.Request to generate a span name for `http.server` span op.
func GetHTTPSpanName(r *http.Request) string { _ = "STUB: not implemented"; return "" }

// If value does not start with HTTP methods, add them.
// The method and the path should be separated by a space.
