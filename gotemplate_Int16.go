package optional

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Int16 struct {
	value *int16
}

// Of wraps the value in an Optional.
func OfInt16(value int16) Int16 {
	return Int16{&value}
}

func OfInt16Ptr(ptr *int16) Int16 {
	return Int16{ptr}
}

// Empty returns an empty Optional.
func EmptyInt16() Int16 {
	return Int16{}
}

// IsPresent returns whether there is a value wrapped by this Optional.
func (o Int16) IsPresent() bool {
	return o.value != nil
}

// IfPresent calls the function if there is a value wrapped by this Optional.
func (o Int16) IfPresent(f func(value int16)) {
	if o.value != nil {
		f(*o.value)
	}
}

// OrElse returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Int16) OrElse(value int16) int16 {
	if o.value != nil {
		return *o.value
	} else {
		return value
	}
}
