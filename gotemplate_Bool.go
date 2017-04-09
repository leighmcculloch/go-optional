package optional

import (
	"reflect"
	"strconv"
)

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Bool map[keyBool]*bool

type keyBool int

const (
	valueKeyBool keyBool = iota
)

// Of wraps the value in an Optional.
func OfBool(value bool) Bool {
	return Bool{valueKeyBool: &value}
}

func OfBoolPtr(ptr *bool) Bool {
	if ptr == nil {
		return EmptyBool()
	} else {
		return OfBool(*ptr)
	}
}

// Empty returns an empty Optional.
func EmptyBool() Bool {
	return nil
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Bool) IsEmpty() bool {
	return o == nil
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Bool) IsPresent() bool {
	return !o.IsEmpty()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Bool) If(f func(value bool)) {
	if o.IsPresent() {
		f(*o[valueKeyBool])
	}
}

func (o Bool) ElseFunc(f func() bool) (value bool) {
	if o.IsEmpty() {
		return f()
	} else {
		o.If(func(v bool) { value = v })
		return
	}
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Bool) Else(elseValue bool) (value bool) {
	return o.ElseFunc(func() bool { return elseValue })
}

func (o Bool) MarshalText() (text []byte, err error) {
	if o == nil {
		return nil, nil
	}
	o.If(func(v bool) {
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Int:
			text = []byte(strconv.FormatInt(rv.Int(), 10))
		}
	})
	return
}

func (o Bool) UnmarshalText(text []byte) error {
	return nil
}
