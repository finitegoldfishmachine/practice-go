package queue

import (
	"testing"
)

func TestPeek(t *testing.T) {
	q := queue{contents: []any{5, 2, 1}}
	want := 5

	got, err := q.Peek()
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Fatalf("Peek() - got: %v, want: %v", got, want)
	}
}

func TestPeekEmpty(t *testing.T) {
	q := queue{}
	_, err := q.Peek()

	if err == nil {
		t.Fatalf("Wanted an error but received nil")
	}
}

func TestInsert(t *testing.T) {
	q := queue{}
	want := 5

	q.Insert(5)

	got, err := q.Peek()
	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Fatalf("Insert(%v) - got: %v, want: %v", want, got, want)
	}
}

func TestPop(t *testing.T) {
	q := queue{}
	want := 4

	for x := 1; x < 5; x++ {
		q.Insert(x)
	}

	for x := 0; x < 3; x++ {
		q.Pop()
	}

	got, err := q.Peek()
	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Fatalf("Pop() - got: %v, want: %v", got, want)
	}
}
