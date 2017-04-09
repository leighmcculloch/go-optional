package optional

import (
	"fmt"
)

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Complex64 map[keyComplex64]complex64

type keyComplex64 int

const (
	valueKeyComplex64 keyComplex64 = iota
)

// Of wraps the value in an Optional.
func OfComplex64(value complex64) Complex64 {
	return Complex64{valueKeyComplex64: value}
}

func OfComplex64Ptr(ptr *complex64) Complex64 {
	if ptr == nil {
		return EmptyComplex64()
	} else {
		return OfComplex64(*ptr)
	}
}

// Empty returns an empty Optional.
func EmptyComplex64() Complex64 {
	return nil
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Complex64) IsEmpty() bool {
	return o == nil
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Complex64) IsPresent() bool {
	return !o.IsEmpty()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Complex64) If(f func(value complex64)) {
	if o.IsPresent() {
		f(o[valueKeyComplex64])
	}
}

func (o Complex64) ElseFunc(f func() complex64) (value complex64) {
	if o.IsPresent() {
		o.If(func(v complex64) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Complex64) Else(elseValue complex64) (value complex64) {
	return o.ElseFunc(func() complex64 { return elseValue })
}

func (o Complex64) String() string {
	if o.IsPresent() {
		var value complex64
		o.If(func(v complex64) { value = v })
		return fmt.Sprintf("%v", value)
	} else {
		return ""
	}
}
