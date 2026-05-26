package sentryfiber

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/getsentry/sentry-go"
)

const (
	// sdkIdentifier is the identifier of the FastHTTP SDK.
	sdkIdentifier = "sentry.go.fiber"

	// valuesKey is used as a key to store the Sentry Hub instance on the fasthttp.RequestCtx.
	valuesKey = "sentry"

	// transactionKey is used as a key to store the Sentry transaction on the fasthttp.RequestCtx.
	transactionKey = "sentry_transaction"
)

type handler struct {
	repanic         bool
	waitForDelivery bool
	timeout         time.Duration
}

type Options struct {
	// Repanic configures whether Sentry should repanic after recovery, in most cases it should be set to false,
	// as fasthttp doesn't include its own Recovery handler.
	Repanic bool
	// WaitForDelivery configures whether you want to block the request before moving forward with the response.
	// Because fasthttp doesn't include its own Recovery handler, it will restart the application,
	// and event won't be delivered otherwise.
	WaitForDelivery bool
	// Timeout for the event delivery requests.
	Timeout time.Duration
}

// New returns a handler struct which satisfies Fiber's middleware interface.
func New(options Options) fiber.Handler { _ = "STUB: not implemented"; return *new(fiber.Handler) }

func (h *handler) handle(ctx *fiber.Ctx) error { _ = "STUB: not implemented"; return nil }

func (h *handler) recoverWithSentry(hub *sentry.Hub, ctx *fiber.Ctx) {
	_ = "STUB: not implemented"
	return
}

// GetHubFromContext retrieves the Hub instance from the *fiber.Ctx.
func GetHubFromContext(ctx *fiber.Ctx) *sentry.Hub { _ = "STUB: not implemented"; return nil }

// SetHubOnContext sets the Hub instance on the *fiber.Ctx.
func SetHubOnContext(ctx *fiber.Ctx, hub *sentry.Hub) { _ = "STUB: not implemented"; return }

// GetSpanFromContext retrieves the Span instance from the *fiber.Ctx.
func GetSpanFromContext(ctx *fiber.Ctx) *sentry.Span { _ = "STUB: not implemented"; return nil }

func convert(ctx *fiber.Ctx) *http.Request { _ = "STUB: not implemented"; return nil }

// Headers

// Cookies
