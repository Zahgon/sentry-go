package protocol

// GenerateEventID generates a random UUID v4 for use as a Sentry event ID.
func GenerateEventID() string { _ = "STUB: not implemented"; return "" }

// Prefer rand.Read over rand.Reader, see https://go-review.googlesource.com/c/go/+/272326/.

// clear version
// set version to 4 (random uuid)
// clear variant
// set to IETF variant
