package sentryfasthttp

import (
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/valyala/fasthttp"
)

const (
	// sdkIdentifier is the identifier of the FastHTTP SDK.
	sdkIdentifier = "sentry.go.fasthttp"

	// valuesKey is used as a key to store the Sentry Hub instance on the  fasthttp.RequestCtx.
	valuesKey = "sentry"

	// transactionKey is used as a key to store the Sentry transaction on the fasthttp.RequestCtx.
	transactionKey = "sentry_transaction"
)

type Handler struct {
	repanic         bool
	waitForDelivery bool
	timeout         time.Duration
}

type Options struct {
	// Repanic configures whether Sentry should repanic after recovery, in most cases it should be set to false,
	// as fasthttp doesn't include it's own Recovery handler.
	Repanic bool
	// WaitForDelivery configures whether you want to block the request before moving forward with the response.
	// Because fasthttp doesn't include it's own Recovery handler, it will restart the application,
	// and event won't be delivered otherwise.
	WaitForDelivery bool
	// Timeout for the event delivery requests.
	Timeout time.Duration
}

// New returns a struct that provides Handle method
// that satisfy fasthttp.RequestHandler interface.
func New(options Options) *Handler { _ = "STUB: not implemented"; return nil }

// Handle wraps fasthttp.RequestHandler and recovers from caught panics.
func (h *Handler) Handle(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	_ = "STUB: not implemented"
	return *new(fasthttp.RequestHandler)
}

func (h *Handler) recoverWithSentry(hub *sentry.Hub, ctx *fasthttp.RequestCtx) {
	_ = "STUB: not implemented"
	return
}

// GetHubFromContext retrieves attached *sentry.Hub instance from fasthttp.RequestCtx.
func GetHubFromContext(ctx *fasthttp.RequestCtx) *sentry.Hub { _ = "STUB: not implemented"; return nil }

// SetHubOnContext attaches the *sentry.Hub instance to the fasthttp.RequestCtx.
func SetHubOnContext(ctx *fasthttp.RequestCtx, hub *sentry.Hub) { _ = "STUB: not implemented"; return }

// GetSpanFromContext retrieves attached *sentry.Span instance from *fasthttp.RequestCtx.
// If there is no transaction on *fasthttp.RequestCtx, it will return nil.
func GetSpanFromContext(ctx *fasthttp.RequestCtx) *sentry.Span {
	_ = "STUB: not implemented"
	return nil
}

func convert(ctx *fasthttp.RequestCtx) *http.Request { _ = "STUB: not implemented"; return nil }

// Headers

// Cookies
