package optional

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Int8 struct {
	value *int8
}

// Of wraps the value in an Optional.
func OfInt8(value int8) Int8 {
	return Int8{&value}
}

func OfNilableInt8(ptr *int8) Int8 {
	return Int8{ptr}
}

// Empty returns an empty Optional.
func EmptyInt8() Int8 {
	return Int8{}
}

// IsPresent returns whether there is a value wrapped by this Optional.
func (o Int8) IsPresent() bool {
	return o.value != nil
}

// IfPresent calls the function if there is a value wrapped by this Optional.
func (o Int8) IfPresent(f func(value int8)) {
	if o.value != nil {
		f(*o.value)
	}
}

// OrElse returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Int8) OrElse(value int8) int8 {
	if o.value != nil {
		return *o.value
	} else {
		return value
	}
}
