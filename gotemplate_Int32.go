package optional

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Int32 struct {
	value *int32
}

// Of wraps the value in an Optional.
func OfInt32(value int32) Int32 {
	return Int32{&value}
}

func OfNilableInt32(ptr *int32) Int32 {
	return Int32{ptr}
}

// Empty returns an empty Optional.
func EmptyInt32() Int32 {
	return Int32{}
}

// IsPresent returns whether there is a value wrapped by this Optional.
func (o Int32) IsPresent() bool {
	return o.value != nil
}

// IfPresent calls the function if there is a value wrapped by this Optional.
func (o Int32) IfPresent(f func(value int32)) {
	if o.value != nil {
		f(*o.value)
	}
}

// OrElse returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Int32) OrElse(value int32) int32 {
	if o.value != nil {
		return *o.value
	} else {
		return value
	}
}
