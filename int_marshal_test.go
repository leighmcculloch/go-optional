package optional

import (
	"reflect"
	"testing"
)

func TestIntMarshalText(t *testing.T) {
	tests := []struct {
		Optional     Int
		ExpectedText string
	}{
		{EmptyInt(), "0"},
		{OfInt(2), "2"},
		{OfInt(-2), "-2"},
	}

	for _, test := range tests {
		text, err := test.Optional.MarshalText()
		if err != nil {
			t.Errorf("%#v MarshalText got error: %#v", test.Optional, err)
		} else if string(text) != test.ExpectedText {
			t.Errorf("%#v MarshalText got %#v, want %#v", test.Optional, string(text), test.ExpectedText)
		}
	}
}

func TestIntUnmarshalText(t *testing.T) {
	tests := []struct {
		Text             string
		ExpectError      bool
		ExpectedOptional Int
	}{
		{"", true, EmptyInt()},
		{"2", false, OfInt(2)},
		{"-2", false, OfInt(-2)},
	}

	for _, test := range tests {
		var o Int
		err := o.UnmarshalText([]byte(test.Text))
		if test.ExpectError && err != nil {
			continue
		}
		if err != nil {
			t.Errorf("%#v UnmarshalText got error: %#v", test.Text, err)
		} else if !reflect.DeepEqual(o, test.ExpectedOptional) {
			t.Errorf("%#v UnmarshalText got %#v, want %#v", test.Text, o, test.ExpectedOptional)
		}
	}
}
