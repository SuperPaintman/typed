package nullx

import (
	"encoding/json"
)

// Nullable represents a T that may be null.
//
// It helps to handle encoding/decoding cases with zero values.
type Nullable[T any] struct {
	Value T
	Valid bool
}

// MarshalJSON implements the json.Marshaler interface.
func (n Nullable[T]) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.Value)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (n *Nullable[T]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*n = Nullable[T]{}
		return nil
	}

	if err := json.Unmarshal(data, &n.Value); err != nil {
		return err
	}
	n.Valid = true

	return nil
}

// NonNullableSlice represents a T slice.
//
// It encodes the nil value as the empty array ([]).
type NonNullableSlice[T any] []T

// MarshalJSON implements the json.Marshaler interface.
func (n NonNullableSlice[T]) MarshalJSON() ([]byte, error) {
	if n == nil {
		return []byte("[]"), nil
	}

	return json.Marshal(([]T)(n))
}

// NonNullableSlice represents a map[K]V.
//
// It encodes the nil value as the empty object ({}).
type NonNullableMap[K comparable, V any] map[K]V

// MarshalJSON implements the json.Marshaler interface.
func (n NonNullableMap[K, V]) MarshalJSON() ([]byte, error) {
	if n == nil {
		return []byte("{}"), nil
	}

	return json.Marshal((map[K]V)(n))
}
