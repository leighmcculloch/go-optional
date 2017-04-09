package optional

import (
	"reflect"
	"strconv"
)

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Uint16 map[keyUint16]uint16

type keyUint16 int

const (
	valueKeyUint16 keyUint16 = iota
)

// Of wraps the value in an Optional.
func OfUint16(value uint16) Uint16 {
	return Uint16{valueKeyUint16: value}
}

func OfUint16Ptr(ptr *uint16) Uint16 {
	if ptr == nil {
		return EmptyUint16()
	} else {
		return OfUint16(*ptr)
	}
}

// Empty returns an empty Optional.
func EmptyUint16() Uint16 {
	return nil
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Uint16) IsEmpty() bool {
	return o == nil
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Uint16) IsPresent() bool {
	return !o.IsEmpty()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Uint16) If(f func(value uint16)) {
	if o.IsPresent() {
		f(o[valueKeyUint16])
	}
}

func (o Uint16) ElseFunc(f func() uint16) (value uint16) {
	if o.IsPresent() {
		o.If(func(v uint16) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Uint16) Else(elseValue uint16) (value uint16) {
	return o.ElseFunc(func() uint16 { return elseValue })
}

// MarshalText returns text for marshaling this Optional.
func (o Uint16) MarshalText() (text []byte, err error) {
	o.If(func(v uint16) {
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Int:
			text = []byte(strconv.FormatInt(rv.Int(), 10))
		}
	})
	return
}

// UnmarshalText returns text for marshaling this Optional.
func (o *Uint16) UnmarshalText(text []byte) error {
	*o = EmptyUint16()
	return nil
}
