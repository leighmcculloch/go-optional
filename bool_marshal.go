package optional

import (
	"strconv"
)

// MarshalText returns text for marshaling this Int.
func (o Bool) MarshalText() (text []byte, err error) {
	o.If(func(v bool) {
		text = []byte(strconv.FormatBool(v))
	})
	return
}

// UnmarshalText parse the text into this Int
func (o *Bool) UnmarshalText(text []byte) error {
	v, err := strconv.ParseBool(string(text))
	if err != nil {
		return err
	}
	*o = OfBool(v)
	return nil
}
