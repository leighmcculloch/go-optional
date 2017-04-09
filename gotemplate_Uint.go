package optional

import (
	"reflect"
	"strconv"
)

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Uint map[keyUint]uint

type keyUint int

const (
	valueKeyUint keyUint = iota
)

// Of wraps the value in an Optional.
func OfUint(value uint) Uint {
	return Uint{valueKeyUint: value}
}

func OfUintPtr(ptr *uint) Uint {
	if ptr == nil {
		return EmptyUint()
	} else {
		return OfUint(*ptr)
	}
}

// Empty returns an empty Optional.
func EmptyUint() Uint {
	return nil
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Uint) IsEmpty() bool {
	return o == nil
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Uint) IsPresent() bool {
	return !o.IsEmpty()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Uint) If(f func(value uint)) {
	if o.IsPresent() {
		f(o[valueKeyUint])
	}
}

func (o Uint) ElseFunc(f func() uint) (value uint) {
	if o.IsPresent() {
		o.If(func(v uint) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Uint) Else(elseValue uint) (value uint) {
	return o.ElseFunc(func() uint { return elseValue })
}

// MarshalText returns text for marshaling this Optional.
func (o Uint) MarshalText() (text []byte, err error) {
	o.If(func(v uint) {
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Int:
			text = []byte(strconv.FormatInt(rv.Int(), 10))
		}
	})
	return
}

// UnmarshalText returns text for marshaling this Optional.
func (o *Uint) UnmarshalText(text []byte) error {
	*o = EmptyUint()
	return nil
}
