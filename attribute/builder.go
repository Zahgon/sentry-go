package attribute

type Builder struct {
	Key   string
	Value Value
}

// String returns a Builder for a string value.
func String(key, value string) Builder { _ = "STUB: not implemented"; return *new(Builder) }

// Int64 returns a Builder for an int64.
func Int64(key string, value int64) Builder { _ = "STUB: not implemented"; return *new(Builder) }

// Int returns a Builder for an int64.
func Int(key string, value int) Builder { _ = "STUB: not implemented"; return *new(Builder) }

// Float64 returns a Builder for a float64.
func Float64(key string, v float64) Builder { _ = "STUB: not implemented"; return *new(Builder) }

// Bool returns a Builder for a boolean.
func Bool(key string, v bool) Builder { _ = "STUB: not implemented"; return *new(Builder) }

// BoolSlice returns a Builder for a bool slice.
func BoolSlice(key string, v []bool) Builder { _ = "STUB: not implemented"; return *new(Builder) }

// IntSlice returns a Builder for an int slice.
func IntSlice(key string, v []int) Builder { _ = "STUB: not implemented"; return *new(Builder) }

// Int64Slice returns a Builder for an int64 slice.
func Int64Slice(key string, v []int64) Builder { _ = "STUB: not implemented"; return *new(Builder) }

// Float64Slice returns a Builder for a float64 slice.
func Float64Slice(key string, v []float64) Builder { _ = "STUB: not implemented"; return *new(Builder) }

// StringSlice returns a Builder for a string slice.
func StringSlice(key string, v []string) Builder { _ = "STUB: not implemented"; return *new(Builder) }

// Valid checks for valid key and type.
func (b *Builder) Valid() bool { _ = "STUB: not implemented"; return false }
