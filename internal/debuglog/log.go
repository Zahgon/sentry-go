package debuglog

import (
	"io"
	"log"
)

// logger is the global debug logger instance.
var logger = log.New(io.Discard, "[Sentry] ", log.LstdFlags)

// SetOutput changes the output destination of the logger.
func SetOutput(w io.Writer) { _ = "STUB: not implemented"; return }

// GetLogger returns the current logger instance.
// This function is thread-safe and can be called concurrently.
func GetLogger() *log.Logger {
	_ = "STUB: not implemented"

	// Printf calls Printf on the underlying logger.
	return nil
}

func Printf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

// Println calls Println on the underlying logger.
func Println(args ...interface{}) { _ = "STUB: not implemented"; return }

// Print calls Print on the underlying logger.
func Print(args ...interface{}) { _ = "STUB: not implemented"; return }
