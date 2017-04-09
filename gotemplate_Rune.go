package optional

import (
	"fmt"
)

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Rune map[keyRune]rune

type keyRune int

const (
	valueKeyRune keyRune = iota
)

// Of wraps the value in an Optional.
func OfRune(value rune) Rune {
	return Rune{valueKeyRune: value}
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
		f(o[valueKeyRune])
	}
}

func (o Rune) ElseFunc(f func() rune) (value rune) {
	if o.IsPresent() {
		o.If(func(v rune) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Rune) Else(elseValue rune) (value rune) {
	return o.ElseFunc(func() rune { return elseValue })
}

func (o Rune) String() string {
	if o.IsPresent() {
		var value rune
		o.If(func(v rune) { value = v })
		return fmt.Sprintf("%v", value)
	} else {
		return ""
	}
}
