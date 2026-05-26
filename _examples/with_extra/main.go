package main

import (
	"fmt"
	"time"

	"github.com/getsentry/sentry-go"
)

func prettyPrint(v interface{}) string { _ = "STUB: not implemented"; return "" }

type devNullTransport struct{}

func (t *devNullTransport) Configure(options sentry.ClientOptions) {
	_ = "STUB: not implemented"
	return
}

func (t *devNullTransport) SendEvent(event *sentry.Event) { _ = "STUB: not implemented"; return }

func (t *devNullTransport) Flush(timeout time.Duration) bool {
	_ = "STUB: not implemented"
	return false
}

type CustomComplexError struct {
	Message  string
	MoreData map[string]string
}

func (e CustomComplexError) Error() string { _ = "STUB: not implemented"; return "" }

func (e CustomComplexError) GimmeMoreData() map[string]string {
	_ = "STUB: not implemented"
	return nil
}

type ExtractExtra struct{}

func (ee ExtractExtra) Name() string { _ = "STUB: not implemented"; return "" }

func (ee ExtractExtra) SetupOnce(client *sentry.Client) { _ = "STUB: not implemented"; return }

func main() {
	if err := sentry.Init(sentry.ClientOptions{
		Debug: true,
		Dsn:   "https://hello@example.com/1337",
		BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			// Solution 1 (use beforeSend, which will be applied to
			// error events and is usually application specific):
			if ex, ok := hint.OriginalException.(CustomComplexError); ok {
				for key, val := range ex.GimmeMoreData() {
					event.Extra[key] = val
				}
			}

			fmt.Printf("%s\n\n", prettyPrint(event.Extra))

			return event
		},
		Transport: &devNullTransport{},

		// Solution 2 (use custom integration, which will be
		// applied to all events, can be extracted as a
		// separate utility, and reused across projects):
		Integrations: func(integrations []sentry.Integration) []sentry.Integration {
			return append(integrations, new(ExtractExtra))
		},
	}); err != nil {
		panic(err)
	}

	// Solution 3 and 4 (use scope event processors, which can be either
	// applied to all events - if used with ConfigureScope or per
	// event/block if used with WithScope):
	sentry.ConfigureScope(func(scope *sentry.Scope) {
		scope.AddEventProcessor(func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			if ex, ok := hint.OriginalException.(CustomComplexError); ok {
				for key, val := range ex.GimmeMoreData() {
					event.Extra[key] = val
				}
			}

			return event
		})
	})

	errWithExtra := CustomComplexError{
		Message: "say what again. SAY WHAT again",
		MoreData: map[string]string{
			"say": "wat",
		},
	}

	sentry.CaptureException(errWithExtra)
}
