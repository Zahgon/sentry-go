package ratelimit

// Reference:
// https://github.com/getsentry/relay/blob/46dfaa850b8717a6e22c3e9a275ba17fe673b9da/relay-base-schema/src/data_category.rs#L231-L271

// Category classifies supported payload types that can be ingested by Sentry
// and, therefore, rate limited.
type Category string

// Known rate limit categories that are specified in rate limit headers.
const (
	CategoryUnknown     Category = "unknown" // Unknown category should not get rate limited
	CategoryAll         Category = ""        // Special category for empty categories (applies to all)
	CategoryError       Category = "error"
	CategoryTransaction Category = "transaction"
	CategorySpan        Category = "span"
	CategoryLog         Category = "log_item"
	CategoryLogByte     Category = "log_byte"
	CategoryMonitor     Category = "monitor"
	CategoryTraceMetric Category = "trace_metric"
)

// knownCategories is the set of currently known categories. Other categories
// are ignored for the purpose of rate-limiting.
var knownCategories = map[Category]struct{}{
	CategoryAll:         {},
	CategoryError:       {},
	CategoryTransaction: {},
	CategoryLog:         {},
	CategoryMonitor:     {},
	CategoryTraceMetric: {},
}

// String returns the category formatted for debugging.
func (c Category) String() string { _ = "STUB: not implemented"; return "" }

// For unknown categories, use the original formatting logic

// Priority represents the importance level of a category for buffer management.
type Priority int

const (
	PriorityCritical Priority = iota + 1
	PriorityHigh
	PriorityMedium
	PriorityLow
	PriorityLowest
)

func (p Priority) String() string { _ = "STUB: not implemented"; return "" }

// GetPriority returns the priority level for this category.
func (c Category) GetPriority() Priority { _ = "STUB: not implemented"; return *new(Priority) }
