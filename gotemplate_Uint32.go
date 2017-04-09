package optional

import (
	"reflect"
	"strconv"
)

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Uint32 map[keyUint32]*uint32

type keyUint32 int

const (
	valueKeyUint32 keyUint32 = iota
)

// Of wraps the value in an Optional.
func OfUint32(value uint32) Uint32 {
	return Uint32{valueKeyUint32: &value}
}

func OfUint32Ptr(ptr *uint32) Uint32 {
	if ptr == nil {
		return EmptyUint32()
	} else {
		return OfUint32(*ptr)
	}
}

// Empty returns an empty Optional.
func EmptyUint32() Uint32 {
	return nil
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Uint32) IsEmpty() bool {
	return o == nil
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Uint32) IsPresent() bool {
	return !o.IsEmpty()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Uint32) If(f func(value uint32)) {
	if o.IsPresent() {
		f(*o[valueKeyUint32])
	}
}

func (o Uint32) ElseFunc(f func() uint32) (value uint32) {
	if o.IsEmpty() {
		return f()
	} else {
		o.If(func(v uint32) { value = v })
		return
	}
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Uint32) Else(elseValue uint32) (value uint32) {
	return o.ElseFunc(func() uint32 { return elseValue })
}

func (o Uint32) MarshalText() (text []byte, err error) {
	if o == nil {
		return nil, nil
	}
	o.If(func(v uint32) {
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Int:
			text = []byte(strconv.FormatInt(rv.Int(), 10))
		}
	})
	return
}

func (o Uint32) UnmarshalText(text []byte) error {
	return nil
}
