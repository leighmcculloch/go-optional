package optional

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Byte struct {
	value *byte
}

// Of wraps the value in an Optional.
func OfByte(value byte) Byte {
	return Byte{&value}
}

func OfNilableByte(ptr *byte) Byte {
	return Byte{ptr}
}

// Empty returns an empty Optional.
func EmptyByte() Byte {
	return Byte{}
}

// IsPresent returns whether there is a value wrapped by this Optional.
func (o Byte) IsPresent() bool {
	return o.value != nil
}

// IfPresent calls the function if there is a value wrapped by this Optional.
func (o Byte) IfPresent(f func(value byte)) {
	if o.value != nil {
		f(*o.value)
	}
}

// OrElse returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Byte) OrElse(value byte) byte {
	if o.value != nil {
		return *o.value
	} else {
		return value
	}
}
