package optional

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Int64 struct {
	value *int64
}

// For wraps the value in an Optional.
func ForInt64(value int64) Int64 {
	return Int64{&value}
}

func ForInt64Ptr(ptr *int64) Int64 {
	return Int64{ptr}
}

// Empty returns an empty Optional.
func EmptyInt64() Int64 {
	return Int64{}
}

// IsPresent returns whether there is a value wrapped by this Optional.
func (o Int64) IsPresent() bool {
	return o.value != nil
}

// IfPresent calls the function if there is a value wrapped by this Optional.
func (o Int64) IfPresent(f func(value int64)) {
	if o.value != nil {
		f(*o.value)
	}
}

// OrElse returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Int64) OrElse(value int64) int64 {
	if o.value != nil {
		return *o.value
	} else {
		return value
	}
}
