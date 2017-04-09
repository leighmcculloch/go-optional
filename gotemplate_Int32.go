package optional

import (
	"reflect"
	"strconv"
)

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Int32 map[keyInt32]*int32

type keyInt32 int

const (
	valueKeyInt32 keyInt32 = iota
)

// Of wraps the value in an Optional.
func OfInt32(value int32) Int32 {
	return Int32{valueKeyInt32: &value}
}

func OfInt32Ptr(ptr *int32) Int32 {
	if ptr == nil {
		return EmptyInt32()
	} else {
		return OfInt32(*ptr)
	}
}

// Empty returns an empty Optional.
func EmptyInt32() Int32 {
	return nil
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Int32) IsEmpty() bool {
	return o == nil
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Int32) IsPresent() bool {
	return !o.IsEmpty()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Int32) If(f func(value int32)) {
	if o.IsPresent() {
		f(*o[valueKeyInt32])
	}
}

func (o Int32) ElseFunc(f func() int32) (value int32) {
	if o.IsEmpty() {
		return f()
	} else {
		o.If(func(v int32) { value = v })
		return
	}
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Int32) Else(elseValue int32) (value int32) {
	return o.ElseFunc(func() int32 { return elseValue })
}

func (o Int32) MarshalText() (text []byte, err error) {
	if o == nil {
		return nil, nil
	}
	o.If(func(v int32) {
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Int:
			text = []byte(strconv.FormatInt(rv.Int(), 10))
		}
	})
	return
}

func (o Int32) UnmarshalText(text []byte) error {
	return nil
}
