package optional

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"
)

var _String = time.Time{}

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type String optionalString

type optionalString []string

const (
	valueKeyString = iota
)

// Of wraps the value in an Optional.
func OfString(value string) String {
	return String{valueKeyString: value}
}

func OfStringPtr(ptr *string) String {
	if ptr == nil {
		return EmptyString()
	} else {
		return OfString(*ptr)
	}
}

// Empty returns an empty Optional.
func EmptyString() String {
	return nil
}

// IsEmpty returns true if there there is no value wrapped by this Optional.
func (o String) IsEmpty() bool {
	return o == nil
}

// IsPresent returns true if there is a value wrapped by this Optional.
func (o String) IsPresent() bool {
	return !o.IsEmpty()
}

// If calls the function if there is a value wrapped by this Optional.
func (o String) If(f func(value string)) {
	if o.IsPresent() {
		f(o[valueKeyString])
	}
}

func (o String) ElseFunc(f func() string) (value string) {
	if o.IsPresent() {
		o.If(func(v string) { value = v })
		return
	} else {
		return f()
	}
}

// Else returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o String) Else(elseValue string) (value string) {
	return o.ElseFunc(func() string { return elseValue })
}

// ElseZero returns the value wrapped by this Optional, or the zero value of
// the type wrapped if there is no value wrapped by this Optional.
func (o String) ElseZero() (value string) {
	var zero string
	return o.Else(zero)
}

// String returns a string representation of the wrapped value if one is present, otherwise an empty string.
func (o String) String() string {
	if o.IsPresent() {
		var value string
		o.If(func(v string) { value = v })
		return fmt.Sprintf("%v", value)
	} else {
		return ""
	}
}

func (o String) MarshalJSON() (data []byte, err error) {
	return json.Marshal(o.ElseZero())
}

func (o *String) UnmarshalJSON(data []byte) error {
	var v string
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	*o = OfString(v)
	return nil
}

func (o String) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(o.ElseZero(), start)
}

func (o *String) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}
	*o = OfString(v)
	return nil
}
