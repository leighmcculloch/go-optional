package template

import "testing"

func TestEmpty(t *testing.T) {
	empty := Empty()
	expected := Optional{}

	if empty != expected {
		t.Errorf("Empty got %#v, want %#v", empty, expected)
	}
}

func TestIsPresent(t *testing.T) {
	s := "ptr to string"
	tests := []struct {
		Optional          Optional
		ExpectedIsPresent bool
	}{
		{Empty(), false},
		{Of(""), true},
		{Of("string"), true},
		{OfNilable((*T)(nil)), false},
		{OfNilable((*T)(&s)), true},
	}

	for _, test := range tests {
		isPresent := test.Optional.IsPresent()

		if isPresent != test.ExpectedIsPresent {
			t.Errorf("%#v IsPresent got %#v, want %#v", test.Optional, isPresent, test.ExpectedIsPresent)
		}
	}
}

func TestIfPresent(t *testing.T) {
	s := "ptr to string"
	tests := []struct {
		Optional       Optional
		ExpectedCalled bool
	}{
		{Empty(), false},
		{Of(""), true},
		{Of("string"), true},
		{OfNilable((*T)(nil)), false},
		{OfNilable((*T)(&s)), true},
	}

	for _, test := range tests {
		called := false
		test.Optional.IfPresent(func(v T) {
			called = true
			if v != *test.Optional.value {
				t.Errorf("%#v IfPresent got %#v, want #%v", test.Optional, v, test.Optional.value)
			}
		})

		if called != test.ExpectedCalled {
			t.Errorf("%#v IfPresent called %#v, want %#v", test.Optional, called, test.ExpectedCalled)
		}
	}
}

func TestOrElse(t *testing.T) {
	s := "ptr to string"
	const orElse = "orelse"
	tests := []struct {
		Optional       Optional
		ExpectedResult T
	}{
		{Empty(), orElse},
		{Of(""), ""},
		{Of("string"), "string"},
		{OfNilable((*T)(nil)), orElse},
		{OfNilable((*T)(&s)), "ptr to string"},
	}

	for _, test := range tests {
		result := test.Optional.OrElse(orElse)

		if result != test.ExpectedResult {
			t.Errorf("%#v OrElse(%#v) got %#v, want %#v", test.Optional, orElse, result, test.ExpectedResult)
		}
	}
}
