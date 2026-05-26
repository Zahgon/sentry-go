package main

import (
	"log"

	"github.com/getsentry/sentry-go"
)

type pickleIntegration struct{}

func (pi *pickleIntegration) Name() string { _ = "STUB: not implemented"; return "" }

func (pi *pickleIntegration) SetupOnce(client *sentry.Client) { _ = "STUB: not implemented"; return }

func (pi *pickleIntegration) processor(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
	_ = "STUB: not implemented"
	return nil
}

func main() {
	scope1 := sentry.NewScope()
	client1, _ := sentry.NewClient(sentry.ClientOptions{
		Dsn: "https://hello@example.com/1",
		BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			log.Println(event.Message)
			return nil
		},
		Integrations: func(integrations []sentry.Integration) []sentry.Integration {
			return append(integrations, &pickleIntegration{})
		},
	})
	hub1 := sentry.NewHub(client1, scope1)

	scope2 := sentry.NewScope()
	client2, _ := sentry.NewClient(sentry.ClientOptions{
		Dsn: "https://hello@example.com/2",
		BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			log.Println(event.Message)
			return nil
		},
	})
	hub2 := sentry.NewHub(client2, scope2)

	hub1.CaptureMessage("Hub: altered message by pickleIntegration")
	hub2.CaptureMessage("Hub: _NOT_ altered message by pickleIntegration")

	client1.CaptureMessage("Client: altered message by pickleIntegration", &sentry.EventHint{}, scope1)
	client2.CaptureMessage("Client: _NOT_ altered message by pickleIntegration", &sentry.EventHint{}, scope2)
}
