package optional

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Uint16 struct {
	value *uint16
}

// For wraps the value in an Optional.
func ForUint16(value uint16) Uint16 {
	return Uint16{&value}
}

func ForUint16Ptr(ptr *uint16) Uint16 {
	return Uint16{ptr}
}

// Empty returns an empty Optional.
func EmptyUint16() Uint16 {
	return Uint16{}
}

// IsPresent returns whether there is a value wrapped by this Optional.
func (o Uint16) IsPresent() bool {
	return o.value != nil
}

// IfPresent calls the function if there is a value wrapped by this Optional.
func (o Uint16) IfPresent(f func(value uint16)) {
	if o.value != nil {
		f(*o.value)
	}
}

// OrElse returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Uint16) OrElse(value uint16) uint16 {
	if o.value != nil {
		return *o.value
	} else {
		return value
	}
}
