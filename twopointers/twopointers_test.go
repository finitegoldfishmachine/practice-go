package twopointers

import "testing"

func TestOneDuplicate(t *testing.T) {
	example := []int{0, 0, 1, 1, 1, 2, 2}
	want := 3

	count := LengthOfUniques(example)

	if count != want {
		t.Fatalf(`LengthOfUniques(%#v), received %v, want  %v`, example, count, want)
	}
}

// Given an array of integers sorted in ascending order, find two numbers that add up to a given target. Return the indices of the two numbers in ascending order. You can assume elements in the array are unique and there is only one solution. Do this in O(n) time and with constant auxiliary space.
func TestTwoSumSorted(t *testing.T) {
	inputFoo := []int{2, 3, 4, 5, 8, 11, 18}
	inputBar := 8
	want := [2]int{1, 3}

	got := TwoSumSorted(inputFoo, inputBar)

	if got != want {
		t.Fatalf(`TwoSumSorted(%#v), received %v, want %v`, [2]any{inputFoo, inputBar}, got, want)
	}
}
