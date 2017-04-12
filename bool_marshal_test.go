package optional

import (
	"reflect"
	"testing"
)

func TestBoolMarshalText(t *testing.T) {
	tests := []struct {
		Optional     Bool
		ExpectedText string
	}{
		{EmptyBool(), "false"},
		{OfBool(true), "true"},
		{OfBool(false), "false"},
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

func TestBoolUnmarshalText(t *testing.T) {
	tests := []struct {
		Text             string
		ExpectError      bool
		ExpectedOptional Bool
	}{
		{"", true, EmptyBool()},
		{"asdf", true, EmptyBool()},
		{"true", false, OfBool(true)},
		{"false", false, OfBool(false)},
	}

	for _, test := range tests {
		var o Bool
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
