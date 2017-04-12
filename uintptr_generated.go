package optional

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"
)

var _Uintptr = time.Time{}

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Uintptr optionalUintptr

type optionalUintptr []uintptr

const (
	valueKeyUintptr = iota
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

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Uintptr) IsPresent() bool {
	return o != nil
}

// If calls the function if there is a value wrapped by this Optional.
func (o Uintptr) If(f func(value uintptr)) {
	if o.IsPresent() {
		f(o[valueKeyUintptr])
	}
}

func (o Uintptr) ElseFunc(f func() uintptr) (value uintptr) {
	if o.IsPresent() {
		o.If(func(v uintptr) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Uintptr) Else(elseValue uintptr) (value uintptr) {
	return o.ElseFunc(func() uintptr { return elseValue })
}

// ElseZero returns the value wrapped by this Optional, or the zero value of
// the type wrapped if there is no value wrapped by this Optional.
func (o Uintptr) ElseZero() (value uintptr) {
	var zero uintptr
	return o.Else(zero)
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o Uintptr) String() string {
	if o.IsPresent() {
		var value uintptr
		o.If(func(v uintptr) { value = v })
		return fmt.Sprintf("%v", value)
	} else {
		return ""
	}
}

func (o Uintptr) MarshalJSON() (data []byte, err error) {
	return json.Marshal(o.ElseZero())
}

func (o *Uintptr) UnmarshalJSON(data []byte) error {
	var v uintptr
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	*o = OfUintptr(v)
	return nil
}

func (o Uintptr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.ElseZero(), start)
}

func (o *Uintptr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v uintptr
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = OfUintptr(v)
	return nil
}
