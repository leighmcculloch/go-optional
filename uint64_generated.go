package optional

import (
	"fmt"
	"time"
)

var _Uint64 = time.Time{}

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Uint64 optionalUint64

type optionalUint64 []uint64

const (
	valueKeyUint64 = iota
)

// Of wraps the value in an Optional.
func OfUint64(value uint64) Uint64 {
	return Uint64{valueKeyUint64: value}
}

func OfUint64Ptr(ptr *uint64) Uint64 {
	if ptr == nil {
		return EmptyUint64()
	} else {
		return OfUint64(*ptr)
	}
}

// Empty returns an empty Optional.
func EmptyUint64() Uint64 {
	return nil
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Uint64) IsEmpty() bool {
	return o == nil
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Uint64) IsPresent() bool {
	return !o.IsEmpty()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Uint64) If(f func(value uint64)) {
	if o.IsPresent() {
		f(o[valueKeyUint64])
	}
}

func (o Uint64) ElseFunc(f func() uint64) (value uint64) {
	if o.IsPresent() {
		o.If(func(v uint64) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Uint64) Else(elseValue uint64) (value uint64) {
	return o.ElseFunc(func() uint64 { return elseValue })
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o Uint64) String() string {
	if o.IsPresent() {
		var value uint64
		o.If(func(v uint64) { value = v })
		return fmt.Sprintf("%v", value)
	} else {
		return ""
	}
}
