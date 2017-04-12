package optional

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"
)

var _Int64 = time.Time{}

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Int64 optionalInt64

type optionalInt64 []int64

const (
	valueKeyInt64 = iota
)

// Of wraps the value in an Optional.
func OfInt64(value int64) Int64 {
	return Int64{valueKeyInt64: value}
}

func OfInt64Ptr(ptr *int64) Int64 {
	if ptr == nil {
		return EmptyInt64()
	} else {
		return OfInt64(*ptr)
	}
}

// Empty returns an empty Optional.
func EmptyInt64() Int64 {
	return nil
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Int64) IsPresent() bool {
	return o != nil
}

// If calls the function if there is a value wrapped by this Optional.
func (o Int64) If(f func(value int64)) {
	if o.IsPresent() {
		f(o[valueKeyInt64])
	}
}

func (o Int64) ElseFunc(f func() int64) (value int64) {
	if o.IsPresent() {
		o.If(func(v int64) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Int64) Else(elseValue int64) (value int64) {
	return o.ElseFunc(func() int64 { return elseValue })
}

// ElseZero returns the value wrapped by this Optional, or the zero value of
// the type wrapped if there is no value wrapped by this Optional.
func (o Int64) ElseZero() (value int64) {
	var zero int64
	return o.Else(zero)
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o Int64) String() string {
	if o.IsPresent() {
		var value int64
		o.If(func(v int64) { value = v })
		return fmt.Sprintf("%v", value)
	} else {
		return ""
	}
}

func (o Int64) MarshalJSON() (data []byte, err error) {
	return json.Marshal(o.ElseZero())
}

func (o *Int64) UnmarshalJSON(data []byte) error {
	var v int64
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	*o = OfInt64(v)
	return nil
}

func (o Int64) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.ElseZero(), start)
}

func (o *Int64) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v int64
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = OfInt64(v)
	return nil
}
