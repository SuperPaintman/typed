package slicesx

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	tt := []struct {
		name   string
		values []int
		fn     func(int) bool
		want   []int
	}{
		{
			name:   "all",
			values: []int{1, 2, 3, 4, 5},
			fn:     func(int) bool { return true },
			want:   []int{1, 2, 3, 4, 5},
		},
		{
			name:   "none",
			values: []int{1, 2, 3, 4, 5},
			fn:     func(int) bool { return false },
			want:   nil,
		},
		{
			name:   "even",
			values: []int{1, 2, 3, 4, 5},
			fn:     func(v int) bool { return v%2 == 0 },
			want:   []int{2, 4},
		},
		{
			name:   "odd",
			values: []int{1, 2, 3, 4, 5},
			fn:     func(v int) bool { return v%2 == 1 },
			want:   []int{1, 3, 5},
		},
		{
			name:   "nil",
			values: nil,
			fn:     func(int) bool { return true },
			want:   nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var origin []int
			if tc.values != nil {
				origin = make([]int, len(tc.values))
				copy(origin, tc.values)
			}

			got := Filter(tc.values, tc.fn)

			if !reflect.DeepEqual(origin, tc.values) {
				t.Errorf("Filter[int](%v) modified the original slice: %v",
					origin,
					tc.values,
				)
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Filter[int](%v) = %v, want %v", origin, got, tc.want)
			}
		})
	}
}

func TestFilterModify(t *testing.T) {
	tt := []struct {
		name   string
		values []int
		fn     func(int) bool
		want   []int
	}{
		{
			name:   "all",
			values: []int{1, 2, 3, 4, 5},
			fn:     func(int) bool { return true },
			want:   []int{1, 2, 3, 4, 5},
		},
		{
			name:   "none",
			values: []int{1, 2, 3, 4, 5},
			fn:     func(int) bool { return false },
			want:   []int{},
		},
		{
			name:   "even",
			values: []int{1, 2, 3, 4, 5},
			fn:     func(v int) bool { return v%2 == 0 },
			want:   []int{2, 4},
		},
		{
			name:   "odd",
			values: []int{1, 2, 3, 4, 5},
			fn:     func(v int) bool { return v%2 == 1 },
			want:   []int{1, 3, 5},
		},
		{
			name:   "nil",
			values: nil,
			fn:     func(int) bool { return true },
			want:   nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := FilterModify(tc.values, tc.fn)

			if !reflect.DeepEqual(got, tc.values[:len(got)]) {
				t.Errorf("FilterModify[int](%v) did not modify the original slice",
					tc.values,
				)
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("FilterModify[int](%v) = %v, want %v", tc.values, got, tc.want)
			}
		})
	}
}
