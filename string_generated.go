package optional

import (
	"fmt"
	"time"
)

var _String = time.Time{}

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type String optionalString

type optionalString []string

const (
	valueKeyString = iota
)

// Of wraps the value in an Optional.
func OfString(value string) String {
	return String{valueKeyString: value}
}

func OfStringPtr(ptr *string) String {
	if ptr == nil {
		return EmptyString()
	} else {
		return OfString(*ptr)
	}
}

// Empty returns an empty Optional.
func EmptyString() String {
	return nil
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o String) IsEmpty() bool {
	return o == nil
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o String) IsPresent() bool {
	return !o.IsEmpty()
}

// If calls the function if there is a value wrapped by this Optional.
func (o String) If(f func(value string)) {
	if o.IsPresent() {
		f(o[valueKeyString])
	}
}

func (o String) ElseFunc(f func() string) (value string) {
	if o.IsPresent() {
		o.If(func(v string) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o String) Else(elseValue string) (value string) {
	return o.ElseFunc(func() string { return elseValue })
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o String) String() string {
	if o.IsPresent() {
		var value string
		o.If(func(v string) { value = v })
		return fmt.Sprintf("%v", value)
	} else {
		return ""
	}
}
