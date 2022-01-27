package mathx

import (
	"constraints"
)

// Max returns the larger of x or y.
//
// For float32 and float64 use math.Max because it covers special cases.
func Max[T constraints.Integer | ~string](x, y T) T {
	if x > y {
		return x
	}

	return y
}

// Min returns the smaller of x or y.
//
// For float32 and float64 use math.Min because it covers special cases.
func Min[T constraints.Integer | ~string](x, y T) T {
	if x < y {
		return x
	}

	return y
}
