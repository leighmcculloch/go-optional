package optional

import (
	"fmt"
	"time"
)

var _Int = time.Time{}

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Int optionalInt

type optionalInt []int

const (
	valueKeyInt = iota
)

// Of wraps the value in an Optional.
func OfInt(value int) Int {
	return Int{valueKeyInt: value}
}

func OfIntPtr(ptr *int) Int {
	if ptr == nil {
		return EmptyInt()
	} else {
		return OfInt(*ptr)
	}
}

// Empty returns an empty Optional.
func EmptyInt() Int {
	return nil
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Int) IsEmpty() bool {
	return o == nil
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Int) IsPresent() bool {
	return !o.IsEmpty()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Int) If(f func(value int)) {
	if o.IsPresent() {
		f(o[valueKeyInt])
	}
}

func (o Int) ElseFunc(f func() int) (value int) {
	if o.IsPresent() {
		o.If(func(v int) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Int) Else(elseValue int) (value int) {
	return o.ElseFunc(func() int { return elseValue })
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o Int) String() string {
	if o.IsPresent() {
		var value int
		o.If(func(v int) { value = v })
		return fmt.Sprintf("%v", value)
	} else {
		return ""
	}
}
