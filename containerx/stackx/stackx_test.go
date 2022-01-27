package stackx

import (
	"testing"
)

func TestStack(t *testing.T) {
	var stack Stack[int]

	v0, ok := stack.Pop()
	checkStackPopValue(t, stack, v0, 0)
	checkStackPopOk(t, stack, ok, false)
	checkStackLen(t, stack, 0)

	stack = append(stack, 1)
	stack = append(stack, 2)
	stack = append(stack, 3)
	checkStackLen(t, stack, 3)

	v1, ok := stack.Pop()
	checkStackPopValue(t, stack, v1, 3)
	checkStackPopOk(t, stack, ok, true)
	checkStackLen(t, stack, 2)

	v2, ok := stack.Pop()
	checkStackPopValue(t, stack, v2, 2)
	checkStackPopOk(t, stack, ok, true)
	checkStackLen(t, stack, 1)

	v3, ok := stack.Pop()
	checkStackPopValue(t, stack, v3, 1)
	checkStackPopOk(t, stack, ok, true)
	checkStackLen(t, stack, 0)

	v4, ok := stack.Pop()
	checkStackPopValue(t, stack, v4, 0)
	checkStackPopOk(t, stack, ok, false)
	checkStackLen(t, stack, 0)
}

func checkStackPopValue[T comparable](t *testing.T, s Stack[T], got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("%T.Pop() = %v, want %v", s, got, want)
	}
}

func checkStackPopOk[T any](t *testing.T, s Stack[T], got, want bool) {
	t.Helper()

	if got != want {
		t.Errorf("%T.Pop() = %v, want %v", s, got, want)
	}
}

func checkStackLen[T any](t *testing.T, s Stack[T], want int) {
	t.Helper()

	if len(s) != want {
		t.Errorf("len(%T) = %v, want %v", s, len(s), want)
	}
}
