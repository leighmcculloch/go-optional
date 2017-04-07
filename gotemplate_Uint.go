package optional

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Uint struct {
	value *uint
}

// Of wraps the value in an Optional.
func OfUint(value uint) Uint {
	return Uint{&value}
}

func OfNilableUint(ptr *uint) Uint {
	return Uint{ptr}
}

// Empty returns an empty Optional.
func EmptyUint() Uint {
	return Uint{}
}

// IsPresent returns whether there is a value wrapped by this Optional.
func (o Uint) IsPresent() bool {
	return o.value != nil
}

// IfPresent calls the function if there is a value wrapped by this Optional.
func (o Uint) IfPresent(f func(value uint)) {
	if o.value != nil {
		f(*o.value)
	}
}

// OrElse returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Uint) OrElse(value uint) uint {
	if o.value != nil {
		return *o.value
	} else {
		return value
	}
}
