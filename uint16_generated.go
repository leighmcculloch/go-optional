package optional

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"
)

var _Uint16 = time.Time{}

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Uint16 optionalUint16

type optionalUint16 []uint16

const (
	valueKeyUint16 = iota
)

// Of wraps the value in an Optional.
func OfUint16(value uint16) Uint16 {
	return Uint16{valueKeyUint16: value}
}

func OfUint16Ptr(ptr *uint16) Uint16 {
	if ptr == nil {
		return EmptyUint16()
	} else {
		return OfUint16(*ptr)
	}
}

// Empty returns an empty Optional.
func EmptyUint16() Uint16 {
	return nil
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Uint16) IsEmpty() bool {
	return o == nil
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Uint16) IsPresent() bool {
	return !o.IsEmpty()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Uint16) If(f func(value uint16)) {
	if o.IsPresent() {
		f(o[valueKeyUint16])
	}
}

func (o Uint16) ElseFunc(f func() uint16) (value uint16) {
	if o.IsPresent() {
		o.If(func(v uint16) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Uint16) Else(elseValue uint16) (value uint16) {
	return o.ElseFunc(func() uint16 { return elseValue })
}

// ElseZero returns the value wrapped by this Optional, or the zero value of
// the type wrapped if there is no value wrapped by this Optional.
func (o Uint16) ElseZero() (value uint16) {
	var zero uint16
	return o.Else(zero)
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o Uint16) String() string {
	if o.IsPresent() {
		var value uint16
		o.If(func(v uint16) { value = v })
		return fmt.Sprintf("%v", value)
	} else {
		return ""
	}
}

func (o Uint16) MarshalJSON() (data []byte, err error) {
	return json.Marshal(o.ElseZero())
}

func (o *Uint16) UnmarshalJSON(data []byte) error {
	var v uint16
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	*o = OfUint16(v)
	return nil
}

func (o Uint16) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.ElseZero(), start)
}

func (o *Uint16) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v uint16
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = OfUint16(v)
	return nil
}
