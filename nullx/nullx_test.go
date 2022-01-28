package nullx

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestNullable_MarshalJSON_int(t *testing.T) {
	tt := []struct {
		name string
		v    Nullable[int]
		want string
	}{
		{
			name: "null",
			v:    Nullable[int]{},
			want: "null",
		},
		{
			name: "0",
			v:    Nullable[int]{Value: 0, Valid: true},
			want: "0",
		},
		{
			name: "1337",
			v:    Nullable[int]{Value: 1337, Valid: true},
			want: "1337",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, err := json.Marshal(tc.v)
			if err != nil {
				t.Fatalf("json.Marshal(%#v) = %s, want nil error",
					tc.v,
					err,
				)
			}

			if string(got) != tc.want {
				t.Fatalf("json.Marshal(%#v) = %s, want %s",
					tc.v,
					got,
					tc.want,
				)
			}
		})
	}
}

func TestNullable_UnmarshalJSON_int(t *testing.T) {
	tt := []struct {
		name string
		v    string
		want Nullable[int]
	}{
		{
			name: "null",
			v:    "null",
			want: Nullable[int]{},
		},
		{
			name: "0",
			v:    "0",
			want: Nullable[int]{Value: 0, Valid: true},
		},
		{
			name: "1337",
			v:    "1337",
			want: Nullable[int]{Value: 1337, Valid: true},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var got Nullable[int]

			if err := json.Unmarshal([]byte(tc.v), &got); err != nil {
				t.Fatalf("json.Unmarshal(%q) = %s, want nil error",
					tc.v,
					err,
				)
			}

			if got != tc.want {
				t.Fatalf("json.Unmarshal(%q) = %#v, want %#v",
					tc.v,
					got,
					tc.want,
				)
			}
		})
	}
}

func TestNullable_MarshalJSON_string(t *testing.T) {
	tt := []struct {
		name string
		v    Nullable[string]
		want string
	}{
		{
			name: "null",
			v:    Nullable[string]{},
			want: "null",
		},
		{
			name: "empty",
			v:    Nullable[string]{Value: "", Valid: true},
			want: "\"\"",
		},
		{
			name: "string",
			v:    Nullable[string]{Value: "hello", Valid: true},
			want: "\"hello\"",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, err := json.Marshal(tc.v)
			if err != nil {
				t.Fatalf("json.Marshal(%#v) = %s, want nil error",
					tc.v,
					err,
				)
			}

			if string(got) != tc.want {
				t.Fatalf("json.Marshal(%#v) = %s, want %s",
					tc.v,
					got,
					tc.want,
				)
			}
		})
	}
}

func TestNullable_UnmarshalJSON_string(t *testing.T) {
	tt := []struct {
		name string
		v    string
		want Nullable[string]
	}{
		{
			name: "null",
			v:    "null",
			want: Nullable[string]{},
		},
		{
			name: "empty",
			v:    "\"\"",
			want: Nullable[string]{Value: "", Valid: true},
		},
		{
			name: "string",
			v:    "\"hello\"",
			want: Nullable[string]{Value: "hello", Valid: true},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var got Nullable[string]

			if err := json.Unmarshal([]byte(tc.v), &got); err != nil {
				t.Fatalf("json.Unmarshal(%q) = %s, want nil error",
					tc.v,
					err,
				)
			}

			if got != tc.want {
				t.Fatalf("json.Unmarshal(%q) = %#v, want %#v",
					tc.v,
					got,
					tc.want,
				)
			}
		})
	}
}

func TestNonNullableSlice_MarshalJSON(t *testing.T) {
	tt := []struct {
		name string
		v    NonNullableSlice[int]
		want string
	}{
		{
			name: "nil",
			v:    nil,
			want: "[]",
		},
		{
			name: "empty",
			v:    []int{},
			want: "[]",
		},
		{
			name: "slice",
			v:    []int{1, 2, 3, 4, 5},
			want: "[1,2,3,4,5]",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, err := json.Marshal(tc.v)
			if err != nil {
				t.Fatalf("json.Marshal(%#v) = %s, want nil error",
					tc.v,
					err,
				)
			}

			if string(got) != tc.want {
				t.Fatalf("json.Marshal(%#v) = %s, want %s",
					tc.v,
					got,
					tc.want,
				)
			}
		})
	}
}

func TestNonNullableSlice_UnmarshalJSON(t *testing.T) {
	tt := []struct {
		name string
		v    string
		want NonNullableSlice[int]
	}{
		{
			name: "nil",
			v:    "null",
			want: nil,
		},
		{
			name: "empty",
			v:    "[]",
			want: []int{},
		},
		{
			name: "slice",
			v:    "[1,2,3,4,5]",
			want: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var got NonNullableSlice[int]

			if err := json.Unmarshal([]byte(tc.v), &got); err != nil {
				t.Fatalf("json.Unmarshal(%q) = %s, want nil error",
					tc.want,
					err,
				)
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("json.Unmarshal(%q) = %#v, want %#v",
					tc.want,
					got,
					tc.want,
				)
			}
		})
	}
}

func TestNonNullableMap_MarshalJSON(t *testing.T) {
	tt := []struct {
		name string
		v    NonNullableMap[string, int]
		want string
	}{
		{
			name: "nil",
			v:    nil,
			want: "{}",
		},
		{
			name: "empty",
			v:    map[string]int{},
			want: "{}",
		},
		{
			name: "slice",
			v:    map[string]int{"a": 1},
			want: "{\"a\":1}",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, err := json.Marshal(tc.v)
			if err != nil {
				t.Fatalf("json.Marshal(%#v) = %s, want nil error",
					tc.v,
					err,
				)
			}

			if string(got) != tc.want {
				t.Fatalf("json.Marshal(%#v) = %s, want %s",
					tc.v,
					got,
					tc.want,
				)
			}
		})
	}
}

func TestNonNullableMap_UnmarshalJSON(t *testing.T) {
	tt := []struct {
		name string
		v    string
		want NonNullableMap[string, int]
	}{
		{
			name: "nil",
			v:    "null",
			want: nil,
		},
		{
			name: "empty",
			v:    "{}",
			want: map[string]int{},
		},
		{
			name: "map",
			v:    "{\"a\":1}",
			want: map[string]int{"a": 1},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var got NonNullableMap[string, int]

			if err := json.Unmarshal([]byte(tc.v), &got); err != nil {
				t.Fatalf("json.Unmarshal(%q) = %s, want nil error",
					tc.want,
					err,
				)
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("json.Unmarshal(%q) = %#v, want %#v",
					tc.want,
					got,
					tc.want,
				)
			}
		})
	}
}
