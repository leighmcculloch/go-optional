package optional

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Uint8 struct {
	value *uint8
}

// Of wraps the value in an Optional.
func OfUint8(value uint8) Uint8 {
	return Uint8{&value}
}

func OfNilableUint8(ptr *uint8) Uint8 {
	return Uint8{ptr}
}

// Empty returns an empty Optional.
func EmptyUint8() Uint8 {
	return Uint8{}
}

// IsPresent returns whether there is a value wrapped by this Optional.
func (o Uint8) IsPresent() bool {
	return o.value != nil
}

// IfPresent calls the function if there is a value wrapped by this Optional.
func (o Uint8) IfPresent(f func(value uint8)) {
	if o.value != nil {
		f(*o.value)
	}
}

// OrElse returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Uint8) OrElse(value uint8) uint8 {
	if o.value != nil {
		return *o.value
	} else {
		return value
	}
}
