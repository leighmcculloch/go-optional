package optional

import (
	"fmt"
	"time"
)

var _Byte = time.Time{}

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Byte optionalByte

type optionalByte []byte

const (
	valueKeyByte = iota
)

// Of wraps the value in an Optional.
func OfByte(value byte) Byte {
	return Byte{valueKeyByte: value}
}

func OfBytePtr(ptr *byte) Byte {
	if ptr == nil {
		return EmptyByte()
	} else {
		return OfByte(*ptr)
	}
}

// Empty returns an empty Optional.
func EmptyByte() Byte {
	return nil
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Byte) IsEmpty() bool {
	return o == nil
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Byte) IsPresent() bool {
	return !o.IsEmpty()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Byte) If(f func(value byte)) {
	if o.IsPresent() {
		f(o[valueKeyByte])
	}
}

func (o Byte) ElseFunc(f func() byte) (value byte) {
	if o.IsPresent() {
		o.If(func(v byte) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Byte) Else(elseValue byte) (value byte) {
	return o.ElseFunc(func() byte { return elseValue })
}

// ElseZero returns the value wrapped by this Optional, or the zero value of
// the type wrapped if there is no value wrapped by this Optional.
func (o Byte) ElseZero() (value byte) {
	var zero byte
	return o.Else(zero)
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o Byte) String() string {
	if o.IsPresent() {
		var value byte
		o.If(func(v byte) { value = v })
		return fmt.Sprintf("%v", value)
	} else {
		return ""
	}
}
