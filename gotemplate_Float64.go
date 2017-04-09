package optional

import (
	"fmt"
)

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Float64 map[keyFloat64]float64

type keyFloat64 int

const (
	valueKeyFloat64 keyFloat64 = iota
)

// Of wraps the value in an Optional.
func OfFloat64(value float64) Float64 {
	return Float64{valueKeyFloat64: value}
}

func OfFloat64Ptr(ptr *float64) Float64 {
	if ptr == nil {
		return EmptyFloat64()
	} else {
		return OfFloat64(*ptr)
	}
}

// Empty returns an empty Optional.
func EmptyFloat64() Float64 {
	return nil
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Float64) IsEmpty() bool {
	return o == nil
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Float64) IsPresent() bool {
	return !o.IsEmpty()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Float64) If(f func(value float64)) {
	if o.IsPresent() {
		f(o[valueKeyFloat64])
	}
}

func (o Float64) ElseFunc(f func() float64) (value float64) {
	if o.IsPresent() {
		o.If(func(v float64) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Float64) Else(elseValue float64) (value float64) {
	return o.ElseFunc(func() float64 { return elseValue })
}

func (o Float64) String() string {
	if o.IsPresent() {
		var value float64
		o.If(func(v float64) { value = v })
		return fmt.Sprintf("%v", value)
	} else {
		return ""
	}
}
