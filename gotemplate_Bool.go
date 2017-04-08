package optional

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Bool struct {
	value *bool
}

// Of wraps the value in an Optional.
func OfBool(value bool) Bool {
	return Bool{&value}
}

func OfBoolPtr(ptr *bool) Bool {
	return Bool{ptr}
}

// Empty returns an empty Optional.
func EmptyBool() Bool {
	return Bool{}
}

// IsPresent returns whether there is a value wrapped by this Optional.
func (o Bool) IsPresent() bool {
	return o.value != nil
}

// IfPresent calls the function if there is a value wrapped by this Optional.
func (o Bool) IfPresent(f func(value bool)) {
	if o.value != nil {
		f(*o.value)
	}
}

// OrElse returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Bool) OrElse(value bool) bool {
	if o.value != nil {
		return *o.value
	} else {
		return value
	}
}
