package main

import (
	"fmt"
	"time"

	"github.com/getsentry/sentry-go"
)

func prettyPrint(v interface{}) string { _ = "STUB: not implemented"; return "" }

func fooErr() { _ = "STUB: not implemented"; return }

func barErr() { _ = "STUB: not implemented"; return }

func bazErr() { _ = "STUB: not implemented"; return }

func fooMsg() { _ = "STUB: not implemented"; return }

func barMsg() { _ = "STUB: not implemented"; return }

func bazMsg() { _ = "STUB: not implemented"; return }

func main() {
	_ = sentry.Init(sentry.ClientOptions{
		Debug:            true,
		Dsn:              "https://hello@example.com/1337",
		AttachStacktrace: true,
		BeforeSend: func(e *sentry.Event, h *sentry.EventHint) *sentry.Event {
			fmt.Println(prettyPrint(e))
			return e
		},
	})

	sentry.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetTag("oristhis", "justfantasy")
		scope.SetTag("isthis", "reallife")
		scope.SetLevel(sentry.LevelFatal)
		scope.SetUser(sentry.User{
			ID: "1337",
		})
	})

	func() {
		defer sentry.Recover()
		fooErr()
	}()

	func() {
		defer sentry.Recover()
		fooMsg()
	}()

	sentry.Flush(time.Second * 5)
}
