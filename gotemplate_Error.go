package optional

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Error struct {
	value *error
}

// Of wraps the value in an Optional.
func OfError(value error) Error {
	return Error{&value}
}

func OfNilableError(ptr *error) Error {
	return Error{ptr}
}

// Empty returns an empty Optional.
func EmptyError() Error {
	return Error{}
}

// IsPresent returns whether there is a value wrapped by this Optional.
func (o Error) IsPresent() bool {
	return o.value != nil
}

// IfPresent calls the function if there is a value wrapped by this Optional.
func (o Error) IfPresent(f func(value error)) {
	if o.value != nil {
		f(*o.value)
	}
}

// OrElse returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Error) OrElse(value error) error {
	if o.value != nil {
		return *o.value
	} else {
		return value
	}
}
