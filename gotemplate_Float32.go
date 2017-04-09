package optional

import (
	"reflect"
	"strconv"
)

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Float32 map[keyFloat32]float32

type keyFloat32 int

const (
	valueKeyFloat32 keyFloat32 = iota
)

// Of wraps the value in an Optional.
func OfFloat32(value float32) Float32 {
	return Float32{valueKeyFloat32: value}
}

func OfFloat32Ptr(ptr *float32) Float32 {
	if ptr == nil {
		return EmptyFloat32()
	} else {
		return OfFloat32(*ptr)
	}
}

// Empty returns an empty Optional.
func EmptyFloat32() Float32 {
	return nil
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Float32) IsEmpty() bool {
	return o == nil
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Float32) IsPresent() bool {
	return !o.IsEmpty()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Float32) If(f func(value float32)) {
	if o.IsPresent() {
		f(o[valueKeyFloat32])
	}
}

func (o Float32) ElseFunc(f func() float32) (value float32) {
	if o.IsEmpty() {
		return f()
	} else {
		o.If(func(v float32) { value = v })
		return
	}
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Float32) Else(elseValue float32) (value float32) {
	return o.ElseFunc(func() float32 { return elseValue })
}

func (o Float32) MarshalText() (text []byte, err error) {
	if o == nil {
		return nil, nil
	}
	o.If(func(v float32) {
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Int:
			text = []byte(strconv.FormatInt(rv.Int(), 10))
		}
	})
	return
}

func (o Float32) UnmarshalText(text []byte) error {
	return nil
}
