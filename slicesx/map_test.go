package slicesx

import (
	"reflect"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	tt := []struct {
		name   string
		values []int
		fn     func(int) string
		want   []string
	}{
		{
			name:   "itoa",
			values: []int{1, 2, 3, 4, 5},
			fn:     func(v int) string { return strconv.Itoa(v) },
			want:   []string{"1", "2", "3", "4", "5"},
		},
		{
			name:   "empty",
			values: []int{},
			fn:     func(v int) string { return strconv.Itoa(v) },
			want:   []string{},
		},
		{
			name:   "nil",
			values: nil,
			fn:     func(v int) string { return strconv.Itoa(v) },
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

			got := Map(tc.values, tc.fn)

			if !reflect.DeepEqual(origin, tc.values) {
				t.Errorf("Map[int, string](%v) modified the original slice: %v",
					origin,
					tc.values,
				)
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Map[int, string](%v) = %v, want %v", origin, got, tc.want)
			}
		})
	}
}

func TestMapModify(t *testing.T) {
	tt := []struct {
		name   string
		values []int
		fn     func(int) int
		want   []int
	}{
		{
			name:   "multiply-by-10",
			values: []int{1, 2, 3, 4, 5},
			fn:     func(v int) int { return v * 10 },
			want:   []int{10, 20, 30, 40, 50},
		},
		{
			name:   "empty",
			values: []int{},
			fn:     func(v int) int { return v * 10 },
			want:   []int{},
		},
		{
			name:   "nil",
			values: nil,
			fn:     func(v int) int { return v * 10 },
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

			got := MapModify(tc.values, tc.fn)

			if len(origin) > 0 {
				if reflect.DeepEqual(origin, tc.values) {
					t.Errorf("MapModify[int](%v) did not modify the original slicev",
						tc.values,
					)
				}
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("MapModify[int](%v) = %v, want %v", origin, got, tc.want)
			}
		})
	}
}
