package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/getsentry/sentry-go"
	sentryhttpclient "github.com/getsentry/sentry-go/httpclient"
)

func main() {
	_ = sentry.Init(sentry.ClientOptions{
		Dsn:              "",
		EnableTracing:    true,
		TracesSampleRate: 1.0,
		BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			fmt.Println(event)
			return event
		},
		BeforeSendTransaction: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			fmt.Println(event)
			return event
		},
		Debug: true,
	})

	// With custom HTTP client
	ctx := sentry.SetHubOnContext(context.Background(), sentry.CurrentHub().Clone())
	httpClient := &http.Client{
		Transport: sentryhttpclient.NewSentryRoundTripper(nil),
	}

	err := getExamplePage(ctx, httpClient)
	if err != nil {
		panic(err)
	}
}

func getExamplePage(ctx context.Context, httpClient *http.Client) error {
	_ = "STUB: not implemented"
	return nil
}
