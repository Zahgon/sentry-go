package sentryzap

import (
	"encoding/json"

	"github.com/getsentry/sentry-go"
	"go.uber.org/zap/zapcore"
)

var encoderConfig = zapcore.EncoderConfig{
	TimeKey:        "time",
	LevelKey:       "level",
	NameKey:        "logger",
	CallerKey:      "caller",
	MessageKey:     "msg",
	StacktraceKey:  "stacktrace",
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    zapcore.LowercaseLevelEncoder,
	EncodeTime:     zapcore.RFC3339TimeEncoder,
	EncodeDuration: zapcore.StringDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
}

// encodeAndExtractValue uses the zapcore.JSONEncoder to serialize custom serializable object/array types.
func encodeAndExtractValue(addToEncoder func(zapcore.Encoder) error) (json.RawMessage, error) {
	_ = "STUB: not implemented"
	return *new(json.RawMessage), nil
}

// stringifyObject uses zap's JSON encoder to safely stringify an object.
func stringifyObject(obj zapcore.ObjectMarshaler) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// stringifyArray uses zap's JSON encoder to safely stringify an array.
func stringifyArray(arr zapcore.ArrayMarshaler) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// isSafeToCall checks that the field.Interface doesn't hold a nil pointer.
func isSafeToCall(v interface{}) bool { _ = "STUB: not implemented"; return false }

// uint64LogEntry is used to pass uint64 values without conversion.
// The concrete sentry.logEntry type satisfies this interface,
// but it is intentionally not part of the public sentry.LogEntry API.
type uint64LogEntry interface {
	Uint64(key string, value uint64) sentry.LogEntry
}

func toUint64(value int64) uint64 {
	_ = "STUB: not implemented"
	//nolint:gosec // zap stores unsigned integers and float bit patterns in Field.Integer.
	return 0
}

func toUint32(value int64) uint32 {
	_ = "STUB: not implemented"
	//nolint:gosec // zap stores float32 bit patterns in the lower 32 bits of Field.Integer.
	return 0
}

// zapFieldToLogEntry converts a zap Field to a sentry LogEntry attribute.
//
//nolint:gocyclo
func zapFieldToLogEntry(entry sentry.LogEntry, field zapcore.Field) sentry.LogEntry {
	_ = "STUB: not implemented"
	return *new(sentry.LogEntry)
}

// Namespace fields are just markers for grouping subsequent fields, so we skip them.

// Fallback for any unknown types
