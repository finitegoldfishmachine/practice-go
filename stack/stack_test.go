package stack

import (
	"testing"
)

func TestPush(t *testing.T) {
	s := Stack{}
	input := 3
	want := 3

	s.Push(input)
	result := s.Peek()

	if result != want {
		t.Fatalf(`s.Push(%#v), received %v, want  %v`, input, result, want)
	}
}

func TestPop(t *testing.T) {
	s := Stack{}
	want := 1

	// Push two items to the stack, Pop one out, Peek the top one
	for _, i := range []int{1, 2} {
		s.Push(i)
	}
	s.Pop()
	result := s.Peek()

	if result != want {
		t.Fatalf(`s.Pop(), received %v, want  %v`, result, want)
	}
}

func TestPeekEmpty(t *testing.T) {
	s := Stack{}
	want := -1

	s.Pop()
	result := s.Peek()

	if result != want {
		t.Fatalf(`s.Pop(), received %v, want %v`, result, want)
	}
}
