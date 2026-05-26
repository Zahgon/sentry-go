package sentrygin

import (
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

const (
	// sdkIdentifier is the identifier of the Gin SDK.
	sdkIdentifier = "sentry.go.gin"

	// valuesKey is used as a key to store the Sentry Hub instance on the gin.Context.
	valuesKey = "sentry"

	// transactionKey is used as a key to store the Sentry transaction on the gin.Context.
	transactionKey = "sentry_transaction"
)

type handler struct {
	repanic         bool
	waitForDelivery bool
	timeout         time.Duration
}

type Options struct {
	// Repanic configures whether Sentry should repanic after recovery, in most cases it should be set to true,
	// as gin.Default includes it's own Recovery middleware what handles http responses.
	Repanic bool
	// WaitForDelivery configures whether you want to block the request before moving forward with the response.
	// Because Gin's default Recovery handler doesn't restart the application,
	// it's safe to either skip this option or set it to false.
	WaitForDelivery bool
	// Timeout for the event delivery requests.
	Timeout time.Duration
}

// New returns a function that satisfies gin.HandlerFunc interface
// It can be used with Use() methods.
func New(options Options) gin.HandlerFunc { _ = "STUB: not implemented"; return *new(gin.HandlerFunc) }

func (h *handler) handle(c *gin.Context) { _ = "STUB: not implemented"; return }

func (h *handler) recoverWithSentry(hub *sentry.Hub, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Check for a broken connection, as this is what Gin does already.
func isBrokenPipeError(err interface{}) bool { _ = "STUB: not implemented"; return false }

// GetHubFromContext retrieves attached *sentry.Hub instance from gin.Context.
func GetHubFromContext(ctx *gin.Context) *sentry.Hub { _ = "STUB: not implemented"; return nil }

// SetHubOnContext sets *sentry.Hub instance to gin.Context.
func SetHubOnContext(ctx *gin.Context, hub *sentry.Hub) { _ = "STUB: not implemented"; return }

// GetSpanFromContext retrieves attached *sentry.Span instance from gin.Context.
// If there is no transaction on echo.Context, it will return nil.
func GetSpanFromContext(ctx *gin.Context) *sentry.Span { _ = "STUB: not implemented"; return nil }
