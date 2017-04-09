package optional

import (
	"reflect"
	"strconv"
)

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Int8 map[keyInt8]int8

type keyInt8 int

const (
	valueKeyInt8 keyInt8 = iota
)

// Of wraps the value in an Optional.
func OfInt8(value int8) Int8 {
	return Int8{valueKeyInt8: value}
}

func OfInt8Ptr(ptr *int8) Int8 {
	if ptr == nil {
		return EmptyInt8()
	} else {
		return OfInt8(*ptr)
	}
}

// Empty returns an empty Optional.
func EmptyInt8() Int8 {
	return nil
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Int8) IsEmpty() bool {
	return o == nil
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Int8) IsPresent() bool {
	return !o.IsEmpty()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Int8) If(f func(value int8)) {
	if o.IsPresent() {
		f(o[valueKeyInt8])
	}
}

func (o Int8) ElseFunc(f func() int8) (value int8) {
	if o.IsPresent() {
		o.If(func(v int8) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Int8) Else(elseValue int8) (value int8) {
	return o.ElseFunc(func() int8 { return elseValue })
}

// MarshalText returns text for marshaling this Optional.
func (o Int8) MarshalText() (text []byte, err error) {
	o.If(func(v int8) {
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Int:
			text = []byte(strconv.FormatInt(rv.Int(), 10))
		}
	})
	return
}

// UnmarshalText returns text for marshaling this Optional.
func (o *Int8) UnmarshalText(text []byte) error {
	*o = EmptyInt8()
	return nil
}
