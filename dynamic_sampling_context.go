package sentry

const (
	sentryPrefix = "sentry-"
)

// DynamicSamplingContext holds information about the current event that can be used to make dynamic sampling decisions.
type DynamicSamplingContext struct {
	Entries map[string]string
	Frozen  bool
}

func DynamicSamplingContextFromHeader(header []byte) (DynamicSamplingContext, error) {
	_ = "STUB: not implemented"
	return *new(DynamicSamplingContext), nil
}

// We only store baggage members if their key starts with "sentry-".

// If there's at least one Sentry value, we consider the DSC frozen

func DynamicSamplingContextFromTransaction(span *Span) DynamicSamplingContext {
	_ = "STUB: not implemented"
	return *new(DynamicSamplingContext)
}

// Only include the transaction name if it's of good quality (not empty and not SourceURL)

func (d DynamicSamplingContext) HasEntries() bool { _ = "STUB: not implemented"; return false }

func (d DynamicSamplingContext) IsFrozen() bool { _ = "STUB: not implemented"; return false }

func (d DynamicSamplingContext) String() string { _ = "STUB: not implemented"; return "" }

// DynamicSamplingContextFromScope Constructs a new DynamicSamplingContext using a scope and client. Accessing
// fields on the scope are not thread safe, and this function should only be
// called within scope methods.
func DynamicSamplingContextFromScope(scope *Scope, client *Client) DynamicSamplingContext {
	_ = "STUB: not implemented"
	return *new(DynamicSamplingContext)
}
