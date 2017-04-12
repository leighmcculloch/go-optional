package optional

import (
	"strconv"
)

func (o Bool) MarshalJSON() (data []byte, err error) {
	return []byte(strconv.FormatBool(o.ElseZero())), nil
}

func (o *Bool) UnmarshalJSON(data []byte) error {
	v, err := strconv.ParseBool(string(data))
	if err != nil {
		return err
	}
	*o = OfBool(v)
	return nil
}

// MarshalText returns text for marshaling this Int.
func (o Bool) MarshalText() (text []byte, err error) {
	return []byte(strconv.FormatBool(o.ElseZero())), nil
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
