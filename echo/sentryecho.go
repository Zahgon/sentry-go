package sentryecho

import (
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v5"
)

const (
	// sdkIdentifier is the identifier of the Echo SDK.
	sdkIdentifier = "sentry.go.echo"

	// valuesKey is used as a key to store the Sentry Hub instance on the  *echo.Context.
	valuesKey = "sentry"

	// transactionKey is used as a key to store the Sentry transaction on the *echo.Context.
	transactionKey = "sentry_transaction"

	// errorKey is used as a key to store the error on the *echo.Context.
	errorKey = "error"
)

type handler struct {
	repanic         bool
	waitForDelivery bool
	timeout         time.Duration
}

type Options struct {
	// Repanic configures whether Sentry should repanic after recovery, in most cases it should be set to true,
	// as Echo includes its own Recover middleware that handles HTTP responses.
	Repanic bool
	// WaitForDelivery configures whether you want to block the request before moving forward with the response.
	// Because Echo's Recover handler doesn't restart the application,
	// it's safe to either skip this option or set it to false.
	WaitForDelivery bool
	// Timeout for the event delivery requests.
	Timeout time.Duration
}

// New returns a function that satisfies echo.HandlerFunc interface
// It can be used with Use() methods.
func New(options Options) echo.MiddlewareFunc {
	_ = "STUB: not implemented"
	return *new(echo.MiddlewareFunc)
}

func (h *handler) handle(next echo.HandlerFunc) echo.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(echo.HandlerFunc)
}

// Store the error so it can be used in the deferred function

func (h *handler) recoverWithSentry(hub *sentry.Hub, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// GetHubFromContext retrieves attached *sentry.Hub instance from *echo.Context.
func GetHubFromContext(ctx *echo.Context) *sentry.Hub { _ = "STUB: not implemented"; return nil }

// SetHubOnContext attaches *sentry.Hub instance to *echo.Context.
func SetHubOnContext(ctx *echo.Context, hub *sentry.Hub) { _ = "STUB: not implemented"; return }

// GetSpanFromContext retrieves attached *sentry.Span instance from *echo.Context.
// If there is no transaction on *echo.Context, it will return nil.
func GetSpanFromContext(ctx *echo.Context) *sentry.Span { _ = "STUB: not implemented"; return nil }
