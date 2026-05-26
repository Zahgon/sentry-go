package sentryiris

import (
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/kataras/iris/v12"
)

// The identifier of the Iris SDK.
const (
	// sdkIdentifier is the identifier of the Iris SDK.
	sdkIdentifier = "sentry.go.iris"

	// valuesKey is used as a key to store the Sentry Hub instance on the iris.Context.
	valuesKey = "sentry"

	// transactionKey is used as a key to store the Sentry transaction on the iris.Context.
	transactionKey = "sentry_transaction"
)

type handler struct {
	repanic         bool
	waitForDelivery bool
	timeout         time.Duration
}

type Options struct {
	// Repanic configures whether Sentry should repanic after recovery, in most cases it should be set to true,
	// as iris.Default includes it's own Recovery middleware what handles http responses.
	Repanic bool
	// WaitForDelivery configures whether you want to block the request before moving forward with the response.
	// Because Iris's default Recovery handler doesn't restart the application,
	// it's safe to either skip this option or set it to false.
	WaitForDelivery bool
	// Timeout for the event delivery requests.
	Timeout time.Duration
}

// New returns a function that satisfies iris.Handler interface
// It can be used with Use() method.
func New(options Options) iris.Handler { _ = "STUB: not implemented"; return *new(iris.Handler) }

func (h *handler) handle(ctx iris.Context) { _ = "STUB: not implemented"; return }

func (h *handler) recoverWithSentry(hub *sentry.Hub, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// GetHubFromContext retrieves attached *sentry.Hub instance from iris.Context.
func GetHubFromContext(ctx iris.Context) *sentry.Hub { _ = "STUB: not implemented"; return nil }

// SetHubOnContext attaches a *sentry.Hub instance to iris.Context.
func SetHubOnContext(ctx iris.Context, hub *sentry.Hub) { _ = "STUB: not implemented"; return }

// GetSpanFromContext retrieves attached *sentry.Span instance from iris.Context.
// If there is no transaction on iris.Context, it will return nil.
func GetSpanFromContext(ctx iris.Context) *sentry.Span { _ = "STUB: not implemented"; return nil }
