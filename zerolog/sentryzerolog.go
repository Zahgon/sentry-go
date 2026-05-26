package sentryzerolog

import (
	"errors"
	"io"
	"time"

	sentry "github.com/getsentry/sentry-go"
	"github.com/rs/zerolog"
)

// A large portion of this implementation has been taken from https://github.com/archdx/zerolog-sentry/blob/master/writer.go

var (
	// ErrFlushTimeout is returned when the flush operation times out.
	ErrFlushTimeout = errors.New("sentryzerolog flush timeout")

	// levels maps zerolog levels to sentry levels.
	levelsMapping = map[zerolog.Level]sentry.Level{
		zerolog.TraceLevel: sentry.LevelDebug,
		zerolog.DebugLevel: sentry.LevelDebug,
		zerolog.InfoLevel:  sentry.LevelInfo,
		zerolog.WarnLevel:  sentry.LevelWarning,
		zerolog.ErrorLevel: sentry.LevelError,
		zerolog.FatalLevel: sentry.LevelFatal,
		zerolog.PanicLevel: sentry.LevelFatal,
	}

	// Ensure that the Writer implements the io.WriteCloser interface.
	_ = io.WriteCloser(new(Writer))

	now = time.Now
)

// The identifier of the Zerolog SDK.
const sdkIdentifier = "sentry.go.zerolog"

// These default log field keys are used to pass specific metadata in a way that
// Sentry understands. If they are found in the log fields, and the value is of
// the expected datatype, it will be converted from a generic field, into Sentry
// metadata.
const (
	// FieldRequest holds an *http.Request.
	FieldRequest = "request"
	// FieldUser holds a User or *User value.
	FieldUser = "user"
	// FieldTransaction holds a transaction ID as a string.
	FieldTransaction = "transaction"
	// FieldFingerprint holds a string slice ([]string), used to dictate the
	// grouping of this event.
	FieldFingerprint = "fingerprint"

	// These fields are simply omitted, as they are duplicated by the Sentry SDK.
	FieldGoVersion = "go_version"
	FieldMaxProcs  = "go_maxprocs"

	// Name of the logger used by the Sentry SDK.
	logger = "zerolog"
)

type Config struct {
	sentry.ClientOptions
	Options
}

type Options struct {
	// Levels specifies the log levels that will trigger event sending to Sentry.
	// Only log messages at these levels will be sent. By default, the levels are
	// Error, Fatal, and Panic.
	Levels []zerolog.Level

	// WithBreadcrumbs, when enabled, adds log entries as breadcrumbs in Sentry.
	// Breadcrumbs provide a trail of events leading up to an error, which can
	// be invaluable for understanding the context of issues.
	WithBreadcrumbs bool

	// FlushTimeout sets the maximum duration allowed for flushing events to Sentry.
	// This is the time limit within which all pending events must be sent to Sentry
	// before the application exits. A typical use is ensuring all logs are sent before
	// application shutdown. The default timeout is usually 3 seconds.
	FlushTimeout time.Duration
}

func (o *Options) SetDefaults() { _ = "STUB: not implemented"; return }

// New creates writer with provided DSN and options.
func New(cfg Config) (*Writer, error) { _ = "STUB: not implemented"; return nil, nil }

// NewWithHub creates a writer using an existing sentry Hub and options.
func NewWithHub(hub *sentry.Hub, opts Options) (*Writer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Writer is a sentry events writer with std io.Writer interface.
type Writer struct {
	hub             *sentry.Hub
	levels          map[zerolog.Level]struct{}
	flushTimeout    time.Duration
	withBreadcrumbs bool
}

// addBreadcrumb adds event as a breadcrumb.
func (w *Writer) addBreadcrumb(event *sentry.Event) { _ = "STUB: not implemented"; return }

// Write handles zerolog's json and sends events to sentry.
func (w *Writer) Write(data []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// if the level is not enabled, add event as a breadcrumb

// should flush before os.Exit

func (w *Writer) WriteLevel(level zerolog.Level, p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// if the level is not enabled, add event as a breadcrumb

// should flush before os.Exit

// Close forces client to flush all pending events.
// Can be useful before application exits.
func (w *Writer) Close() error { _ = "STUB: not implemented"; return nil }

func parseLogLevel(data []byte) (zerolog.Level, error) {
	_ = "STUB: not implemented"
	return *new(zerolog.Level), nil
}

func parseLogEvent(data []byte) (*sentry.Event, bool) { _ = "STUB: not implemented"; return nil, false }
