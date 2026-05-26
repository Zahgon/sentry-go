package sentry

// MergeBaggage merges an existing baggage header with a Sentry-generated one.
//
// Existing third-party members are preserved. If both baggage strings contain
// the same member key, the Sentry-generated member wins. The helper is best-effort
// and only keeps the sentry baggage in case the existing one is malformed.
func MergeBaggage(existingHeader, sentryHeader string) (string, error) {
	_ = "STUB: not implemented"
	// TODO: we are reparsing the headers here, because we currently don't
	// expose a method to get only DSC or its baggage members.
	return "", nil
}

// in case that the incoming header is malformed we should only
// care about merging sentry related baggage information for distributed tracing.
