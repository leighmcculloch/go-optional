package optional

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Uint32 struct {
	value *uint32
}

// Of wraps the value in an Optional.
func OfUint32(value uint32) Uint32 {
	return Uint32{&value}
}

func OfUint32Ptr(ptr *uint32) Uint32 {
	return Uint32{ptr}
}

// Empty returns an empty Optional.
func EmptyUint32() Uint32 {
	return Uint32{}
}

// IsPresent returns whether there is a value wrapped by this Optional.
func (o Uint32) IsPresent() bool {
	return o.value != nil
}

// IfPresent calls the function if there is a value wrapped by this Optional.
func (o Uint32) IfPresent(f func(value uint32)) {
	if o.value != nil {
		f(*o.value)
	}
}

// OrElse returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Uint32) OrElse(value uint32) uint32 {
	if o.value != nil {
		return *o.value
	} else {
		return value
	}
}
