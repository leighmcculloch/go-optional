package optional

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Uintptr struct {
	value *uintptr
}

// Of wraps the value in an Optional.
func OfUintptr(value uintptr) Uintptr {
	return Uintptr{&value}
}

func OfNilableUintptr(ptr *uintptr) Uintptr {
	return Uintptr{ptr}
}

// Empty returns an empty Optional.
func EmptyUintptr() Uintptr {
	return Uintptr{}
}

// IsPresent returns whether there is a value wrapped by this Optional.
func (o Uintptr) IsPresent() bool {
	return o.value != nil
}

// IfPresent calls the function if there is a value wrapped by this Optional.
func (o Uintptr) IfPresent(f func(value uintptr)) {
	if o.value != nil {
		f(*o.value)
	}
}

// OrElse returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Uintptr) OrElse(value uintptr) uintptr {
	if o.value != nil {
		return *o.value
	} else {
		return value
	}
}
