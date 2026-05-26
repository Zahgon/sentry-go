// Adapted from https://github.com/open-telemetry/opentelemetry-go/blob/cc43e01c27892252aac9a8f20da28cdde957a289/attribute/value.go
//
// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package attribute

// Type describes the type of the data Value holds.
type Type int // redefines builtin Type.

// Value represents the value part in key-value pairs.
type Value struct {
	vtype    Type
	numeric  uint64
	stringly string
	slice    any
}

const (
	// INVALID is used for a Value with no value set.
	INVALID Type = iota
	// BOOL is a boolean Type Value.
	BOOL
	// INT64 is a 64-bit signed integral Type Value.
	INT64
	// FLOAT64 is a 64-bit floating point Type Value.
	FLOAT64
	// STRING is a string Type Value.
	STRING
	// BOOLSLICE is a slice of booleans Type Value.
	BOOLSLICE
	// INT64SLICE is a slice of 64-bit signed integral numbers Type Value.
	INT64SLICE
	// FLOAT64SLICE is a slice of 64-bit floating point numbers Type Value.
	FLOAT64SLICE
	// STRINGSLICE is a slice of strings Type Value.
	STRINGSLICE
	// UINT64 is a 64-bit unsigned integral Type Value.
	//
	// This type is intentionally not exposed through the Builder API.
	UINT64
)

// BoolValue creates a BOOL Value.
func BoolValue(v bool) Value { _ = "STUB: not implemented"; return *new(Value) }

// BoolSliceValue creates a BOOLSLICE Value.
func BoolSliceValue(v []bool) Value { _ = "STUB: not implemented"; return *new(Value) }

// IntValue creates an INT64 Value.
func IntValue(v int) Value { _ = "STUB: not implemented"; return *new(Value) }

// IntSliceValue creates an INTSLICE Value.
func IntSliceValue(v []int) Value { _ = "STUB: not implemented"; return *new(Value) }

// Int64Value creates an INT64 Value.
func Int64Value(v int64) Value { _ = "STUB: not implemented"; return *new(Value) }

// Int64SliceValue creates an INT64SLICE Value.
func Int64SliceValue(v []int64) Value { _ = "STUB: not implemented"; return *new(Value) }

// Float64Value creates a FLOAT64 Value.
func Float64Value(v float64) Value { _ = "STUB: not implemented"; return *new(Value) }

// Float64SliceValue creates a FLOAT64SLICE Value.
func Float64SliceValue(v []float64) Value { _ = "STUB: not implemented"; return *new(Value) }

// StringValue creates a STRING Value.
func StringValue(v string) Value { _ = "STUB: not implemented"; return *new(Value) }

// StringSliceValue creates a STRINGSLICE Value.
func StringSliceValue(v []string) Value { _ = "STUB: not implemented"; return *new(Value) }

// Uint64Value creates a UINT64 Value.
//
// This constructor is intentionally not exposed through the Builder API.
func Uint64Value(v uint64) Value { _ = "STUB: not implemented"; return *new(Value) }

// Type returns a type of the Value.
func (v Value) Type() Type {
	_ = "STUB: not implemented"

	// AsBool returns the bool value. Make sure that the Value's type is
	// BOOL.
	return *new(Type)
}

func (v Value) AsBool() bool { _ = "STUB: not implemented"; return false }

// AsBoolSlice returns the []bool value. Make sure that the Value's type is
// BOOLSLICE.
func (v Value) AsBoolSlice() []bool { _ = "STUB: not implemented"; return nil }

// AsInt64 returns the int64 value. Make sure that the Value's type is
// INT64.
func (v Value) AsInt64() int64 { _ = "STUB: not implemented"; return 0 }

// AsInt64Slice returns the []int64 value. Make sure that the Value's type is
// INT64SLICE.
func (v Value) AsInt64Slice() []int64 { _ = "STUB: not implemented"; return nil }

// AsFloat64 returns the float64 value. Make sure that the Value's
// type is FLOAT64.
func (v Value) AsFloat64() float64 { _ = "STUB: not implemented"; return 0 }

// AsFloat64Slice returns the []float64 value. Make sure that the Value's type is
// FLOAT64SLICE.
func (v Value) AsFloat64Slice() []float64 { _ = "STUB: not implemented"; return nil }

// AsString returns the string value. Make sure that the Value's type
// is STRING.
func (v Value) AsString() string {
	_ = "STUB: not implemented"

	// AsStringSlice returns the []string value. Make sure that the Value's type is
	// STRINGSLICE.
	return ""
}

func (v Value) AsStringSlice() []string { _ = "STUB: not implemented"; return nil }

// AsUint64 returns the uint64 value. Make sure that the Value's type is
// UINT64.
func (v Value) AsUint64() uint64 { _ = "STUB: not implemented"; return 0 }

type unknownValueType struct{}

// AsInterface returns Value's data as interface{}.
func (v Value) AsInterface() interface{} { _ = "STUB: not implemented"; return nil }

// String returns a string representation of Value's data.
func (v Value) String() string { _ = "STUB: not implemented"; return "" }

// MarshalJSON returns the JSON encoding of the Value.
func (v Value) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (t Type) String() string { _ = "STUB: not implemented"; return "" }

// mapTypesToStr is a map from attribute.Type to the primitive types the server understands.
// https://develop.sentry.dev/sdk/foundations/data-model/attributes/#primitive-types
var mapTypesToStr = map[Type]string{
	INVALID:      "",
	BOOL:         "boolean",
	INT64:        "integer",
	FLOAT64:      "double",
	STRING:       "string",
	BOOLSLICE:    "array",
	INT64SLICE:   "array",
	FLOAT64SLICE: "array",
	STRINGSLICE:  "array",
	UINT64:       "integer", // wire format: same "integer" type
}
