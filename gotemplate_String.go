package optional

import (
	"reflect"
	"strconv"
)

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type String map[keyString]string

type keyString int

const (
	valueKeyString keyString = iota
)

// Of wraps the value in an Optional.
func OfString(value string) String {
	return String{valueKeyString: value}
}

func OfStringPtr(ptr *string) String {
	if ptr == nil {
		return EmptyString()
	} else {
		return OfString(*ptr)
	}
}

// Empty returns an empty Optional.
func EmptyString() String {
	return nil
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o String) IsEmpty() bool {
	return o == nil
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o String) IsPresent() bool {
	return !o.IsEmpty()
}

// If calls the function if there is a value wrapped by this Optional.
func (o String) If(f func(value string)) {
	if o.IsPresent() {
		f(o[valueKeyString])
	}
}

func (o String) ElseFunc(f func() string) (value string) {
	if o.IsPresent() {
		o.If(func(v string) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o String) Else(elseValue string) (value string) {
	return o.ElseFunc(func() string { return elseValue })
}

// MarshalText returns text for marshaling this Optional.
func (o String) MarshalText() (text []byte, err error) {
	o.If(func(v string) {
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Int:
			text = []byte(strconv.FormatInt(rv.Int(), 10))
		}
	})
	return
}

// UnmarshalText returns text for marshaling this Optional.
func (o *String) UnmarshalText(text []byte) error {
	*o = EmptyString()
	return nil
}
