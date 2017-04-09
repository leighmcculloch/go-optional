package optional

import (
	"reflect"
	"strconv"
)

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Error map[keyError]error

type keyError int

const (
	valueKeyError keyError = iota
)

// Of wraps the value in an Optional.
func OfError(value error) Error {
	return Error{valueKeyError: value}
}

func OfErrorPtr(ptr *error) Error {
	if ptr == nil {
		return EmptyError()
	} else {
		return OfError(*ptr)
	}
}

// Empty returns an empty Optional.
func EmptyError() Error {
	return nil
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Error) IsEmpty() bool {
	return o == nil
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Error) IsPresent() bool {
	return !o.IsEmpty()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Error) If(f func(value error)) {
	if o.IsPresent() {
		f(o[valueKeyError])
	}
}

func (o Error) ElseFunc(f func() error) (value error) {
	if o.IsEmpty() {
		return f()
	} else {
		o.If(func(v error) { value = v })
		return
	}
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Error) Else(elseValue error) (value error) {
	return o.ElseFunc(func() error { return elseValue })
}

func (o Error) MarshalText() (text []byte, err error) {
	if o == nil {
		return nil, nil
	}
	o.If(func(v error) {
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Int:
			text = []byte(strconv.FormatInt(rv.Int(), 10))
		}
	})
	return
}

func (o Error) UnmarshalText(text []byte) error {
	return nil
}
