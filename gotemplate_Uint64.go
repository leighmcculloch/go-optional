package optional

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Uint64 struct {
	value *uint64
}

// Of wraps the value in an Optional.
func OfUint64(value uint64) Uint64 {
	return Uint64{&value}
}

func OfUint64Ptr(ptr *uint64) Uint64 {
	return Uint64{ptr}
}

// Empty returns an empty Optional.
func EmptyUint64() Uint64 {
	return Uint64{}
}

// IsPresent returns whether there is a value wrapped by this Optional.
func (o Uint64) IsPresent() bool {
	return o.value != nil
}

// IfPresent calls the function if there is a value wrapped by this Optional.
func (o Uint64) IfPresent(f func(value uint64)) {
	if o.value != nil {
		f(*o.value)
	}
}

// OrElse returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Uint64) OrElse(value uint64) uint64 {
	if o.value != nil {
		return *o.value
	} else {
		return value
	}
}
