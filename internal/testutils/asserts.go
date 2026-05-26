package testutils

import (
	"testing"
)

func AssertEqual(t *testing.T, got, want interface{}, userMessage ...interface{}) {
	_ = "STUB: not implemented"

	// Ideally, we would switch to cmp.Diff. However, in a general case, cmp.Diff
	// is not able to compare structs with unexported (private) fields by default. There
	// are ways to override modify that behaviour (e.g. by passing AllowUnexported), but
	// it significantly complicates its usage.
	return
}

func AssertNotEqual(t *testing.T, got, want interface{}, userMessage ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func AssertTrue(t *testing.T, condition bool, userMessage ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func AssertFalse(t *testing.T, condition bool, userMessage ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func logFailedAssertion(t *testing.T, summary string, userMessage ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func formatUnequalValues(got, want interface{}) string { _ = "STUB: not implemented"; return "" }

func AssertBaggageStringsEqual(t *testing.T, got, want string) { _ = "STUB: not implemented"; return }
