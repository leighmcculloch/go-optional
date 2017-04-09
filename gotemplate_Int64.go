package optional

import (
	"reflect"
	"strconv"
)

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Int64 map[keyInt64]int64

type keyInt64 int

const (
	valueKeyInt64 keyInt64 = iota
)

// Of wraps the value in an Optional.
func OfInt64(value int64) Int64 {
	return Int64{valueKeyInt64: value}
}

func OfInt64Ptr(ptr *int64) Int64 {
	if ptr == nil {
		return EmptyInt64()
	} else {
		return OfInt64(*ptr)
	}
}

// Empty returns an empty Optional.
func EmptyInt64() Int64 {
	return nil
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Int64) IsEmpty() bool {
	return o == nil
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Int64) IsPresent() bool {
	return !o.IsEmpty()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Int64) If(f func(value int64)) {
	if o.IsPresent() {
		f(o[valueKeyInt64])
	}
}

func (o Int64) ElseFunc(f func() int64) (value int64) {
	if o.IsPresent() {
		o.If(func(v int64) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Int64) Else(elseValue int64) (value int64) {
	return o.ElseFunc(func() int64 { return elseValue })
}

// MarshalText returns text for marshaling this Optional.
func (o Int64) MarshalText() (text []byte, err error) {
	o.If(func(v int64) {
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Int:
			text = []byte(strconv.FormatInt(rv.Int(), 10))
		}
	})
	return
}

// UnmarshalText returns text for marshaling this Optional.
func (o *Int64) UnmarshalText(text []byte) error {
	*o = EmptyInt64()
	return nil
}
