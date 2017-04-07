package optional

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Complex128 struct {
	value *complex128
}

// Of wraps the value in an Optional.
func OfComplex128(value complex128) Complex128 {
	return Complex128{&value}
}

func OfNilableComplex128(ptr *complex128) Complex128 {
	return Complex128{ptr}
}

// Empty returns an empty Optional.
func EmptyComplex128() Complex128 {
	return Complex128{}
}

// IsPresent returns whether there is a value wrapped by this Optional.
func (o Complex128) IsPresent() bool {
	return o.value != nil
}

// IfPresent calls the function if there is a value wrapped by this Optional.
func (o Complex128) IfPresent(f func(value complex128)) {
	if o.value != nil {
		f(*o.value)
	}
}

// OrElse returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Complex128) OrElse(value complex128) complex128 {
	if o.value != nil {
		return *o.value
	} else {
		return value
	}
}
