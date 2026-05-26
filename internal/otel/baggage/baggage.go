// Adapted from https://github.com/open-telemetry/opentelemetry-go/blob/c21b6b6bb31a2f74edd06e262f1690f3f6ea3d5c/baggage/baggage.go
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

package baggage

import (
	"errors"
	"regexp"

	"github.com/getsentry/sentry-go/internal/otel/baggage/internal/baggage"
)

const (
	maxMembers               = 180
	maxBytesPerMembers       = 4096
	maxBytesPerBaggageString = 8192

	listDelimiter     = ","
	keyValueDelimiter = "="
	propertyDelimiter = ";"

	keyDef      = `([\x21\x23-\x27\x2A\x2B\x2D\x2E\x30-\x39\x41-\x5a\x5e-\x7a\x7c\x7e]+)`
	valueDef    = `([\x21\x23-\x2b\x2d-\x3a\x3c-\x5B\x5D-\x7e]*)`
	keyValueDef = `\s*` + keyDef + `\s*` + keyValueDelimiter + `\s*` + valueDef + `\s*`
)

var (
	keyRe      = regexp.MustCompile(`^` + keyDef + `$`)
	valueRe    = regexp.MustCompile(`^` + valueDef + `$`)
	propertyRe = regexp.MustCompile(`^(?:\s*` + keyDef + `\s*|` + keyValueDef + `)$`)
)

var (
	errInvalidKey      = errors.New("invalid key")
	errInvalidValue    = errors.New("invalid value")
	errInvalidProperty = errors.New("invalid baggage list-member property")
	errInvalidMember   = errors.New("invalid baggage list-member")
	errMemberNumber    = errors.New("too many list-members in baggage-string")
	errMemberBytes     = errors.New("list-member too large")
	errBaggageBytes    = errors.New("baggage-string too large")
)

// Property is an additional metadata entry for a baggage list-member.
type Property struct {
	key, value string

	// hasValue indicates if a zero-value value means the property does not
	// have a value or if it was the zero-value.
	hasValue bool

	// hasData indicates whether the created property contains data or not.
	// Properties that do not contain data are invalid with no other check
	// required.
	hasData bool
}

// NewKeyProperty returns a new Property for key.
//
// If key is invalid, an error will be returned.
func NewKeyProperty(key string) (Property, error) {
	_ = "STUB: not implemented"
	return *new(Property), nil
}

// NewKeyValueProperty returns a new Property for key with value.
//
// If key or value are invalid, an error will be returned.
func NewKeyValueProperty(key, value string) (Property, error) {
	_ = "STUB: not implemented"
	return *new(Property), nil
}

func newInvalidProperty() Property {
	_ = "STUB: not implemented"

	// parseProperty attempts to decode a Property from the passed string. It
	// returns an error if the input is invalid according to the W3C Baggage
	// specification.
	return *new(Property)
}

func parseProperty(property string) (Property, error) {
	_ = "STUB: not implemented"
	return *new(Property), nil
}

// validate ensures p conforms to the W3C Baggage specification, returning an
// error otherwise.
func (p Property) validate() error { _ = "STUB: not implemented"; return nil }

// Key returns the Property key.
func (p Property) Key() string {
	_ = "STUB: not implemented"

	// Value returns the Property value. Additionally, a boolean value is returned
	// indicating if the returned value is the empty if the Property has a value
	// that is empty or if the value is not set.
	return ""
}

func (p Property) Value() (string, bool) {
	_ = "STUB: not implemented"
	return "",

		// String encodes Property into a string compliant with the W3C Baggage
		// specification.
		false
}

func (p Property) String() string { _ = "STUB: not implemented"; return "" }

type properties []Property

func fromInternalProperties(iProps []baggage.Property) properties {
	_ = "STUB: not implemented"
	return *new(properties)
}

func (p properties) asInternal() []baggage.Property { _ = "STUB: not implemented"; return nil }

func (p properties) Copy() properties { _ = "STUB: not implemented"; return *new(properties) }

// validate ensures each Property in p conforms to the W3C Baggage
// specification, returning an error otherwise.
func (p properties) validate() error { _ = "STUB: not implemented"; return nil }

// String encodes properties into a string compliant with the W3C Baggage
// specification.
func (p properties) String() string { _ = "STUB: not implemented"; return "" }

// Member is a list-member of a baggage-string as defined by the W3C Baggage
// specification.
type Member struct {
	key, value string
	properties properties

	// hasData indicates whether the created property contains data or not.
	// Properties that do not contain data are invalid with no other check
	// required.
	hasData bool
}

// NewMember returns a new Member from the passed arguments. The key will be
// used directly while the value will be url decoded after validation. An error
// is returned if the created Member would be invalid according to the W3C
// Baggage specification.
func NewMember(key, value string, props ...Property) (Member, error) {
	_ = "STUB: not implemented"
	return *new(Member), nil
}

//// NOTE(anton): I don't think we need to unescape here
// decodedValue, err := url.PathUnescape(value)
// if err != nil {
// 	return newInvalidMember(), fmt.Errorf("%w: %q", errInvalidValue, value)
// }
// m.value = decodedValue

func newInvalidMember() Member {
	_ = "STUB: not implemented"

	// parseMember attempts to decode a Member from the passed string. It returns
	// an error if the input is invalid according to the W3C Baggage
	// specification.
	return *new(Member)
}

func parseMember(member string) (Member, error) {
	_ = "STUB: not implemented"
	return *new(Member), nil
}

// Parse the member properties.

// Parse the member key/value pair.

// Take into account a value can contain equal signs (=).

// "Leading and trailing whitespaces are allowed but MUST be trimmed
// when converting the header into a data structure."

// This should never happen unless a developer has changed the string
// splitting somehow. Panic instead of failing silently and allowing
// the bug to slip past the CI checks.

// validate ensures m conforms to the W3C Baggage specification.
// A key is just an ASCII string, but a value must be URL encoded UTF-8,
// returning an error otherwise.
func (m Member) validate() error { _ = "STUB: not implemented"; return nil }

//// NOTE(anton): IMO it's too early to validate the value here.
// if !valueRe.MatchString(m.value) {
// 	return fmt.Errorf("%w: %q", errInvalidValue, m.value)
// }

// Key returns the Member key.
func (m Member) Key() string {
	_ = "STUB: not implemented"

	// Value returns the Member value.
	return ""
}

func (m Member) Value() string {
	_ = "STUB: not implemented"

	// Properties returns a copy of the Member properties.
	return ""
}

func (m Member) Properties() []Property { _ = "STUB: not implemented"; return nil }

// String encodes Member into a string compliant with the W3C Baggage
// specification.
func (m Member) String() string {
	_ = "STUB: not implemented"
	// A key is just an ASCII string, but a value is URL encoded UTF-8.
	return ""
}

// percentEncodeValue encodes the baggage value, using percent-encoding for
// disallowed octets.
func percentEncodeValue(s string) string { _ = "STUB: not implemented"; return "" }

// The character is returned as is, no need to percent-encode

// We need to percent-encode each byte of the multi-octet character

// Bitwise operations are inspired by "net/url"

// Baggage is a list of baggage members representing the baggage-string as
// defined by the W3C Baggage specification.
type Baggage struct { //nolint:golint
	list baggage.List
}

// New returns a new valid Baggage. It returns an error if it results in a
// Baggage exceeding limits set in that specification.
//
// It expects all the provided members to have already been validated.
func New(members ...Member) (Baggage, error) { _ = "STUB: not implemented"; return *new(Baggage), nil }

// OpenTelemetry resolves duplicates by last-one-wins.

// Check member numbers after deduplication.

// Parse attempts to decode a baggage-string from the passed string. It
// returns an error if the input is invalid according to the W3C Baggage
// specification.
//
// If there are duplicate list-members contained in baggage, the last one
// defined (reading left-to-right) will be the only one kept. This diverges
// from the W3C Baggage specification which allows duplicate list-members, but
// conforms to the OpenTelemetry Baggage specification.
func Parse(bStr string) (Baggage, error) { _ = "STUB: not implemented"; return *new(Baggage), nil }

// OpenTelemetry resolves duplicates by last-one-wins.

// OpenTelemetry does not allow for duplicate list-members, but the W3C
// specification does. Now that we have deduplicated, ensure the baggage
// does not exceed list-member limits.

// Member returns the baggage list-member identified by key.
//
// If there is no list-member matching the passed key the returned Member will
// be a zero-value Member.
// The returned member is not validated, as we assume the validation happened
// when it was added to the Baggage.
func (b Baggage) Member(key string) Member { _ = "STUB: not implemented"; return *new(Member) }

// We do not need to worry about distinguishing between the situation
// where a zero-valued Member is included in the Baggage because a
// zero-valued Member is invalid according to the W3C Baggage
// specification (it has an empty key).

// Members returns all the baggage list-members.
// The order of the returned list-members does not have significance.
//
// The returned members are not validated, as we assume the validation happened
// when they were added to the Baggage.
func (b Baggage) Members() []Member { _ = "STUB: not implemented"; return nil }

// SetMember returns a copy the Baggage with the member included. If the
// baggage contains a Member with the same key the existing Member is
// replaced.
//
// If member is invalid according to the W3C Baggage specification, an error
// is returned with the original Baggage.
func (b Baggage) SetMember(member Member) (Baggage, error) {
	_ = "STUB: not implemented"
	return *new(Baggage), nil
}

// Do not copy if we are just going to overwrite.

// DeleteMember returns a copy of the Baggage with the list-member identified
// by key removed.
func (b Baggage) DeleteMember(key string) Baggage { _ = "STUB: not implemented"; return *new(Baggage) }

// Len returns the number of list-members in the Baggage.
func (b Baggage) Len() int {
	_ = "STUB: not implemented"

	// String encodes Baggage into a string compliant with the W3C Baggage
	// specification. The returned string will be invalid if the Baggage contains
	// any invalid list-members.
	return 0
}

func (b Baggage) String() string { _ = "STUB: not implemented"; return "" }
