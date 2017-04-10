package optional

import (
	"fmt"
	"time"
)

var _Uintptr = time.Time{}

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Uintptr optionalUintptr

type optionalUintptr []uintptr

const (
	valueKeyUintptr = iota
)

// Of wraps the value in an Optional.
func OfUintptr(value uintptr) Uintptr {
	return Uintptr{valueKeyUintptr: value}
}

func OfUintptrPtr(ptr *uintptr) Uintptr {
	if ptr == nil {
		return EmptyUintptr()
	} else {
		return OfUintptr(*ptr)
	}
}

// Empty returns an empty Optional.
func EmptyUintptr() Uintptr {
	return nil
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Uintptr) IsEmpty() bool {
	return o == nil
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Uintptr) IsPresent() bool {
	return !o.IsEmpty()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Uintptr) If(f func(value uintptr)) {
	if o.IsPresent() {
		f(o[valueKeyUintptr])
	}
}

func (o Uintptr) ElseFunc(f func() uintptr) (value uintptr) {
	if o.IsPresent() {
		o.If(func(v uintptr) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Uintptr) Else(elseValue uintptr) (value uintptr) {
	return o.ElseFunc(func() uintptr { return elseValue })
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o Uintptr) String() string {
	if o.IsPresent() {
		var value uintptr
		o.If(func(v uintptr) { value = v })
		return fmt.Sprintf("%v", value)
	} else {
		return ""
	}
}
