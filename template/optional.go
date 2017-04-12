package template

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"
)

var _ = time.Time{}

// template type Optional(T)

type T string

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Optional optional

type optional []T

const (
	valueKey = iota
)

// Of wraps the value in an optional.
func Of(value T) Optional {
	return Optional{valueKey: value}
}

func OfOptionalPtr(ptr *T) Optional {
	if ptr == nil {
		return Empty()
	} else {
		return Of(*ptr)
	}
}

// Empty returns an empty optional.
func Empty() Optional {
	return nil
}

// Get returns the value wrapped by this optional, and an ok signal for whether a value was wrapped.
func (o Optional) Get() (value T, ok bool) {
	o.If(func(v T) {
		value = v
		ok = true
	})
	return
}

// IsPresent returns true if there is a value wrapped by this optional.
func (o Optional) IsPresent() bool {
	return o != nil
}

// If calls the function if there is a value wrapped by this optional.
func (o Optional) If(f func(value T)) {
	if o.IsPresent() {
		f(o[valueKey])
	}
}

func (o Optional) ElseFunc(f func() T) (value T) {
	if o.IsPresent() {
		o.If(func(v T) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this optional, or the value passed in if
// there is no value wrapped by this optional.
func (o Optional) Else(elseValue T) (value T) {
	return o.ElseFunc(func() T { return elseValue })
}

// ElseZero returns the value wrapped by this optional, or the zero value of
// the type wrapped if there is no value wrapped by this optional.
func (o Optional) ElseZero() (value T) {
	var zero T
	return o.Else(zero)
}

// String returns the string representation of the wrapped value, or the string
// representation of the zero value of the type wrapped if there is no value
// wrapped by this optional.
func (o Optional) String() string {
	return fmt.Sprintf("%v", o.ElseZero())
}

// MarshalJSON marshals the value being wrapped to JSON. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Optional) MarshalJSON() (data []byte, err error) {
	return json.Marshal(o.ElseZero())
}

// UnmarshalJSON unmarshals the JSON into a value wrapped by this optional.
func (o *Optional) UnmarshalJSON(data []byte) error {
	var v T
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	*o = Of(v)
	return nil
}

// MarshalXML marshals the value being wrapped to XML. If there is no vale
// being wrapped, the zero value of its type is marshaled.
func (o Optional) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.ElseZero(), start)
}

// UnmarshalXML unmarshals the XML into a value wrapped by this optional.
func (o *Optional) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v T
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = Of(v)
	return nil
}
