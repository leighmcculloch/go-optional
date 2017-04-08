package template

// template type Optional(T)

type T string

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Optional struct {
	value *T
}

// Of wraps the value in an Optional.
func Of(value T) Optional {
	return Optional{&value}
}

func OfOptionalPtr(ptr *T) Optional {
	return Optional{ptr}
}

// Empty returns an empty Optional.
func Empty() Optional {
	return Optional{}
}

// IsPresent returns whether there is a value wrapped by this Optional.
func (o Optional) IsPresent() bool {
	return o.value != nil
}

// IfPresent calls the function if there is a value wrapped by this Optional.
func (o Optional) IfPresent(f func(value T)) {
	if o.value != nil {
		f(*o.value)
	}
}

// OrElse returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Optional) OrElse(value T) T {
	if o.value != nil {
		return *o.value
	} else {
		return value
	}
}
