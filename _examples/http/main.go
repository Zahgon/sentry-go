// This is an example web server to demonstrate how to instrument error and
// performance monitoring with Sentry.
//
// Try it by running:
//
//	go run main.go
//
// To actually report events to Sentry, set the DSN either by editing the
// appropriate line below or setting the environment variable SENTRY_DSN to
// match the DSN of your Sentry project.
package main

import (
	"context"
	"flag"
	"html/template"
	"image"
	"log"
)

var addr = flag.String("addr", "127.0.0.1:3000", "bind address")

func main() {
	flag.Parse()
	configureLoggers()
	// The helper run function does not call log.Fatal, otherwise deferred
	// function calls would not be executed when the program exits.
	log.Fatal(run())
}

// run runs a web server. As with http.ListenAndServe, the returned error is
// always non-nil.
func run() error { _ = "STUB: not implemented"; return nil }

// Either set your DSN here or set the SENTRY_DSN environment variable.

// Enable printing of SDK debug messages.
// Useful when getting started or trying to figure something out.

// Here you can inspect/modify non-transaction events (for example, errors) before they are sent.
// Returning nil drops the event.

// Here you can inspect/modify transaction events before they are sent.
// Returning nil drops the event.

// Drop the transaction

// Enable tracing

// Specify either a TracesSampleRate...

// ... or a TracesSampler

// Don't sample health checks.

// Flush buffered events before the program terminates.
// Set the timeout to the maximum duration the program can afford to wait.

// Main HTTP handler, renders an HTML page with a random image.
//
// A new transaction is automatically sent to Sentry when the handler is
// invoked.

// Use GetHubFromContext to get a hub associated with the
// current request. Hubs provide data isolation, such that tags,
// breadcrumbs and other attributes are never mixed up across
// requests.

// Set a custom transaction name: use "Home" instead of the
// default "/" based on r.URL.Path.

// The next block of code shows how to instrument concurrent
// tasks.

// For demonstration only, ensure homepage loading takes
// at least 40ms.

// HTTP handler for the random image.
//
// A new transaction is automatically sent to Sentry when the handler is
// invoked. We use sentry.StartSpan and span.Finish to create additional
// child spans measuring specific parts of the image computation.
//
// In general, wrap potentially slow parts of your handlers (external
// network calls, CPU-intensive tasks, etc) to help identify where time
// is spent.

// this line will panic

// Wrap the default mux with Sentry to capture panics, report errors and
// measure performance.
//
// Alternatively, you can also wrap individual handlers if you need to
// use different options for different parts of your app.

var t = template.Must(template.New("").Parse(`<!DOCTYPE html>
<html lang="en">
<head>
<title>Random Image</title>
<style>
img {
	width: 128px;
	height: 128px;
}
</style>
<base target="_blank">
</head>
<body>
<h1>Random Image</h1>
<img src="/random.png?q={{.}}" alt="Random Image">
<h2>Click one of these links to send an event to Sentry</h2>
<ul>
<li><a href="/random.png?q={{.}}&timeout=20ms">Open random image and abort if it takes longer than 20ms</a></li>
<li><a href="/404">Trigger 404 not found error</a></li>
<li><a href="/panic/">Trigger server-side panic</a></li>
</ul>
</body>
</html>`))

// NewImage returns a random image based on seed, with the given width and
// height.
//
// NewImage uses the context to create spans that measure the performance of its
// internal parts.
func NewImage(ctx context.Context, width, height int, seed []byte) image.Image {
	_ = "STUB: not implemented"
	return *new(image.Image)
}

// Context canceled, abort image generation.

// Set a tag on the current span.

// Set a tag on the current transaction.
//
// Note that spans are not designed to be mutated from
// concurrent goroutines. If multiple goroutines may try
// to mutate a span/transaction, for example to set
// tags, use a mutex to synchronize changes, or use a
// channel to communicate the desired changes back into
// the goroutine where the span was created.

// Spot the bug: the returned image cannot be encoded as
// PNG and will cause an error that will be reported to
// Sentry.

// configureLoggers configures the standard logger and the logger used by the
// Sentry SDK.
//
// The only reason to change logger configuration in this example is aesthetics.
func configureLoggers() { _ = "STUB: not implemented"; return }
