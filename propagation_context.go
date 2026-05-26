package sentry

type PropagationContext struct {
	TraceID                TraceID                `json:"trace_id"`
	SpanID                 SpanID                 `json:"span_id"`
	ParentSpanID           SpanID                 `json:"parent_span_id,omitzero"`
	DynamicSamplingContext DynamicSamplingContext `json:"-"`
}

func (p PropagationContext) Map() map[string]interface{} { _ = "STUB: not implemented"; return nil }

func NewPropagationContext() PropagationContext {
	_ = "STUB: not implemented"
	return *new(PropagationContext)
}

func PropagationContextFromHeaders(trace, baggage string) (PropagationContext, error) {
	_ = "STUB: not implemented"
	return *new(PropagationContext), nil
}

// In case a sentry-trace header is present but there are no sentry-related
// values in the baggage, create an empty, frozen DynamicSamplingContext.
