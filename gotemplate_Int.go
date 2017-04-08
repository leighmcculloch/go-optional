package optional

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Int struct {
	value *int
}

// For wraps the value in an Optional.
func ForInt(value int) Int {
	return Int{&value}
}

func ForIntPtr(ptr *int) Int {
	return Int{ptr}
}

// Empty returns an empty Optional.
func EmptyInt() Int {
	return Int{}
}

// IsPresent returns whether there is a value wrapped by this Optional.
func (o Int) IsPresent() bool {
	return o.value != nil
}

// IfPresent calls the function if there is a value wrapped by this Optional.
func (o Int) IfPresent(f func(value int)) {
	if o.value != nil {
		f(*o.value)
	}
}

// OrElse returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Int) OrElse(value int) int {
	if o.value != nil {
		return *o.value
	} else {
		return value
	}
}
