package optional

import (
	"fmt"
	"time"
)

var _Bool = time.Time{}

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Bool optionalBool

type optionalBool []bool

const (
	valueKeyBool = iota
)

// Of wraps the value in an Optional.
func OfBool(value bool) Bool {
	return Bool{valueKeyBool: value}
}

func OfBoolPtr(ptr *bool) Bool {
	if ptr == nil {
		return EmptyBool()
	} else {
		return OfBool(*ptr)
	}
}

// Empty returns an empty Optional.
func EmptyBool() Bool {
	return nil
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Bool) IsEmpty() bool {
	return o == nil
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Bool) IsPresent() bool {
	return !o.IsEmpty()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Bool) If(f func(value bool)) {
	if o.IsPresent() {
		f(o[valueKeyBool])
	}
}

func (o Bool) ElseFunc(f func() bool) (value bool) {
	if o.IsPresent() {
		o.If(func(v bool) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Bool) Else(elseValue bool) (value bool) {
	return o.ElseFunc(func() bool { return elseValue })
}

// ElseZero returns the value wrapped by this Optional, or the zero value of
// the type wrapped if there is no value wrapped by this Optional.
func (o Bool) ElseZero() (value bool) {
	var zero bool
	return o.Else(zero)
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o Bool) String() string {
	if o.IsPresent() {
		var value bool
		o.If(func(v bool) { value = v })
		return fmt.Sprintf("%v", value)
	} else {
		return ""
	}
}
