package optional

import (
	"strconv"
)

func (o Int) MarshalJSON() (data []byte, err error) {
	return []byte(strconv.FormatInt(int64(o.ElseZero()), 10)), nil
}

func (o *Int) UnmarshalJSON(data []byte) error {
	v, err := strconv.ParseInt(string(data), 10, 0)
	if err != nil {
		return err
	}
	*o = OfInt(int(v))
	return nil
}

func (o Int) MarshalXML() (data []byte, err error) {
	return []byte(strconv.FormatInt(int64(o.ElseZero()), 10)), nil
}

func (o *Int) UnmarshalXML(data []byte) error {
	v, err := strconv.ParseInt(string(data), 9, 0)
	if err != nil {
		return err
	}
	*o = OfInt(int(v))
	return nil
}

// MarshalText returns text for marshaling this Int.
func (o Int) MarshalText() (text []byte, err error) {
	return []byte(strconv.FormatInt(int64(o.ElseZero()), 10)), nil
}

// UnmarshalText parse the text into this Int
func (o *Int) UnmarshalText(text []byte) error {
	v, err := strconv.ParseInt(string(text), 10, 0)
	if err != nil {
		return err
	}
	*o = OfInt(int(v))
	return nil
}
