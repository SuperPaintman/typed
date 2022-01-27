package slicesx

// Filter returns a slice of all elements for which fn returned true.
//
// Filter does not modify the original slice and returns a new allocated slice.
func Filter[T any](values []T, fn func(T) bool) []T {
	var res []T
	for _, v := range values {
		if fn(v) {
			res = append(res, v)
		}
	}
	return res
}

// FilterModify returns a slice of all elements for which fn returned true.
//
// FilterModify does not allocate a new slice, but it modifies (sort) the values
// in place.
func FilterModify[T any](values []T, fn func(T) bool) []T {
	var j int
	for i := 0; i < len(values); i++ {
		if fn(values[i]) {
			values[i], values[j] = values[j], values[i]
			j++
		}
	}

	return values[:j]
}
