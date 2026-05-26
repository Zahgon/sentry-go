package main

import (
	"time"

	"github.com/getsentry/sentry-go"
)

func runTask(monitorSlug string, duration time.Duration, shouldFail bool) {
	_ = "STUB: not implemented"
	return
}

func main() {
	_ = sentry.Init(sentry.ClientOptions{
		Dsn:   "",
		Debug: true,
	})

	// Start a task that runs every minute and always succeeds
	go func() {
		for {
			go runTask("sentry-go-periodic-task-success", time.Second, false)
			time.Sleep(time.Minute)
		}
	}()

	time.Sleep(3 * time.Second)

	// Start a task that runs every minute and fails every second time
	go func() {
		shouldFail := true
		for {
			go runTask("sentry-go-periodic-task-sometimes-fail", 2*time.Second, shouldFail)
			time.Sleep(time.Minute)
			shouldFail = !shouldFail
		}
	}()

	select {}
}
