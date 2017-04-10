package optional

import (
	"fmt"
)

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Complex128 optionalComplex128

type optionalComplex128 []complex128

const (
	valueKeyComplex128 = iota
)

// Of wraps the value in an Optional.
func OfComplex128(value complex128) Complex128 {
	return Complex128{valueKeyComplex128: value}
}

func OfComplex128Ptr(ptr *complex128) Complex128 {
	if ptr == nil {
		return EmptyComplex128()
	} else {
		return OfComplex128(*ptr)
	}
}

// Empty returns an empty Optional.
func EmptyComplex128() Complex128 {
	return nil
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Complex128) IsEmpty() bool {
	return o == nil
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Complex128) IsPresent() bool {
	return !o.IsEmpty()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Complex128) If(f func(value complex128)) {
	if o.IsPresent() {
		f(o[valueKeyComplex128])
	}
}

func (o Complex128) ElseFunc(f func() complex128) (value complex128) {
	if o.IsPresent() {
		o.If(func(v complex128) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Complex128) Else(elseValue complex128) (value complex128) {
	return o.ElseFunc(func() complex128 { return elseValue })
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o Complex128) String() string {
	if o.IsPresent() {
		var value complex128
		o.If(func(v complex128) { value = v })
		return fmt.Sprintf("%v", value)
	} else {
		return ""
	}
}
