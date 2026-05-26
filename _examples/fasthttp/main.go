package main

import (
	"fmt"

	"github.com/getsentry/sentry-go"
	sentryfasthttp "github.com/getsentry/sentry-go/fasthttp"
	"github.com/valyala/fasthttp"
)

func enhanceSentryEvent(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	_ = "STUB: not implemented"
	return *new(fasthttp.RequestHandler)
}

// do resource intensive thing

// Acquire transaction on current hub that's created by the SDK.
// Be careful, it might be a nil value if you didn't set up sentryecho middleware.

// Pass in the `.Context()` method from `*sentry.Span` struct.
// The `context.Context` instance inherits the context from `echo.Context`.

func main() {
	_ = sentry.Init(sentry.ClientOptions{
		Dsn: "",
		BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			if hint.Context != nil {
				if ctx, ok := hint.Context.Value(sentry.RequestContextKey).(*fasthttp.RequestCtx); ok {
					// You have access to the original Context if it panicked
					fmt.Println(string(ctx.Request.Host()))
				}
			}
			fmt.Println(event)
			return event
		},
		Debug:            true,
		AttachStacktrace: true,
	})

	sentryHandler := sentryfasthttp.New(sentryfasthttp.Options{})

	defaultHandler := func(ctx *fasthttp.RequestCtx) {
		if hub := sentryfasthttp.GetHubFromContext(ctx); hub != nil {
			hub.WithScope(func(scope *sentry.Scope) {
				scope.SetTag("unwantedQuery", "someQueryDataMaybe")
				hub.CaptureMessage("User provided unwanted query string, but we recovered just fine")
			})
		}
		ctx.SetStatusCode(fasthttp.StatusOK)
	}

	fooHandler := enhanceSentryEvent(func(ctx *fasthttp.RequestCtx) {
		panic("y tho")
	})

	fastHTTPHandler := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/foo":
			fooHandler(ctx)
		default:
			defaultHandler(ctx)
		}
	}

	fmt.Println("Listening and serving HTTP on :3000")

	if err := fasthttp.ListenAndServe(":3000", sentryHandler.Handle(fastHTTPHandler)); err != nil {
		panic(err)
	}
}
