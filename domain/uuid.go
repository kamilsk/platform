package domain

import "regexp"

var uuid = regexp.MustCompile(`(?i:^[0-9A-F]{8}-[0-9A-F]{4}-[4][0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}$)`)

// UUID wraps built-in string type and provides useful methods above it.
type UUID string

// IsEmpty returns true if the UUID has empty value.
func (value UUID) IsEmpty() bool {
	return value == ""
}

// IsValid returns true if the UUID is compatible with RFC 4122.
func (value UUID) IsValid() bool {
	return !(value == "") && uuid.MatchString(string(value)) // IsEmpty and String were inlined manually
}

// String implements built-in fmt.Stringer interface and returns string representation of the UUID.
func (value UUID) String() string {
	return string(value)
}
