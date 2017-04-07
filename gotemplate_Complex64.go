package optional

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Complex64 struct {
	value *complex64
}

// Of wraps the value in an Optional.
func OfComplex64(value complex64) Complex64 {
	return Complex64{&value}
}

func OfNilableComplex64(ptr *complex64) Complex64 {
	return Complex64{ptr}
}

// Empty returns an empty Optional.
func EmptyComplex64() Complex64 {
	return Complex64{}
}

// IsPresent returns whether there is a value wrapped by this Optional.
func (o Complex64) IsPresent() bool {
	return o.value != nil
}

// IfPresent calls the function if there is a value wrapped by this Optional.
func (o Complex64) IfPresent(f func(value complex64)) {
	if o.value != nil {
		f(*o.value)
	}
}

// OrElse returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Complex64) OrElse(value complex64) complex64 {
	if o.value != nil {
		return *o.value
	} else {
		return value
	}
}
