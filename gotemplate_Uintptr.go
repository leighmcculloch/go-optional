package optional

import (
	"reflect"
	"strconv"
)

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Uintptr map[keyUintptr]uintptr

type keyUintptr int

const (
	valueKeyUintptr keyUintptr = iota
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
	if o.IsEmpty() {
		return f()
	} else {
		o.If(func(v uintptr) { value = v })
		return
	}
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Uintptr) Else(elseValue uintptr) (value uintptr) {
	return o.ElseFunc(func() uintptr { return elseValue })
}

func (o Uintptr) MarshalText() (text []byte, err error) {
	if o == nil {
		return nil, nil
	}
	o.If(func(v uintptr) {
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Int:
			text = []byte(strconv.FormatInt(rv.Int(), 10))
		}
	})
	return
}

func (o Uintptr) UnmarshalText(text []byte) error {
	return nil
}
