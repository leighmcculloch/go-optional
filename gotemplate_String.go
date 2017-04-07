package optional

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type String struct {
	value *string
}

// Of wraps the value in an Optional.
func OfString(value string) String {
	return String{&value}
}

func OfNilableString(ptr *string) String {
	return String{ptr}
}

// Empty returns an empty Optional.
func EmptyString() String {
	return String{}
}

// IsPresent returns whether there is a value wrapped by this Optional.
func (o String) IsPresent() bool {
	return o.value != nil
}

// IfPresent calls the function if there is a value wrapped by this Optional.
func (o String) IfPresent(f func(value string)) {
	if o.value != nil {
		f(*o.value)
	}
}

// OrElse returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o String) OrElse(value string) string {
	if o.value != nil {
		return *o.value
	} else {
		return value
	}
}
