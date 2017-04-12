package optional

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"
)

var _Complex64 = time.Time{}

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Complex64 optionalComplex64

type optionalComplex64 []complex64

const (
	valueKeyComplex64 = iota
)

// Of wraps the value in an Optional.
func OfComplex64(value complex64) Complex64 {
	return Complex64{valueKeyComplex64: value}
}

func OfComplex64Ptr(ptr *complex64) Complex64 {
	if ptr == nil {
		return EmptyComplex64()
	} else {
		return OfComplex64(*ptr)
	}
}

// Empty returns an empty Optional.
func EmptyComplex64() Complex64 {
	return nil
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o Complex64) IsEmpty() bool {
	return o == nil
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Complex64) IsPresent() bool {
	return !o.IsEmpty()
}

// If calls the function if there is a value wrapped by this Optional.
func (o Complex64) If(f func(value complex64)) {
	if o.IsPresent() {
		f(o[valueKeyComplex64])
	}
}

func (o Complex64) ElseFunc(f func() complex64) (value complex64) {
	if o.IsPresent() {
		o.If(func(v complex64) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Complex64) Else(elseValue complex64) (value complex64) {
	return o.ElseFunc(func() complex64 { return elseValue })
}

// ElseZero returns the value wrapped by this Optional, or the zero value of
// the type wrapped if there is no value wrapped by this Optional.
func (o Complex64) ElseZero() (value complex64) {
	var zero complex64
	return o.Else(zero)
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o Complex64) String() string {
	if o.IsPresent() {
		var value complex64
		o.If(func(v complex64) { value = v })
		return fmt.Sprintf("%v", value)
	} else {
		return ""
	}
}

func (o Complex64) MarshalJSON() (data []byte, err error) {
	return json.Marshal(o.ElseZero())
}

func (o *Complex64) UnmarshalJSON(data []byte) error {
	var v complex64
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	*o = OfComplex64(v)
	return nil
}

func (o Complex64) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.ElseZero(), start)
}

func (o *Complex64) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v complex64
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = OfComplex64(v)
	return nil
}
