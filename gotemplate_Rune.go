package optional

import (
	"reflect"
	"strconv"
)

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Rune map[keyRune]*rune

type keyRune int

const (
	valueKeyRune keyRune = iota
)

// Of wraps the value in an Optional.
func OfRune(value rune) Rune {
	return Rune{valueKeyRune: &value}
}

func OfRunePtr(ptr *rune) Rune {
	if ptr == nil {
		return EmptyRune()
	} else {
		return OfRune(*ptr)
	}
}

// Empty returns an empty Optional.
func EmptyRune() Rune {
	return nil
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Rune) IsEmpty() bool {
	return o == nil
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Rune) IsPresent() bool {
	return !o.IsEmpty()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Rune) If(f func(value rune)) {
	if o.IsPresent() {
		f(*o[valueKeyRune])
	}
}

func (o Rune) ElseFunc(f func() rune) (value rune) {
	if o.IsEmpty() {
		return f()
	} else {
		o.If(func(v rune) { value = v })
		return
	}
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Rune) Else(elseValue rune) (value rune) {
	return o.ElseFunc(func() rune { return elseValue })
}

func (o Rune) MarshalText() (text []byte, err error) {
	if o == nil {
		return nil, nil
	}
	o.If(func(v rune) {
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Int:
			text = []byte(strconv.FormatInt(rv.Int(), 10))
		}
	})
	return
}

func (o Rune) UnmarshalText(text []byte) error {
	return nil
}
