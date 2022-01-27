package mathx

import (
	"constraints"
)

// Max returns the larger of x or y.
func Max[T constraints.Ordered](x, y T) T {
	if x > y {
		return x
	}

	return y
}

// Min returns the smaller of x or y.
func Min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}

	return y
}
