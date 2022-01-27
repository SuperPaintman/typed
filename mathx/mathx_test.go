package mathx

import (
	"constraints"
	"fmt"
	"testing"
)

type maxMinTestCase[T constraints.Ordered] struct {
	x, y T
	want T
}

func TestMax_int(t *testing.T) {
	tt := []maxMinTestCase[int]{
		{
			x:    0,
			y:    0,
			want: 0,
		},
		{
			x:    0,
			y:    -1,
			want: 0,
		},
		{
			x:    1000,
			y:    2000,
			want: 2000,
		},
		{
			x:    2000,
			y:    1000,
			want: 2000,
		},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%d > %d", tc.x, tc.y), func(t *testing.T) {
			got := Max(tc.x, tc.y)

			if got != tc.want {
				t.Errorf("Max[int](%d, %d) = %d, want = %d",
					tc.x,
					tc.y,
					got,
					tc.want,
				)
			}
		})
	}
}

func TestMax_string(t *testing.T) {
	tt := []maxMinTestCase[string]{
		{
			x:    "",
			y:    "",
			want: "",
		},
		{
			x:    "a",
			y:    "b",
			want: "b",
		},
		{
			x:    "b",
			y:    "a",
			want: "b",
		},
		{
			x:    "ab",
			y:    "aa",
			want: "ab",
		},
		{
			x:    "ab",
			y:    "b",
			want: "b",
		},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%q > %q", tc.x, tc.y), func(t *testing.T) {
			got := Max(tc.x, tc.y)

			if got != tc.want {
				t.Errorf("Max[string](%q, %q) = %q, want = %q",
					tc.x,
					tc.y,
					got,
					tc.want,
				)
			}
		})
	}
}

func TestMin_int(t *testing.T) {
	tt := []maxMinTestCase[int]{
		{
			x:    0,
			y:    0,
			want: 0,
		},
		{
			x:    0,
			y:    -1,
			want: -1,
		},
		{
			x:    1000,
			y:    2000,
			want: 1000,
		},
		{
			x:    2000,
			y:    1000,
			want: 1000,
		},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%d < %d", tc.x, tc.y), func(t *testing.T) {
			got := Min(tc.x, tc.y)

			if got != tc.want {
				t.Errorf("Min[int](%d, %d) = %d, want = %d",
					tc.x,
					tc.y,
					got,
					tc.want,
				)
			}
		})
	}
}

func TestMin_string(t *testing.T) {
	tt := []maxMinTestCase[string]{
		{
			x:    "",
			y:    "",
			want: "",
		},
		{
			x:    "a",
			y:    "b",
			want: "a",
		},
		{
			x:    "b",
			y:    "a",
			want: "a",
		},
		{
			x:    "ab",
			y:    "aa",
			want: "aa",
		},
		{
			x:    "ab",
			y:    "b",
			want: "ab",
		},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%q < %q", tc.x, tc.y), func(t *testing.T) {
			got := Min(tc.x, tc.y)

			if got != tc.want {
				t.Errorf("Min[string](%q, %q) = %q, want = %q",
					tc.x,
					tc.y,
					got,
					tc.want,
				)
			}
		})
	}
}
