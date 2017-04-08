package optional

// template type Optional(T)

// Optional wraps a value that may or may not be nil.
// If a value is present, it may be unwrapped to expose the underlying value.
type Float32 struct {
	value *float32
}

// For wraps the value in an Optional.
func ForFloat32(value float32) Float32 {
	return Float32{&value}
}

func ForFloat32Ptr(ptr *float32) Float32 {
	return Float32{ptr}
}

// Empty returns an empty Optional.
func EmptyFloat32() Float32 {
	return Float32{}
}

// IsPresent returns whether there is a value wrapped by this Optional.
func (o Float32) IsPresent() bool {
	return o.value != nil
}

// IfPresent calls the function if there is a value wrapped by this Optional.
func (o Float32) IfPresent(f func(value float32)) {
	if o.value != nil {
		f(*o.value)
	}
}

// OrElse returns the value wrapped by this Optional, or the value passed in if
// there is no value wrapped by this Optional.
func (o Float32) OrElse(value float32) float32 {
	if o.value != nil {
		return *o.value
	} else {
		return value
	}
}
