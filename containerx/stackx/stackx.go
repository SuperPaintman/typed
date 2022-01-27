package stackx

// Stack is like a Go slice but with the functionality of the stack.
//
// For pushing new elements and getting the length of the Stack, use standard
// Go functions (like append, len, cap, etc.).
type Stack[T any] []T

// Pop returns the value stored on the top of the stack and pops it, or the
// default value for type if the stack is empty.
//
// The ok result indicates whether the value is present (false if the stack is
// empty).
func (s *Stack[T]) Pop() (value T, ok bool) {
	if len(*s) == 0 {
		return
	}

	value = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return value, true
}
