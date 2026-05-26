package debug

import (
	"io"
	"net/http"
)

// Transport implements http.RoundTripper and can be used to wrap other HTTP
// transports for debugging, normally http.DefaultTransport.
type Transport struct {
	http.RoundTripper
	Output io.Writer
	// Dump controls whether to dump HTTP request and responses.
	Dump bool
	// Trace enables usage of net/http/httptrace.
	Trace bool
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ensureTrailingNewline(b []byte) []byte { _ = "STUB: not implemented"; return nil }
