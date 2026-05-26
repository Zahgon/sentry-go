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

func recoverHandler() { _ = "STUB: not implemented"; return }

func beforeSend() { _ = "STUB: not implemented"; return }

func captureMessage() { _ = "STUB: not implemented"; return }

func configureScope() { _ = "STUB: not implemented"; return }

func withScope() { _ = "STUB: not implemented"; return }

func addBreadcrumbs() { _ = "STUB: not implemented"; return }

func withScopeAndConfigureScope() { _ = "STUB: not implemented"; return }

type CustomComplexError struct {
	Message      string
	AnswerToLife int
}

func (e CustomComplexError) Error() string { _ = "STUB: not implemented"; return "" }

func (e CustomComplexError) GimmeMoreData() string { _ = "STUB: not implemented"; return "" }

func eventHint() { _ = "STUB: not implemented"; return }

func main() {
	if err := sentry.Init(sentry.ClientOptions{
		Debug:        true,
		Dsn:          "https://hello@example.com/1337",
		IgnoreErrors: []string{"^(?i)drop me"},
		BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			if ex, ok := hint.OriginalException.(CustomComplexError); ok {
				event.Message = event.Message + " - " + ex.GimmeMoreData()
			}

			fmt.Printf("%s\n\n", prettyPrint(event))

			return event
		},
		BeforeBreadcrumb: func(breadcrumb *sentry.Breadcrumb, _ *sentry.BreadcrumbHint) *sentry.Breadcrumb {
			if breadcrumb.Message == "Random breadcrumb 3" {
				breadcrumb.Message = "Not so random breadcrumb 3"
			}

			fmt.Printf("%s\n\n", prettyPrint(breadcrumb))

			return breadcrumb
		},
		SampleRate: 1,
		Transport:  &devNullTransport{},
		Integrations: func(integrations []sentry.Integration) []sentry.Integration {
			return append(integrations, integrations[1])
		},
	}); err != nil {
		panic(err)
	}

	beforeSend()
	configureScope()
	withScope()
	captureMessage()
	addBreadcrumbs()
	withScopeAndConfigureScope()
	recoverHandler()
	eventHint()
}
