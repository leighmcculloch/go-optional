package optional

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Float64 struct {
	value *float64
}

// Of wraps the value in an Optional.
func OfFloat64(value float64) Float64 {
	return Float64{&value}
}

func OfNilableFloat64(ptr *float64) Float64 {
	return Float64{ptr}
}

// Empty returns an empty Optional.
func EmptyFloat64() Float64 {
	return Float64{}
}

// IsPresent returns whether there is a value wrapped by this Optional.
func (o Float64) IsPresent() bool {
	return o.value != nil
}

// IfPresent calls the function if there is a value wrapped by this Optional.
func (o Float64) IfPresent(f func(value float64)) {
	if o.value != nil {
		f(*o.value)
	}
}

// OrElse returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Float64) OrElse(value float64) float64 {
	if o.value != nil {
		return *o.value
	} else {
		return value
	}
}
