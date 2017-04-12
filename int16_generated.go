package optional

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"
)

var _Int16 = time.Time{}

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Int16 optionalInt16

type optionalInt16 []int16

const (
	valueKeyInt16 = iota
)

// Of wraps the value in an Optional.
func OfInt16(value int16) Int16 {
	return Int16{valueKeyInt16: value}
}

func OfInt16Ptr(ptr *int16) Int16 {
	if ptr == nil {
		return EmptyInt16()
	} else {
		return OfInt16(*ptr)
	}
}

// Empty returns an empty Optional.
func EmptyInt16() Int16 {
	return nil
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Int16) IsPresent() bool {
	return o != nil
}

// If calls the function if there is a value wrapped by this Optional.
func (o Int16) If(f func(value int16)) {
	if o.IsPresent() {
		f(o[valueKeyInt16])
	}
}

func (o Int16) ElseFunc(f func() int16) (value int16) {
	if o.IsPresent() {
		o.If(func(v int16) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Int16) Else(elseValue int16) (value int16) {
	return o.ElseFunc(func() int16 { return elseValue })
}

// ElseZero returns the value wrapped by this Optional, or the zero value of
// the type wrapped if there is no value wrapped by this Optional.
func (o Int16) ElseZero() (value int16) {
	var zero int16
	return o.Else(zero)
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o Int16) String() string {
	if o.IsPresent() {
		var value int16
		o.If(func(v int16) { value = v })
		return fmt.Sprintf("%v", value)
	} else {
		return ""
	}
}

func (o Int16) MarshalJSON() (data []byte, err error) {
	return json.Marshal(o.ElseZero())
}

func (o *Int16) UnmarshalJSON(data []byte) error {
	var v int16
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	*o = OfInt16(v)
	return nil
}

func (o Int16) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.ElseZero(), start)
}

func (o *Int16) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v int16
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = OfInt16(v)
	return nil
}
