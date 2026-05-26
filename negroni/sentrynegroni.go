package sentrynegroni

import (
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/urfave/negroni/v3"
)

// The identifier of the Negroni SDK.
const sdkIdentifier = "sentry.go.negroni"

type handler struct {
	repanic         bool
	waitForDelivery bool
	timeout         time.Duration
}

type Options struct {
	// Repanic configures whether Sentry should repanic after recovery, in most cases it should be set to true,
	// as negroni.Classic includes it's own Recovery middleware that handles http responses.
	Repanic bool
	// WaitForDelivery configures whether you want to block the request before moving forward with the response.
	// Because Negroni's default Recovery handler doesn't restart the application,
	// it's safe to either skip this option or set it to false.
	WaitForDelivery bool
	// Timeout for the event delivery requests.
	Timeout time.Duration
}

// New returns a handler struct which satisfies Negroni's middleware interface
// It can be used with New(), Use() or With() methods.
func New(options Options) negroni.Handler { _ = "STUB: not implemented"; return *new(negroni.Handler) }

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	_ = "STUB: not implemented"
	return
}

func (h *handler) recoverWithSentry(hub *sentry.Hub, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// PanicHandlerFunc can be used for Negroni's default Recovery middleware option called PanicHandlerFunc,
// which let you "plug-in" to its own handler.
func PanicHandlerFunc(info *negroni.PanicInformation) { _ = "STUB: not implemented"; return }
