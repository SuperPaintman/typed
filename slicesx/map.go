package slicesx

// Map returns a slice of all elements modified through fn.
//
// Map might also change the type of elements.
//
// Map does not modify the original slice and returns a new allocated slice.
func Map[T, P any](values []T, fn func(T) P) []P {
	if values == nil {
		return nil
	}

	res := make([]P, len(values))
	for i, v := range values {
		res[i] = fn(v)
	}
	return res
}

// MapModify returns a slice of all elements modified through fn.
//
// MapModify does not allocate a new slice, but it modifies the values
// in place.
func MapModify[T any](values []T, fn func(T) T) []T {
	for i, v := range values {
		values[i] = fn(v)
	}
	return values
}
