package optional

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"
)

var _Time = time.Time{}

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Time optionalTime

type optionalTime []time.Time

const (
	valueKeyTime = iota
)

// Of wraps the value in an Optional.
func OfTime(value time.Time) Time {
	return Time{valueKeyTime: value}
}

func OfTimePtr(ptr *time.Time) Time {
	if ptr == nil {
		return EmptyTime()
	} else {
		return OfTime(*ptr)
	}
}

// Empty returns an empty Optional.
func EmptyTime() Time {
	return nil
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o Time) IsPresent() bool {
	return o != nil
}

// If calls the function if there is a value wrapped by this Optional.
func (o Time) If(f func(value time.Time)) {
	if o.IsPresent() {
		f(o[valueKeyTime])
	}
}

func (o Time) ElseFunc(f func() time.Time) (value time.Time) {
	if o.IsPresent() {
		o.If(func(v time.Time) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Time) Else(elseValue time.Time) (value time.Time) {
	return o.ElseFunc(func() time.Time { return elseValue })
}

// ElseZero returns the value wrapped by this Optional, or the zero value of
// the type wrapped if there is no value wrapped by this Optional.
func (o Time) ElseZero() (value time.Time) {
	var zero time.Time
	return o.Else(zero)
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o Time) String() string {
	if o.IsPresent() {
		var value time.Time
		o.If(func(v time.Time) { value = v })
		return fmt.Sprintf("%v", value)
	} else {
		return ""
	}
}

func (o Time) MarshalJSON() (data []byte, err error) {
	return json.Marshal(o.ElseZero())
}

func (o *Time) UnmarshalJSON(data []byte) error {
	var v time.Time
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	*o = OfTime(v)
	return nil
}

func (o Time) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.ElseZero(), start)
}

func (o *Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v time.Time
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = OfTime(v)
	return nil
}
