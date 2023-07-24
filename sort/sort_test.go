package sort

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	testCases := []struct {
		desc string
		nums []int
		want []int
	}{
		{
			desc: "list 1-5",
			nums: []int{5, 3, 1, 2, 4},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			desc: "list 1-10",
			nums: []int{8, 10, 1, 3, 4, 6, 9, 2, 7, 5},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			desc: "list bar",
			nums: []int{5, 10, 31, 5, 51, 25, 214, 176, 3, 353, 5, 499},
			want: []int{3, 5, 5, 5, 10, 25, 31, 51, 176, 214, 353, 499},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := SortList(tC.nums)
			if !reflect.DeepEqual(got, tC.want) {
				t.Fatalf("SortList(%#v) - got: %#v, want: %#v", tC.nums, got, tC.want)
			}
		})
	}
}

func TestInsertionSort(t *testing.T) {
	testCases := []struct {
		desc string
		nums []int
		want []int
	}{
		{
			desc: "list 1-5",
			nums: []int{5, 3, 1, 2, 4},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			desc: "list 1-10",
			nums: []int{8, 10, 1, 3, 4, 6, 9, 2, 7, 5},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			desc: "list bar",
			nums: []int{5, 10, 31, 5, 51, 25, 214, 176, 3, 353, 5, 499},
			want: []int{3, 5, 5, 5, 10, 25, 31, 51, 176, 214, 353, 499},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := InsertionSort(tC.nums)
			if !reflect.DeepEqual(got, tC.want) {
				t.Fatalf("SortList(%#v) - got: %#v, want: %#v", tC.nums, got, tC.want)
			}
		})
	}
}

func TestInsertionSortCondensed(t *testing.T) {
	testCases := []struct {
		desc string
		nums []int
		want []int
	}{
		{
			desc: "list 1-5",
			nums: []int{5, 3, 1, 2, 4},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			desc: "list 1-10",
			nums: []int{8, 10, 1, 3, 4, 6, 9, 2, 7, 5},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			desc: "list bar",
			nums: []int{5, 10, 31, 5, 51, 25, 214, 176, 3, 353, 5, 499},
			want: []int{3, 5, 5, 5, 10, 25, 31, 51, 176, 214, 353, 499},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := InsertionSortCondensed(tC.nums)
			if !reflect.DeepEqual(got, tC.want) {
				t.Fatalf("SortList(%#v) - got: %#v, want: %#v", tC.nums, got, tC.want)
			}
		})
	}
}
